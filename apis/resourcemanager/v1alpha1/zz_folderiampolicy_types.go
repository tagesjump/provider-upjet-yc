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

type FolderIAMPolicyInitParameters struct {

	// The yandex_iam_policy data source that represents
	// the IAM policy that will be applied to the folder. This policy overrides any existing policy applied to the folder.
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`
}

type FolderIAMPolicyObservation struct {

	// ID of the folder that the policy is attached to.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The yandex_iam_policy data source that represents
	// the IAM policy that will be applied to the folder. This policy overrides any existing policy applied to the folder.
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`
}

type FolderIAMPolicyParameters struct {

	// ID of the folder that the policy is attached to.
	// +crossplane:generate:reference:type=Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// The yandex_iam_policy data source that represents
	// the IAM policy that will be applied to the folder. This policy overrides any existing policy applied to the folder.
	// +kubebuilder:validation:Optional
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`
}

// FolderIAMPolicySpec defines the desired state of FolderIAMPolicy
type FolderIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderIAMPolicyParameters `json:"forProvider"`
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
	InitProvider FolderIAMPolicyInitParameters `json:"initProvider,omitempty"`
}

// FolderIAMPolicyStatus defines the observed state of FolderIAMPolicy.
type FolderIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FolderIAMPolicy is the Schema for the FolderIAMPolicys API. Allows management of the IAM policy for a Yandex Resource Manager folder.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud.upjet}
type FolderIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyData) || (has(self.initProvider) && has(self.initProvider.policyData))",message="spec.forProvider.policyData is a required parameter"
	Spec   FolderIAMPolicySpec   `json:"spec"`
	Status FolderIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderIAMPolicyList contains a list of FolderIAMPolicys
type FolderIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FolderIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	FolderIAMPolicy_Kind             = "FolderIAMPolicy"
	FolderIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FolderIAMPolicy_Kind}.String()
	FolderIAMPolicy_KindAPIVersion   = FolderIAMPolicy_Kind + "." + CRDGroupVersion.String()
	FolderIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(FolderIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&FolderIAMPolicy{}, &FolderIAMPolicyList{})
}
