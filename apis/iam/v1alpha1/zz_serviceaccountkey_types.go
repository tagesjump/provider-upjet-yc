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

type ServiceAccountKeyInitParameters struct {

	// The description of the key pair.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The output format of the keys. PEM_FILE is the default format.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The algorithm used to generate the key. RSA_2048 is the default algorithm.
	// Valid values are listed in the API reference.
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty" tf:"key_algorithm,omitempty"`

	// An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a keybase username in the form keybase:keybaseusername.
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`
}

type ServiceAccountKeyObservation struct {

	// Creation timestamp of the static access key.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The description of the key pair.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The encrypted private key, base64 encoded. This is only populated when pgp_key is supplied.
	EncryptedPrivateKey *string `json:"encryptedPrivateKey,omitempty" tf:"encrypted_private_key,omitempty"`

	// The output format of the keys. PEM_FILE is the default format.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The algorithm used to generate the key. RSA_2048 is the default algorithm.
	// Valid values are listed in the API reference.
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty" tf:"key_algorithm,omitempty"`

	// The fingerprint of the PGP key used to encrypt the private key. This is only populated when pgp_key is supplied.
	KeyFingerprint *string `json:"keyFingerprint,omitempty" tf:"key_fingerprint,omitempty"`

	// An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a keybase username in the form keybase:keybaseusername.
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// The public key.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// ID of the service account to create a pair for.
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`
}

type ServiceAccountKeyParameters struct {

	// The description of the key pair.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The output format of the keys. PEM_FILE is the default format.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The algorithm used to generate the key. RSA_2048 is the default algorithm.
	// Valid values are listed in the API reference.
	// +kubebuilder:validation:Optional
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty" tf:"key_algorithm,omitempty"`

	// An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a keybase username in the form keybase:keybaseusername.
	// +kubebuilder:validation:Optional
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// ID of the service account to create a pair for.
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

// ServiceAccountKeySpec defines the desired state of ServiceAccountKey
type ServiceAccountKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAccountKeyParameters `json:"forProvider"`
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
	InitProvider ServiceAccountKeyInitParameters `json:"initProvider,omitempty"`
}

// ServiceAccountKeyStatus defines the observed state of ServiceAccountKey.
type ServiceAccountKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAccountKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountKey is the Schema for the ServiceAccountKeys API. Allows management of a Yandex.Cloud IAM service account key.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type ServiceAccountKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountKeySpec   `json:"spec"`
	Status            ServiceAccountKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountKeyList contains a list of ServiceAccountKeys
type ServiceAccountKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccountKey `json:"items"`
}

// Repository type metadata.
var (
	ServiceAccountKey_Kind             = "ServiceAccountKey"
	ServiceAccountKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAccountKey_Kind}.String()
	ServiceAccountKey_KindAPIVersion   = ServiceAccountKey_Kind + "." + CRDGroupVersion.String()
	ServiceAccountKey_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAccountKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccountKey{}, &ServiceAccountKeyList{})
}
