// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PolicyBindingsInitParameters struct {

	// — Compute Cloud instance ID.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/compute/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in compute to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in compute to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// — Backup Policy ID.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`
}

type PolicyBindingsObservation struct {

	// Creation timestamp of the Yandex Cloud Policy Bindings.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Boolean flag that specifies whether the policy application is enabled. May be false if processing flag is true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// — Compute Cloud instance ID.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// — Backup Policy ID.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Boolean flag that specifies whether the policy is in the process of binding to an instance.
	Processing *bool `json:"processing,omitempty" tf:"processing,omitempty"`
}

type PolicyBindingsParameters struct {

	// — Compute Cloud instance ID.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/compute/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in compute to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in compute to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// — Backup Policy ID.
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`
}

// PolicyBindingsSpec defines the desired state of PolicyBindings
type PolicyBindingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyBindingsParameters `json:"forProvider"`
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
	InitProvider PolicyBindingsInitParameters `json:"initProvider,omitempty"`
}

// PolicyBindingsStatus defines the observed state of PolicyBindings.
type PolicyBindingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyBindingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PolicyBindings is the Schema for the PolicyBindingss API. Allows to bind compute instance with backup policy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type PolicyBindings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyId) || (has(self.initProvider) && has(self.initProvider.policyId))",message="spec.forProvider.policyId is a required parameter"
	Spec   PolicyBindingsSpec   `json:"spec"`
	Status PolicyBindingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyBindingsList contains a list of PolicyBindingss
type PolicyBindingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyBindings `json:"items"`
}

// Repository type metadata.
var (
	PolicyBindings_Kind             = "PolicyBindings"
	PolicyBindings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyBindings_Kind}.String()
	PolicyBindings_KindAPIVersion   = PolicyBindings_Kind + "." + CRDGroupVersion.String()
	PolicyBindings_GroupVersionKind = CRDGroupVersion.WithKind(PolicyBindings_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyBindings{}, &PolicyBindingsList{})
}