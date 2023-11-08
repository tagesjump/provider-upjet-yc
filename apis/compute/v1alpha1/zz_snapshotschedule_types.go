// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SchedulePolicyInitParameters struct {

	// Cron expression to schedule snapshots (in cron format "* * * * *").
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Time to start the snapshot schedule (in format RFC3339 "2006-01-02T15:04:05Z07:00"). If empty current time will be used. Unlike an expression that specifies regularity rules, the start_at parameter determines from what point these rules will be applied.
	StartAt *string `json:"startAt,omitempty" tf:"start_at,omitempty"`
}

type SchedulePolicyObservation struct {

	// Cron expression to schedule snapshots (in cron format "* * * * *").
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Time to start the snapshot schedule (in format RFC3339 "2006-01-02T15:04:05Z07:00"). If empty current time will be used. Unlike an expression that specifies regularity rules, the start_at parameter determines from what point these rules will be applied.
	StartAt *string `json:"startAt,omitempty" tf:"start_at,omitempty"`
}

type SchedulePolicyParameters struct {

	// Cron expression to schedule snapshots (in cron format "* * * * *").
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Time to start the snapshot schedule (in format RFC3339 "2006-01-02T15:04:05Z07:00"). If empty current time will be used. Unlike an expression that specifies regularity rules, the start_at parameter determines from what point these rules will be applied.
	// +kubebuilder:validation:Optional
	StartAt *string `json:"startAt,omitempty" tf:"start_at,omitempty"`
}

type SnapshotScheduleInitParameters struct {

	// Description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of key/value label pairs to assign to the snapshot schedule.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A name for the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Time duration applied to snapshots created by this snapshot schedule. This is a signed sequence of decimal numbers, each with optional fraction and a unit suffix. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Examples: "300ms", "1.5h" or "2h45m".
	RetentionPeriod *string `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`

	// Schedule policy of the snapshot schedule.
	SchedulePolicy []SchedulePolicyInitParameters `json:"schedulePolicy,omitempty" tf:"schedule_policy,omitempty"`

	// Maximum number of snapshots for every disk of the snapshot schedule.
	SnapshotCount *float64 `json:"snapshotCount,omitempty" tf:"snapshot_count,omitempty"`

	// Additional attributes for snapshots created by this snapshot schedule.
	SnapshotSpec []SnapshotSpecInitParameters `json:"snapshotSpec,omitempty" tf:"snapshot_spec,omitempty"`
}

type SnapshotScheduleObservation struct {

	// Creation timestamp of the snapshot schedule.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// IDs of the disk for snapshot schedule.
	DiskIds []*string `json:"diskIds,omitempty" tf:"disk_ids,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the snapshot schedule.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A name for the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Time duration applied to snapshots created by this snapshot schedule. This is a signed sequence of decimal numbers, each with optional fraction and a unit suffix. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Examples: "300ms", "1.5h" or "2h45m".
	RetentionPeriod *string `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`

	// Schedule policy of the snapshot schedule.
	SchedulePolicy []SchedulePolicyObservation `json:"schedulePolicy,omitempty" tf:"schedule_policy,omitempty"`

	// Maximum number of snapshots for every disk of the snapshot schedule.
	SnapshotCount *float64 `json:"snapshotCount,omitempty" tf:"snapshot_count,omitempty"`

	// Additional attributes for snapshots created by this snapshot schedule.
	SnapshotSpec []SnapshotSpecObservation `json:"snapshotSpec,omitempty" tf:"snapshot_spec,omitempty"`

	// The status of the snapshot schedule.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SnapshotScheduleParameters struct {

	// Description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// IDs of the disk for snapshot schedule.
	// +crossplane:generate:reference:type=Disk
	// +kubebuilder:validation:Optional
	DiskIds []*string `json:"diskIds,omitempty" tf:"disk_ids,omitempty"`

	// References to Disk to populate diskIds.
	// +kubebuilder:validation:Optional
	DiskIdsRefs []v1.Reference `json:"diskIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Disk to populate diskIds.
	// +kubebuilder:validation:Optional
	DiskIdsSelector *v1.Selector `json:"diskIdsSelector,omitempty" tf:"-"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the snapshot schedule.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A name for the resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Time duration applied to snapshots created by this snapshot schedule. This is a signed sequence of decimal numbers, each with optional fraction and a unit suffix. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Examples: "300ms", "1.5h" or "2h45m".
	// +kubebuilder:validation:Optional
	RetentionPeriod *string `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`

	// Schedule policy of the snapshot schedule.
	// +kubebuilder:validation:Optional
	SchedulePolicy []SchedulePolicyParameters `json:"schedulePolicy,omitempty" tf:"schedule_policy,omitempty"`

	// Maximum number of snapshots for every disk of the snapshot schedule.
	// +kubebuilder:validation:Optional
	SnapshotCount *float64 `json:"snapshotCount,omitempty" tf:"snapshot_count,omitempty"`

	// Additional attributes for snapshots created by this snapshot schedule.
	// +kubebuilder:validation:Optional
	SnapshotSpec []SnapshotSpecParameters `json:"snapshotSpec,omitempty" tf:"snapshot_spec,omitempty"`
}

type SnapshotSpecInitParameters struct {

	// Description to assign to snapshots created by this snapshot schedule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of key/value label pairs to assign to snapshots created by this snapshot schedule.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type SnapshotSpecObservation struct {

	// Description to assign to snapshots created by this snapshot schedule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of key/value label pairs to assign to snapshots created by this snapshot schedule.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type SnapshotSpecParameters struct {

	// Description to assign to snapshots created by this snapshot schedule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of key/value label pairs to assign to snapshots created by this snapshot schedule.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

// SnapshotScheduleSpec defines the desired state of SnapshotSchedule
type SnapshotScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotScheduleParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SnapshotScheduleInitParameters `json:"initProvider,omitempty"`
}

// SnapshotScheduleStatus defines the observed state of SnapshotSchedule.
type SnapshotScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotSchedule is the Schema for the SnapshotSchedules API. Creates a new snapshot schedule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type SnapshotSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotScheduleSpec   `json:"spec"`
	Status            SnapshotScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotScheduleList contains a list of SnapshotSchedules
type SnapshotScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotSchedule `json:"items"`
}

// Repository type metadata.
var (
	SnapshotSchedule_Kind             = "SnapshotSchedule"
	SnapshotSchedule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnapshotSchedule_Kind}.String()
	SnapshotSchedule_KindAPIVersion   = SnapshotSchedule_Kind + "." + CRDGroupVersion.String()
	SnapshotSchedule_GroupVersionKind = CRDGroupVersion.WithKind(SnapshotSchedule_Kind)
)

func init() {
	SchemeBuilder.Register(&SnapshotSchedule{}, &SnapshotScheduleList{})
}
