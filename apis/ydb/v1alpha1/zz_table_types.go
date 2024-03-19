// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ColumnInitParameters struct {

	// Column group
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// Column name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A column cannot have the NULL data type. (	Default: false	)
	NotNull *bool `json:"notNull,omitempty" tf:"not_null,omitempty"`

	// Column data type. YQL data types are used.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ColumnObservation struct {

	// Column group
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// Column name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A column cannot have the NULL data type. (	Default: false	)
	NotNull *bool `json:"notNull,omitempty" tf:"not_null,omitempty"`

	// Column data type. YQL data types are used.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ColumnParameters struct {

	// Column group
	// +kubebuilder:validation:Optional
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// Column name
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A column cannot have the NULL data type. (	Default: false	)
	// +kubebuilder:validation:Optional
	NotNull *bool `json:"notNull,omitempty" tf:"not_null,omitempty"`

	// Column data type. YQL data types are used.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type FamilyInitParameters struct {

	// Data codec (acceptable values: off, lz4).
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	// Type of storage device for column data in this group (acceptable values: ssd, rot (from HDD spindle rotation)).
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// Column family name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FamilyObservation struct {

	// Data codec (acceptable values: off, lz4).
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	// Type of storage device for column data in this group (acceptable values: ssd, rot (from HDD spindle rotation)).
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// Column family name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FamilyParameters struct {

	// Data codec (acceptable values: off, lz4).
	// +kubebuilder:validation:Optional
	Compression *string `json:"compression" tf:"compression,omitempty"`

	// Type of storage device for column data in this group (acceptable values: ssd, rot (from HDD spindle rotation)).
	// +kubebuilder:validation:Optional
	Data *string `json:"data" tf:"data,omitempty"`

	// Column family name
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type PartitionAtKeysInitParameters struct {
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`
}

type PartitionAtKeysObservation struct {
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`
}

type PartitionAtKeysParameters struct {

	// +kubebuilder:validation:Optional
	Keys []*string `json:"keys" tf:"keys,omitempty"`
}

type PartitioningSettingsInitParameters struct {
	AutoPartitioningByLoad *bool `json:"autoPartitioningByLoad,omitempty" tf:"auto_partitioning_by_load,omitempty"`

	AutoPartitioningBySizeEnabled *bool `json:"autoPartitioningBySizeEnabled,omitempty" tf:"auto_partitioning_by_size_enabled,omitempty"`

	AutoPartitioningMaxPartitionsCount *float64 `json:"autoPartitioningMaxPartitionsCount,omitempty" tf:"auto_partitioning_max_partitions_count,omitempty"`

	AutoPartitioningMinPartitionsCount *float64 `json:"autoPartitioningMinPartitionsCount,omitempty" tf:"auto_partitioning_min_partitions_count,omitempty"`

	AutoPartitioningPartitionSizeMb *float64 `json:"autoPartitioningPartitionSizeMb,omitempty" tf:"auto_partitioning_partition_size_mb,omitempty"`

	PartitionAtKeys []PartitionAtKeysInitParameters `json:"partitionAtKeys,omitempty" tf:"partition_at_keys,omitempty"`

	UniformPartitions *float64 `json:"uniformPartitions,omitempty" tf:"uniform_partitions,omitempty"`
}

type PartitioningSettingsObservation struct {
	AutoPartitioningByLoad *bool `json:"autoPartitioningByLoad,omitempty" tf:"auto_partitioning_by_load,omitempty"`

	AutoPartitioningBySizeEnabled *bool `json:"autoPartitioningBySizeEnabled,omitempty" tf:"auto_partitioning_by_size_enabled,omitempty"`

	AutoPartitioningMaxPartitionsCount *float64 `json:"autoPartitioningMaxPartitionsCount,omitempty" tf:"auto_partitioning_max_partitions_count,omitempty"`

	AutoPartitioningMinPartitionsCount *float64 `json:"autoPartitioningMinPartitionsCount,omitempty" tf:"auto_partitioning_min_partitions_count,omitempty"`

	AutoPartitioningPartitionSizeMb *float64 `json:"autoPartitioningPartitionSizeMb,omitempty" tf:"auto_partitioning_partition_size_mb,omitempty"`

	PartitionAtKeys []PartitionAtKeysObservation `json:"partitionAtKeys,omitempty" tf:"partition_at_keys,omitempty"`

	UniformPartitions *float64 `json:"uniformPartitions,omitempty" tf:"uniform_partitions,omitempty"`
}

type PartitioningSettingsParameters struct {

	// +kubebuilder:validation:Optional
	AutoPartitioningByLoad *bool `json:"autoPartitioningByLoad,omitempty" tf:"auto_partitioning_by_load,omitempty"`

	// +kubebuilder:validation:Optional
	AutoPartitioningBySizeEnabled *bool `json:"autoPartitioningBySizeEnabled,omitempty" tf:"auto_partitioning_by_size_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	AutoPartitioningMaxPartitionsCount *float64 `json:"autoPartitioningMaxPartitionsCount,omitempty" tf:"auto_partitioning_max_partitions_count,omitempty"`

	// +kubebuilder:validation:Optional
	AutoPartitioningMinPartitionsCount *float64 `json:"autoPartitioningMinPartitionsCount,omitempty" tf:"auto_partitioning_min_partitions_count,omitempty"`

	// +kubebuilder:validation:Optional
	AutoPartitioningPartitionSizeMb *float64 `json:"autoPartitioningPartitionSizeMb,omitempty" tf:"auto_partitioning_partition_size_mb,omitempty"`

	// +kubebuilder:validation:Optional
	PartitionAtKeys []PartitionAtKeysParameters `json:"partitionAtKeys,omitempty" tf:"partition_at_keys,omitempty"`

	// +kubebuilder:validation:Optional
	UniformPartitions *float64 `json:"uniformPartitions,omitempty" tf:"uniform_partitions,omitempty"`
}

type TTLInitParameters struct {

	// Column name for TTL
	ColumnName *string `json:"columnName,omitempty" tf:"column_name,omitempty"`

	// Interval in the ISO 8601 format
	ExpireInterval *string `json:"expireInterval,omitempty" tf:"expire_interval,omitempty"`

	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type TTLObservation struct {

	// Column name for TTL
	ColumnName *string `json:"columnName,omitempty" tf:"column_name,omitempty"`

	// Interval in the ISO 8601 format
	ExpireInterval *string `json:"expireInterval,omitempty" tf:"expire_interval,omitempty"`

	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type TTLParameters struct {

	// Column name for TTL
	// +kubebuilder:validation:Optional
	ColumnName *string `json:"columnName" tf:"column_name,omitempty"`

	// Interval in the ISO 8601 format
	// +kubebuilder:validation:Optional
	ExpireInterval *string `json:"expireInterval" tf:"expire_interval,omitempty"`

	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type TableInitParameters struct {

	// A map of table attributes.
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// A list of column configuration options.
	// The structure is documented below.
	Column []ColumnInitParameters `json:"column,omitempty" tf:"column,omitempty"`

	// Connection string for database.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// A list of column group configuration options.
	// The structure is documented below.
	Family []FamilyInitParameters `json:"family,omitempty" tf:"family,omitempty"`

	// Use the Bloom filter for the primary key
	KeyBloomFilter *bool `json:"keyBloomFilter,omitempty" tf:"key_bloom_filter,omitempty"`

	// Table partiotioning settings
	// The structure is documented below.
	PartitioningSettings []PartitioningSettingsInitParameters `json:"partitioningSettings,omitempty" tf:"partitioning_settings,omitempty"`

	// Table path.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// A list of table columns to be uased as primary key.
	PrimaryKey []*string `json:"primaryKey,omitempty" tf:"primary_key,omitempty"`

	// Read replication settings
	ReadReplicasSettings *string `json:"readReplicasSettings,omitempty" tf:"read_replicas_settings,omitempty"`

	// ttl		TTL settings
	// The structure is documented below.
	TTL []TTLInitParameters `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type TableObservation struct {

	// A map of table attributes.
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// A list of column configuration options.
	// The structure is documented below.
	Column []ColumnObservation `json:"column,omitempty" tf:"column,omitempty"`

	// Connection string for database.
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// A list of column group configuration options.
	// The structure is documented below.
	Family []FamilyObservation `json:"family,omitempty" tf:"family,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Use the Bloom filter for the primary key
	KeyBloomFilter *bool `json:"keyBloomFilter,omitempty" tf:"key_bloom_filter,omitempty"`

	// Table partiotioning settings
	// The structure is documented below.
	PartitioningSettings []PartitioningSettingsObservation `json:"partitioningSettings,omitempty" tf:"partitioning_settings,omitempty"`

	// Table path.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// A list of table columns to be uased as primary key.
	PrimaryKey []*string `json:"primaryKey,omitempty" tf:"primary_key,omitempty"`

	// Read replication settings
	ReadReplicasSettings *string `json:"readReplicasSettings,omitempty" tf:"read_replicas_settings,omitempty"`

	// ttl		TTL settings
	// The structure is documented below.
	TTL []TTLObservation `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type TableParameters struct {

	// A map of table attributes.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// A list of column configuration options.
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	Column []ColumnParameters `json:"column,omitempty" tf:"column,omitempty"`

	// Connection string for database.
	// +kubebuilder:validation:Optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// A list of column group configuration options.
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	Family []FamilyParameters `json:"family,omitempty" tf:"family,omitempty"`

	// Use the Bloom filter for the primary key
	// +kubebuilder:validation:Optional
	KeyBloomFilter *bool `json:"keyBloomFilter,omitempty" tf:"key_bloom_filter,omitempty"`

	// Table partiotioning settings
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	PartitioningSettings []PartitioningSettingsParameters `json:"partitioningSettings,omitempty" tf:"partitioning_settings,omitempty"`

	// Table path.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// A list of table columns to be uased as primary key.
	// +kubebuilder:validation:Optional
	PrimaryKey []*string `json:"primaryKey,omitempty" tf:"primary_key,omitempty"`

	// Read replication settings
	// +kubebuilder:validation:Optional
	ReadReplicasSettings *string `json:"readReplicasSettings,omitempty" tf:"read_replicas_settings,omitempty"`

	// ttl		TTL settings
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	TTL []TTLParameters `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableParameters `json:"forProvider"`
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
	InitProvider TableInitParameters `json:"initProvider,omitempty"`
}

// TableStatus defines the observed state of Table.
type TableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Table is the Schema for the Tables API. Manages Yandex Database dedicated cluster.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.column) || (has(self.initProvider) && has(self.initProvider.column))",message="spec.forProvider.column is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionString) || (has(self.initProvider) && has(self.initProvider.connectionString))",message="spec.forProvider.connectionString is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.path) || (has(self.initProvider) && has(self.initProvider.path))",message="spec.forProvider.path is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.primaryKey) || (has(self.initProvider) && has(self.initProvider.primaryKey))",message="spec.forProvider.primaryKey is a required parameter"
	Spec   TableSpec   `json:"spec"`
	Status TableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableList contains a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

// Repository type metadata.
var (
	Table_Kind             = "Table"
	Table_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Table_Kind}.String()
	Table_KindAPIVersion   = Table_Kind + "." + CRDGroupVersion.String()
	Table_GroupVersionKind = CRDGroupVersion.WithKind(Table_Kind)
)

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}
