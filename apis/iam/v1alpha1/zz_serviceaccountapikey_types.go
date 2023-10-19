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

type ServiceAccountAPIKeyInitParameters struct {

	// The description of the key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form keybase:keybaseusername.
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`
}

type ServiceAccountAPIKeyObservation struct {

	// Creation timestamp of the static access key.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The description of the key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The encrypted secret key, base64 encoded. This is only populated when pgp_key is supplied.
	EncryptedSecretKey *string `json:"encryptedSecretKey,omitempty" tf:"encrypted_secret_key,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The fingerprint of the PGP key used to encrypt the secret key. This is only populated when pgp_key is supplied.
	KeyFingerprint *string `json:"keyFingerprint,omitempty" tf:"key_fingerprint,omitempty"`

	// An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form keybase:keybaseusername.
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// ID of the service account to an API key for.
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`
}

type ServiceAccountAPIKeyParameters struct {

	// The description of the key.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form keybase:keybaseusername.
	// +kubebuilder:validation:Optional
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// ID of the service account to an API key for.
	// +crossplane:generate:reference:type=ServiceAccount
	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// Reference to a ServiceAccount to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`
}

// ServiceAccountAPIKeySpec defines the desired state of ServiceAccountAPIKey
type ServiceAccountAPIKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAccountAPIKeyParameters `json:"forProvider"`
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
	InitProvider ServiceAccountAPIKeyInitParameters `json:"initProvider,omitempty"`
}

// ServiceAccountAPIKeyStatus defines the observed state of ServiceAccountAPIKey.
type ServiceAccountAPIKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAccountAPIKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountAPIKey is the Schema for the ServiceAccountAPIKeys API. Allows management of a Yandex.Cloud IAM service account API key.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud.upjet}
type ServiceAccountAPIKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountAPIKeySpec   `json:"spec"`
	Status            ServiceAccountAPIKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountAPIKeyList contains a list of ServiceAccountAPIKeys
type ServiceAccountAPIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccountAPIKey `json:"items"`
}

// Repository type metadata.
var (
	ServiceAccountAPIKey_Kind             = "ServiceAccountAPIKey"
	ServiceAccountAPIKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAccountAPIKey_Kind}.String()
	ServiceAccountAPIKey_KindAPIVersion   = ServiceAccountAPIKey_Kind + "." + CRDGroupVersion.String()
	ServiceAccountAPIKey_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAccountAPIKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccountAPIKey{}, &ServiceAccountAPIKeyList{})
}