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

type AsymmetricSignatureKeyIAMBindingInitParameters struct {

	// The role that should be applied. See roles.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type AsymmetricSignatureKeyIAMBindingObservation struct {

	// The Yandex Key Management Service Asymmetric Signature Key ID to apply a binding to.
	AsymmetricSignatureKeyID *string `json:"asymmetricSignatureKeyId,omitempty" tf:"asymmetric_signature_key_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Identities that will be granted the privilege in role.
	// Each entry can have one of the following values:
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be applied. See roles.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type AsymmetricSignatureKeyIAMBindingParameters struct {

	// The Yandex Key Management Service Asymmetric Signature Key ID to apply a binding to.
	// +crossplane:generate:reference:type=AsymmetricSignatureKey
	// +kubebuilder:validation:Optional
	AsymmetricSignatureKeyID *string `json:"asymmetricSignatureKeyId,omitempty" tf:"asymmetric_signature_key_id,omitempty"`

	// Reference to a AsymmetricSignatureKey to populate asymmetricSignatureKeyId.
	// +kubebuilder:validation:Optional
	AsymmetricSignatureKeyIDRef *v1.Reference `json:"asymmetricSignatureKeyIdRef,omitempty" tf:"-"`

	// Selector for a AsymmetricSignatureKey to populate asymmetricSignatureKeyId.
	// +kubebuilder:validation:Optional
	AsymmetricSignatureKeyIDSelector *v1.Selector `json:"asymmetricSignatureKeyIdSelector,omitempty" tf:"-"`

	// Identities that will be granted the privilege in role.
	// Each entry can have one of the following values:
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/iam/v1alpha1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/tagesjump/provider-upjet-yc/config/iam.ServiceAccountRefValue()
	// +crossplane:generate:reference:refFieldName=ServiceAccountRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountSelector
	// +kubebuilder:validation:Optional
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
}

// AsymmetricSignatureKeyIAMBindingSpec defines the desired state of AsymmetricSignatureKeyIAMBinding
type AsymmetricSignatureKeyIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AsymmetricSignatureKeyIAMBindingParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AsymmetricSignatureKeyIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// AsymmetricSignatureKeyIAMBindingStatus defines the observed state of AsymmetricSignatureKeyIAMBinding.
type AsymmetricSignatureKeyIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AsymmetricSignatureKeyIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AsymmetricSignatureKeyIAMBinding is the Schema for the AsymmetricSignatureKeyIAMBindings API. Allows management of a single IAM binding for a
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud.upjet}
type AsymmetricSignatureKeyIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   AsymmetricSignatureKeyIAMBindingSpec   `json:"spec"`
	Status AsymmetricSignatureKeyIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AsymmetricSignatureKeyIAMBindingList contains a list of AsymmetricSignatureKeyIAMBindings
type AsymmetricSignatureKeyIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AsymmetricSignatureKeyIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	AsymmetricSignatureKeyIAMBinding_Kind             = "AsymmetricSignatureKeyIAMBinding"
	AsymmetricSignatureKeyIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AsymmetricSignatureKeyIAMBinding_Kind}.String()
	AsymmetricSignatureKeyIAMBinding_KindAPIVersion   = AsymmetricSignatureKeyIAMBinding_Kind + "." + CRDGroupVersion.String()
	AsymmetricSignatureKeyIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(AsymmetricSignatureKeyIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&AsymmetricSignatureKeyIAMBinding{}, &AsymmetricSignatureKeyIAMBindingList{})
}