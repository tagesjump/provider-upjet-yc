

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

)




type AttachedTargetGroupInitParameters struct {


// A HealthCheck resource. The structure is documented below.
Healthcheck []HealthcheckInitParameters `json:"healthcheck,omitempty" tf:"healthcheck,omitempty"`

// ID of the target group.
// +crossplane:generate:reference:type=TargetGroup
TargetGroupID *string `json:"targetGroupId,omitempty" tf:"target_group_id,omitempty"`

// Reference to a TargetGroup to populate targetGroupId.
// +kubebuilder:validation:Optional
TargetGroupIDRef *v1.Reference `json:"targetGroupIdRef,omitempty" tf:"-"`

// Selector for a TargetGroup to populate targetGroupId.
// +kubebuilder:validation:Optional
TargetGroupIDSelector *v1.Selector `json:"targetGroupIdSelector,omitempty" tf:"-"`
}


type AttachedTargetGroupObservation struct {


// A HealthCheck resource. The structure is documented below.
Healthcheck []HealthcheckObservation `json:"healthcheck,omitempty" tf:"healthcheck,omitempty"`

// ID of the target group.
TargetGroupID *string `json:"targetGroupId,omitempty" tf:"target_group_id,omitempty"`
}


type AttachedTargetGroupParameters struct {


// A HealthCheck resource. The structure is documented below.
// +kubebuilder:validation:Optional
Healthcheck []HealthcheckParameters `json:"healthcheck" tf:"healthcheck,omitempty"`

// ID of the target group.
// +crossplane:generate:reference:type=TargetGroup
// +kubebuilder:validation:Optional
TargetGroupID *string `json:"targetGroupId,omitempty" tf:"target_group_id,omitempty"`

// Reference to a TargetGroup to populate targetGroupId.
// +kubebuilder:validation:Optional
TargetGroupIDRef *v1.Reference `json:"targetGroupIdRef,omitempty" tf:"-"`

// Selector for a TargetGroup to populate targetGroupId.
// +kubebuilder:validation:Optional
TargetGroupIDSelector *v1.Selector `json:"targetGroupIdSelector,omitempty" tf:"-"`
}


type ExternalAddressSpecInitParameters struct {


// Internal IP address for a listener. Must belong to the subnet that is referenced in subnet_id. IP address will be allocated if it wasn't been set.
Address *string `json:"address,omitempty" tf:"address,omitempty"`

// IP version of the internal addresses that the load balancer works with. Must be one of ipv4 or ipv6. The default is ipv4.
IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`
}


type ExternalAddressSpecObservation struct {


// Internal IP address for a listener. Must belong to the subnet that is referenced in subnet_id. IP address will be allocated if it wasn't been set.
Address *string `json:"address,omitempty" tf:"address,omitempty"`

// IP version of the internal addresses that the load balancer works with. Must be one of ipv4 or ipv6. The default is ipv4.
IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`
}


type ExternalAddressSpecParameters struct {


// Internal IP address for a listener. Must belong to the subnet that is referenced in subnet_id. IP address will be allocated if it wasn't been set.
// +kubebuilder:validation:Optional
Address *string `json:"address,omitempty" tf:"address,omitempty"`

// IP version of the internal addresses that the load balancer works with. Must be one of ipv4 or ipv6. The default is ipv4.
// +kubebuilder:validation:Optional
IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`
}


type HTTPOptionsInitParameters struct {


// URL path to set for health checking requests for every target in the target group. For example /ping. The default path is /.
Path *string `json:"path,omitempty" tf:"path,omitempty"`

// Port to use for TCP health checks.
Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}


type HTTPOptionsObservation struct {


// URL path to set for health checking requests for every target in the target group. For example /ping. The default path is /.
Path *string `json:"path,omitempty" tf:"path,omitempty"`

// Port to use for TCP health checks.
Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}


type HTTPOptionsParameters struct {


// URL path to set for health checking requests for every target in the target group. For example /ping. The default path is /.
// +kubebuilder:validation:Optional
Path *string `json:"path,omitempty" tf:"path,omitempty"`

// Port to use for TCP health checks.
// +kubebuilder:validation:Optional
Port *float64 `json:"port" tf:"port,omitempty"`
}


type HealthcheckInitParameters struct {


// Options for HTTP health check. The structure is documented below.
HTTPOptions []HTTPOptionsInitParameters `json:"httpOptions,omitempty" tf:"http_options,omitempty"`

// Number of successful health checks required in order to set the HEALTHY status for the target.
HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

// The interval between health checks. The default is 2 seconds.
Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

// Name of the network load balancer. Provided by the client when the network load balancer is created.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// Options for TCP health check. The structure is documented below.
TCPOptions []TCPOptionsInitParameters `json:"tcpOptions,omitempty" tf:"tcp_options,omitempty"`

// Timeout for a target to return a response for the health check. The default is 1 second.
Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

// Number of failed health checks before changing the status to UNHEALTHY. The default is 2.
UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}


type HealthcheckObservation struct {


// Options for HTTP health check. The structure is documented below.
HTTPOptions []HTTPOptionsObservation `json:"httpOptions,omitempty" tf:"http_options,omitempty"`

// Number of successful health checks required in order to set the HEALTHY status for the target.
HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

// The interval between health checks. The default is 2 seconds.
Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

// Name of the network load balancer. Provided by the client when the network load balancer is created.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// Options for TCP health check. The structure is documented below.
TCPOptions []TCPOptionsObservation `json:"tcpOptions,omitempty" tf:"tcp_options,omitempty"`

// Timeout for a target to return a response for the health check. The default is 1 second.
Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

// Number of failed health checks before changing the status to UNHEALTHY. The default is 2.
UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}


type HealthcheckParameters struct {


// Options for HTTP health check. The structure is documented below.
// +kubebuilder:validation:Optional
HTTPOptions []HTTPOptionsParameters `json:"httpOptions,omitempty" tf:"http_options,omitempty"`

// Number of successful health checks required in order to set the HEALTHY status for the target.
// +kubebuilder:validation:Optional
HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

// The interval between health checks. The default is 2 seconds.
// +kubebuilder:validation:Optional
Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

// Name of the network load balancer. Provided by the client when the network load balancer is created.
// +kubebuilder:validation:Optional
Name *string `json:"name" tf:"name,omitempty"`

// Options for TCP health check. The structure is documented below.
// +kubebuilder:validation:Optional
TCPOptions []TCPOptionsParameters `json:"tcpOptions,omitempty" tf:"tcp_options,omitempty"`

// Timeout for a target to return a response for the health check. The default is 1 second.
// +kubebuilder:validation:Optional
Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

// Number of failed health checks before changing the status to UNHEALTHY. The default is 2.
// +kubebuilder:validation:Optional
UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}


type InternalAddressSpecInitParameters struct {


// Internal IP address for a listener. Must belong to the subnet that is referenced in subnet_id. IP address will be allocated if it wasn't been set.
Address *string `json:"address,omitempty" tf:"address,omitempty"`

// IP version of the internal addresses that the load balancer works with. Must be one of ipv4 or ipv6. The default is ipv4.
IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

// ID of the subnet to which the internal IP address belongs.
// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Subnet
SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

// Reference to a Subnet in vpc to populate subnetId.
// +kubebuilder:validation:Optional
SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

// Selector for a Subnet in vpc to populate subnetId.
// +kubebuilder:validation:Optional
SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}


type InternalAddressSpecObservation struct {


// Internal IP address for a listener. Must belong to the subnet that is referenced in subnet_id. IP address will be allocated if it wasn't been set.
Address *string `json:"address,omitempty" tf:"address,omitempty"`

// IP version of the internal addresses that the load balancer works with. Must be one of ipv4 or ipv6. The default is ipv4.
IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

// ID of the subnet to which the internal IP address belongs.
SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}


type InternalAddressSpecParameters struct {


// Internal IP address for a listener. Must belong to the subnet that is referenced in subnet_id. IP address will be allocated if it wasn't been set.
// +kubebuilder:validation:Optional
Address *string `json:"address,omitempty" tf:"address,omitempty"`

// IP version of the internal addresses that the load balancer works with. Must be one of ipv4 or ipv6. The default is ipv4.
// +kubebuilder:validation:Optional
IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

// ID of the subnet to which the internal IP address belongs.
// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Subnet
// +kubebuilder:validation:Optional
SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

// Reference to a Subnet in vpc to populate subnetId.
// +kubebuilder:validation:Optional
SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

// Selector for a Subnet in vpc to populate subnetId.
// +kubebuilder:validation:Optional
SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}


type ListenerInitParameters struct {


// External IP address specification. The structure is documented below.
ExternalAddressSpec []ExternalAddressSpecInitParameters `json:"externalAddressSpec,omitempty" tf:"external_address_spec,omitempty"`

// Internal IP address specification. The structure is documented below.
InternalAddressSpec []InternalAddressSpecInitParameters `json:"internalAddressSpec,omitempty" tf:"internal_address_spec,omitempty"`

// Name of the listener. The name must be unique for each listener on a single load balancer.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// Port for incoming traffic.
Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

// Protocol for incoming traffic. TCP or UDP and the default is TCP.
Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

// Port of a target. The default is the same as listener's port.
TargetPort *float64 `json:"targetPort,omitempty" tf:"target_port,omitempty"`
}


type ListenerObservation struct {


// External IP address specification. The structure is documented below.
ExternalAddressSpec []ExternalAddressSpecObservation `json:"externalAddressSpec,omitempty" tf:"external_address_spec,omitempty"`

// Internal IP address specification. The structure is documented below.
InternalAddressSpec []InternalAddressSpecObservation `json:"internalAddressSpec,omitempty" tf:"internal_address_spec,omitempty"`

// Name of the listener. The name must be unique for each listener on a single load balancer.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// Port for incoming traffic.
Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

// Protocol for incoming traffic. TCP or UDP and the default is TCP.
Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

// Port of a target. The default is the same as listener's port.
TargetPort *float64 `json:"targetPort,omitempty" tf:"target_port,omitempty"`
}


type ListenerParameters struct {


// External IP address specification. The structure is documented below.
// +kubebuilder:validation:Optional
ExternalAddressSpec []ExternalAddressSpecParameters `json:"externalAddressSpec,omitempty" tf:"external_address_spec,omitempty"`

// Internal IP address specification. The structure is documented below.
// +kubebuilder:validation:Optional
InternalAddressSpec []InternalAddressSpecParameters `json:"internalAddressSpec,omitempty" tf:"internal_address_spec,omitempty"`

// Name of the listener. The name must be unique for each listener on a single load balancer.
// +kubebuilder:validation:Optional
Name *string `json:"name" tf:"name,omitempty"`

// Port for incoming traffic.
// +kubebuilder:validation:Optional
Port *float64 `json:"port" tf:"port,omitempty"`

// Protocol for incoming traffic. TCP or UDP and the default is TCP.
// +kubebuilder:validation:Optional
Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

// Port of a target. The default is the same as listener's port.
// +kubebuilder:validation:Optional
TargetPort *float64 `json:"targetPort,omitempty" tf:"target_port,omitempty"`
}


type NetworkLoadBalancerInitParameters struct {


// An AttachedTargetGroup resource. The structure is documented below.
AttachedTargetGroup []AttachedTargetGroupInitParameters `json:"attachedTargetGroup,omitempty" tf:"attached_target_group,omitempty"`

// Flag that protects the network load balancer from accidental deletion.
DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

// An optional description of the network load balancer. Provide this property when you create the resource.
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// The ID of the folder to which the resource belongs. If omitted, the provider folder is used.
// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

// Reference to a Folder in resourcemanager to populate folderId.
// +kubebuilder:validation:Optional
FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

// Selector for a Folder in resourcemanager to populate folderId.
// +kubebuilder:validation:Optional
FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

// Labels to assign to this network load balancer. A list of key/value pairs.
// +mapType=granular
Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

// Listener specification that will be used by a network load balancer. The structure is documented below.
Listener []ListenerInitParameters `json:"listener,omitempty" tf:"listener,omitempty"`

// Name of the network load balancer. Provided by the client when the network load balancer is created.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// ID of the availability zone where the network load balancer resides. If omitted, default region is being used.
RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

// Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.
Type *string `json:"type,omitempty" tf:"type,omitempty"`
}


type NetworkLoadBalancerObservation struct {


// An AttachedTargetGroup resource. The structure is documented below.
AttachedTargetGroup []AttachedTargetGroupObservation `json:"attachedTargetGroup,omitempty" tf:"attached_target_group,omitempty"`

// The network load balancer creation timestamp.
CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

// Flag that protects the network load balancer from accidental deletion.
DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

// An optional description of the network load balancer. Provide this property when you create the resource.
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// The ID of the folder to which the resource belongs. If omitted, the provider folder is used.
FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

// The ID of the network load balancer.
ID *string `json:"id,omitempty" tf:"id,omitempty"`

// Labels to assign to this network load balancer. A list of key/value pairs.
// +mapType=granular
Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

// Listener specification that will be used by a network load balancer. The structure is documented below.
Listener []ListenerObservation `json:"listener,omitempty" tf:"listener,omitempty"`

// Name of the network load balancer. Provided by the client when the network load balancer is created.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// ID of the availability zone where the network load balancer resides. If omitted, default region is being used.
RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

// Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.
Type *string `json:"type,omitempty" tf:"type,omitempty"`
}


type NetworkLoadBalancerParameters struct {


// An AttachedTargetGroup resource. The structure is documented below.
// +kubebuilder:validation:Optional
AttachedTargetGroup []AttachedTargetGroupParameters `json:"attachedTargetGroup,omitempty" tf:"attached_target_group,omitempty"`

// Flag that protects the network load balancer from accidental deletion.
// +kubebuilder:validation:Optional
DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

// An optional description of the network load balancer. Provide this property when you create the resource.
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// The ID of the folder to which the resource belongs. If omitted, the provider folder is used.
// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
// +kubebuilder:validation:Optional
FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

// Reference to a Folder in resourcemanager to populate folderId.
// +kubebuilder:validation:Optional
FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

// Selector for a Folder in resourcemanager to populate folderId.
// +kubebuilder:validation:Optional
FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

// Labels to assign to this network load balancer. A list of key/value pairs.
// +kubebuilder:validation:Optional
// +mapType=granular
Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

// Listener specification that will be used by a network load balancer. The structure is documented below.
// +kubebuilder:validation:Optional
Listener []ListenerParameters `json:"listener,omitempty" tf:"listener,omitempty"`

// Name of the network load balancer. Provided by the client when the network load balancer is created.
// +kubebuilder:validation:Optional
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// ID of the availability zone where the network load balancer resides. If omitted, default region is being used.
// +kubebuilder:validation:Optional
RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

// Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.
// +kubebuilder:validation:Optional
Type *string `json:"type,omitempty" tf:"type,omitempty"`
}


type TCPOptionsInitParameters struct {


// Port to use for TCP health checks.
Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}


type TCPOptionsObservation struct {


// Port to use for TCP health checks.
Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}


type TCPOptionsParameters struct {


// Port to use for TCP health checks.
// +kubebuilder:validation:Optional
Port *float64 `json:"port" tf:"port,omitempty"`
}

// NetworkLoadBalancerSpec defines the desired state of NetworkLoadBalancer
type NetworkLoadBalancerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       NetworkLoadBalancerParameters `json:"forProvider"`
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
	InitProvider       NetworkLoadBalancerInitParameters `json:"initProvider,omitempty"`
}

// NetworkLoadBalancerStatus defines the observed state of NetworkLoadBalancer.
type NetworkLoadBalancerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          NetworkLoadBalancerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion


// NetworkLoadBalancer is the Schema for the NetworkLoadBalancers API. A network load balancer is used to evenly distribute the load across cloud resources.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type NetworkLoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkLoadBalancerSpec   `json:"spec"`
	Status            NetworkLoadBalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkLoadBalancerList contains a list of NetworkLoadBalancers
type NetworkLoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkLoadBalancer `json:"items"`
}

// Repository type metadata.
var (
	NetworkLoadBalancer_Kind             = "NetworkLoadBalancer"
	NetworkLoadBalancer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkLoadBalancer_Kind}.String()
	NetworkLoadBalancer_KindAPIVersion   = NetworkLoadBalancer_Kind + "." + CRDGroupVersion.String()
	NetworkLoadBalancer_GroupVersionKind = CRDGroupVersion.WithKind(NetworkLoadBalancer_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkLoadBalancer{}, &NetworkLoadBalancerList{})
}
