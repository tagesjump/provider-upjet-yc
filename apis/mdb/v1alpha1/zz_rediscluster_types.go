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

type RedisClusterConfigInitParameters struct {

	// Normal clients output buffer limits.
	// See redis config file.
	ClientOutputBufferLimitNormal *string `json:"clientOutputBufferLimitNormal,omitempty" tf:"client_output_buffer_limit_normal,omitempty"`

	// Pubsub clients output buffer limits.
	// See redis config file.
	ClientOutputBufferLimitPubsub *string `json:"clientOutputBufferLimitPubsub,omitempty" tf:"client_output_buffer_limit_pubsub,omitempty"`

	// Number of databases (changing requires redis-server restart).
	Databases *float64 `json:"databases,omitempty" tf:"databases,omitempty"`

	// Redis maxmemory usage in percent
	MaxmemoryPercent *float64 `json:"maxmemoryPercent,omitempty" tf:"maxmemory_percent,omitempty"`

	// Redis key eviction policy for a dataset that reaches maximum memory.
	// Can be any of the listed in the official RedisDB documentation.
	MaxmemoryPolicy *string `json:"maxmemoryPolicy,omitempty" tf:"maxmemory_policy,omitempty"`

	// Select the events that Redis will notify among a set of classes.
	NotifyKeyspaceEvents *string `json:"notifyKeyspaceEvents,omitempty" tf:"notify_keyspace_events,omitempty"`

	// Log slow queries below this number in microseconds.
	SlowlogLogSlowerThan *float64 `json:"slowlogLogSlowerThan,omitempty" tf:"slowlog_log_slower_than,omitempty"`

	// Slow queries log length.
	SlowlogMaxLen *float64 `json:"slowlogMaxLen,omitempty" tf:"slowlog_max_len,omitempty"`

	// Close the connection after a client is idle for N seconds.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Version of Redis (6.2).
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type RedisClusterConfigObservation struct {

	// Normal clients output buffer limits.
	// See redis config file.
	ClientOutputBufferLimitNormal *string `json:"clientOutputBufferLimitNormal,omitempty" tf:"client_output_buffer_limit_normal,omitempty"`

	// Pubsub clients output buffer limits.
	// See redis config file.
	ClientOutputBufferLimitPubsub *string `json:"clientOutputBufferLimitPubsub,omitempty" tf:"client_output_buffer_limit_pubsub,omitempty"`

	// Number of databases (changing requires redis-server restart).
	Databases *float64 `json:"databases,omitempty" tf:"databases,omitempty"`

	// Redis maxmemory usage in percent
	MaxmemoryPercent *float64 `json:"maxmemoryPercent,omitempty" tf:"maxmemory_percent,omitempty"`

	// Redis key eviction policy for a dataset that reaches maximum memory.
	// Can be any of the listed in the official RedisDB documentation.
	MaxmemoryPolicy *string `json:"maxmemoryPolicy,omitempty" tf:"maxmemory_policy,omitempty"`

	// Select the events that Redis will notify among a set of classes.
	NotifyKeyspaceEvents *string `json:"notifyKeyspaceEvents,omitempty" tf:"notify_keyspace_events,omitempty"`

	// Log slow queries below this number in microseconds.
	SlowlogLogSlowerThan *float64 `json:"slowlogLogSlowerThan,omitempty" tf:"slowlog_log_slower_than,omitempty"`

	// Slow queries log length.
	SlowlogMaxLen *float64 `json:"slowlogMaxLen,omitempty" tf:"slowlog_max_len,omitempty"`

	// Close the connection after a client is idle for N seconds.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Version of Redis (6.2).
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type RedisClusterConfigParameters struct {

	// Normal clients output buffer limits.
	// See redis config file.
	// +kubebuilder:validation:Optional
	ClientOutputBufferLimitNormal *string `json:"clientOutputBufferLimitNormal,omitempty" tf:"client_output_buffer_limit_normal,omitempty"`

	// Pubsub clients output buffer limits.
	// See redis config file.
	// +kubebuilder:validation:Optional
	ClientOutputBufferLimitPubsub *string `json:"clientOutputBufferLimitPubsub,omitempty" tf:"client_output_buffer_limit_pubsub,omitempty"`

	// Number of databases (changing requires redis-server restart).
	// +kubebuilder:validation:Optional
	Databases *float64 `json:"databases,omitempty" tf:"databases,omitempty"`

	// Redis maxmemory usage in percent
	// +kubebuilder:validation:Optional
	MaxmemoryPercent *float64 `json:"maxmemoryPercent,omitempty" tf:"maxmemory_percent,omitempty"`

	// Redis key eviction policy for a dataset that reaches maximum memory.
	// Can be any of the listed in the official RedisDB documentation.
	// +kubebuilder:validation:Optional
	MaxmemoryPolicy *string `json:"maxmemoryPolicy,omitempty" tf:"maxmemory_policy,omitempty"`

	// Select the events that Redis will notify among a set of classes.
	// +kubebuilder:validation:Optional
	NotifyKeyspaceEvents *string `json:"notifyKeyspaceEvents,omitempty" tf:"notify_keyspace_events,omitempty"`

	// Password for the Redis cluster.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Log slow queries below this number in microseconds.
	// +kubebuilder:validation:Optional
	SlowlogLogSlowerThan *float64 `json:"slowlogLogSlowerThan,omitempty" tf:"slowlog_log_slower_than,omitempty"`

	// Slow queries log length.
	// +kubebuilder:validation:Optional
	SlowlogMaxLen *float64 `json:"slowlogMaxLen,omitempty" tf:"slowlog_max_len,omitempty"`

	// Close the connection after a client is idle for N seconds.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Version of Redis (6.2).
	// +kubebuilder:validation:Optional
	Version *string `json:"version" tf:"version,omitempty"`
}

type RedisClusterHostInitParameters struct {

	// Sets whether the host should get a public IP address or not.
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// Replica priority of a current replica (usable for non-sharded only).
	ReplicaPriority *float64 `json:"replicaPriority,omitempty" tf:"replica_priority,omitempty"`

	// The name of the shard to which the host belongs.
	ShardName *string `json:"shardName,omitempty" tf:"shard_name,omitempty"`

	// The ID of the subnet, to which the host belongs. The subnet must
	// be a part of the network to which the cluster belongs.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The availability zone where the Redis host will be created.
	// For more information see the official documentation.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type RedisClusterHostObservation struct {

	// Sets whether the host should get a public IP address or not.
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// The fully qualified domain name of the host.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// Replica priority of a current replica (usable for non-sharded only).
	ReplicaPriority *float64 `json:"replicaPriority,omitempty" tf:"replica_priority,omitempty"`

	// The name of the shard to which the host belongs.
	ShardName *string `json:"shardName,omitempty" tf:"shard_name,omitempty"`

	// The ID of the subnet, to which the host belongs. The subnet must
	// be a part of the network to which the cluster belongs.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The availability zone where the Redis host will be created.
	// For more information see the official documentation.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type RedisClusterHostParameters struct {

	// Sets whether the host should get a public IP address or not.
	// +kubebuilder:validation:Optional
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// Replica priority of a current replica (usable for non-sharded only).
	// +kubebuilder:validation:Optional
	ReplicaPriority *float64 `json:"replicaPriority,omitempty" tf:"replica_priority,omitempty"`

	// The name of the shard to which the host belongs.
	// +kubebuilder:validation:Optional
	ShardName *string `json:"shardName,omitempty" tf:"shard_name,omitempty"`

	// The ID of the subnet, to which the host belongs. The subnet must
	// be a part of the network to which the cluster belongs.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The availability zone where the Redis host will be created.
	// For more information see the official documentation.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type RedisClusterInitParameters struct {

	// Announce fqdn instead of ip address.
	AnnounceHostnames *bool `json:"announceHostnames,omitempty" tf:"announce_hostnames,omitempty"`

	// Configuration of the Redis cluster. The structure is documented below.
	Config []RedisClusterConfigInitParameters `json:"config,omitempty" tf:"config,omitempty"`

	// Inhibits deletion of the cluster.  Can be either true or false.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Description of the Redis cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Deployment environment of the Redis cluster. Can be either PRESTABLE or PRODUCTION.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A host of the Redis cluster. The structure is documented below.
	Host []RedisClusterHostInitParameters `json:"host,omitempty" tf:"host,omitempty"`

	// A set of key/value label pairs to assign to the Redis cluster.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	MaintenanceWindow []RedisClusterMaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// Name of the Redis cluster. Provided by the client when the cluster is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network, to which the Redis cluster belongs.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Network
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Persistence mode.
	PersistenceMode *string `json:"persistenceMode,omitempty" tf:"persistence_mode,omitempty"`

	// Resources allocated to hosts of the Redis cluster. The structure is documented below.
	Resources []RedisClusterResourcesInitParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// A set of ids of security groups assigned to hosts of the cluster.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// Redis Cluster mode enabled/disabled.
	Sharded *bool `json:"sharded,omitempty" tf:"sharded,omitempty"`

	// TLS support mode enabled/disabled.
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`
}

type RedisClusterMaintenanceWindowInitParameters struct {

	// Day of week for maintenance window if window type is weekly. Possible values: MON, TUE, WED, THU, FRI, SAT, SUN.
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Hour of day in UTC time zone (1-24) for maintenance window if window type is weekly.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Type of maintenance window. Can be either ANYTIME or WEEKLY. A day and hour of window need to be specified with weekly window.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RedisClusterMaintenanceWindowObservation struct {

	// Day of week for maintenance window if window type is weekly. Possible values: MON, TUE, WED, THU, FRI, SAT, SUN.
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Hour of day in UTC time zone (1-24) for maintenance window if window type is weekly.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Type of maintenance window. Can be either ANYTIME or WEEKLY. A day and hour of window need to be specified with weekly window.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RedisClusterMaintenanceWindowParameters struct {

	// Day of week for maintenance window if window type is weekly. Possible values: MON, TUE, WED, THU, FRI, SAT, SUN.
	// +kubebuilder:validation:Optional
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Hour of day in UTC time zone (1-24) for maintenance window if window type is weekly.
	// +kubebuilder:validation:Optional
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Type of maintenance window. Can be either ANYTIME or WEEKLY. A day and hour of window need to be specified with weekly window.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type RedisClusterObservation struct {

	// Announce fqdn instead of ip address.
	AnnounceHostnames *bool `json:"announceHostnames,omitempty" tf:"announce_hostnames,omitempty"`

	// Configuration of the Redis cluster. The structure is documented below.
	Config []RedisClusterConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	// Creation timestamp of the key.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Inhibits deletion of the cluster.  Can be either true or false.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Description of the Redis cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Deployment environment of the Redis cluster. Can be either PRESTABLE or PRODUCTION.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Aggregated health of the cluster. Can be either ALIVE, DEGRADED, DEAD or HEALTH_UNKNOWN.
	// For more information see health field of JSON representation in the official documentation.
	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	// A host of the Redis cluster. The structure is documented below.
	Host []RedisClusterHostObservation `json:"host,omitempty" tf:"host,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the Redis cluster.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	MaintenanceWindow []RedisClusterMaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// Name of the Redis cluster. Provided by the client when the cluster is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network, to which the Redis cluster belongs.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Persistence mode.
	PersistenceMode *string `json:"persistenceMode,omitempty" tf:"persistence_mode,omitempty"`

	// Resources allocated to hosts of the Redis cluster. The structure is documented below.
	Resources []RedisClusterResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// A set of ids of security groups assigned to hosts of the cluster.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Redis Cluster mode enabled/disabled.
	Sharded *bool `json:"sharded,omitempty" tf:"sharded,omitempty"`

	// Status of the cluster. Can be either CREATING, STARTING, RUNNING, UPDATING, STOPPING, STOPPED, ERROR or STATUS_UNKNOWN.
	// For more information see status field of JSON representation in the official documentation.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// TLS support mode enabled/disabled.
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`
}

type RedisClusterParameters struct {

	// Announce fqdn instead of ip address.
	// +kubebuilder:validation:Optional
	AnnounceHostnames *bool `json:"announceHostnames,omitempty" tf:"announce_hostnames,omitempty"`

	// Configuration of the Redis cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	Config []RedisClusterConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// Inhibits deletion of the cluster.  Can be either true or false.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Description of the Redis cluster.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Deployment environment of the Redis cluster. Can be either PRESTABLE or PRODUCTION.
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A host of the Redis cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	Host []RedisClusterHostParameters `json:"host,omitempty" tf:"host,omitempty"`

	// A set of key/value label pairs to assign to the Redis cluster.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow []RedisClusterMaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// Name of the Redis cluster. Provided by the client when the cluster is created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network, to which the Redis cluster belongs.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Persistence mode.
	// +kubebuilder:validation:Optional
	PersistenceMode *string `json:"persistenceMode,omitempty" tf:"persistence_mode,omitempty"`

	// Resources allocated to hosts of the Redis cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	Resources []RedisClusterResourcesParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// A set of ids of security groups assigned to hosts of the cluster.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// Redis Cluster mode enabled/disabled.
	// +kubebuilder:validation:Optional
	Sharded *bool `json:"sharded,omitempty" tf:"sharded,omitempty"`

	// TLS support mode enabled/disabled.
	// +kubebuilder:validation:Optional
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`
}

type RedisClusterResourcesInitParameters struct {

	// Volume of the storage available to a host, in gigabytes.
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Type of the storage of Redis hosts - environment default is used if missing.
	DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

	ResourcePresetID *string `json:"resourcePresetId,omitempty" tf:"resource_preset_id,omitempty"`
}

type RedisClusterResourcesObservation struct {

	// Volume of the storage available to a host, in gigabytes.
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Type of the storage of Redis hosts - environment default is used if missing.
	DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

	ResourcePresetID *string `json:"resourcePresetId,omitempty" tf:"resource_preset_id,omitempty"`
}

type RedisClusterResourcesParameters struct {

	// Volume of the storage available to a host, in gigabytes.
	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize" tf:"disk_size,omitempty"`

	// Type of the storage of Redis hosts - environment default is used if missing.
	// +kubebuilder:validation:Optional
	DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

// RedisClusterSpec defines the desired state of RedisCluster
type RedisClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedisClusterParameters `json:"forProvider"`
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
	InitProvider RedisClusterInitParameters `json:"initProvider,omitempty"`
}

// RedisClusterStatus defines the observed state of RedisCluster.
type RedisClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedisClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RedisCluster is the Schema for the RedisClusters API. Manages a Redis cluster within Yandex.Cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type RedisCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.config) || (has(self.initProvider) && has(self.initProvider.config))",message="spec.forProvider.config is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.environment) || (has(self.initProvider) && has(self.initProvider.environment))",message="spec.forProvider.environment is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.host) || (has(self.initProvider) && has(self.initProvider.host))",message="spec.forProvider.host is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resources) || (has(self.initProvider) && has(self.initProvider.resources))",message="spec.forProvider.resources is a required parameter"
	Spec   RedisClusterSpec   `json:"spec"`
	Status RedisClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisClusterList contains a list of RedisClusters
type RedisClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisCluster `json:"items"`
}

// Repository type metadata.
var (
	RedisCluster_Kind             = "RedisCluster"
	RedisCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedisCluster_Kind}.String()
	RedisCluster_KindAPIVersion   = RedisCluster_Kind + "." + CRDGroupVersion.String()
	RedisCluster_GroupVersionKind = CRDGroupVersion.WithKind(RedisCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&RedisCluster{}, &RedisClusterList{})
}
