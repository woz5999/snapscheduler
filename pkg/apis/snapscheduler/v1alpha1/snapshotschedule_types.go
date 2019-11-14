/*
Copyright (C) 2019  The snapscheduler authors

This file may be used, at your option, according to either the GNU AGPL 3.0 or
the Apache V2 license.

---
This program is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, either version 3 of the License, or (at your option) any
later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.  See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along
with this program.  If not, see <https://www.gnu.org/licenses/>.

---
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
	conditionsv1 "github.com/openshift/custom-resource-status/conditions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

// SnapshotRetentionSpec defines how long snapshots should be kept.
// +k8s:openapi-gen=true
type SnapshotRetentionSpec struct {
	// Expires is the length of time (time.Duration) after which a given
	// Snapshot will be deleted.
	// +kubebuilder:validation:Pattern=^\d+(h|m|s)$
	// +optional
	Expires string `json:"expires,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +optional
	MaxCount *int32 `json:"maxCount,omitempty"`
}

// SnapshotTemplateSpec defines the template for Snapshot objects
// +k8s:openapi-gen=true
type SnapshotTemplateSpec struct {
	// Labels is a list of labels that should be added to each Snapshot
	// created by this schedule.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	// SnapshotClassName is the name of the VSC to be used when creating
	// Snapshots.
	// +optional
	SnapshotClassName *string `json:"snapshotClassName,omitempty"`
}

// SnapshotScheduleSpec defines the desired state of SnapshotSchedule
// +k8s:openapi-gen=true
type SnapshotScheduleSpec struct {
	// ClaimSelector selects which PVCs will be snapshotted according to
	// this schedule.
	// +optional
	ClaimSelector metav1.LabelSelector `json:"claimSelector,omitempty"`
	// Retention determines how long this schedule's snapshots will be kept.
	// +optional
	Retention SnapshotRetentionSpec `json:"retention,omitempty"`
	// Schedule is a Cronspec specifying when snapshots should be taken. See
	// https://en.wikipedia.org/wiki/Cron for a description of the format.
	// +kubebuilder:validation:Pattern=^((\d+|\*)(/\d+)?(\s+(\d+|\*)(/\d+)?){4}|@(hourly|daily|weekly|monthly|yearly))$
	Schedule string `json:"schedule,omitempty"`
	// Disabled determines whether this schedule is currently disabled.
	// +optional
	Disabled bool `json:"disabled,omitempty"`
	// SnapshotTemplate is a template description of the Snapshots to be created.
	SnapshotTemplate *SnapshotTemplateSpec `json:"snapshotTemplate,omitempty"`
}

// SnapshotScheduleStatus defines the observed state of SnapshotSchedule
// +k8s:openapi-gen=true
type SnapshotScheduleStatus struct {
	// Conditions is a list of conditions related to operator reconciliation.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []conditionsv1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	// LastSnapshotTime is the time of the most recent set of snapshots
	// generated by this schedule.
	// +optional
	LastSnapshotTime *metav1.Time `json:"lastSnapshotTime,omitempty"`
	// NextSnapshotTime is the time when this schedule should create the
	// next set of snapshots.
	// +optional
	NextSnapshotTime *metav1.Time `json:"nextSnapshotTime,omitempty"`
}

const (
	// ConditionReconciled is a Condition indicating whether the object is fully
	// reconciled.
	ConditionReconciled conditionsv1.ConditionType = "Reconciled"
	// ReconciledReasonError indicates there was an error while attempting to reconcile.
	ReconciledReasonError = "ReconcileError"
	// ReconciledReasonComplete indicates reconcile was successful
	ReconciledReasonComplete = "ReconcileComplete"
)

// SnapshotSchedule is the Schema for the snapshotschedules API
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Schedule",type=string,JSONPath=".spec.schedule"
// +kubebuilder:printcolumn:name="Max age",type=string,JSONPath=".spec.retention.expires"
// +kubebuilder:printcolumn:name="Max num",type=integer,JSONPath=".spec.retention.maxCount"
// +kubebuilder:printcolumn:name="Disabled",type=boolean,JSONPath=".spec.disabled"
// +kubebuilder:printcolumn:name="Next snapshot",type=string,JSONPath=".status.nextSnapshotTime"
type SnapshotSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SnapshotScheduleSpec   `json:"spec,omitempty"`
	Status SnapshotScheduleStatus `json:"status,omitempty"`
}

// SnapshotScheduleList contains a list of SnapshotSchedule
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SnapshotScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotSchedule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SnapshotSchedule{}, &SnapshotScheduleList{})
}
