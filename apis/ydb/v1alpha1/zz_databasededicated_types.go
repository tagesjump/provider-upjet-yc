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

type DatabaseDedicatedInitParameters struct {

	// Whether public IP addresses should be assigned to the Yandex Database cluster.
	AssignPublicIps *bool `json:"assignPublicIps,omitempty" tf:"assign_public_ips,omitempty"`

	// Inhibits deletion of the database. Can be either true or false
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A description for the Yandex Database cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder that the Yandex Database cluster belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the Yandex Database cluster.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Location for the Yandex Database cluster.
	// The structure is documented below.
	Location []LocationInitParameters `json:"location,omitempty" tf:"location,omitempty"`

	// Location ID for the Yandex Database cluster.
	LocationID *string `json:"locationId,omitempty" tf:"location_id,omitempty"`

	// Name of the Yandex Database cluster.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network to attach the Yandex Database cluster to.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Network
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// The Yandex Database cluster preset.
	// Available presets can be obtained via yc ydb resource-preset list command.
	ResourcePresetID *string `json:"resourcePresetId,omitempty" tf:"resource_preset_id,omitempty"`

	// Scaling policy for the Yandex Database cluster.
	// The structure is documented below.
	ScalePolicy []ScalePolicyInitParameters `json:"scalePolicy,omitempty" tf:"scale_policy,omitempty"`

	// A list of storage configuration options for the Yandex Database cluster.
	// The structure is documented below.
	StorageConfig []StorageConfigInitParameters `json:"storageConfig,omitempty" tf:"storage_config,omitempty"`

	// List of subnet IDs to attach the Yandex Database cluster to.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Subnet
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// References to Subnet in vpc to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsRefs []v1.Reference `json:"subnetIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in vpc to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsSelector *v1.Selector `json:"subnetIdsSelector,omitempty" tf:"-"`
}

type DatabaseDedicatedObservation struct {

	// Whether public IP addresses should be assigned to the Yandex Database cluster.
	AssignPublicIps *bool `json:"assignPublicIps,omitempty" tf:"assign_public_ips,omitempty"`

	// The Yandex Database cluster creation timestamp.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Full database path of the Yandex Database cluster.
	// Useful for SDK configuration.
	DatabasePath *string `json:"databasePath,omitempty" tf:"database_path,omitempty"`

	// Inhibits deletion of the database. Can be either true or false
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A description for the Yandex Database cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder that the Yandex Database cluster belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// ID of the Yandex Database cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the Yandex Database cluster.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Location for the Yandex Database cluster.
	// The structure is documented below.
	Location []LocationObservation `json:"location,omitempty" tf:"location,omitempty"`

	// Location ID for the Yandex Database cluster.
	LocationID *string `json:"locationId,omitempty" tf:"location_id,omitempty"`

	// Name of the Yandex Database cluster.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network to attach the Yandex Database cluster to.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// The Yandex Database cluster preset.
	// Available presets can be obtained via yc ydb resource-preset list command.
	ResourcePresetID *string `json:"resourcePresetId,omitempty" tf:"resource_preset_id,omitempty"`

	// Scaling policy for the Yandex Database cluster.
	// The structure is documented below.
	ScalePolicy []ScalePolicyObservation `json:"scalePolicy,omitempty" tf:"scale_policy,omitempty"`

	// Status of the Yandex Database cluster.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A list of storage configuration options for the Yandex Database cluster.
	// The structure is documented below.
	StorageConfig []StorageConfigObservation `json:"storageConfig,omitempty" tf:"storage_config,omitempty"`

	// List of subnet IDs to attach the Yandex Database cluster to.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Whether TLS is enabled for the Yandex Database cluster.
	// Useful for SDK configuration.
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`

	// API endpoint of the Yandex Database cluster.
	// Useful for SDK configuration.
	YdbAPIEndpoint *string `json:"ydbApiEndpoint,omitempty" tf:"ydb_api_endpoint,omitempty"`

	// Full endpoint of the Yandex Database cluster.
	YdbFullEndpoint *string `json:"ydbFullEndpoint,omitempty" tf:"ydb_full_endpoint,omitempty"`
}

type DatabaseDedicatedParameters struct {

	// Whether public IP addresses should be assigned to the Yandex Database cluster.
	// +kubebuilder:validation:Optional
	AssignPublicIps *bool `json:"assignPublicIps,omitempty" tf:"assign_public_ips,omitempty"`

	// Inhibits deletion of the database. Can be either true or false
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A description for the Yandex Database cluster.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder that the Yandex Database cluster belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the Yandex Database cluster.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Location for the Yandex Database cluster.
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	Location []LocationParameters `json:"location,omitempty" tf:"location,omitempty"`

	// Location ID for the Yandex Database cluster.
	// +kubebuilder:validation:Optional
	LocationID *string `json:"locationId,omitempty" tf:"location_id,omitempty"`

	// Name of the Yandex Database cluster.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network to attach the Yandex Database cluster to.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// The Yandex Database cluster preset.
	// Available presets can be obtained via yc ydb resource-preset list command.
	// +kubebuilder:validation:Optional
	ResourcePresetID *string `json:"resourcePresetId,omitempty" tf:"resource_preset_id,omitempty"`

	// Scaling policy for the Yandex Database cluster.
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	ScalePolicy []ScalePolicyParameters `json:"scalePolicy,omitempty" tf:"scale_policy,omitempty"`

	// A list of storage configuration options for the Yandex Database cluster.
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	StorageConfig []StorageConfigParameters `json:"storageConfig,omitempty" tf:"storage_config,omitempty"`

	// List of subnet IDs to attach the Yandex Database cluster to.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// References to Subnet in vpc to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsRefs []v1.Reference `json:"subnetIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in vpc to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsSelector *v1.Selector `json:"subnetIdsSelector,omitempty" tf:"-"`
}

type FixedScaleInitParameters struct {

	// Number of instances for the Yandex Database cluster.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

type FixedScaleObservation struct {

	// Number of instances for the Yandex Database cluster.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

type FixedScaleParameters struct {

	// Number of instances for the Yandex Database cluster.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size" tf:"size,omitempty"`
}

type LocationInitParameters struct {

	// Region for the Yandex Database cluster.
	// The structure is documented below.
	Region []RegionInitParameters `json:"region,omitempty" tf:"region,omitempty"`
}

type LocationObservation struct {

	// Region for the Yandex Database cluster.
	// The structure is documented below.
	Region []RegionObservation `json:"region,omitempty" tf:"region,omitempty"`
}

type LocationParameters struct {

	// Region for the Yandex Database cluster.
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	Region []RegionParameters `json:"region,omitempty" tf:"region,omitempty"`
}

type RegionInitParameters struct {

	// Region ID for the Yandex Database cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RegionObservation struct {

	// Region ID for the Yandex Database cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RegionParameters struct {

	// Region ID for the Yandex Database cluster.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`
}

type ScalePolicyInitParameters struct {

	// Fixed scaling policy for the Yandex Database cluster.
	// The structure is documented below.
	FixedScale []FixedScaleInitParameters `json:"fixedScale,omitempty" tf:"fixed_scale,omitempty"`
}

type ScalePolicyObservation struct {

	// Fixed scaling policy for the Yandex Database cluster.
	// The structure is documented below.
	FixedScale []FixedScaleObservation `json:"fixedScale,omitempty" tf:"fixed_scale,omitempty"`
}

type ScalePolicyParameters struct {

	// Fixed scaling policy for the Yandex Database cluster.
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	FixedScale []FixedScaleParameters `json:"fixedScale" tf:"fixed_scale,omitempty"`
}

type StorageConfigInitParameters struct {

	// Amount of storage groups of selected type for the Yandex Database cluster.
	GroupCount *float64 `json:"groupCount,omitempty" tf:"group_count,omitempty"`

	// Storage type ID for the Yandex Database cluster.
	// Available presets can be obtained via yc ydb storage-type list command.
	StorageTypeID *string `json:"storageTypeId,omitempty" tf:"storage_type_id,omitempty"`
}

type StorageConfigObservation struct {

	// Amount of storage groups of selected type for the Yandex Database cluster.
	GroupCount *float64 `json:"groupCount,omitempty" tf:"group_count,omitempty"`

	// Storage type ID for the Yandex Database cluster.
	// Available presets can be obtained via yc ydb storage-type list command.
	StorageTypeID *string `json:"storageTypeId,omitempty" tf:"storage_type_id,omitempty"`
}

type StorageConfigParameters struct {

	// Amount of storage groups of selected type for the Yandex Database cluster.
	// +kubebuilder:validation:Optional
	GroupCount *float64 `json:"groupCount" tf:"group_count,omitempty"`

	// Storage type ID for the Yandex Database cluster.
	// Available presets can be obtained via yc ydb storage-type list command.
	// +kubebuilder:validation:Optional
	StorageTypeID *string `json:"storageTypeId" tf:"storage_type_id,omitempty"`
}

// DatabaseDedicatedSpec defines the desired state of DatabaseDedicated
type DatabaseDedicatedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseDedicatedParameters `json:"forProvider"`
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
	InitProvider DatabaseDedicatedInitParameters `json:"initProvider,omitempty"`
}

// DatabaseDedicatedStatus defines the observed state of DatabaseDedicated.
type DatabaseDedicatedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseDedicatedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DatabaseDedicated is the Schema for the DatabaseDedicateds API. Manages Yandex Database dedicated cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type DatabaseDedicated struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourcePresetId) || (has(self.initProvider) && has(self.initProvider.resourcePresetId))",message="spec.forProvider.resourcePresetId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scalePolicy) || (has(self.initProvider) && has(self.initProvider.scalePolicy))",message="spec.forProvider.scalePolicy is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageConfig) || (has(self.initProvider) && has(self.initProvider.storageConfig))",message="spec.forProvider.storageConfig is a required parameter"
	Spec   DatabaseDedicatedSpec   `json:"spec"`
	Status DatabaseDedicatedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseDedicatedList contains a list of DatabaseDedicateds
type DatabaseDedicatedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseDedicated `json:"items"`
}

// Repository type metadata.
var (
	DatabaseDedicated_Kind             = "DatabaseDedicated"
	DatabaseDedicated_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseDedicated_Kind}.String()
	DatabaseDedicated_KindAPIVersion   = DatabaseDedicated_Kind + "." + CRDGroupVersion.String()
	DatabaseDedicated_GroupVersionKind = CRDGroupVersion.WithKind(DatabaseDedicated_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabaseDedicated{}, &DatabaseDedicatedList{})
}
