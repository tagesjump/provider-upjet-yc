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

type TargetGroupInitParameters struct {

	// An optional description of the target group. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Labels to assign to this target group. A list of key/value pairs.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the target group. Provided by the client when the target group is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the availability zone where the target group resides.
	// If omitted, default region is being used.
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// A Target resource. The structure is documented below.
	Target []TargetInitParameters `json:"target,omitempty" tf:"target,omitempty"`
}

type TargetGroupObservation struct {

	// The target group creation timestamp.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// An optional description of the target group. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder to which the resource belongs.
	// If omitted, the provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// The ID of the target group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Labels to assign to this target group. A list of key/value pairs.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the target group. Provided by the client when the target group is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the availability zone where the target group resides.
	// If omitted, default region is being used.
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// A Target resource. The structure is documented below.
	Target []TargetObservation `json:"target,omitempty" tf:"target,omitempty"`
}

type TargetGroupParameters struct {

	// An optional description of the target group. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder to which the resource belongs.
	// If omitted, the provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// Labels to assign to this target group. A list of key/value pairs.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the target group. Provided by the client when the target group is created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the availability zone where the target group resides.
	// If omitted, default region is being used.
	// +kubebuilder:validation:Optional
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// A Target resource. The structure is documented below.
	// +kubebuilder:validation:Optional
	Target []TargetParameters `json:"target,omitempty" tf:"target,omitempty"`
}

type TargetInitParameters struct {

	// IP address of the target.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`
}

type TargetObservation struct {

	// IP address of the target.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// ID of the subnet that targets are connected to.
	// All targets in the target group must be connected to the same subnet within a single availability zone.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type TargetParameters struct {

	// IP address of the target.
	// +kubebuilder:validation:Optional
	Address *string `json:"address" tf:"address,omitempty"`

	// ID of the subnet that targets are connected to.
	// All targets in the target group must be connected to the same subnet within a single availability zone.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// TargetGroupSpec defines the desired state of TargetGroup
type TargetGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetGroupParameters `json:"forProvider"`
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
	InitProvider TargetGroupInitParameters `json:"initProvider,omitempty"`
}

// TargetGroupStatus defines the observed state of TargetGroup.
type TargetGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TargetGroup is the Schema for the TargetGroups API. A load balancer distributes the load across cloud resources that are combined into a target group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type TargetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TargetGroupSpec   `json:"spec"`
	Status            TargetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetGroupList contains a list of TargetGroups
type TargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetGroup `json:"items"`
}

// Repository type metadata.
var (
	TargetGroup_Kind             = "TargetGroup"
	TargetGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TargetGroup_Kind}.String()
	TargetGroup_KindAPIVersion   = TargetGroup_Kind + "." + CRDGroupVersion.String()
	TargetGroup_GroupVersionKind = CRDGroupVersion.WithKind(TargetGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&TargetGroup{}, &TargetGroupList{})
}
