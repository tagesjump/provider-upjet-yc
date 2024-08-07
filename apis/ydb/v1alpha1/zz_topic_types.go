// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TopicConsumerInitParameters struct {
	Important *bool `json:"important,omitempty" tf:"important,omitempty"`

	// Topic name. Type: string, required. Default value: "".
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Timestamp in UNIX timestamp format from which the reader will start reading data. Type: integer, optional. Default value: 0.
	StartingMessageTimestampMs *float64 `json:"startingMessageTimestampMs,omitempty" tf:"starting_message_timestamp_ms,omitempty"`

	// Supported data encodings. Types: array[string]. Default value: ["gzip", "raw", "zstd"].
	// +listType=set
	SupportedCodecs []*string `json:"supportedCodecs,omitempty" tf:"supported_codecs,omitempty"`
}

type TopicConsumerObservation struct {
	Important *bool `json:"important,omitempty" tf:"important,omitempty"`

	// Topic name. Type: string, required. Default value: "".
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Timestamp in UNIX timestamp format from which the reader will start reading data. Type: integer, optional. Default value: 0.
	StartingMessageTimestampMs *float64 `json:"startingMessageTimestampMs,omitempty" tf:"starting_message_timestamp_ms,omitempty"`

	// Supported data encodings. Types: array[string]. Default value: ["gzip", "raw", "zstd"].
	// +listType=set
	SupportedCodecs []*string `json:"supportedCodecs,omitempty" tf:"supported_codecs,omitempty"`
}

type TopicConsumerParameters struct {

	// +kubebuilder:validation:Optional
	Important *bool `json:"important,omitempty" tf:"important,omitempty"`

	// Topic name. Type: string, required. Default value: "".
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Timestamp in UNIX timestamp format from which the reader will start reading data. Type: integer, optional. Default value: 0.
	// +kubebuilder:validation:Optional
	StartingMessageTimestampMs *float64 `json:"startingMessageTimestampMs,omitempty" tf:"starting_message_timestamp_ms,omitempty"`

	// Supported data encodings. Types: array[string]. Default value: ["gzip", "raw", "zstd"].
	// +kubebuilder:validation:Optional
	// +listType=set
	SupportedCodecs []*string `json:"supportedCodecs,omitempty" tf:"supported_codecs,omitempty"`
}

type TopicInitParameters struct {

	// Topic Readers. Types: array[consumer], optional. Default value: null.
	Consumer []TopicConsumerInitParameters `json:"consumer,omitempty" tf:"consumer,omitempty"`

	// YDB database endpoint. Types: string, required. Default value: "".
	DatabaseEndpoint *string `json:"databaseEndpoint,omitempty" tf:"database_endpoint,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	MeteringMode *string `json:"meteringMode,omitempty" tf:"metering_mode,omitempty"`

	// Topic name. Type: string, required. Default value: "".
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	PartitionWriteSpeedKbps *float64 `json:"partitionWriteSpeedKbps,omitempty" tf:"partition_write_speed_kbps,omitempty"`

	// Number of partitions. Types: integer, optional. Default value: 2.
	PartitionsCount *float64 `json:"partitionsCount,omitempty" tf:"partitions_count,omitempty"`

	RetentionPeriodHours *float64 `json:"retentionPeriodHours,omitempty" tf:"retention_period_hours,omitempty"`

	RetentionStorageMb *float64 `json:"retentionStorageMb,omitempty" tf:"retention_storage_mb,omitempty"`

	// Supported data encodings. Types: array[string]. Default value: ["gzip", "raw", "zstd"].
	// +listType=set
	SupportedCodecs []*string `json:"supportedCodecs,omitempty" tf:"supported_codecs,omitempty"`
}

type TopicObservation struct {

	// Topic Readers. Types: array[consumer], optional. Default value: null.
	Consumer []TopicConsumerObservation `json:"consumer,omitempty" tf:"consumer,omitempty"`

	// YDB database endpoint. Types: string, required. Default value: "".
	DatabaseEndpoint *string `json:"databaseEndpoint,omitempty" tf:"database_endpoint,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MeteringMode *string `json:"meteringMode,omitempty" tf:"metering_mode,omitempty"`

	// Topic name. Type: string, required. Default value: "".
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	PartitionWriteSpeedKbps *float64 `json:"partitionWriteSpeedKbps,omitempty" tf:"partition_write_speed_kbps,omitempty"`

	// Number of partitions. Types: integer, optional. Default value: 2.
	PartitionsCount *float64 `json:"partitionsCount,omitempty" tf:"partitions_count,omitempty"`

	RetentionPeriodHours *float64 `json:"retentionPeriodHours,omitempty" tf:"retention_period_hours,omitempty"`

	RetentionStorageMb *float64 `json:"retentionStorageMb,omitempty" tf:"retention_storage_mb,omitempty"`

	// Supported data encodings. Types: array[string]. Default value: ["gzip", "raw", "zstd"].
	// +listType=set
	SupportedCodecs []*string `json:"supportedCodecs,omitempty" tf:"supported_codecs,omitempty"`
}

type TopicParameters struct {

	// Topic Readers. Types: array[consumer], optional. Default value: null.
	// +kubebuilder:validation:Optional
	Consumer []TopicConsumerParameters `json:"consumer,omitempty" tf:"consumer,omitempty"`

	// YDB database endpoint. Types: string, required. Default value: "".
	// +kubebuilder:validation:Optional
	DatabaseEndpoint *string `json:"databaseEndpoint,omitempty" tf:"database_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	MeteringMode *string `json:"meteringMode,omitempty" tf:"metering_mode,omitempty"`

	// Topic name. Type: string, required. Default value: "".
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PartitionWriteSpeedKbps *float64 `json:"partitionWriteSpeedKbps,omitempty" tf:"partition_write_speed_kbps,omitempty"`

	// Number of partitions. Types: integer, optional. Default value: 2.
	// +kubebuilder:validation:Optional
	PartitionsCount *float64 `json:"partitionsCount,omitempty" tf:"partitions_count,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionPeriodHours *float64 `json:"retentionPeriodHours,omitempty" tf:"retention_period_hours,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionStorageMb *float64 `json:"retentionStorageMb,omitempty" tf:"retention_storage_mb,omitempty"`

	// Supported data encodings. Types: array[string]. Default value: ["gzip", "raw", "zstd"].
	// +kubebuilder:validation:Optional
	// +listType=set
	SupportedCodecs []*string `json:"supportedCodecs,omitempty" tf:"supported_codecs,omitempty"`
}

// TopicSpec defines the desired state of Topic
type TopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicParameters `json:"forProvider"`
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
	InitProvider TopicInitParameters `json:"initProvider,omitempty"`
}

// TopicStatus defines the observed state of Topic.
type TopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Topic is the Schema for the Topics API. Get information about a Yandex YDB Topics.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.databaseEndpoint) || (has(self.initProvider) && has(self.initProvider.databaseEndpoint))",message="spec.forProvider.databaseEndpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   TopicSpec   `json:"spec"`
	Status TopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicList contains a list of Topics
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Topic `json:"items"`
}

// Repository type metadata.
var (
	Topic_Kind             = "Topic"
	Topic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Topic_Kind}.String()
	Topic_KindAPIVersion   = Topic_Kind + "." + CRDGroupVersion.String()
	Topic_GroupVersionKind = CRDGroupVersion.WithKind(Topic_Kind)
)

func init() {
	SchemeBuilder.Register(&Topic{}, &TopicList{})
}
