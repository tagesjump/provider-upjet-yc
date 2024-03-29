// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GroupIAMMemberInitParameters struct {

	// ID of the organization to attach a policy to.
	// +crossplane:generate:reference:type=Group
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// The identity that will be granted the privilege that is specified in the role field.
	// This field can have one of the following values:
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// The role that should be assigned.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type GroupIAMMemberObservation struct {

	// ID of the organization to attach a policy to.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The identity that will be granted the privilege that is specified in the role field.
	// This field can have one of the following values:
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// The role that should be assigned.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type GroupIAMMemberParameters struct {

	// ID of the organization to attach a policy to.
	// +crossplane:generate:reference:type=Group
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// The identity that will be granted the privilege that is specified in the role field.
	// This field can have one of the following values:
	// +kubebuilder:validation:Optional
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// The role that should be assigned.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// GroupIAMMemberSpec defines the desired state of GroupIAMMember
type GroupIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupIAMMemberParameters `json:"forProvider"`
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
	InitProvider GroupIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// GroupIAMMemberStatus defines the observed state of GroupIAMMember.
type GroupIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GroupIAMMember is the Schema for the GroupIAMMembers API. Allows management of a single member for a single IAM binding on a Yandex.Cloud Organization Manager Group.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type GroupIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.member) || (has(self.initProvider) && has(self.initProvider.member))",message="spec.forProvider.member is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   GroupIAMMemberSpec   `json:"spec"`
	Status GroupIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupIAMMemberList contains a list of GroupIAMMembers
type GroupIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupIAMMember `json:"items"`
}

// Repository type metadata.
var (
	GroupIAMMember_Kind             = "GroupIAMMember"
	GroupIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupIAMMember_Kind}.String()
	GroupIAMMember_KindAPIVersion   = GroupIAMMember_Kind + "." + CRDGroupVersion.String()
	GroupIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(GroupIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupIAMMember{}, &GroupIAMMemberList{})
}
