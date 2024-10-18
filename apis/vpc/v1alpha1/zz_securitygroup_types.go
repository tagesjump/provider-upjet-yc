// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecurityGroupEgressInitParameters struct {

	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Minimum port number.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Labels to assign to this security group.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets. self_security_group refers to this particular security group. loadbalancer_healthchecks represents loadbalancer health check nodes.
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Target security group ID for this rule.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Maximum port number.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type SecurityGroupEgressObservation struct {

	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Minimum port number.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Id of the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Labels to assign to this security group.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets. self_security_group refers to this particular security group. loadbalancer_healthchecks represents loadbalancer health check nodes.
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Target security group ID for this rule.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Maximum port number.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type SecurityGroupEgressParameters struct {

	// Description of the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Minimum port number.
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Labels to assign to this security group.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets. self_security_group refers to this particular security group. loadbalancer_healthchecks represents loadbalancer health check nodes.
	// +kubebuilder:validation:Optional
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Target security group ID for this rule.
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Maximum port number.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	// +kubebuilder:validation:Optional
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	// +kubebuilder:validation:Optional
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type SecurityGroupIngressInitParameters struct {

	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Minimum port number.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Labels to assign to this rule.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets. self_security_group refers to this particular security group. loadbalancer_healthchecks represents loadbalancer health check nodes.
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Target security group ID for this rule.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Maximum port number.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type SecurityGroupIngressObservation struct {

	// Description of the rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Minimum port number.
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Id of the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Labels to assign to this rule.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets. self_security_group refers to this particular security group. loadbalancer_healthchecks represents loadbalancer health check nodes.
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Target security group ID for this rule.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Maximum port number.
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type SecurityGroupIngressParameters struct {

	// Description of the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Minimum port number.
	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Labels to assign to this rule.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Port number (if applied to a single port).
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Special-purpose targets. self_security_group refers to this particular security group. loadbalancer_healthchecks represents loadbalancer health check nodes.
	// +kubebuilder:validation:Optional
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// One of ANY, TCP, UDP, ICMP, IPV6_ICMP.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Target security group ID for this rule.
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Maximum port number.
	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// The blocks of IPv4 addresses for this rule.
	// +kubebuilder:validation:Optional
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// The blocks of IPv6 addresses for this rule. v6_cidr_blocks argument is currently not supported. It will be available in the future.
	// +kubebuilder:validation:Optional
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type SecurityGroupInitParameters struct {

	// Description of the security group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of egress rules. The structure is documented below.
	Egress []SecurityGroupEgressInitParameters `json:"egress,omitempty" tf:"egress,omitempty"`

	// ID of the folder this security group belongs to.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A list of ingress rules.
	Ingress []SecurityGroupIngressInitParameters `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// Labels to assign to this security group.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the security group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network this security group belongs to.
	// +crossplane:generate:reference:type=Network
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`
}

type SecurityGroupObservation struct {

	// Creation timestamp of this security group.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Description of the security group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of egress rules. The structure is documented below.
	Egress []SecurityGroupEgressObservation `json:"egress,omitempty" tf:"egress,omitempty"`

	// ID of the folder this security group belongs to.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Id of the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of ingress rules.
	Ingress []SecurityGroupIngressObservation `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// Labels to assign to this security group.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the security group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network this security group belongs to.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Status of this security group.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SecurityGroupParameters struct {

	// Description of the security group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of egress rules. The structure is documented below.
	// +kubebuilder:validation:Optional
	Egress []SecurityGroupEgressParameters `json:"egress,omitempty" tf:"egress,omitempty"`

	// ID of the folder this security group belongs to.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A list of ingress rules.
	// +kubebuilder:validation:Optional
	Ingress []SecurityGroupIngressParameters `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// Labels to assign to this security group.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the security group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network this security group belongs to.
	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`
}

// SecurityGroupSpec defines the desired state of SecurityGroup
type SecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupParameters `json:"forProvider"`
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
	InitProvider SecurityGroupInitParameters `json:"initProvider,omitempty"`
}

// SecurityGroupStatus defines the observed state of SecurityGroup.
type SecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecurityGroup is the Schema for the SecurityGroups API. Yandex VPC Security Group.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupSpec   `json:"spec"`
	Status            SecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupList contains a list of SecurityGroups
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroup_Kind             = "SecurityGroup"
	SecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroup_Kind}.String()
	SecurityGroup_KindAPIVersion   = SecurityGroup_Kind + "." + CRDGroupVersion.String()
	SecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}
