// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SymmetricKeyIAMBindingInitParameters struct {

	// Identities that will be granted the privilege in role.
	// Each entry can have one of the following values:
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/iam/v1alpha1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/tagesjump/provider-upjet-yc/config/iam.ServiceAccountRefValue()
	// +crossplane:generate:reference:refFieldName=ServiceAccountRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountSelector
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be applied. See roles.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// References to ServiceAccount in iam to populate members.
	// +kubebuilder:validation:Optional
	ServiceAccountRef []v1.Reference `json:"serviceAccountRef,omitempty" tf:"-"`

	// Selector for a list of ServiceAccount in iam to populate members.
	// +kubebuilder:validation:Optional
	ServiceAccountSelector *v1.Selector `json:"serviceAccountSelector,omitempty" tf:"-"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`

	// The Yandex Key Management Service Symmetric Key ID to apply a binding to.
	// +crossplane:generate:reference:type=SymmetricKey
	SymmetricKeyID *string `json:"symmetricKeyId,omitempty" tf:"symmetric_key_id,omitempty"`

	// Reference to a SymmetricKey to populate symmetricKeyId.
	// +kubebuilder:validation:Optional
	SymmetricKeyIDRef *v1.Reference `json:"symmetricKeyIdRef,omitempty" tf:"-"`

	// Selector for a SymmetricKey to populate symmetricKeyId.
	// +kubebuilder:validation:Optional
	SymmetricKeyIDSelector *v1.Selector `json:"symmetricKeyIdSelector,omitempty" tf:"-"`
}

type SymmetricKeyIAMBindingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Identities that will be granted the privilege in role.
	// Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be applied. See roles.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`

	// The Yandex Key Management Service Symmetric Key ID to apply a binding to.
	SymmetricKeyID *string `json:"symmetricKeyId,omitempty" tf:"symmetric_key_id,omitempty"`
}

type SymmetricKeyIAMBindingParameters struct {

	// Identities that will be granted the privilege in role.
	// Each entry can have one of the following values:
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/iam/v1alpha1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/tagesjump/provider-upjet-yc/config/iam.ServiceAccountRefValue()
	// +crossplane:generate:reference:refFieldName=ServiceAccountRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be applied. See roles.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// References to ServiceAccount in iam to populate members.
	// +kubebuilder:validation:Optional
	ServiceAccountRef []v1.Reference `json:"serviceAccountRef,omitempty" tf:"-"`

	// Selector for a list of ServiceAccount in iam to populate members.
	// +kubebuilder:validation:Optional
	ServiceAccountSelector *v1.Selector `json:"serviceAccountSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`

	// The Yandex Key Management Service Symmetric Key ID to apply a binding to.
	// +crossplane:generate:reference:type=SymmetricKey
	// +kubebuilder:validation:Optional
	SymmetricKeyID *string `json:"symmetricKeyId,omitempty" tf:"symmetric_key_id,omitempty"`

	// Reference to a SymmetricKey to populate symmetricKeyId.
	// +kubebuilder:validation:Optional
	SymmetricKeyIDRef *v1.Reference `json:"symmetricKeyIdRef,omitempty" tf:"-"`

	// Selector for a SymmetricKey to populate symmetricKeyId.
	// +kubebuilder:validation:Optional
	SymmetricKeyIDSelector *v1.Selector `json:"symmetricKeyIdSelector,omitempty" tf:"-"`
}

// SymmetricKeyIAMBindingSpec defines the desired state of SymmetricKeyIAMBinding
type SymmetricKeyIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SymmetricKeyIAMBindingParameters `json:"forProvider"`
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
	InitProvider SymmetricKeyIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// SymmetricKeyIAMBindingStatus defines the observed state of SymmetricKeyIAMBinding.
type SymmetricKeyIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SymmetricKeyIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SymmetricKeyIAMBinding is the Schema for the SymmetricKeyIAMBindings API. Allows management of a single IAM binding for a
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type SymmetricKeyIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   SymmetricKeyIAMBindingSpec   `json:"spec"`
	Status SymmetricKeyIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SymmetricKeyIAMBindingList contains a list of SymmetricKeyIAMBindings
type SymmetricKeyIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SymmetricKeyIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	SymmetricKeyIAMBinding_Kind             = "SymmetricKeyIAMBinding"
	SymmetricKeyIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SymmetricKeyIAMBinding_Kind}.String()
	SymmetricKeyIAMBinding_KindAPIVersion   = SymmetricKeyIAMBinding_Kind + "." + CRDGroupVersion.String()
	SymmetricKeyIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(SymmetricKeyIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&SymmetricKeyIAMBinding{}, &SymmetricKeyIAMBindingList{})
}
