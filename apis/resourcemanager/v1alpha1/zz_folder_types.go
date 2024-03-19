// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FolderInitParameters struct {

	// Cloud that the resource belongs to. If value is omitted, the default provider Cloud ID is used.
	// +crossplane:generate:reference:type=Cloud
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	// Reference to a Cloud to populate cloudId.
	// +kubebuilder:validation:Optional
	CloudIDRef *v1.Reference `json:"cloudIdRef,omitempty" tf:"-"`

	// Selector for a Cloud to populate cloudId.
	// +kubebuilder:validation:Optional
	CloudIDSelector *v1.Selector `json:"cloudIdSelector,omitempty" tf:"-"`

	// A description of the Folder.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of key/value label pairs to assign to the Folder.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the Folder.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FolderObservation struct {

	// Cloud that the resource belongs to. If value is omitted, the default provider Cloud ID is used.
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// A description of the Folder.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the Folder.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the Folder.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FolderParameters struct {

	// Cloud that the resource belongs to. If value is omitted, the default provider Cloud ID is used.
	// +crossplane:generate:reference:type=Cloud
	// +kubebuilder:validation:Optional
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	// Reference to a Cloud to populate cloudId.
	// +kubebuilder:validation:Optional
	CloudIDRef *v1.Reference `json:"cloudIdRef,omitempty" tf:"-"`

	// Selector for a Cloud to populate cloudId.
	// +kubebuilder:validation:Optional
	CloudIDSelector *v1.Selector `json:"cloudIdSelector,omitempty" tf:"-"`

	// A description of the Folder.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of key/value label pairs to assign to the Folder.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the Folder.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// FolderSpec defines the desired state of Folder
type FolderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderParameters `json:"forProvider"`
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
	InitProvider FolderInitParameters `json:"initProvider,omitempty"`
}

// FolderStatus defines the observed state of Folder.
type FolderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Folder is the Schema for the Folders API. Allows management of the Cloud Folder.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Folder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderSpec   `json:"spec"`
	Status            FolderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderList contains a list of Folders
type FolderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Folder `json:"items"`
}

// Repository type metadata.
var (
	Folder_Kind             = "Folder"
	Folder_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Folder_Kind}.String()
	Folder_KindAPIVersion   = Folder_Kind + "." + CRDGroupVersion.String()
	Folder_GroupVersionKind = CRDGroupVersion.WithKind(Folder_Kind)
)

func init() {
	SchemeBuilder.Register(&Folder{}, &FolderList{})
}
