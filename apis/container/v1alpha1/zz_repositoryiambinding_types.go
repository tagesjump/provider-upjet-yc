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

type RepositoryIAMBindingInitParameters struct {

	// The role that should be applied. See roles.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type RepositoryIAMBindingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Identities that will be granted the privilege in role.
	// Each entry can have one of the following values:
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The Yandex Container Repository ID to apply a binding to.
	RepositoryID *string `json:"repositoryId,omitempty" tf:"repository_id,omitempty"`

	// The role that should be applied. See roles.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type RepositoryIAMBindingParameters struct {

	// Identities that will be granted the privilege in role.
	// Each entry can have one of the following values:
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/iam/v1alpha1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/tagesjump/provider-upjet-yc/config/iam.ServiceAccountRefValue()
	// +crossplane:generate:reference:refFieldName=ServiceAccountRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountSelector
	// +kubebuilder:validation:Optional
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The Yandex Container Repository ID to apply a binding to.
	// +crossplane:generate:reference:type=Repository
	// +kubebuilder:validation:Optional
	RepositoryID *string `json:"repositoryId,omitempty" tf:"repository_id,omitempty"`

	// Reference to a Repository to populate repositoryId.
	// +kubebuilder:validation:Optional
	RepositoryIDRef *v1.Reference `json:"repositoryIdRef,omitempty" tf:"-"`

	// Selector for a Repository to populate repositoryId.
	// +kubebuilder:validation:Optional
	RepositoryIDSelector *v1.Selector `json:"repositoryIdSelector,omitempty" tf:"-"`

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

// RepositoryIAMBindingSpec defines the desired state of RepositoryIAMBinding
type RepositoryIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryIAMBindingParameters `json:"forProvider"`
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
	InitProvider RepositoryIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// RepositoryIAMBindingStatus defines the observed state of RepositoryIAMBinding.
type RepositoryIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryIAMBinding is the Schema for the RepositoryIAMBindings API. Allows management of a single IAM binding for a
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud.upjet}
type RepositoryIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   RepositoryIAMBindingSpec   `json:"spec"`
	Status RepositoryIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryIAMBindingList contains a list of RepositoryIAMBindings
type RepositoryIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepositoryIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	RepositoryIAMBinding_Kind             = "RepositoryIAMBinding"
	RepositoryIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RepositoryIAMBinding_Kind}.String()
	RepositoryIAMBinding_KindAPIVersion   = RepositoryIAMBinding_Kind + "." + CRDGroupVersion.String()
	RepositoryIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(RepositoryIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&RepositoryIAMBinding{}, &RepositoryIAMBindingList{})
}
