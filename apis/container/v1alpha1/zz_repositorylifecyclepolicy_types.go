// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RepositoryLifecyclePolicyInitParameters struct {

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=Repository
	RepositoryID *string `json:"repositoryId,omitempty" tf:"repository_id,omitempty"`

	// Reference to a Repository to populate repositoryId.
	// +kubebuilder:validation:Optional
	RepositoryIDRef *v1.Reference `json:"repositoryIdRef,omitempty" tf:"-"`

	// Selector for a Repository to populate repositoryId.
	// +kubebuilder:validation:Optional
	RepositoryIDSelector *v1.Selector `json:"repositoryIdSelector,omitempty" tf:"-"`

	// (Block List) (see below for nested schema)
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// (String)
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RepositoryLifecyclePolicyObservation struct {

	// (String)
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	RepositoryID *string `json:"repositoryId,omitempty" tf:"repository_id,omitempty"`

	// (Block List) (see below for nested schema)
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`

	// (String)
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RepositoryLifecyclePolicyParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=Repository
	// +kubebuilder:validation:Optional
	RepositoryID *string `json:"repositoryId,omitempty" tf:"repository_id,omitempty"`

	// Reference to a Repository to populate repositoryId.
	// +kubebuilder:validation:Optional
	RepositoryIDRef *v1.Reference `json:"repositoryIdRef,omitempty" tf:"-"`

	// Selector for a Repository to populate repositoryId.
	// +kubebuilder:validation:Optional
	RepositoryIDSelector *v1.Selector `json:"repositoryIdSelector,omitempty" tf:"-"`

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RuleInitParameters struct {

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	ExpirePeriod *string `json:"expirePeriod,omitempty" tf:"expire_period,omitempty"`

	// (Number)
	RetainedTop *float64 `json:"retainedTop,omitempty" tf:"retained_top,omitempty"`

	// (String)
	TagRegexp *string `json:"tagRegexp,omitempty" tf:"tag_regexp,omitempty"`

	// (Boolean)
	Untagged *bool `json:"untagged,omitempty" tf:"untagged,omitempty"`
}

type RuleObservation struct {

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	ExpirePeriod *string `json:"expirePeriod,omitempty" tf:"expire_period,omitempty"`

	// (Number)
	RetainedTop *float64 `json:"retainedTop,omitempty" tf:"retained_top,omitempty"`

	// (String)
	TagRegexp *string `json:"tagRegexp,omitempty" tf:"tag_regexp,omitempty"`

	// (Boolean)
	Untagged *bool `json:"untagged,omitempty" tf:"untagged,omitempty"`
}

type RuleParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ExpirePeriod *string `json:"expirePeriod,omitempty" tf:"expire_period,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	RetainedTop *float64 `json:"retainedTop,omitempty" tf:"retained_top,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	TagRegexp *string `json:"tagRegexp,omitempty" tf:"tag_regexp,omitempty"`

	// (Boolean)
	// +kubebuilder:validation:Optional
	Untagged *bool `json:"untagged,omitempty" tf:"untagged,omitempty"`
}

// RepositoryLifecyclePolicySpec defines the desired state of RepositoryLifecyclePolicy
type RepositoryLifecyclePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryLifecyclePolicyParameters `json:"forProvider"`
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
	InitProvider RepositoryLifecyclePolicyInitParameters `json:"initProvider,omitempty"`
}

// RepositoryLifecyclePolicyStatus defines the observed state of RepositoryLifecyclePolicy.
type RepositoryLifecyclePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryLifecyclePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RepositoryLifecyclePolicy is the Schema for the RepositoryLifecyclePolicys API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type RepositoryLifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.status) || (has(self.initProvider) && has(self.initProvider.status))",message="spec.forProvider.status is a required parameter"
	Spec   RepositoryLifecyclePolicySpec   `json:"spec"`
	Status RepositoryLifecyclePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryLifecyclePolicyList contains a list of RepositoryLifecyclePolicys
type RepositoryLifecyclePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepositoryLifecyclePolicy `json:"items"`
}

// Repository type metadata.
var (
	RepositoryLifecyclePolicy_Kind             = "RepositoryLifecyclePolicy"
	RepositoryLifecyclePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RepositoryLifecyclePolicy_Kind}.String()
	RepositoryLifecyclePolicy_KindAPIVersion   = RepositoryLifecyclePolicy_Kind + "." + CRDGroupVersion.String()
	RepositoryLifecyclePolicy_GroupVersionKind = CRDGroupVersion.WithKind(RepositoryLifecyclePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&RepositoryLifecyclePolicy{}, &RepositoryLifecyclePolicyList{})
}
