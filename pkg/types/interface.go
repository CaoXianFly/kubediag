/*
Copyright 2020 The Kube Diagnoser Authors.

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

package types

import (
	"context"
	"net/http"

	"github.com/go-logr/logr"

	diagnosisv1 "netease.com/k8s/kube-diagnoser/api/v1"
)

// AbnormalProcessor manages http requests for processing abnormals.
type AbnormalProcessor interface {
	// Context carries values across API boundaries.
	context.Context
	// Logger represents the ability to log messages.
	logr.Logger
	// Handler handles http requests.
	Handler(http.ResponseWriter, *http.Request)
}

// AbnormalManager manages processors for processing abnormals.
type AbnormalManager interface {
	// Context carries values across API boundaries.
	context.Context
	// Logger represents the ability to log messages.
	logr.Logger
	// Run runs the AbnormalManager.
	Run(<-chan struct{})
	// SyncAbnormal syncs abnormals.
	SyncAbnormal(diagnosisv1.Abnormal) (diagnosisv1.Abnormal, error)
	// Handler handles http requests.
	Handler(http.ResponseWriter, *http.Request)
}