

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

)




type IAMBindingInitParameters struct {


// The Yandex Cloud Function ID to apply a binding to.
FunctionID *string `json:"functionId,omitempty" tf:"function_id,omitempty"`

// Identities that will be granted the privilege in role. Each entry can have one of the following values:
// +listType=set
Members []*string `json:"members,omitempty" tf:"members,omitempty"`

// The role that should be applied. See roles
Role *string `json:"role,omitempty" tf:"role,omitempty"`

SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}


type IAMBindingObservation struct {


// The Yandex Cloud Function ID to apply a binding to.
FunctionID *string `json:"functionId,omitempty" tf:"function_id,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

// Identities that will be granted the privilege in role. Each entry can have one of the following values:
// +listType=set
Members []*string `json:"members,omitempty" tf:"members,omitempty"`

// The role that should be applied. See roles
Role *string `json:"role,omitempty" tf:"role,omitempty"`

SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}


type IAMBindingParameters struct {


// The Yandex Cloud Function ID to apply a binding to.
// +kubebuilder:validation:Optional
FunctionID *string `json:"functionId,omitempty" tf:"function_id,omitempty"`

// Identities that will be granted the privilege in role. Each entry can have one of the following values:
// +kubebuilder:validation:Optional
// +listType=set
Members []*string `json:"members,omitempty" tf:"members,omitempty"`

// The role that should be applied. See roles
// +kubebuilder:validation:Optional
Role *string `json:"role,omitempty" tf:"role,omitempty"`

// +kubebuilder:validation:Optional
SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// IAMBindingSpec defines the desired state of IAMBinding
type IAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       IAMBindingParameters `json:"forProvider"`
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
	InitProvider       IAMBindingInitParameters `json:"initProvider,omitempty"`
}

// IAMBindingStatus defines the observed state of IAMBinding.
type IAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          IAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion


// IAMBinding is the Schema for the IAMBindings API. Allows management of a single IAM binding for a
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type IAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.functionId) || (has(self.initProvider) && has(self.initProvider.functionId))",message="spec.forProvider.functionId is a required parameter"
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.members) || (has(self.initProvider) && has(self.initProvider.members))",message="spec.forProvider.members is a required parameter"
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec              IAMBindingSpec   `json:"spec"`
	Status            IAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IAMBindingList contains a list of IAMBindings
type IAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMBinding `json:"items"`
}

// Repository type metadata.
var (
	IAMBinding_Kind             = "IAMBinding"
	IAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IAMBinding_Kind}.String()
	IAMBinding_KindAPIVersion   = IAMBinding_Kind + "." + CRDGroupVersion.String()
	IAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(IAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&IAMBinding{}, &IAMBindingList{})
}
