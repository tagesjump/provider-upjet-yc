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

type AsymmetricEncryptionKeyIAMBindingInitParameters struct {

	// The Yandex Key Management Service Asymmetric Encryption Key ID to apply a binding to.
	// +crossplane:generate:reference:type=AsymmetricEncryptionKey
	AsymmetricEncryptionKeyID *string `json:"asymmetricEncryptionKeyId,omitempty" tf:"asymmetric_encryption_key_id,omitempty"`

	// Reference to a AsymmetricEncryptionKey to populate asymmetricEncryptionKeyId.
	// +kubebuilder:validation:Optional
	AsymmetricEncryptionKeyIDRef *v1.Reference `json:"asymmetricEncryptionKeyIdRef,omitempty" tf:"-"`

	// Selector for a AsymmetricEncryptionKey to populate asymmetricEncryptionKeyId.
	// +kubebuilder:validation:Optional
	AsymmetricEncryptionKeyIDSelector *v1.Selector `json:"asymmetricEncryptionKeyIdSelector,omitempty" tf:"-"`

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
}

type AsymmetricEncryptionKeyIAMBindingObservation struct {

	// The Yandex Key Management Service Asymmetric Encryption Key ID to apply a binding to.
	AsymmetricEncryptionKeyID *string `json:"asymmetricEncryptionKeyId,omitempty" tf:"asymmetric_encryption_key_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Identities that will be granted the privilege in role.
	// Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be applied. See roles.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type AsymmetricEncryptionKeyIAMBindingParameters struct {

	// The Yandex Key Management Service Asymmetric Encryption Key ID to apply a binding to.
	// +crossplane:generate:reference:type=AsymmetricEncryptionKey
	// +kubebuilder:validation:Optional
	AsymmetricEncryptionKeyID *string `json:"asymmetricEncryptionKeyId,omitempty" tf:"asymmetric_encryption_key_id,omitempty"`

	// Reference to a AsymmetricEncryptionKey to populate asymmetricEncryptionKeyId.
	// +kubebuilder:validation:Optional
	AsymmetricEncryptionKeyIDRef *v1.Reference `json:"asymmetricEncryptionKeyIdRef,omitempty" tf:"-"`

	// Selector for a AsymmetricEncryptionKey to populate asymmetricEncryptionKeyId.
	// +kubebuilder:validation:Optional
	AsymmetricEncryptionKeyIDSelector *v1.Selector `json:"asymmetricEncryptionKeyIdSelector,omitempty" tf:"-"`

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
}

// AsymmetricEncryptionKeyIAMBindingSpec defines the desired state of AsymmetricEncryptionKeyIAMBinding
type AsymmetricEncryptionKeyIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AsymmetricEncryptionKeyIAMBindingParameters `json:"forProvider"`
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
	InitProvider AsymmetricEncryptionKeyIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// AsymmetricEncryptionKeyIAMBindingStatus defines the observed state of AsymmetricEncryptionKeyIAMBinding.
type AsymmetricEncryptionKeyIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AsymmetricEncryptionKeyIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AsymmetricEncryptionKeyIAMBinding is the Schema for the AsymmetricEncryptionKeyIAMBindings API. Allows management of a single IAM binding for a
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type AsymmetricEncryptionKeyIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   AsymmetricEncryptionKeyIAMBindingSpec   `json:"spec"`
	Status AsymmetricEncryptionKeyIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AsymmetricEncryptionKeyIAMBindingList contains a list of AsymmetricEncryptionKeyIAMBindings
type AsymmetricEncryptionKeyIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AsymmetricEncryptionKeyIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	AsymmetricEncryptionKeyIAMBinding_Kind             = "AsymmetricEncryptionKeyIAMBinding"
	AsymmetricEncryptionKeyIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AsymmetricEncryptionKeyIAMBinding_Kind}.String()
	AsymmetricEncryptionKeyIAMBinding_KindAPIVersion   = AsymmetricEncryptionKeyIAMBinding_Kind + "." + CRDGroupVersion.String()
	AsymmetricEncryptionKeyIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(AsymmetricEncryptionKeyIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&AsymmetricEncryptionKeyIAMBinding{}, &AsymmetricEncryptionKeyIAMBindingList{})
}
