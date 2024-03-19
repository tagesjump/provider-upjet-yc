// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DiskPlacementGroupInitParameters struct {

	// A description of the Disk Placement Group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the Disk Placement Group.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the Disk Placement Group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the zone where the Disk Placement Group resides.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DiskPlacementGroupObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// A description of the Disk Placement Group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the Disk Placement Group.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the Disk Placement Group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Status of the Disk Placement Group.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// ID of the zone where the Disk Placement Group resides.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DiskPlacementGroupParameters struct {

	// A description of the Disk Placement Group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Folder that the resource belongs to. If value is omitted, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the Disk Placement Group.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the Disk Placement Group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the zone where the Disk Placement Group resides.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// DiskPlacementGroupSpec defines the desired state of DiskPlacementGroup
type DiskPlacementGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskPlacementGroupParameters `json:"forProvider"`
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
	InitProvider DiskPlacementGroupInitParameters `json:"initProvider,omitempty"`
}

// DiskPlacementGroupStatus defines the observed state of DiskPlacementGroup.
type DiskPlacementGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskPlacementGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DiskPlacementGroup is the Schema for the DiskPlacementGroups API. Manages a Disk Placement Group resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type DiskPlacementGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskPlacementGroupSpec   `json:"spec"`
	Status            DiskPlacementGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskPlacementGroupList contains a list of DiskPlacementGroups
type DiskPlacementGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskPlacementGroup `json:"items"`
}

// Repository type metadata.
var (
	DiskPlacementGroup_Kind             = "DiskPlacementGroup"
	DiskPlacementGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiskPlacementGroup_Kind}.String()
	DiskPlacementGroup_KindAPIVersion   = DiskPlacementGroup_Kind + "." + CRDGroupVersion.String()
	DiskPlacementGroup_GroupVersionKind = CRDGroupVersion.WithKind(DiskPlacementGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&DiskPlacementGroup{}, &DiskPlacementGroupList{})
}
