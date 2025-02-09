// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CanaryInitParameters struct {

	// A set of values for variables in gateway specification.
	// +mapType=granular
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`

	// Percentage of requests, which will be processed by canary release.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type CanaryObservation struct {

	// A set of values for variables in gateway specification.
	// +mapType=granular
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`

	// Percentage of requests, which will be processed by canary release.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type CanaryParameters struct {

	// A set of values for variables in gateway specification.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`

	// Percentage of requests, which will be processed by canary release.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type ConnectivityInitParameters struct {

	// Network the gateway will have access to. It's essential to specify network with subnets in all availability zones.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`
}

type ConnectivityObservation struct {

	// Network the gateway will have access to. It's essential to specify network with subnets in all availability zones.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`
}

type ConnectivityParameters struct {

	// Network the gateway will have access to. It's essential to specify network with subnets in all availability zones.
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId" tf:"network_id,omitempty"`
}

type CustomDomainsInitParameters struct {
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
}

type CustomDomainsObservation struct {
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
}

type CustomDomainsParameters struct {

	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId" tf:"certificate_id,omitempty"`

	// +kubebuilder:validation:Optional
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// +kubebuilder:validation:Optional
	Fqdn *string `json:"fqdn" tf:"fqdn,omitempty"`
}

type GatewayInitParameters struct {

	// Canary release settings of gateway.
	Canary []CanaryInitParameters `json:"canary,omitempty" tf:"canary,omitempty"`

	// Gateway connectivity. If specified the gateway will be attached to specified network.
	Connectivity []ConnectivityInitParameters `json:"connectivity,omitempty" tf:"connectivity,omitempty"`

	// Set of custom domains to be attached to Yandex API Gateway.
	CustomDomains []CustomDomainsInitParameters `json:"customDomains,omitempty" tf:"custom_domains,omitempty"`

	// Description of the Yandex Cloud API Gateway.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Execution timeout in seconds for the Yandex Cloud API Gateway.
	ExecutionTimeout *string `json:"executionTimeout,omitempty" tf:"execution_timeout,omitempty"`

	// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the Yandex Cloud API Gateway.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Options for logging from Yandex Cloud API Gateway.
	LogOptions []LogOptionsInitParameters `json:"logOptions,omitempty" tf:"log_options,omitempty"`

	// Yandex Cloud API Gateway name used to define API Gateway.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// OpenAPI specification for Yandex API Gateway.
	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	// A set of values for variables in gateway specification.
	// +mapType=granular
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`
}

type GatewayObservation struct {

	// Canary release settings of gateway.
	Canary []CanaryObservation `json:"canary,omitempty" tf:"canary,omitempty"`

	// Gateway connectivity. If specified the gateway will be attached to specified network.
	Connectivity []ConnectivityObservation `json:"connectivity,omitempty" tf:"connectivity,omitempty"`

	// Creation timestamp of the Yandex Cloud API Gateway.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Set of custom domains to be attached to Yandex API Gateway.
	CustomDomains []CustomDomainsObservation `json:"customDomains,omitempty" tf:"custom_domains,omitempty"`

	// Description of the Yandex Cloud API Gateway.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Default domain for the Yandex API Gateway. Generated at creation time.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Execution timeout in seconds for the Yandex Cloud API Gateway.
	ExecutionTimeout *string `json:"executionTimeout,omitempty" tf:"execution_timeout,omitempty"`

	// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the Yandex Cloud API Gateway.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Log entries are written to specified log group
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	// Options for logging from Yandex Cloud API Gateway.
	LogOptions []LogOptionsObservation `json:"logOptions,omitempty" tf:"log_options,omitempty"`

	// Yandex Cloud API Gateway name used to define API Gateway.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// OpenAPI specification for Yandex API Gateway.
	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	// Status of the Yandex API Gateway.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (DEPRECATED, use custom_domains instead) Set of user domains attached to Yandex API Gateway.
	// +listType=set
	UserDomains []*string `json:"userDomains,omitempty" tf:"user_domains,omitempty"`

	// A set of values for variables in gateway specification.
	// +mapType=granular
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`
}

type GatewayParameters struct {

	// Canary release settings of gateway.
	// +kubebuilder:validation:Optional
	Canary []CanaryParameters `json:"canary,omitempty" tf:"canary,omitempty"`

	// Gateway connectivity. If specified the gateway will be attached to specified network.
	// +kubebuilder:validation:Optional
	Connectivity []ConnectivityParameters `json:"connectivity,omitempty" tf:"connectivity,omitempty"`

	// Set of custom domains to be attached to Yandex API Gateway.
	// +kubebuilder:validation:Optional
	CustomDomains []CustomDomainsParameters `json:"customDomains,omitempty" tf:"custom_domains,omitempty"`

	// Description of the Yandex Cloud API Gateway.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Execution timeout in seconds for the Yandex Cloud API Gateway.
	// +kubebuilder:validation:Optional
	ExecutionTimeout *string `json:"executionTimeout,omitempty" tf:"execution_timeout,omitempty"`

	// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the Yandex Cloud API Gateway.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Options for logging from Yandex Cloud API Gateway.
	// +kubebuilder:validation:Optional
	LogOptions []LogOptionsParameters `json:"logOptions,omitempty" tf:"log_options,omitempty"`

	// Yandex Cloud API Gateway name used to define API Gateway.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// OpenAPI specification for Yandex API Gateway.
	// +kubebuilder:validation:Optional
	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	// A set of values for variables in gateway specification.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`
}

type LogOptionsInitParameters struct {

	// Is logging from API Gateway disabled
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Log entries are written to specified log group
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	// Minimum log entry level
	MinLevel *string `json:"minLevel,omitempty" tf:"min_level,omitempty"`
}

type LogOptionsObservation struct {

	// Is logging from API Gateway disabled
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Log entries are written to specified log group
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	// Minimum log entry level
	MinLevel *string `json:"minLevel,omitempty" tf:"min_level,omitempty"`
}

type LogOptionsParameters struct {

	// Is logging from API Gateway disabled
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Folder ID for the Yandex Cloud API Gateway. If it is not provided, the default provider folder is used.
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Log entries are written to specified log group
	// +kubebuilder:validation:Optional
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	// Minimum log entry level
	// +kubebuilder:validation:Optional
	MinLevel *string `json:"minLevel,omitempty" tf:"min_level,omitempty"`
}

// GatewaySpec defines the desired state of Gateway
type GatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayParameters `json:"forProvider"`
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
	InitProvider GatewayInitParameters `json:"initProvider,omitempty"`
}

// GatewayStatus defines the observed state of Gateway.
type GatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Gateway is the Schema for the Gateways API. Allows management of a Yandex Cloud API Gateway.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.spec) || (has(self.initProvider) && has(self.initProvider.spec))",message="spec.forProvider.spec is a required parameter"
	Spec   GatewaySpec   `json:"spec"`
	Status GatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayList contains a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// Repository type metadata.
var (
	Gateway_Kind             = "Gateway"
	Gateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Gateway_Kind}.String()
	Gateway_KindAPIVersion   = Gateway_Kind + "." + CRDGroupVersion.String()
	Gateway_GroupVersionKind = CRDGroupVersion.WithKind(Gateway_Kind)
)

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}
