// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SnapshotScheduleIAMBindingInitParameters struct {

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_compute_snapshot_schedule_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// ID of the snapshot schedule to attach the policy to.
	SnapshotScheduleID *string `json:"snapshotScheduleId,omitempty" tf:"snapshot_schedule_id,omitempty"`
}

type SnapshotScheduleIAMBindingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_compute_snapshot_schedule_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// ID of the snapshot schedule to attach the policy to.
	SnapshotScheduleID *string `json:"snapshotScheduleId,omitempty" tf:"snapshot_schedule_id,omitempty"`
}

type SnapshotScheduleIAMBindingParameters struct {

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +kubebuilder:validation:Optional
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_compute_snapshot_schedule_iam_binding can be used per role.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// ID of the snapshot schedule to attach the policy to.
	// +kubebuilder:validation:Optional
	SnapshotScheduleID *string `json:"snapshotScheduleId,omitempty" tf:"snapshot_schedule_id,omitempty"`
}

// SnapshotScheduleIAMBindingSpec defines the desired state of SnapshotScheduleIAMBinding
type SnapshotScheduleIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotScheduleIAMBindingParameters `json:"forProvider"`
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
	InitProvider SnapshotScheduleIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// SnapshotScheduleIAMBindingStatus defines the observed state of SnapshotScheduleIAMBinding.
type SnapshotScheduleIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotScheduleIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SnapshotScheduleIAMBinding is the Schema for the SnapshotScheduleIAMBindings API. Allows management of a single IAM binding for a Snapshot Schedule.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type SnapshotScheduleIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.members) || (has(self.initProvider) && has(self.initProvider.members))",message="spec.forProvider.members is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.snapshotScheduleId) || (has(self.initProvider) && has(self.initProvider.snapshotScheduleId))",message="spec.forProvider.snapshotScheduleId is a required parameter"
	Spec   SnapshotScheduleIAMBindingSpec   `json:"spec"`
	Status SnapshotScheduleIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotScheduleIAMBindingList contains a list of SnapshotScheduleIAMBindings
type SnapshotScheduleIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotScheduleIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	SnapshotScheduleIAMBinding_Kind             = "SnapshotScheduleIAMBinding"
	SnapshotScheduleIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnapshotScheduleIAMBinding_Kind}.String()
	SnapshotScheduleIAMBinding_KindAPIVersion   = SnapshotScheduleIAMBinding_Kind + "." + CRDGroupVersion.String()
	SnapshotScheduleIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(SnapshotScheduleIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&SnapshotScheduleIAMBinding{}, &SnapshotScheduleIAMBindingList{})
}
