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
	corev1 "k8s.io/api/core/v1"
)

const (
	// MaxHealthScore is the max number of health score.
	MaxHealthScore = 100
)

// ClusterHealth represents the health of kubernetes cluster.
type ClusterHealth struct {
	// Score is a weighted score of kubernetes cluster health.
	Score int
	// WorkloadHealth represents the health of workloads in kubernetes cluster.
	WorkloadHealth WorkloadHealth
	// NodeHealth represents the health of nodes in kubernetes cluster.
	NodeHealth NodeHealth
	// ControlPlaneComponentHealth represents the health of control plane components in kubernetes cluster.
	ControlPlaneComponentHealth ControlPlaneComponentHealth
	// NodeComponentHealth represents the health of node components in kubernetes cluster.
	NodeComponentHealth NodeComponentHealth
}

// WorkloadHealth represents the health of workloads in kubernetes cluster.
type WorkloadHealth struct {
	// Score is a weighted score of workload health.
	Score int
	// PodHealth represents the health of pods in kubernetes cluster.
	PodHealth PodHealth
}

// PodHealth represents the health of pods in kubernetes cluster.
type PodHealth struct {
	// Score is a weighted score of pod health.
	Score int
	// Statistics contains information about healthy and unhealthy pods.
	Statistics PodStatistics
}

// PodStatistics contains information about healthy and unhealthy pods.
type PodStatistics struct {
	// Total is the total number of pods in kubernetes cluster.
	Total int
	// Healthy contains information about healthy pods.
	Healthy HealthyPodStatistics
	// Unhealthy contains information about unhealthy pods.
	Unhealthy UnhealthyPodStatistics
}

// HealthyPodStatistics contains information about healthy pods. The are two types of healthy pods:
//
// Ready: The pod is in Running phase and all containers in the pod are ready.
// Succeeded: The pod is in Succeeded phase.
type HealthyPodStatistics struct {
	// Ready is the number of pods in Running phase and all containers in the pod are ready.
	Ready int
	// Succeeded is the number of pods in Succeeded phase.
	Succeeded int
}

// UnhealthyPodStatistics contains information about unhealthy pods. The are five types of unhealthy pods:
//
// Unready: The pod is in Running phase and some container in the pod is not ready.
// Terminating: The pod in Running phase and has a DeletionTimestamp.
// Pending: The pod is in Pending phase.
// Failed: The pod is in Failed phase.
// Unknown: The pod is in Unknown phase.
type UnhealthyPodStatistics struct {
	// Unready is the number of pods in Running phase and some container in the pod is not ready.
	Unready int
	// Terminating is the number of pods in Running phase and has a DeletionTimestamp.
	Terminating int
	// Pending is the number of pods in Pending phase.
	Pending int
	// Failed is the number of pods in Failed phase.
	Failed int
	// Unknown is the number of pods in Unknown phase.
	Unknown int
	// ContainerStateReasons statisticizes reasons of unhealthy containers in pods. The key is the reason
	// of the first waiting or terminated container and the value is the count of that reason.
	// The following are possible reasons of a waiting or terminated container:
	//
	// CrashLoopBackOff
	// RunContainerError
	// KillContainerError
	// VerifyNonRootError
	// RunInitContainerError
	// CreatePodSandboxError
	// ConfigPodSandboxError
	// KillPodSandboxError
	// SetupNetworkError
	// TeardownNetworkError
	// OOMKilled
	// Error
	// ContainerCannotRun
	ContainerStateReasons map[string]int
}

// NodeHealth represents the health of nodes in kubernetes cluster.
type NodeHealth struct {
	// Score is a weighted score of node health.
	Score int
	// Statistics contains information about healthy and unhealthy nodes.
	Statistics NodeStatistics
}

// NodeStatistics contains information about healthy and unhealthy nodes.
type NodeStatistics struct {
	// Total is the total number of nodes in kubernetes cluster.
	Total int
	// Healthy contains information about healthy nodes. The is one condition type of a healthy node:
	//
	// Ready: The node is in Ready condition.
	Healthy int
	// Unhealthy contains information about unhealthy nodes. The key is the first unhealthy condition type and
	// the value is the count of that condition type. The following are possible types of an unhealthy node:
	//
	// OutOfDisk: The node is in OutOfDisk condition.
	// MemoryPressure: The node is in MemoryPressure condition.
	// DiskPressure: The node is in DiskPressure condition.
	// PIDPressure: The node is in PIDPressure condition.
	// NetworkUnavailable: The node is in NetworkUnavailable condition.
	// Unknown: The node does not report any condition.
	Unhealthy map[corev1.NodeConditionType]int
}

// ControlPlaneComponentHealth represents the health of control plane components.
type ControlPlaneComponentHealth struct {
	// Score is a weighted score of control plane component health.
	Score int
	// EtcdHealth represents the health of etcd in kubernetes cluster.
	EtcdHealth EtcdHealth
	// APIServerHealth represents the health of apiserver in kubernetes cluster.
	APIServerHealth APIServerHealth
	// ControllerManagerHealth represents the health of controller manager in kubernetes cluster.
	ControllerManagerHealth ControllerManagerHealth
	// SchedulerHealth represents the health of scheduler in kubernetes cluster.
	SchedulerHealth SchedulerHealth
}

// EtcdHealth represents the health of etcd in kubernetes cluster.
type EtcdHealth struct {
	// Score is a weighted score of etcd health.
	Score int
}

// APIServerHealth represents the health of apiserver in kubernetes cluster.
type APIServerHealth struct {
	// Score is a weighted score of apiserver health.
	Score int
}

// ControllerManagerHealth represents the health of controller manager in kubernetes cluster.
type ControllerManagerHealth struct {
	// Score is a weighted score of controller manager health.
	Score int
}

// SchedulerHealth represents the health of scheduler in kubernetes cluster.
type SchedulerHealth struct {
	// Score is a weighted score of scheduler health.
	Score int
}

// NodeComponentHealth represents the health of node components.
type NodeComponentHealth struct {
	// Score is a weighted score of node component health.
	Score int
	// KubeletHealth represents the health of kubelet in kubernetes cluster.
	KubeletHealth KubeletHealth
	// KubeProxyHealth represents the health of kube proxy in kubernetes cluster.
	KubeProxyHealth KubeProxyHealth
}

// KubeletHealth represents the health of kubelet in kubernetes cluster.
type KubeletHealth struct {
	// Score is a weighted score of kubelet health.
	Score int
}

// KubeProxyHealth represents the health of kube proxy in kubernetes cluster.
type KubeProxyHealth struct {
	// Score is a weighted score of kube proxy health.
	Score int
}