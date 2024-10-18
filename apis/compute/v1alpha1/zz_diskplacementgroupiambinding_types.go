// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DiskPlacementGroupIAMBindingInitParameters struct {

	// ID of the disk placement group to attach the policy to.
	DiskPlacementGroupID *string `json:"diskPlacementGroupId,omitempty" tf:"disk_placement_group_id,omitempty"`

	// An array of identities that will be granted the privilege in the role. Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one yandex_compute_disk_placement_group_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type DiskPlacementGroupIAMBindingObservation struct {

	// ID of the disk placement group to attach the policy to.
	DiskPlacementGroupID *string `json:"diskPlacementGroupId,omitempty" tf:"disk_placement_group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An array of identities that will be granted the privilege in the role. Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one yandex_compute_disk_placement_group_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type DiskPlacementGroupIAMBindingParameters struct {

	// ID of the disk placement group to attach the policy to.
	// +kubebuilder:validation:Optional
	DiskPlacementGroupID *string `json:"diskPlacementGroupId,omitempty" tf:"disk_placement_group_id,omitempty"`

	// An array of identities that will be granted the privilege in the role. Each entry can have one of the following values:
	// +kubebuilder:validation:Optional
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one yandex_compute_disk_placement_group_iam_binding can be used per role.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// DiskPlacementGroupIAMBindingSpec defines the desired state of DiskPlacementGroupIAMBinding
type DiskPlacementGroupIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskPlacementGroupIAMBindingParameters `json:"forProvider"`
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
	InitProvider DiskPlacementGroupIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// DiskPlacementGroupIAMBindingStatus defines the observed state of DiskPlacementGroupIAMBinding.
type DiskPlacementGroupIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskPlacementGroupIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DiskPlacementGroupIAMBinding is the Schema for the DiskPlacementGroupIAMBindings API. Allows management of a single IAM binding for a Disk Placement Group.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type DiskPlacementGroupIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.diskPlacementGroupId) || (has(self.initProvider) && has(self.initProvider.diskPlacementGroupId))",message="spec.forProvider.diskPlacementGroupId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.members) || (has(self.initProvider) && has(self.initProvider.members))",message="spec.forProvider.members is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   DiskPlacementGroupIAMBindingSpec   `json:"spec"`
	Status DiskPlacementGroupIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskPlacementGroupIAMBindingList contains a list of DiskPlacementGroupIAMBindings
type DiskPlacementGroupIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskPlacementGroupIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	DiskPlacementGroupIAMBinding_Kind             = "DiskPlacementGroupIAMBinding"
	DiskPlacementGroupIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiskPlacementGroupIAMBinding_Kind}.String()
	DiskPlacementGroupIAMBinding_KindAPIVersion   = DiskPlacementGroupIAMBinding_Kind + "." + CRDGroupVersion.String()
	DiskPlacementGroupIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(DiskPlacementGroupIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&DiskPlacementGroupIAMBinding{}, &DiskPlacementGroupIAMBindingList{})
}
