/*
Copyright 2025.

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

package v1alpha1

import (
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ScheduledJobSpec defines the desired state of ScheduledJob
type ScheduledJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Schedule string `json:"schedule"`

	// The time zone name for the given schedule, see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.
	// If not specified, this will default to the time zone of the kube-controller-manager process.
	// The set of valid time zone names and the time zone offset is loaded from the system-wide time zone
	// database by the API server during ScheduledJob validation and the controller manager during execution.
	// If no system-wide time zone database can be found a bundled version of the database is used instead.
	// If the time zone name becomes invalid during the lifetime of a ScheduledJob or due to a change in host
	// configuration, the controller will stop creating new new Jobs and will create a system event with the
	// reason UnknownTimeZone.
	// More information can be found in https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#time-zones
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	TimeZone *string `json:"timeZone,omitempty"`

	// This flag tells the controller to suspend subsequent executions, it does
	// not apply to already started executions.  Defaults to false.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Suspend *bool `json:"suspend,omitempty"`

	// Optional deadline in seconds for starting the job if it misses scheduled
	// time for any reason.  Missed jobs executions will be counted as failed ones.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds,omitempty"`

	// Specifies the job that will be created when executing a ScheduledJob.
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	JobTemplate v1.JobTemplateSpec `json:"jobTemplate"`

	// The number of successful finished jobs to retain. Value must be non-negative integer.
	// Defaults to 1.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit,omitempty"`

	// The number of failed finished jobs to retain. Value must be non-negative integer.
	// Defaults to 1.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit,omitempty"`
}

// ScheduledJobStatus defines the observed state of ScheduledJob
type ScheduledJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Conditions store the status conditions of created Jobs
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ScheduledJob is the Schema for the scheduledjobs API
// +kubebuilder:subresource:status
type ScheduledJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScheduledJobSpec   `json:"spec,omitempty"`
	Status ScheduledJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduledJobList contains a list of ScheduledJob
type ScheduledJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScheduledJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ScheduledJob{}, &ScheduledJobList{})
}
