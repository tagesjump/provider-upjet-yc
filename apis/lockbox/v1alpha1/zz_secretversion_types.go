// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CommandInitParameters struct {

	// List of arguments to be passed to the script/command.
	Args []*string `json:"args,omitempty" tf:"args,omitempty"`

	// Map of environment variables to set before calling the script/command.
	// +mapType=granular
	Env map[string]*string `json:"env,omitempty" tf:"env,omitempty"`

	// The path to the script or command to execute.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type CommandObservation struct {

	// List of arguments to be passed to the script/command.
	Args []*string `json:"args,omitempty" tf:"args,omitempty"`

	// Map of environment variables to set before calling the script/command.
	// +mapType=granular
	Env map[string]*string `json:"env,omitempty" tf:"env,omitempty"`

	// The path to the script or command to execute.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type CommandParameters struct {

	// List of arguments to be passed to the script/command.
	// +kubebuilder:validation:Optional
	Args []*string `json:"args,omitempty" tf:"args,omitempty"`

	// Map of environment variables to set before calling the script/command.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Env map[string]*string `json:"env,omitempty" tf:"env,omitempty"`

	// The path to the script or command to execute.
	// +kubebuilder:validation:Optional
	Path *string `json:"path" tf:"path,omitempty"`
}

type EntriesInitParameters struct {

	// The command that generates the text value of the entry.
	Command []CommandInitParameters `json:"command,omitempty" tf:"command,omitempty"`

	// The key of the entry.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The text value of the entry.
	TextValueSecretRef *v1.SecretKeySelector `json:"textValueSecretRef,omitempty" tf:"-"`
}

type EntriesObservation struct {

	// The command that generates the text value of the entry.
	Command []CommandObservation `json:"command,omitempty" tf:"command,omitempty"`

	// The key of the entry.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type EntriesParameters struct {

	// The command that generates the text value of the entry.
	// +kubebuilder:validation:Optional
	Command []CommandParameters `json:"command,omitempty" tf:"command,omitempty"`

	// The key of the entry.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// The text value of the entry.
	// +kubebuilder:validation:Optional
	TextValueSecretRef *v1.SecretKeySelector `json:"textValueSecretRef,omitempty" tf:"-"`
}

type SecretVersionInitParameters struct {

	// The Yandex Cloud Lockbox secret version description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of entries in the Yandex Cloud Lockbox secret version.
	Entries []EntriesInitParameters `json:"entries,omitempty" tf:"entries,omitempty"`

	// The Yandex Cloud Lockbox secret ID where to add the version.
	// +crossplane:generate:reference:type=Secret
	SecretID *string `json:"secretId,omitempty" tf:"secret_id,omitempty"`

	// Reference to a Secret to populate secretId.
	// +kubebuilder:validation:Optional
	SecretIDRef *v1.Reference `json:"secretIdRef,omitempty" tf:"-"`

	// Selector for a Secret to populate secretId.
	// +kubebuilder:validation:Optional
	SecretIDSelector *v1.Selector `json:"secretIdSelector,omitempty" tf:"-"`
}

type SecretVersionObservation struct {

	// The Yandex Cloud Lockbox secret version description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of entries in the Yandex Cloud Lockbox secret version.
	Entries []EntriesObservation `json:"entries,omitempty" tf:"entries,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Yandex Cloud Lockbox secret ID where to add the version.
	SecretID *string `json:"secretId,omitempty" tf:"secret_id,omitempty"`
}

type SecretVersionParameters struct {

	// The Yandex Cloud Lockbox secret version description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of entries in the Yandex Cloud Lockbox secret version.
	// +kubebuilder:validation:Optional
	Entries []EntriesParameters `json:"entries,omitempty" tf:"entries,omitempty"`

	// The Yandex Cloud Lockbox secret ID where to add the version.
	// +crossplane:generate:reference:type=Secret
	// +kubebuilder:validation:Optional
	SecretID *string `json:"secretId,omitempty" tf:"secret_id,omitempty"`

	// Reference to a Secret to populate secretId.
	// +kubebuilder:validation:Optional
	SecretIDRef *v1.Reference `json:"secretIdRef,omitempty" tf:"-"`

	// Selector for a Secret to populate secretId.
	// +kubebuilder:validation:Optional
	SecretIDSelector *v1.Selector `json:"secretIdSelector,omitempty" tf:"-"`
}

// SecretVersionSpec defines the desired state of SecretVersion
type SecretVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretVersionParameters `json:"forProvider"`
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
	InitProvider SecretVersionInitParameters `json:"initProvider,omitempty"`
}

// SecretVersionStatus defines the observed state of SecretVersion.
type SecretVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecretVersion is the Schema for the SecretVersions API. Manages Yandex Cloud Lockbox secret version.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type SecretVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.entries) || (has(self.initProvider) && has(self.initProvider.entries))",message="spec.forProvider.entries is a required parameter"
	Spec   SecretVersionSpec   `json:"spec"`
	Status SecretVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretVersionList contains a list of SecretVersions
type SecretVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretVersion `json:"items"`
}

// Repository type metadata.
var (
	SecretVersion_Kind             = "SecretVersion"
	SecretVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretVersion_Kind}.String()
	SecretVersion_KindAPIVersion   = SecretVersion_Kind + "." + CRDGroupVersion.String()
	SecretVersion_GroupVersionKind = CRDGroupVersion.WithKind(SecretVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretVersion{}, &SecretVersionList{})
}
