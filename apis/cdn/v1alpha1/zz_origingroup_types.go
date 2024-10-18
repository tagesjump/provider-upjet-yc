// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OriginGroupInitParameters struct {

	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// CDN Origin Group name used to define device.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Origin []OriginInitParameters `json:"origin,omitempty" tf:"origin,omitempty"`

	// If the option is active (has true value), in case the origin responds with 4XX or 5XX codes, use the next origin from the list.
	UseNext *bool `json:"useNext,omitempty" tf:"use_next,omitempty"`
}

type OriginGroupObservation struct {
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// CDN Origin Group name used to define device.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Origin []OriginObservation `json:"origin,omitempty" tf:"origin,omitempty"`

	// If the option is active (has true value), in case the origin responds with 4XX or 5XX codes, use the next origin from the list.
	UseNext *bool `json:"useNext,omitempty" tf:"use_next,omitempty"`
}

type OriginGroupParameters struct {

	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// CDN Origin Group name used to define device.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Origin []OriginParameters `json:"origin,omitempty" tf:"origin,omitempty"`

	// If the option is active (has true value), in case the origin responds with 4XX or 5XX codes, use the next origin from the list.
	// +kubebuilder:validation:Optional
	UseNext *bool `json:"useNext,omitempty" tf:"use_next,omitempty"`
}

type OriginInitParameters struct {
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type OriginObservation struct {
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	OriginGroupID *float64 `json:"originGroupId,omitempty" tf:"origin_group_id,omitempty"`

	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type OriginParameters struct {

	// +kubebuilder:validation:Optional
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Source *string `json:"source" tf:"source,omitempty"`
}

// OriginGroupSpec defines the desired state of OriginGroup
type OriginGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OriginGroupParameters `json:"forProvider"`
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
	InitProvider OriginGroupInitParameters `json:"initProvider,omitempty"`
}

// OriginGroupStatus defines the observed state of OriginGroup.
type OriginGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OriginGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OriginGroup is the Schema for the OriginGroups API. Allows management of a Yandex.Cloud CDN Origin Groups.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type OriginGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.origin) || (has(self.initProvider) && has(self.initProvider.origin))",message="spec.forProvider.origin is a required parameter"
	Spec   OriginGroupSpec   `json:"spec"`
	Status OriginGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OriginGroupList contains a list of OriginGroups
type OriginGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OriginGroup `json:"items"`
}

// Repository type metadata.
var (
	OriginGroup_Kind             = "OriginGroup"
	OriginGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OriginGroup_Kind}.String()
	OriginGroup_KindAPIVersion   = OriginGroup_Kind + "." + CRDGroupVersion.String()
	OriginGroup_GroupVersionKind = CRDGroupVersion.WithKind(OriginGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&OriginGroup{}, &OriginGroupList{})
}