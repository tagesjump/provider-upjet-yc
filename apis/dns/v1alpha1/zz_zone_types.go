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

type ZoneInitParameters struct {

	// Description of the DNS zone.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of key/value label pairs to assign to the DNS zone.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// User assigned name of a specific resource. Must be unique within the folder.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// The DNS name of this zone, e.g. "example.com.". Must ends with dot.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ZoneObservation struct {

	// (Computed) The DNS zone creation timestamp.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Description of the DNS zone.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder to create a zone in. If it is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// (Computed) ID of a new DNS zone.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the DNS zone.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// User assigned name of a specific resource. Must be unique within the folder.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
	PrivateNetworks []*string `json:"privateNetworks,omitempty" tf:"private_networks,omitempty"`

	// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// The DNS name of this zone, e.g. "example.com.". Must ends with dot.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ZoneParameters struct {

	// Description of the DNS zone.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder to create a zone in. If it is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the DNS zone.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// User assigned name of a specific resource. Must be unique within the folder.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	PrivateNetworks []*string `json:"privateNetworks,omitempty" tf:"private_networks,omitempty"`

	// References to Network in vpc to populate privateNetworks.
	// +kubebuilder:validation:Optional
	PrivateNetworksRefs []v1.Reference `json:"privateNetworksRefs,omitempty" tf:"-"`

	// Selector for a list of Network in vpc to populate privateNetworks.
	// +kubebuilder:validation:Optional
	PrivateNetworksSelector *v1.Selector `json:"privateNetworksSelector,omitempty" tf:"-"`

	// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.
	// +kubebuilder:validation:Optional
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// The DNS name of this zone, e.g. "example.com.". Must ends with dot.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// ZoneSpec defines the desired state of Zone
type ZoneSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ZoneParameters `json:"forProvider"`
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
	InitProvider ZoneInitParameters `json:"initProvider,omitempty"`
}

// ZoneStatus defines the observed state of Zone.
type ZoneStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Zone is the Schema for the Zones API. Manages a DNS Zone within Yandex.Cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Zone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zone) || (has(self.initProvider) && has(self.initProvider.zone))",message="spec.forProvider.zone is a required parameter"
	Spec   ZoneSpec   `json:"spec"`
	Status ZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneList contains a list of Zones
type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Zone `json:"items"`
}

// Repository type metadata.
var (
	Zone_Kind             = "Zone"
	Zone_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Zone_Kind}.String()
	Zone_KindAPIVersion   = Zone_Kind + "." + CRDGroupVersion.String()
	Zone_GroupVersionKind = CRDGroupVersion.WithKind(Zone_Kind)
)

func init() {
	SchemeBuilder.Register(&Zone{}, &ZoneList{})
}
