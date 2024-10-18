// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CoreDeviceInitParameters struct {

	// A set of key/value aliases pairs to assign to the IoT Core Device
	// +mapType=granular
	Aliases map[string]*string `json:"aliases,omitempty" tf:"aliases,omitempty"`

	// A set of certificate's fingerprints for the IoT Core Device
	// +listType=set
	Certificates []*string `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// Description of the IoT Core Device
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// IoT Core Device name used to define device
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Passwords []*string `json:"passwordsSecretRef,omitempty" tf:"-"`

	// IoT Core Registry ID for the IoT Core Device
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`
}

type CoreDeviceObservation struct {

	// A set of key/value aliases pairs to assign to the IoT Core Device
	// +mapType=granular
	Aliases map[string]*string `json:"aliases,omitempty" tf:"aliases,omitempty"`

	// A set of certificate's fingerprints for the IoT Core Device
	// +listType=set
	Certificates []*string `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// Creation timestamp of the IoT Core Device
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Description of the IoT Core Device
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IoT Core Device name used to define device
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// IoT Core Registry ID for the IoT Core Device
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`
}

type CoreDeviceParameters struct {

	// A set of key/value aliases pairs to assign to the IoT Core Device
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Aliases map[string]*string `json:"aliases,omitempty" tf:"aliases,omitempty"`

	// A set of certificate's fingerprints for the IoT Core Device
	// +kubebuilder:validation:Optional
	// +listType=set
	Certificates []*string `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// Description of the IoT Core Device
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// IoT Core Device name used to define device
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A set of passwords's id for the IoT Core Device
	// +kubebuilder:validation:Optional
	PasswordsSecretRef *[]v1.SecretKeySelector `json:"passwordsSecretRef,omitempty" tf:"-"`

	// IoT Core Registry ID for the IoT Core Device
	// +kubebuilder:validation:Optional
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`
}

// CoreDeviceSpec defines the desired state of CoreDevice
type CoreDeviceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CoreDeviceParameters `json:"forProvider"`
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
	InitProvider CoreDeviceInitParameters `json:"initProvider,omitempty"`
}

// CoreDeviceStatus defines the observed state of CoreDevice.
type CoreDeviceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CoreDeviceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CoreDevice is the Schema for the CoreDevices API. Allows management of a Yandex.Cloud IoT Core Device.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type CoreDevice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.registryId) || (has(self.initProvider) && has(self.initProvider.registryId))",message="spec.forProvider.registryId is a required parameter"
	Spec   CoreDeviceSpec   `json:"spec"`
	Status CoreDeviceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CoreDeviceList contains a list of CoreDevices
type CoreDeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CoreDevice `json:"items"`
}

// Repository type metadata.
var (
	CoreDevice_Kind             = "CoreDevice"
	CoreDevice_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CoreDevice_Kind}.String()
	CoreDevice_KindAPIVersion   = CoreDevice_Kind + "." + CRDGroupVersion.String()
	CoreDevice_GroupVersionKind = CRDGroupVersion.WithKind(CoreDevice_Kind)
)

func init() {
	SchemeBuilder.Register(&CoreDevice{}, &CoreDeviceList{})
}