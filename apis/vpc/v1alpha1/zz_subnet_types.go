// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type DHCPOptionsInitParameters struct {

	// Domain name.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Domain name server IP addresses.
	DomainNameServers []*string `json:"domainNameServers,omitempty" tf:"domain_name_servers,omitempty"`

	// NTP server IP addresses.
	NtpServers []*string `json:"ntpServers,omitempty" tf:"ntp_servers,omitempty"`
}

type DHCPOptionsObservation struct {

	// Domain name.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Domain name server IP addresses.
	DomainNameServers []*string `json:"domainNameServers,omitempty" tf:"domain_name_servers,omitempty"`

	// NTP server IP addresses.
	NtpServers []*string `json:"ntpServers,omitempty" tf:"ntp_servers,omitempty"`
}

type DHCPOptionsParameters struct {

	// Domain name.
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Domain name server IP addresses.
	// +kubebuilder:validation:Optional
	DomainNameServers []*string `json:"domainNameServers,omitempty" tf:"domain_name_servers,omitempty"`

	// NTP server IP addresses.
	// +kubebuilder:validation:Optional
	NtpServers []*string `json:"ntpServers,omitempty" tf:"ntp_servers,omitempty"`
}

type SubnetInitParameters struct {

	// Options for DHCP client. The structure is documented below.
	DHCPOptions []DHCPOptionsInitParameters `json:"dhcpOptions,omitempty" tf:"dhcp_options,omitempty"`

	// An optional description of the subnet. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Labels to assign to this subnet. A list of key/value pairs.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the subnet. Provided by the client when the subnet is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the route table to assign to this subnet. Assigned route table should
	// belong to the same network as this subnet.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// A list of blocks of internal IPv4 addresses that are owned by this subnet.
	// Provide this property when you create the subnet. For example, 10.0.0.0/22 or 192.168.0.0/16.
	// Blocks of addresses must be unique and non-overlapping within a network.
	// Minimum subnet size is /28, and maximum subnet size is /16. Only IPv4 is supported.
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// Name of the Yandex.Cloud zone for this subnet.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type SubnetObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Options for DHCP client. The structure is documented below.
	DHCPOptions []DHCPOptionsObservation `json:"dhcpOptions,omitempty" tf:"dhcp_options,omitempty"`

	// An optional description of the subnet. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder to which the resource belongs.
	// If omitted, the provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Labels to assign to this subnet. A list of key/value pairs.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the subnet. Provided by the client when the subnet is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnets.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// The ID of the route table to assign to this subnet. Assigned route table should
	// belong to the same network as this subnet.
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// A list of blocks of internal IPv4 addresses that are owned by this subnet.
	// Provide this property when you create the subnet. For example, 10.0.0.0/22 or 192.168.0.0/16.
	// Blocks of addresses must be unique and non-overlapping within a network.
	// Minimum subnet size is /28, and maximum subnet size is /16. Only IPv4 is supported.
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// An optional list of blocks of IPv6 addresses that are owned by this subnet.
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`

	// Name of the Yandex.Cloud zone for this subnet.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type SubnetParameters struct {

	// Options for DHCP client. The structure is documented below.
	// +kubebuilder:validation:Optional
	DHCPOptions []DHCPOptionsParameters `json:"dhcpOptions,omitempty" tf:"dhcp_options,omitempty"`

	// An optional description of the subnet. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder to which the resource belongs.
	// If omitted, the provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// Labels to assign to this subnet. A list of key/value pairs.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the subnet. Provided by the client when the subnet is created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnets.
	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// The ID of the route table to assign to this subnet. Assigned route table should
	// belong to the same network as this subnet.
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// A list of blocks of internal IPv4 addresses that are owned by this subnet.
	// Provide this property when you create the subnet. For example, 10.0.0.0/22 or 192.168.0.0/16.
	// Blocks of addresses must be unique and non-overlapping within a network.
	// Minimum subnet size is /28, and maximum subnet size is /16. Only IPv4 is supported.
	// +kubebuilder:validation:Optional
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// Name of the Yandex.Cloud zone for this subnet.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// SubnetSpec defines the desired state of Subnet
type SubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetParameters `json:"forProvider"`
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
	InitProvider SubnetInitParameters `json:"initProvider,omitempty"`
}

// SubnetStatus defines the observed state of Subnet.
type SubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subnet is the Schema for the Subnets API. A VPC network is a virtual version of the traditional physical networks that exist within and between physical data centers.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.v4CidrBlocks) || (has(self.initProvider) && has(self.initProvider.v4CidrBlocks))",message="spec.forProvider.v4CidrBlocks is a required parameter"
	Spec   SubnetSpec   `json:"spec"`
	Status SubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetList contains a list of Subnets
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

// Repository type metadata.
var (
	Subnet_Kind             = "Subnet"
	Subnet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subnet_Kind}.String()
	Subnet_KindAPIVersion   = Subnet_Kind + "." + CRDGroupVersion.String()
	Subnet_GroupVersionKind = CRDGroupVersion.WithKind(Subnet_Kind)
)

func init() {
	SchemeBuilder.Register(&Subnet{}, &SubnetList{})
}
