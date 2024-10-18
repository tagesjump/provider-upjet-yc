// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AgentInitParameters struct {

	// The template for creating new compute instance running load testing agent. The structure is documented below.
	ComputeInstance []ComputeInstanceInitParameters `json:"computeInstance,omitempty" tf:"compute_instance,omitempty"`

	// A description of the load testing agent.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder that the resources belong to.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the agent.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the load testing agent. Must be unique within folder.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AgentObservation struct {

	// The template for creating new compute instance running load testing agent. The structure is documented below.
	ComputeInstance []ComputeInstanceObservation `json:"computeInstance,omitempty" tf:"compute_instance,omitempty"`

	ComputeInstanceID *string `json:"computeInstanceId,omitempty" tf:"compute_instance_id,omitempty"`

	// A description of the load testing agent.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder that the resources belong to.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the agent.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the load testing agent. Must be unique within folder.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AgentParameters struct {

	// The template for creating new compute instance running load testing agent. The structure is documented below.
	// +kubebuilder:validation:Optional
	ComputeInstance []ComputeInstanceParameters `json:"computeInstance,omitempty" tf:"compute_instance,omitempty"`

	// A description of the load testing agent.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder that the resources belong to.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the agent.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the load testing agent. Must be unique within folder.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type BootDiskInitParameters struct {

	// Whether the disk is auto-deleted when the instance is deleted. The default value is true.
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// This value can be used to reference the device under /dev/disk/by-id/.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Parameters for creating a disk alongside the instance. The structure is documented below.
	InitializeParams []InitializeParamsInitParameters `json:"initializeParams,omitempty" tf:"initialize_params,omitempty"`
}

type BootDiskObservation struct {

	// Whether the disk is auto-deleted when the instance is deleted. The default value is true.
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// This value can be used to reference the device under /dev/disk/by-id/.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// (Computed) The ID of created disk.
	DiskID *string `json:"diskId,omitempty" tf:"disk_id,omitempty"`

	// Parameters for creating a disk alongside the instance. The structure is documented below.
	InitializeParams []InitializeParamsObservation `json:"initializeParams,omitempty" tf:"initialize_params,omitempty"`
}

type BootDiskParameters struct {

	// Whether the disk is auto-deleted when the instance is deleted. The default value is true.
	// +kubebuilder:validation:Optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// This value can be used to reference the device under /dev/disk/by-id/.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Parameters for creating a disk alongside the instance. The structure is documented below.
	// +kubebuilder:validation:Optional
	InitializeParams []InitializeParamsParameters `json:"initializeParams" tf:"initialize_params,omitempty"`
}

type ComputeInstanceInitParameters struct {

	// Boot disk specifications for the instance. The structure is documented below.
	BootDisk []BootDiskInitParameters `json:"bootDisk,omitempty" tf:"boot_disk,omitempty"`

	// A set of key/value label pairs to assign to the instance.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A set of metadata key/value pairs to make available from within the instance.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Network specifications for the instance. This can be used multiple times for adding multiple interfaces. The structure is documented below.
	NetworkInterface []NetworkInterfaceInitParameters `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// The Compute platform of virtual machine. If it is not provided, the standard-v2 platform will be used.
	PlatformID *string `json:"platformId,omitempty" tf:"platform_id,omitempty"`

	// Compute resource specifications for the instance. The structure is documented below.
	Resources []ResourcesInitParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// The ID of the service account authorized for this load testing agent. Service account should have loadtesting.generatorClient or loadtesting.externalAgent role in the folder.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.SecurityGroup
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// Reference to a SecurityGroup in vpc to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in vpc to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`

	// The availability zone where the virtual machine will be created. If it is not provided, the default provider folder is used.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ComputeInstanceObservation struct {

	// Boot disk specifications for the instance. The structure is documented below.
	BootDisk []BootDiskObservation `json:"bootDisk,omitempty" tf:"boot_disk,omitempty"`

	// (Computed) The set of labels key:value pairs assigned to this instance. This includes user custom labels and predefined items created by Yandex Cloud Load Testing.
	// +mapType=granular
	ComputedLabels map[string]*string `json:"computedLabels,omitempty" tf:"computed_labels,omitempty"`

	// (Computed) The set of metadata key:value pairs assigned to this instance. This includes user custom metadata, and predefined items created by Yandex Cloud Load Testing.
	// +mapType=granular
	ComputedMetadata map[string]*string `json:"computedMetadata,omitempty" tf:"computed_metadata,omitempty"`

	// A set of key/value label pairs to assign to the instance.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A set of metadata key/value pairs to make available from within the instance.
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Network specifications for the instance. This can be used multiple times for adding multiple interfaces. The structure is documented below.
	NetworkInterface []NetworkInterfaceObservation `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// The Compute platform of virtual machine. If it is not provided, the standard-v2 platform will be used.
	PlatformID *string `json:"platformId,omitempty" tf:"platform_id,omitempty"`

	// Compute resource specifications for the instance. The structure is documented below.
	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// The ID of the service account authorized for this load testing agent. Service account should have loadtesting.generatorClient or loadtesting.externalAgent role in the folder.
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// The availability zone where the virtual machine will be created. If it is not provided, the default provider folder is used.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type ComputeInstanceParameters struct {

	// Boot disk specifications for the instance. The structure is documented below.
	// +kubebuilder:validation:Optional
	BootDisk []BootDiskParameters `json:"bootDisk" tf:"boot_disk,omitempty"`

	// A set of key/value label pairs to assign to the instance.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A set of metadata key/value pairs to make available from within the instance.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Network specifications for the instance. This can be used multiple times for adding multiple interfaces. The structure is documented below.
	// +kubebuilder:validation:Optional
	NetworkInterface []NetworkInterfaceParameters `json:"networkInterface" tf:"network_interface,omitempty"`

	// The Compute platform of virtual machine. If it is not provided, the standard-v2 platform will be used.
	// +kubebuilder:validation:Optional
	PlatformID *string `json:"platformId,omitempty" tf:"platform_id,omitempty"`

	// Compute resource specifications for the instance. The structure is documented below.
	// +kubebuilder:validation:Optional
	Resources []ResourcesParameters `json:"resources" tf:"resources,omitempty"`

	// The ID of the service account authorized for this load testing agent. Service account should have loadtesting.generatorClient or loadtesting.externalAgent role in the folder.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// Reference to a SecurityGroup in vpc to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in vpc to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`

	// The availability zone where the virtual machine will be created. If it is not provided, the default provider folder is used.
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type InitializeParamsInitParameters struct {

	// Block size of the disk, specified in bytes.
	BlockSize *float64 `json:"blockSize,omitempty" tf:"block_size,omitempty"`

	// A description of the boot disk.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the load testing agent. Must be unique within folder.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The size of the disk in GB. Defaults to 15 GB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The disk type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InitializeParamsObservation struct {

	// Block size of the disk, specified in bytes.
	BlockSize *float64 `json:"blockSize,omitempty" tf:"block_size,omitempty"`

	// A description of the boot disk.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the load testing agent. Must be unique within folder.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The size of the disk in GB. Defaults to 15 GB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The disk type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InitializeParamsParameters struct {

	// Block size of the disk, specified in bytes.
	// +kubebuilder:validation:Optional
	BlockSize *float64 `json:"blockSize,omitempty" tf:"block_size,omitempty"`

	// A description of the boot disk.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the load testing agent. Must be unique within folder.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The size of the disk in GB. Defaults to 15 GB.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The disk type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkInterfaceInitParameters struct {

	// Manual set static IP address.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Flag for allocating IPv4 address for the network interface.
	IPv4 *bool `json:"ipv4,omitempty" tf:"ipv4,omitempty"`

	// Flag for allocating IPv6 address for the network interface.
	IPv6 *bool `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// Manual set static IPv6 address.
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// Flag for using NAT.
	NAT *bool `json:"nat,omitempty" tf:"nat,omitempty"`

	// A public address that can be used to access the internet over NAT.
	NATIPAddress *string `json:"natIpAddress,omitempty" tf:"nat_ip_address,omitempty"`

	// Security group ids for network interface.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The ID of the subnet to attach this interface to. The subnet must reside in the same zone where this instance was created.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type NetworkInterfaceObservation struct {

	// Manual set static IP address.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Flag for allocating IPv4 address for the network interface.
	IPv4 *bool `json:"ipv4,omitempty" tf:"ipv4,omitempty"`

	// Flag for allocating IPv6 address for the network interface.
	IPv6 *bool `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// Manual set static IPv6 address.
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	Index *float64 `json:"index,omitempty" tf:"index,omitempty"`

	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// Flag for using NAT.
	NAT *bool `json:"nat,omitempty" tf:"nat,omitempty"`

	// A public address that can be used to access the internet over NAT.
	NATIPAddress *string `json:"natIpAddress,omitempty" tf:"nat_ip_address,omitempty"`

	NATIPVersion *string `json:"natIpVersion,omitempty" tf:"nat_ip_version,omitempty"`

	// Security group ids for network interface.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The ID of the subnet to attach this interface to. The subnet must reside in the same zone where this instance was created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type NetworkInterfaceParameters struct {

	// Manual set static IP address.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Flag for allocating IPv4 address for the network interface.
	// +kubebuilder:validation:Optional
	IPv4 *bool `json:"ipv4,omitempty" tf:"ipv4,omitempty"`

	// Flag for allocating IPv6 address for the network interface.
	// +kubebuilder:validation:Optional
	IPv6 *bool `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// Manual set static IPv6 address.
	// +kubebuilder:validation:Optional
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// Flag for using NAT.
	// +kubebuilder:validation:Optional
	NAT *bool `json:"nat,omitempty" tf:"nat,omitempty"`

	// A public address that can be used to access the internet over NAT.
	// +kubebuilder:validation:Optional
	NATIPAddress *string `json:"natIpAddress,omitempty" tf:"nat_ip_address,omitempty"`

	// Security group ids for network interface.
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The ID of the subnet to attach this interface to. The subnet must reside in the same zone where this instance was created.
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

type ResourcesInitParameters struct {

	// If provided, specifies baseline core performance as a percent.
	CoreFraction *float64 `json:"coreFraction,omitempty" tf:"core_fraction,omitempty"`

	// The number of CPU cores for the instance. Defaults to 2 cores.
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// The memory size in GB. Defaults to 2 GB.
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`
}

type ResourcesObservation struct {

	// If provided, specifies baseline core performance as a percent.
	CoreFraction *float64 `json:"coreFraction,omitempty" tf:"core_fraction,omitempty"`

	// The number of CPU cores for the instance. Defaults to 2 cores.
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// The memory size in GB. Defaults to 2 GB.
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`
}

type ResourcesParameters struct {

	// If provided, specifies baseline core performance as a percent.
	// +kubebuilder:validation:Optional
	CoreFraction *float64 `json:"coreFraction,omitempty" tf:"core_fraction,omitempty"`

	// The number of CPU cores for the instance. Defaults to 2 cores.
	// +kubebuilder:validation:Optional
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// The memory size in GB. Defaults to 2 GB.
	// +kubebuilder:validation:Optional
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`
}

// AgentSpec defines the desired state of Agent
type AgentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AgentParameters `json:"forProvider"`
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
	InitProvider AgentInitParameters `json:"initProvider,omitempty"`
}

// AgentStatus defines the observed state of Agent.
type AgentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AgentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Agent is the Schema for the Agents API. Manages an Yandex Cloud Load Testing Agent resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Agent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.computeInstance) || (has(self.initProvider) && has(self.initProvider.computeInstance))",message="spec.forProvider.computeInstance is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AgentSpec   `json:"spec"`
	Status AgentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgentList contains a list of Agents
type AgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Agent `json:"items"`
}

// Repository type metadata.
var (
	Agent_Kind             = "Agent"
	Agent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Agent_Kind}.String()
	Agent_KindAPIVersion   = Agent_Kind + "." + CRDGroupVersion.String()
	Agent_GroupVersionKind = CRDGroupVersion.WithKind(Agent_Kind)
)

func init() {
	SchemeBuilder.Register(&Agent{}, &AgentList{})
}
