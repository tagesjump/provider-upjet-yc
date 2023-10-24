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

type CloudInitParameters struct {

	// A description of the Cloud.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of key/value label pairs to assign to the Cloud.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the Cloud.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Yandex.Cloud Organization that the cloud belongs to. If value is omitted, the default provider Organization ID is used.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`
}

type CloudObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// A description of the Cloud.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the Cloud.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the Cloud.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Yandex.Cloud Organization that the cloud belongs to. If value is omitted, the default provider Organization ID is used.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`
}

type CloudParameters struct {

	// A description of the Cloud.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of key/value label pairs to assign to the Cloud.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the Cloud.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Yandex.Cloud Organization that the cloud belongs to. If value is omitted, the default provider Organization ID is used.
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`
}

// CloudSpec defines the desired state of Cloud
type CloudSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudParameters `json:"forProvider"`
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
	InitProvider CloudInitParameters `json:"initProvider,omitempty"`
}

// CloudStatus defines the observed state of Cloud.
type CloudStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cloud is the Schema for the Clouds API. Allows management of the Cloud resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Cloud struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudSpec   `json:"spec"`
	Status            CloudStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudList contains a list of Clouds
type CloudList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cloud `json:"items"`
}

// Repository type metadata.
var (
	Cloud_Kind             = "Cloud"
	Cloud_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cloud_Kind}.String()
	Cloud_KindAPIVersion   = Cloud_Kind + "." + CRDGroupVersion.String()
	Cloud_GroupVersionKind = CRDGroupVersion.WithKind(Cloud_Kind)
)

func init() {
	SchemeBuilder.Register(&Cloud{}, &CloudList{})
}
