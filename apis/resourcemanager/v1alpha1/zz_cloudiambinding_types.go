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

type CloudIAMBindingInitParameters struct {

	// ID of the cloud to attach the policy to.
	// +crossplane:generate:reference:type=Cloud
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	// Reference to a Cloud to populate cloudId.
	// +kubebuilder:validation:Optional
	CloudIDRef *v1.Reference `json:"cloudIdRef,omitempty" tf:"-"`

	// Selector for a Cloud to populate cloudId.
	// +kubebuilder:validation:Optional
	CloudIDSelector *v1.Selector `json:"cloudIdSelector,omitempty" tf:"-"`

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_resourcemanager_cloud_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type CloudIAMBindingObservation struct {

	// ID of the cloud to attach the policy to.
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_resourcemanager_cloud_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type CloudIAMBindingParameters struct {

	// ID of the cloud to attach the policy to.
	// +crossplane:generate:reference:type=Cloud
	// +kubebuilder:validation:Optional
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	// Reference to a Cloud to populate cloudId.
	// +kubebuilder:validation:Optional
	CloudIDRef *v1.Reference `json:"cloudIdRef,omitempty" tf:"-"`

	// Selector for a Cloud to populate cloudId.
	// +kubebuilder:validation:Optional
	CloudIDSelector *v1.Selector `json:"cloudIdSelector,omitempty" tf:"-"`

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +kubebuilder:validation:Optional
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_resourcemanager_cloud_iam_binding can be used per role.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// CloudIAMBindingSpec defines the desired state of CloudIAMBinding
type CloudIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudIAMBindingParameters `json:"forProvider"`
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
	InitProvider CloudIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// CloudIAMBindingStatus defines the observed state of CloudIAMBinding.
type CloudIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudIAMBinding is the Schema for the CloudIAMBindings API. Allows management of a single IAM binding for a Yandex Resource Manager cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type CloudIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.members) || (has(self.initProvider) && has(self.initProvider.members))",message="spec.forProvider.members is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   CloudIAMBindingSpec   `json:"spec"`
	Status CloudIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudIAMBindingList contains a list of CloudIAMBindings
type CloudIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	CloudIAMBinding_Kind             = "CloudIAMBinding"
	CloudIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudIAMBinding_Kind}.String()
	CloudIAMBinding_KindAPIVersion   = CloudIAMBinding_Kind + "." + CRDGroupVersion.String()
	CloudIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(CloudIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudIAMBinding{}, &CloudIAMBindingList{})
}
