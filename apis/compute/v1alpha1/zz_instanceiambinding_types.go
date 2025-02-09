// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstanceIAMBindingInitParameters struct {

	// ID of the instance to attach the policy to.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_compute_instance_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type InstanceIAMBindingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the instance to attach the policy to.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_compute_instance_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type InstanceIAMBindingParameters struct {

	// ID of the instance to attach the policy to.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +kubebuilder:validation:Optional
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_compute_instance_iam_binding can be used per role.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// InstanceIAMBindingSpec defines the desired state of InstanceIAMBinding
type InstanceIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceIAMBindingParameters `json:"forProvider"`
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
	InitProvider InstanceIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// InstanceIAMBindingStatus defines the observed state of InstanceIAMBinding.
type InstanceIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InstanceIAMBinding is the Schema for the InstanceIAMBindings API. Allows management of a single IAM binding for an instance.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type InstanceIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceId) || (has(self.initProvider) && has(self.initProvider.instanceId))",message="spec.forProvider.instanceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.members) || (has(self.initProvider) && has(self.initProvider.members))",message="spec.forProvider.members is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   InstanceIAMBindingSpec   `json:"spec"`
	Status InstanceIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceIAMBindingList contains a list of InstanceIAMBindings
type InstanceIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	InstanceIAMBinding_Kind             = "InstanceIAMBinding"
	InstanceIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceIAMBinding_Kind}.String()
	InstanceIAMBinding_KindAPIVersion   = InstanceIAMBinding_Kind + "." + CRDGroupVersion.String()
	InstanceIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(InstanceIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceIAMBinding{}, &InstanceIAMBindingList{})
}
