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

type DiskInitParameters struct {

	// Default is 5 minutes.
	AllowRecreate *bool `json:"allowRecreate,omitempty" tf:"allow_recreate,omitempty"`

	// Block size of the disk, specified in bytes.
	BlockSize *float64 `json:"blockSize,omitempty" tf:"block_size,omitempty"`

	// Description of the disk. Provide this property when
	// you create a resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disk placement policy configuration. The structure is documented below.
	DiskPlacementPolicy []DiskPlacementPolicyInitParameters `json:"diskPlacementPolicy,omitempty" tf:"disk_placement_policy,omitempty"`

	// The ID of the folder that the disk belongs to.
	// If it is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// The source image to use for disk creation.
	// +crossplane:generate:reference:type=Image
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// Reference to a Image to populate imageId.
	// +kubebuilder:validation:Optional
	ImageIDRef *v1.Reference `json:"imageIdRef,omitempty" tf:"-"`

	// Selector for a Image to populate imageId.
	// +kubebuilder:validation:Optional
	ImageIDSelector *v1.Selector `json:"imageIdSelector,omitempty" tf:"-"`

	// Labels to assign to this disk. A list of key/value pairs.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the disk. Provide this property when
	// you create a resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the image_id or snapshot_id
	// parameter, or specify it alone to create an empty persistent disk.
	// If you specify this field along with image_id or snapshot_id,
	// the size value must not be less than the size of the source image
	// or the size of the snapshot.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The source snapshot to use for disk creation.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// Type of disk to create. Provide this when creating a disk.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Availability zone where the disk will reside.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DiskObservation struct {

	// Default is 5 minutes.
	AllowRecreate *bool `json:"allowRecreate,omitempty" tf:"allow_recreate,omitempty"`

	// Block size of the disk, specified in bytes.
	BlockSize *float64 `json:"blockSize,omitempty" tf:"block_size,omitempty"`

	// Creation timestamp of the disk.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Description of the disk. Provide this property when
	// you create a resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disk placement policy configuration. The structure is documented below.
	DiskPlacementPolicy []DiskPlacementPolicyObservation `json:"diskPlacementPolicy,omitempty" tf:"disk_placement_policy,omitempty"`

	// The ID of the folder that the disk belongs to.
	// If it is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The source image to use for disk creation.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// Labels to assign to this disk. A list of key/value pairs.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the disk. Provide this property when
	// you create a resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	ProductIds []*string `json:"productIds,omitempty" tf:"product_ids,omitempty"`

	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the image_id or snapshot_id
	// parameter, or specify it alone to create an empty persistent disk.
	// If you specify this field along with image_id or snapshot_id,
	// the size value must not be less than the size of the source image
	// or the size of the snapshot.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The source snapshot to use for disk creation.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// The status of the disk.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Type of disk to create. Provide this when creating a disk.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Availability zone where the disk will reside.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DiskParameters struct {

	// Default is 5 minutes.
	// +kubebuilder:validation:Optional
	AllowRecreate *bool `json:"allowRecreate,omitempty" tf:"allow_recreate,omitempty"`

	// Block size of the disk, specified in bytes.
	// +kubebuilder:validation:Optional
	BlockSize *float64 `json:"blockSize,omitempty" tf:"block_size,omitempty"`

	// Description of the disk. Provide this property when
	// you create a resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disk placement policy configuration. The structure is documented below.
	// +kubebuilder:validation:Optional
	DiskPlacementPolicy []DiskPlacementPolicyParameters `json:"diskPlacementPolicy,omitempty" tf:"disk_placement_policy,omitempty"`

	// The ID of the folder that the disk belongs to.
	// If it is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// The source image to use for disk creation.
	// +crossplane:generate:reference:type=Image
	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// Reference to a Image to populate imageId.
	// +kubebuilder:validation:Optional
	ImageIDRef *v1.Reference `json:"imageIdRef,omitempty" tf:"-"`

	// Selector for a Image to populate imageId.
	// +kubebuilder:validation:Optional
	ImageIDSelector *v1.Selector `json:"imageIdSelector,omitempty" tf:"-"`

	// Labels to assign to this disk. A list of key/value pairs.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the disk. Provide this property when
	// you create a resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Size of the persistent disk, specified in GB. You can specify this
	// field when creating a persistent disk using the image_id or snapshot_id
	// parameter, or specify it alone to create an empty persistent disk.
	// If you specify this field along with image_id or snapshot_id,
	// the size value must not be less than the size of the source image
	// or the size of the snapshot.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The source snapshot to use for disk creation.
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// Type of disk to create. Provide this when creating a disk.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Availability zone where the disk will reside.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DiskPlacementPolicyInitParameters struct {

	// Specifies Disk Placement Group id.
	DiskPlacementGroupID *string `json:"diskPlacementGroupId,omitempty" tf:"disk_placement_group_id,omitempty"`
}

type DiskPlacementPolicyObservation struct {

	// Specifies Disk Placement Group id.
	DiskPlacementGroupID *string `json:"diskPlacementGroupId,omitempty" tf:"disk_placement_group_id,omitempty"`
}

type DiskPlacementPolicyParameters struct {

	// Specifies Disk Placement Group id.
	// +kubebuilder:validation:Optional
	DiskPlacementGroupID *string `json:"diskPlacementGroupId" tf:"disk_placement_group_id,omitempty"`
}

// DiskSpec defines the desired state of Disk
type DiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskParameters `json:"forProvider"`
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
	InitProvider DiskInitParameters `json:"initProvider,omitempty"`
}

// DiskStatus defines the observed state of Disk.
type DiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Disk is the Schema for the Disks API. Persistent disks are durable storage devices that function similarly to the physical disks in a desktop or a server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Disk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskSpec   `json:"spec"`
	Status            DiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskList contains a list of Disks
type DiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Disk `json:"items"`
}

// Repository type metadata.
var (
	Disk_Kind             = "Disk"
	Disk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Disk_Kind}.String()
	Disk_KindAPIVersion   = Disk_Kind + "." + CRDGroupVersion.String()
	Disk_GroupVersionKind = CRDGroupVersion.WithKind(Disk_Kind)
)

func init() {
	SchemeBuilder.Register(&Disk{}, &DiskList{})
}
