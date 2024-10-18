

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

)




type DataNodeInitParameters struct {


// Resources allocated to hosts of the Elasticsearch master nodes subcluster. The structure is documented below.
Resources []DataNodeResourcesInitParameters `json:"resources,omitempty" tf:"resources,omitempty"`
}


type DataNodeObservation struct {


// Resources allocated to hosts of the Elasticsearch master nodes subcluster. The structure is documented below.
Resources []DataNodeResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`
}


type DataNodeParameters struct {


// Resources allocated to hosts of the Elasticsearch master nodes subcluster. The structure is documented below.
// +kubebuilder:validation:Optional
Resources []DataNodeResourcesParameters `json:"resources" tf:"resources,omitempty"`
}


type DataNodeResourcesInitParameters struct {


// Volume of the storage available to a host, in gigabytes.
DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

// Type of the storage of Elasticsearch hosts.
DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

ResourcePresetID *string `json:"resourcePresetId,omitempty" tf:"resource_preset_id,omitempty"`
}


type DataNodeResourcesObservation struct {


// Volume of the storage available to a host, in gigabytes.
DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

// Type of the storage of Elasticsearch hosts.
DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

ResourcePresetID *string `json:"resourcePresetId,omitempty" tf:"resource_preset_id,omitempty"`
}


type DataNodeResourcesParameters struct {


// Volume of the storage available to a host, in gigabytes.
// +kubebuilder:validation:Optional
DiskSize *float64 `json:"diskSize" tf:"disk_size,omitempty"`

// Type of the storage of Elasticsearch hosts.
// +kubebuilder:validation:Optional
DiskTypeID *string `json:"diskTypeId" tf:"disk_type_id,omitempty"`

// +kubebuilder:validation:Optional
ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}


type ElasticsearchClusterConfigInitParameters struct {


// Password for admin user of Elasticsearch.
AdminPasswordSecretRef v1.SecretKeySelector `json:"adminPasswordSecretRef" tf:"-"`

// Configuration for Elasticsearch data nodes subcluster. The structure is documented below.
DataNode []DataNodeInitParameters `json:"dataNode,omitempty" tf:"data_node,omitempty"`

// Edition of Elasticsearch. For more information, see the official documentation.
Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

// Configuration for Elasticsearch master nodes subcluster. The structure is documented below.
MasterNode []MasterNodeInitParameters `json:"masterNode,omitempty" tf:"master_node,omitempty"`

// A set of Elasticsearch plugins to install.
// +listType=set
Plugins []*string `json:"plugins,omitempty" tf:"plugins,omitempty"`

// Version of Elasticsearch.
Version *string `json:"version,omitempty" tf:"version,omitempty"`
}


type ElasticsearchClusterConfigObservation struct {


// Configuration for Elasticsearch data nodes subcluster. The structure is documented below.
DataNode []DataNodeObservation `json:"dataNode,omitempty" tf:"data_node,omitempty"`

// Edition of Elasticsearch. For more information, see the official documentation.
Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

// Configuration for Elasticsearch master nodes subcluster. The structure is documented below.
MasterNode []MasterNodeObservation `json:"masterNode,omitempty" tf:"master_node,omitempty"`

// A set of Elasticsearch plugins to install.
// +listType=set
Plugins []*string `json:"plugins,omitempty" tf:"plugins,omitempty"`

// Version of Elasticsearch.
Version *string `json:"version,omitempty" tf:"version,omitempty"`
}


type ElasticsearchClusterConfigParameters struct {


// Password for admin user of Elasticsearch.
// +kubebuilder:validation:Optional
AdminPasswordSecretRef v1.SecretKeySelector `json:"adminPasswordSecretRef" tf:"-"`

// Configuration for Elasticsearch data nodes subcluster. The structure is documented below.
// +kubebuilder:validation:Optional
DataNode []DataNodeParameters `json:"dataNode" tf:"data_node,omitempty"`

// Edition of Elasticsearch. For more information, see the official documentation.
// +kubebuilder:validation:Optional
Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

// Configuration for Elasticsearch master nodes subcluster. The structure is documented below.
// +kubebuilder:validation:Optional
MasterNode []MasterNodeParameters `json:"masterNode,omitempty" tf:"master_node,omitempty"`

// A set of Elasticsearch plugins to install.
// +kubebuilder:validation:Optional
// +listType=set
Plugins []*string `json:"plugins,omitempty" tf:"plugins,omitempty"`

// Version of Elasticsearch.
// +kubebuilder:validation:Optional
Version *string `json:"version,omitempty" tf:"version,omitempty"`
}


type ElasticsearchClusterHostInitParameters struct {


// Sets whether the host should get a public IP address on creation. Can be either true or false.
AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

// User defined host name.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.
// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Subnet
SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

// Reference to a Subnet in vpc to populate subnetId.
// +kubebuilder:validation:Optional
SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

// Selector for a Subnet in vpc to populate subnetId.
// +kubebuilder:validation:Optional
SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

// The type of the host to be deployed. Can be either DATA_NODE or MASTER_NODE.
Type *string `json:"type,omitempty" tf:"type,omitempty"`

// The availability zone where the Elasticsearch host will be created. For more information see the official documentation.
Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}


type ElasticsearchClusterHostObservation struct {


// Sets whether the host should get a public IP address on creation. Can be either true or false.
AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

// The fully qualified domain name of the host.
Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

// User defined host name.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.
SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

// The type of the host to be deployed. Can be either DATA_NODE or MASTER_NODE.
Type *string `json:"type,omitempty" tf:"type,omitempty"`

// The availability zone where the Elasticsearch host will be created. For more information see the official documentation.
Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}


type ElasticsearchClusterHostParameters struct {


// Sets whether the host should get a public IP address on creation. Can be either true or false.
// +kubebuilder:validation:Optional
AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

// User defined host name.
// +kubebuilder:validation:Optional
Name *string `json:"name" tf:"name,omitempty"`

// The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.
// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Subnet
// +kubebuilder:validation:Optional
SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

// Reference to a Subnet in vpc to populate subnetId.
// +kubebuilder:validation:Optional
SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

// Selector for a Subnet in vpc to populate subnetId.
// +kubebuilder:validation:Optional
SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

// The type of the host to be deployed. Can be either DATA_NODE or MASTER_NODE.
// +kubebuilder:validation:Optional
Type *string `json:"type" tf:"type,omitempty"`

// The availability zone where the Elasticsearch host will be created. For more information see the official documentation.
// +kubebuilder:validation:Optional
Zone *string `json:"zone" tf:"zone,omitempty"`
}


type ElasticsearchClusterInitParameters struct {


// Configuration of the Elasticsearch cluster. The structure is documented below.
Config []ElasticsearchClusterConfigInitParameters `json:"config,omitempty" tf:"config,omitempty"`

// Inhibits deletion of the cluster. Can be either true or false.
DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

// Description of the Elasticsearch cluster.
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Deployment environment of the Elasticsearch cluster. Can be either PRESTABLE or PRODUCTION.
Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

// Reference to a Folder in resourcemanager to populate folderId.
// +kubebuilder:validation:Optional
FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

// Selector for a Folder in resourcemanager to populate folderId.
// +kubebuilder:validation:Optional
FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

// A host of the Elasticsearch cluster. The structure is documented below.
Host []ElasticsearchClusterHostInitParameters `json:"host,omitempty" tf:"host,omitempty"`

// A set of key/value label pairs to assign to the Elasticsearch cluster.
// +mapType=granular
Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

MaintenanceWindow []ElasticsearchClusterMaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

// Name of the Elasticsearch cluster. Provided by the client when the cluster is created.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// ID of the network, to which the Elasticsearch cluster belongs.
// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Network
NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

// Reference to a Network in vpc to populate networkId.
// +kubebuilder:validation:Optional
NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

// Selector for a Network in vpc to populate networkId.
// +kubebuilder:validation:Optional
NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

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

// ID of the service account authorized for this cluster.
// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/iam/v1alpha1.ServiceAccount
ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

// Reference to a ServiceAccount in iam to populate serviceAccountId.
// +kubebuilder:validation:Optional
ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

// Selector for a ServiceAccount in iam to populate serviceAccountId.
// +kubebuilder:validation:Optional
ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`
}


type ElasticsearchClusterMaintenanceWindowInitParameters struct {


// Day of week for maintenance window if window type is weekly. Possible values: MON, TUE, WED, THU, FRI, SAT, SUN.
Day *string `json:"day,omitempty" tf:"day,omitempty"`

// Hour of day in UTC time zone (1-24) for maintenance window if window type is weekly.
Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

// Type of maintenance window. Can be either ANYTIME or WEEKLY. A day and hour of window need to be specified with weekly window.
Type *string `json:"type,omitempty" tf:"type,omitempty"`
}


type ElasticsearchClusterMaintenanceWindowObservation struct {


// Day of week for maintenance window if window type is weekly. Possible values: MON, TUE, WED, THU, FRI, SAT, SUN.
Day *string `json:"day,omitempty" tf:"day,omitempty"`

// Hour of day in UTC time zone (1-24) for maintenance window if window type is weekly.
Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

// Type of maintenance window. Can be either ANYTIME or WEEKLY. A day and hour of window need to be specified with weekly window.
Type *string `json:"type,omitempty" tf:"type,omitempty"`
}


type ElasticsearchClusterMaintenanceWindowParameters struct {


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


type ElasticsearchClusterObservation struct {


// Configuration of the Elasticsearch cluster. The structure is documented below.
Config []ElasticsearchClusterConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

// Creation timestamp of the key.
CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

// Inhibits deletion of the cluster. Can be either true or false.
DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

// Description of the Elasticsearch cluster.
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Deployment environment of the Elasticsearch cluster. Can be either PRESTABLE or PRODUCTION.
Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

// Aggregated health of the cluster. Can be either ALIVE, DEGRADED, DEAD or HEALTH_UNKNOWN. For more information see health field of JSON representation in the official documentation.
Health *string `json:"health,omitempty" tf:"health,omitempty"`

// A host of the Elasticsearch cluster. The structure is documented below.
Host []ElasticsearchClusterHostObservation `json:"host,omitempty" tf:"host,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

// A set of key/value label pairs to assign to the Elasticsearch cluster.
// +mapType=granular
Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

MaintenanceWindow []ElasticsearchClusterMaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

// Name of the Elasticsearch cluster. Provided by the client when the cluster is created.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// ID of the network, to which the Elasticsearch cluster belongs.
NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

// A set of ids of security groups assigned to hosts of the cluster.
// +listType=set
SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

// ID of the service account authorized for this cluster.
ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

// Status of the cluster. Can be either CREATING, STARTING, RUNNING, UPDATING, STOPPING, STOPPED, ERROR or STATUS_UNKNOWN. For more information see status field of JSON representation in the official documentation.
Status *string `json:"status,omitempty" tf:"status,omitempty"`
}


type ElasticsearchClusterParameters struct {


// Configuration of the Elasticsearch cluster. The structure is documented below.
// +kubebuilder:validation:Optional
Config []ElasticsearchClusterConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

// Inhibits deletion of the cluster. Can be either true or false.
// +kubebuilder:validation:Optional
DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

// Description of the Elasticsearch cluster.
// +kubebuilder:validation:Optional
Description *string `json:"description,omitempty" tf:"description,omitempty"`

// Deployment environment of the Elasticsearch cluster. Can be either PRESTABLE or PRODUCTION.
// +kubebuilder:validation:Optional
Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

// The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
// +kubebuilder:validation:Optional
FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

// Reference to a Folder in resourcemanager to populate folderId.
// +kubebuilder:validation:Optional
FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

// Selector for a Folder in resourcemanager to populate folderId.
// +kubebuilder:validation:Optional
FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

// A host of the Elasticsearch cluster. The structure is documented below.
// +kubebuilder:validation:Optional
Host []ElasticsearchClusterHostParameters `json:"host,omitempty" tf:"host,omitempty"`

// A set of key/value label pairs to assign to the Elasticsearch cluster.
// +kubebuilder:validation:Optional
// +mapType=granular
Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

// +kubebuilder:validation:Optional
MaintenanceWindow []ElasticsearchClusterMaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

// Name of the Elasticsearch cluster. Provided by the client when the cluster is created.
// +kubebuilder:validation:Optional
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// ID of the network, to which the Elasticsearch cluster belongs.
// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1.Network
// +kubebuilder:validation:Optional
NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

// Reference to a Network in vpc to populate networkId.
// +kubebuilder:validation:Optional
NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

// Selector for a Network in vpc to populate networkId.
// +kubebuilder:validation:Optional
NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

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

// ID of the service account authorized for this cluster.
// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/iam/v1alpha1.ServiceAccount
// +kubebuilder:validation:Optional
ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

// Reference to a ServiceAccount in iam to populate serviceAccountId.
// +kubebuilder:validation:Optional
ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

// Selector for a ServiceAccount in iam to populate serviceAccountId.
// +kubebuilder:validation:Optional
ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`
}


type MasterNodeInitParameters struct {


// Resources allocated to hosts of the Elasticsearch master nodes subcluster. The structure is documented below.
Resources []MasterNodeResourcesInitParameters `json:"resources,omitempty" tf:"resources,omitempty"`
}


type MasterNodeObservation struct {


// Resources allocated to hosts of the Elasticsearch master nodes subcluster. The structure is documented below.
Resources []MasterNodeResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`
}


type MasterNodeParameters struct {


// Resources allocated to hosts of the Elasticsearch master nodes subcluster. The structure is documented below.
// +kubebuilder:validation:Optional
Resources []MasterNodeResourcesParameters `json:"resources" tf:"resources,omitempty"`
}


type MasterNodeResourcesInitParameters struct {


// Volume of the storage available to a host, in gigabytes.
DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

// Type of the storage of Elasticsearch hosts.
DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

ResourcePresetID *string `json:"resourcePresetId,omitempty" tf:"resource_preset_id,omitempty"`
}


type MasterNodeResourcesObservation struct {


// Volume of the storage available to a host, in gigabytes.
DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

// Type of the storage of Elasticsearch hosts.
DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

ResourcePresetID *string `json:"resourcePresetId,omitempty" tf:"resource_preset_id,omitempty"`
}


type MasterNodeResourcesParameters struct {


// Volume of the storage available to a host, in gigabytes.
// +kubebuilder:validation:Optional
DiskSize *float64 `json:"diskSize" tf:"disk_size,omitempty"`

// Type of the storage of Elasticsearch hosts.
// +kubebuilder:validation:Optional
DiskTypeID *string `json:"diskTypeId" tf:"disk_type_id,omitempty"`

// +kubebuilder:validation:Optional
ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

// ElasticsearchClusterSpec defines the desired state of ElasticsearchCluster
type ElasticsearchClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       ElasticsearchClusterParameters `json:"forProvider"`
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
	InitProvider       ElasticsearchClusterInitParameters `json:"initProvider,omitempty"`
}

// ElasticsearchClusterStatus defines the observed state of ElasticsearchCluster.
type ElasticsearchClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          ElasticsearchClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion


// ElasticsearchCluster is the Schema for the ElasticsearchClusters API. Manages a Elasticsearch cluster within Yandex.Cloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type ElasticsearchCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.config) || (has(self.initProvider) && has(self.initProvider.config))",message="spec.forProvider.config is a required parameter"
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.environment) || (has(self.initProvider) && has(self.initProvider.environment))",message="spec.forProvider.environment is a required parameter"
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec              ElasticsearchClusterSpec   `json:"spec"`
	Status            ElasticsearchClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchClusterList contains a list of ElasticsearchClusters
type ElasticsearchClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticsearchCluster `json:"items"`
}

// Repository type metadata.
var (
	ElasticsearchCluster_Kind             = "ElasticsearchCluster"
	ElasticsearchCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ElasticsearchCluster_Kind}.String()
	ElasticsearchCluster_KindAPIVersion   = ElasticsearchCluster_Kind + "." + CRDGroupVersion.String()
	ElasticsearchCluster_GroupVersionKind = CRDGroupVersion.WithKind(ElasticsearchCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&ElasticsearchCluster{}, &ElasticsearchClusterList{})
}
