/*
Copyright 2022 The KubeDiag Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package diagnoser

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"github.com/vmware-tanzu/sonobuoy/pkg/client/results"
	"gopkg.in/yaml.v2"

	"github.com/kubediag/kubediag/pkg/executor"
	"github.com/kubediag/kubediag/pkg/processors"
	"github.com/kubediag/kubediag/pkg/processors/utils"
	"github.com/kubediag/kubediag/pkg/util"
)

const (
	SonobuoyResultDiagnoserTemporaryResultDirectory = "sonobuoy_result_diagnoser.temporary_result_directory"
	SonobuoyResultDiagnoserE2eFile                  = "sonobuoy_result_diagnoser.e2e_file"

	SonobuoyResultDiagnoserSummary         = "sonobuoy_result_diagnoser.summary"
	SonobuoyResultDiagnoserResultDirectory = "sonobuoy_result_diagnoser.result_directory"
)

// Summary contains the categories of sonobuoy result items.
type Summary []Category

// Category contains sonobuoy result items under the same category.
type Category struct {
	CategoryName string       `json:"category_name"`
	StatusCounts StatusCounts `json:"status_counts"`
	Items        []Item       `json:"items"`
}

// StatusCounts contains the count of different status.
type StatusCounts map[string]int

// Item is the shortened version of results.Item https://github.com/vmware-tanzu/sonobuoy/blob/v0.17.2/pkg/client/results/processing.go#L94.
type Item struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// sonobuoyResultDiagnoser analyze the e2e test result generated by sonobuoy.
type sonobuoyResultDiagnoser struct {
	// Context carries values across API boundaries.
	context.Context
	// Logger represents the ability to log messages.
	logr.Logger
	// dataRoot is root directory of persistent kubediag data.
	dataRoot string
	// bindAddress is the address on which to advertise.
	bindAddress string
	// sonobuoyResultDiagnoserEnabled indicates whether sonobuoyResultDiagnoser is enabled.
	sonobuoyResultDiagnoserEnabled bool
}

// NewSonobuoyResultDiagnoser creates a new sonobuoyResultDiagnoser.
func NewSonobuoyResultDiagnoser(
	ctx context.Context,
	logger logr.Logger,
	dataRoot string,
	bindAddress string,
	sonobuoyResultDiagnoserEnabled bool,
) processors.Processor {
	return &sonobuoyResultDiagnoser{
		Context:                        ctx,
		Logger:                         logger,
		dataRoot:                       dataRoot,
		bindAddress:                    bindAddress,
		sonobuoyResultDiagnoserEnabled: sonobuoyResultDiagnoserEnabled,
	}
}

// Handler handles http requests for sonobuoy result diagnoser.
func (s *sonobuoyResultDiagnoser) Handler(w http.ResponseWriter, r *http.Request) {
	if !s.sonobuoyResultDiagnoserEnabled {
		http.Error(w, "sonobuoy result diagnoser is not enabled", http.StatusUnprocessableEntity)
		return
	}

	switch r.Method {
	case "POST":
		s.Info("handle POST request")
		contexts, err := utils.ExtractParametersFromHTTPContext(r)
		if err != nil {
			s.Error(err, "extract contexts failed")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Handle sonobuoy dump result.
		tmpResultDir := contexts[SonobuoyResultDiagnoserTemporaryResultDirectory]
		e2eFile := contexts[SonobuoyResultDiagnoserE2eFile]
		summary := s.getResultSummary(tmpResultDir, e2eFile)

		// TODO: Functionalize diagnosis result directory name generating.
		diagnosisNamespace := contexts[executor.DiagnosisNamespaceTelemetryKey]
		diagnosisName := contexts[executor.DiagnosisNameTelemetryKey]
		timestamp := strconv.Itoa(int(time.Now().Unix()))
		diagnosisResultDir := strings.Join([]string{diagnosisNamespace, diagnosisName, timestamp}, "_")

		resultDir := filepath.Join(s.dataRoot, "diagnoses", diagnosisResultDir)
		err = util.MoveFiles(tmpResultDir, resultDir)
		if err != nil {
			s.Error(err, "move files failed")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		raw, err := json.Marshal(summary)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to marshal sonobuoy dump result: %v", err), http.StatusInternalServerError)
			return
		}

		result := make(map[string]string)
		result[SonobuoyResultDiagnoserSummary] = string(raw)
		result[SonobuoyResultDiagnoserResultDirectory] = resultDir
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to marshal sonobuoy result diagnoser result: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	default:
		http.Error(w, fmt.Sprintf("method %s is not supported", r.Method), http.StatusMethodNotAllowed)
	}
}

// getResultSummary gets sonobuoy dump result and store processed result in Summary.
func (s *sonobuoyResultDiagnoser) getResultSummary(tmpResultDir string, e2eFile string) Summary {
	// Unmarshal dump mode e2e result file.
	var item results.Item
	filePath := filepath.Join(tmpResultDir, e2eFile)
	byteValue, err := ioutil.ReadFile(filePath)
	if err != nil {
		s.Logger.Error(err, fmt.Sprintf("failed to find file: %s", filePath))
		return nil
	}
	err = yaml.Unmarshal([]byte(byteValue), &item)
	if err != nil {
		s.Logger.Error(err, "failed to unmarshal yaml file")
		return nil
	}

	categoryMap := make(map[string][]Item)
	statusCountMap := make(map[string]StatusCounts)
	categoryMap, statusCountMap = walkForSummary(&item, categoryMap, statusCountMap)

	summary := make(Summary, 0)
	for categoryName, items := range categoryMap {
		category := Category{
			CategoryName: categoryName,
			Items:        items,
		}

		statusCounts, ok := statusCountMap[categoryName]
		if ok {
			category.StatusCounts = statusCounts
		}

		summary = append(summary, category)
	}

	return summary
}

// walkForSummary walk for summary of status.
func walkForSummary(item *results.Item, categoryMap map[string][]Item, statusCountMap map[string]StatusCounts) (map[string][]Item, map[string]StatusCounts) {
	if len(item.Items) > 0 {
		for _, item := range item.Items {
			categoryMap, statusCountMap = walkForSummary(&item, categoryMap, statusCountMap)
		}
		return categoryMap, statusCountMap
	}

	if item.Status != results.StatusSkipped {
		categoryName := getCategoryNameFromItem(item)
		if _, ok := categoryMap[categoryName]; !ok {
			categoryMap[categoryName] = make([]Item, 0)
		}
		categoryMap[categoryName] = append(categoryMap[categoryName], Item{Name: item.Name, Status: item.Status})

		if _, ok := statusCountMap[categoryName]; !ok {
			statusCountMap[categoryName] = make(StatusCounts)
		}
		statusCountMap[categoryName][item.Status]++
	}

	return categoryMap, statusCountMap
}

// getCategoryNameFromItem retrieves the category name of an item.
// It returns the first string slice starts with "[" and ends with "]".
// Otherwise, it returns "other" as the category.
func getCategoryNameFromItem(item *results.Item) string {
	if strings.HasPrefix(item.Name, "[") {
		strs := strings.Split(item.Name, "]")
		if len(strs) > 1 {
			return strs[0] + "]"
		}
	}

	return "other"
}