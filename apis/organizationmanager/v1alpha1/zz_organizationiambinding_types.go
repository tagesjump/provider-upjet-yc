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

type OrganizationIAMBindingInitParameters struct {

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// ID of the organization to attach the policy to.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// The role that should be assigned. Only one
	// yandex_organizationmanager_organization_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type OrganizationIAMBindingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// ID of the organization to attach the policy to.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// The role that should be assigned. Only one
	// yandex_organizationmanager_organization_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type OrganizationIAMBindingParameters struct {

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +kubebuilder:validation:Optional
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// ID of the organization to attach the policy to.
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// The role that should be assigned. Only one
	// yandex_organizationmanager_organization_iam_binding can be used per role.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// OrganizationIAMBindingSpec defines the desired state of OrganizationIAMBinding
type OrganizationIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationIAMBindingParameters `json:"forProvider"`
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
	InitProvider OrganizationIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// OrganizationIAMBindingStatus defines the observed state of OrganizationIAMBinding.
type OrganizationIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationIAMBinding is the Schema for the OrganizationIAMBindings API. Allows management of a single IAM binding for a Yandex Organization Manager organization.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type OrganizationIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.members) || (has(self.initProvider) && has(self.initProvider.members))",message="spec.forProvider.members is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.organizationId) || (has(self.initProvider) && has(self.initProvider.organizationId))",message="spec.forProvider.organizationId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   OrganizationIAMBindingSpec   `json:"spec"`
	Status OrganizationIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationIAMBindingList contains a list of OrganizationIAMBindings
type OrganizationIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	OrganizationIAMBinding_Kind             = "OrganizationIAMBinding"
	OrganizationIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationIAMBinding_Kind}.String()
	OrganizationIAMBinding_KindAPIVersion   = OrganizationIAMBinding_Kind + "." + CRDGroupVersion.String()
	OrganizationIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationIAMBinding{}, &OrganizationIAMBindingList{})
}
