// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecretInitParameters struct {

	// Whether the Yandex Cloud Lockbox secret is protected from deletion.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A description for the Yandex Cloud Lockbox secret.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder that the Yandex Cloud Lockbox secret belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// The KMS key used to encrypt the Yandex Cloud Lockbox secret.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/kms/v1alpha1.SymmetricKey
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a SymmetricKey in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a SymmetricKey in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the Yandex Cloud Lockbox secret.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name for the Yandex Cloud Lockbox secret.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SecretObservation struct {

	// The Yandex Cloud Lockbox secret creation timestamp.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Whether the Yandex Cloud Lockbox secret is protected from deletion.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A description for the Yandex Cloud Lockbox secret.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder that the Yandex Cloud Lockbox secret belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The KMS key used to encrypt the Yandex Cloud Lockbox secret.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// A set of key/value label pairs to assign to the Yandex Cloud Lockbox secret.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name for the Yandex Cloud Lockbox secret.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Yandex Cloud Lockbox secret status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SecretParameters struct {

	// Whether the Yandex Cloud Lockbox secret is protected from deletion.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A description for the Yandex Cloud Lockbox secret.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder that the Yandex Cloud Lockbox secret belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// The KMS key used to encrypt the Yandex Cloud Lockbox secret.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/kms/v1alpha1.SymmetricKey
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a SymmetricKey in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a SymmetricKey in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the Yandex Cloud Lockbox secret.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name for the Yandex Cloud Lockbox secret.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// SecretSpec defines the desired state of Secret
type SecretSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretParameters `json:"forProvider"`
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
	InitProvider SecretInitParameters `json:"initProvider,omitempty"`
}

// SecretStatus defines the observed state of Secret.
type SecretStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Secret is the Schema for the Secrets API. Manages Yandex Cloud Lockbox secret.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretSpec   `json:"spec"`
	Status            SecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretList contains a list of Secrets
type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Secret `json:"items"`
}

// Repository type metadata.
var (
	Secret_Kind             = "Secret"
	Secret_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Secret_Kind}.String()
	Secret_KindAPIVersion   = Secret_Kind + "." + CRDGroupVersion.String()
	Secret_GroupVersionKind = CRDGroupVersion.WithKind(Secret_Kind)
)

func init() {
	SchemeBuilder.Register(&Secret{}, &SecretList{})
}
