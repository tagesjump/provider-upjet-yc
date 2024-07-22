// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConnectorConfigMirrormakerInitParameters struct {

	// Replication factor for topics created in target cluster
	ReplicationFactor *float64 `json:"replicationFactor,omitempty" tf:"replication_factor,omitempty"`

	// Settings for source cluster. The structure is documented below.
	SourceCluster []SourceClusterInitParameters `json:"sourceCluster,omitempty" tf:"source_cluster,omitempty"`

	// Settings for target cluster. The structure is documented below.
	TargetCluster []TargetClusterInitParameters `json:"targetCluster,omitempty" tf:"target_cluster,omitempty"`

	// The pattern for topic names to be replicated.
	Topics *string `json:"topics,omitempty" tf:"topics,omitempty"`
}

type ConnectorConfigMirrormakerObservation struct {

	// Replication factor for topics created in target cluster
	ReplicationFactor *float64 `json:"replicationFactor,omitempty" tf:"replication_factor,omitempty"`

	// Settings for source cluster. The structure is documented below.
	SourceCluster []SourceClusterObservation `json:"sourceCluster,omitempty" tf:"source_cluster,omitempty"`

	// Settings for target cluster. The structure is documented below.
	TargetCluster []TargetClusterObservation `json:"targetCluster,omitempty" tf:"target_cluster,omitempty"`

	// The pattern for topic names to be replicated.
	Topics *string `json:"topics,omitempty" tf:"topics,omitempty"`
}

type ConnectorConfigMirrormakerParameters struct {

	// Replication factor for topics created in target cluster
	// +kubebuilder:validation:Optional
	ReplicationFactor *float64 `json:"replicationFactor" tf:"replication_factor,omitempty"`

	// Settings for source cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	SourceCluster []SourceClusterParameters `json:"sourceCluster" tf:"source_cluster,omitempty"`

	// Settings for target cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	TargetCluster []TargetClusterParameters `json:"targetCluster" tf:"target_cluster,omitempty"`

	// The pattern for topic names to be replicated.
	// +kubebuilder:validation:Optional
	Topics *string `json:"topics" tf:"topics,omitempty"`
}

type ConnectorConfigS3SinkInitParameters struct {

	// Сompression type for messages. Cannot be changed.
	FileCompressionType *string `json:"fileCompressionType,omitempty" tf:"file_compression_type,omitempty"`

	// Max records per file.
	FileMaxRecords *float64 `json:"fileMaxRecords,omitempty" tf:"file_max_records,omitempty"`

	// Settings for connection to s3-compatible storage. The structure is documented below.
	S3Connection []S3ConnectionInitParameters `json:"s3Connection,omitempty" tf:"s3_connection,omitempty"`

	// The pattern for topic names to be copied to s3 bucket.
	Topics *string `json:"topics,omitempty" tf:"topics,omitempty"`
}

type ConnectorConfigS3SinkObservation struct {

	// Сompression type for messages. Cannot be changed.
	FileCompressionType *string `json:"fileCompressionType,omitempty" tf:"file_compression_type,omitempty"`

	// Max records per file.
	FileMaxRecords *float64 `json:"fileMaxRecords,omitempty" tf:"file_max_records,omitempty"`

	// Settings for connection to s3-compatible storage. The structure is documented below.
	S3Connection []S3ConnectionObservation `json:"s3Connection,omitempty" tf:"s3_connection,omitempty"`

	// The pattern for topic names to be copied to s3 bucket.
	Topics *string `json:"topics,omitempty" tf:"topics,omitempty"`
}

type ConnectorConfigS3SinkParameters struct {

	// Сompression type for messages. Cannot be changed.
	// +kubebuilder:validation:Optional
	FileCompressionType *string `json:"fileCompressionType" tf:"file_compression_type,omitempty"`

	// Max records per file.
	// +kubebuilder:validation:Optional
	FileMaxRecords *float64 `json:"fileMaxRecords,omitempty" tf:"file_max_records,omitempty"`

	// Settings for connection to s3-compatible storage. The structure is documented below.
	// +kubebuilder:validation:Optional
	S3Connection []S3ConnectionParameters `json:"s3Connection" tf:"s3_connection,omitempty"`

	// The pattern for topic names to be copied to s3 bucket.
	// +kubebuilder:validation:Optional
	Topics *string `json:"topics" tf:"topics,omitempty"`
}

type ExternalClusterInitParameters struct {

	// List of bootstrap servers to connect to cluster
	BootstrapServers *string `json:"bootstrapServers,omitempty" tf:"bootstrap_servers,omitempty"`

	// Type of SASL authentification mechanism to use
	SaslMechanism *string `json:"saslMechanism,omitempty" tf:"sasl_mechanism,omitempty"`

	// Password to use in SASL authentification mechanism
	SaslPasswordSecretRef *v1.SecretKeySelector `json:"saslPasswordSecretRef,omitempty" tf:"-"`

	// Username to use in SASL authentification mechanism
	SaslUsername *string `json:"saslUsername,omitempty" tf:"sasl_username,omitempty"`

	// Security protocol to use
	SecurityProtocol *string `json:"securityProtocol,omitempty" tf:"security_protocol,omitempty"`
}

type ExternalClusterObservation struct {

	// List of bootstrap servers to connect to cluster
	BootstrapServers *string `json:"bootstrapServers,omitempty" tf:"bootstrap_servers,omitempty"`

	// Type of SASL authentification mechanism to use
	SaslMechanism *string `json:"saslMechanism,omitempty" tf:"sasl_mechanism,omitempty"`

	// Username to use in SASL authentification mechanism
	SaslUsername *string `json:"saslUsername,omitempty" tf:"sasl_username,omitempty"`

	// Security protocol to use
	SecurityProtocol *string `json:"securityProtocol,omitempty" tf:"security_protocol,omitempty"`
}

type ExternalClusterParameters struct {

	// List of bootstrap servers to connect to cluster
	// +kubebuilder:validation:Optional
	BootstrapServers *string `json:"bootstrapServers" tf:"bootstrap_servers,omitempty"`

	// Type of SASL authentification mechanism to use
	// +kubebuilder:validation:Optional
	SaslMechanism *string `json:"saslMechanism,omitempty" tf:"sasl_mechanism,omitempty"`

	// Password to use in SASL authentification mechanism
	// +kubebuilder:validation:Optional
	SaslPasswordSecretRef *v1.SecretKeySelector `json:"saslPasswordSecretRef,omitempty" tf:"-"`

	// Username to use in SASL authentification mechanism
	// +kubebuilder:validation:Optional
	SaslUsername *string `json:"saslUsername,omitempty" tf:"sasl_username,omitempty"`

	// Security protocol to use
	// +kubebuilder:validation:Optional
	SecurityProtocol *string `json:"securityProtocol,omitempty" tf:"security_protocol,omitempty"`
}

type ExternalS3InitParameters struct {

	// ID of aws-compatible static key.
	AccessKeyID *string `json:"accessKeyId,omitempty" tf:"access_key_id,omitempty"`

	// URL of s3-compatible storage.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// region of s3-compatible storage. Available region list.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Secret key of aws-compatible static key.
	SecretAccessKeySecretRef *v1.SecretKeySelector `json:"secretAccessKeySecretRef,omitempty" tf:"-"`
}

type ExternalS3Observation struct {

	// ID of aws-compatible static key.
	AccessKeyID *string `json:"accessKeyId,omitempty" tf:"access_key_id,omitempty"`

	// URL of s3-compatible storage.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// region of s3-compatible storage. Available region list.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type ExternalS3Parameters struct {

	// ID of aws-compatible static key.
	// +kubebuilder:validation:Optional
	AccessKeyID *string `json:"accessKeyId,omitempty" tf:"access_key_id,omitempty"`

	// URL of s3-compatible storage.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`

	// region of s3-compatible storage. Available region list.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Secret key of aws-compatible static key.
	// +kubebuilder:validation:Optional
	SecretAccessKeySecretRef *v1.SecretKeySelector `json:"secretAccessKeySecretRef,omitempty" tf:"-"`
}

type KafkaConnectorInitParameters struct {

	// +crossplane:generate:reference:type=KafkaCluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a KafkaCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a KafkaCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Params for MirrorMaker2 connector. The structure is documented below.
	ConnectorConfigMirrormaker []ConnectorConfigMirrormakerInitParameters `json:"connectorConfigMirrormaker,omitempty" tf:"connector_config_mirrormaker,omitempty"`

	// Params for S3 Sink connector. The structure is documented below.
	ConnectorConfigS3Sink []ConnectorConfigS3SinkInitParameters `json:"connectorConfigS3Sink,omitempty" tf:"connector_config_s3_sink,omitempty"`

	// The name of the connector.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Additional properties for connector.
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// The number of the connector's parallel working tasks. Default is the number of brokers
	TasksMax *float64 `json:"tasksMax,omitempty" tf:"tasks_max,omitempty"`
}

type KafkaConnectorObservation struct {
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Params for MirrorMaker2 connector. The structure is documented below.
	ConnectorConfigMirrormaker []ConnectorConfigMirrormakerObservation `json:"connectorConfigMirrormaker,omitempty" tf:"connector_config_mirrormaker,omitempty"`

	// Params for S3 Sink connector. The structure is documented below.
	ConnectorConfigS3Sink []ConnectorConfigS3SinkObservation `json:"connectorConfigS3Sink,omitempty" tf:"connector_config_s3_sink,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the connector.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Additional properties for connector.
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// The number of the connector's parallel working tasks. Default is the number of brokers
	TasksMax *float64 `json:"tasksMax,omitempty" tf:"tasks_max,omitempty"`
}

type KafkaConnectorParameters struct {

	// +crossplane:generate:reference:type=KafkaCluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a KafkaCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a KafkaCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Params for MirrorMaker2 connector. The structure is documented below.
	// +kubebuilder:validation:Optional
	ConnectorConfigMirrormaker []ConnectorConfigMirrormakerParameters `json:"connectorConfigMirrormaker,omitempty" tf:"connector_config_mirrormaker,omitempty"`

	// Params for S3 Sink connector. The structure is documented below.
	// +kubebuilder:validation:Optional
	ConnectorConfigS3Sink []ConnectorConfigS3SinkParameters `json:"connectorConfigS3Sink,omitempty" tf:"connector_config_s3_sink,omitempty"`

	// The name of the connector.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Additional properties for connector.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// The number of the connector's parallel working tasks. Default is the number of brokers
	// +kubebuilder:validation:Optional
	TasksMax *float64 `json:"tasksMax,omitempty" tf:"tasks_max,omitempty"`
}

type S3ConnectionInitParameters struct {

	// Name of the bucket in s3-compatible storage.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Connection params for external s3-compatible storage. The structure is documented below.
	ExternalS3 []ExternalS3InitParameters `json:"externalS3,omitempty" tf:"external_s3,omitempty"`
}

type S3ConnectionObservation struct {

	// Name of the bucket in s3-compatible storage.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Connection params for external s3-compatible storage. The structure is documented below.
	ExternalS3 []ExternalS3Observation `json:"externalS3,omitempty" tf:"external_s3,omitempty"`
}

type S3ConnectionParameters struct {

	// Name of the bucket in s3-compatible storage.
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`

	// Connection params for external s3-compatible storage. The structure is documented below.
	// +kubebuilder:validation:Optional
	ExternalS3 []ExternalS3Parameters `json:"externalS3" tf:"external_s3,omitempty"`
}

type SourceClusterInitParameters struct {

	// Name of the cluster. Used also as a topic prefix
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Connection params for external cluster
	ExternalCluster []ExternalClusterInitParameters `json:"externalCluster,omitempty" tf:"external_cluster,omitempty"`

	// Using this section in the cluster definition (source or target) means it's this cluster
	ThisCluster []ThisClusterInitParameters `json:"thisCluster,omitempty" tf:"this_cluster,omitempty"`
}

type SourceClusterObservation struct {

	// Name of the cluster. Used also as a topic prefix
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Connection params for external cluster
	ExternalCluster []ExternalClusterObservation `json:"externalCluster,omitempty" tf:"external_cluster,omitempty"`

	// Using this section in the cluster definition (source or target) means it's this cluster
	ThisCluster []ThisClusterParameters `json:"thisCluster,omitempty" tf:"this_cluster,omitempty"`
}

type SourceClusterParameters struct {

	// Name of the cluster. Used also as a topic prefix
	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Connection params for external cluster
	// +kubebuilder:validation:Optional
	ExternalCluster []ExternalClusterParameters `json:"externalCluster,omitempty" tf:"external_cluster,omitempty"`

	// Using this section in the cluster definition (source or target) means it's this cluster
	// +kubebuilder:validation:Optional
	ThisCluster []ThisClusterParameters `json:"thisCluster,omitempty" tf:"this_cluster,omitempty"`
}

type TargetClusterExternalClusterInitParameters struct {

	// List of bootstrap servers to connect to cluster
	BootstrapServers *string `json:"bootstrapServers,omitempty" tf:"bootstrap_servers,omitempty"`

	// Type of SASL authentification mechanism to use
	SaslMechanism *string `json:"saslMechanism,omitempty" tf:"sasl_mechanism,omitempty"`

	// Password to use in SASL authentification mechanism
	SaslPasswordSecretRef *v1.SecretKeySelector `json:"saslPasswordSecretRef,omitempty" tf:"-"`

	// Username to use in SASL authentification mechanism
	SaslUsername *string `json:"saslUsername,omitempty" tf:"sasl_username,omitempty"`

	// Security protocol to use
	SecurityProtocol *string `json:"securityProtocol,omitempty" tf:"security_protocol,omitempty"`
}

type TargetClusterExternalClusterObservation struct {

	// List of bootstrap servers to connect to cluster
	BootstrapServers *string `json:"bootstrapServers,omitempty" tf:"bootstrap_servers,omitempty"`

	// Type of SASL authentification mechanism to use
	SaslMechanism *string `json:"saslMechanism,omitempty" tf:"sasl_mechanism,omitempty"`

	// Username to use in SASL authentification mechanism
	SaslUsername *string `json:"saslUsername,omitempty" tf:"sasl_username,omitempty"`

	// Security protocol to use
	SecurityProtocol *string `json:"securityProtocol,omitempty" tf:"security_protocol,omitempty"`
}

type TargetClusterExternalClusterParameters struct {

	// List of bootstrap servers to connect to cluster
	// +kubebuilder:validation:Optional
	BootstrapServers *string `json:"bootstrapServers" tf:"bootstrap_servers,omitempty"`

	// Type of SASL authentification mechanism to use
	// +kubebuilder:validation:Optional
	SaslMechanism *string `json:"saslMechanism,omitempty" tf:"sasl_mechanism,omitempty"`

	// Password to use in SASL authentification mechanism
	// +kubebuilder:validation:Optional
	SaslPasswordSecretRef *v1.SecretKeySelector `json:"saslPasswordSecretRef,omitempty" tf:"-"`

	// Username to use in SASL authentification mechanism
	// +kubebuilder:validation:Optional
	SaslUsername *string `json:"saslUsername,omitempty" tf:"sasl_username,omitempty"`

	// Security protocol to use
	// +kubebuilder:validation:Optional
	SecurityProtocol *string `json:"securityProtocol,omitempty" tf:"security_protocol,omitempty"`
}

type TargetClusterInitParameters struct {

	// Name of the cluster. Used also as a topic prefix
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Connection params for external cluster
	ExternalCluster []TargetClusterExternalClusterInitParameters `json:"externalCluster,omitempty" tf:"external_cluster,omitempty"`

	// Using this section in the cluster definition (source or target) means it's this cluster
	ThisCluster []TargetClusterThisClusterInitParameters `json:"thisCluster,omitempty" tf:"this_cluster,omitempty"`
}

type TargetClusterObservation struct {

	// Name of the cluster. Used also as a topic prefix
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Connection params for external cluster
	ExternalCluster []TargetClusterExternalClusterObservation `json:"externalCluster,omitempty" tf:"external_cluster,omitempty"`

	// Using this section in the cluster definition (source or target) means it's this cluster
	ThisCluster []TargetClusterThisClusterParameters `json:"thisCluster,omitempty" tf:"this_cluster,omitempty"`
}

type TargetClusterParameters struct {

	// Name of the cluster. Used also as a topic prefix
	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Connection params for external cluster
	// +kubebuilder:validation:Optional
	ExternalCluster []TargetClusterExternalClusterParameters `json:"externalCluster,omitempty" tf:"external_cluster,omitempty"`

	// Using this section in the cluster definition (source or target) means it's this cluster
	// +kubebuilder:validation:Optional
	ThisCluster []TargetClusterThisClusterParameters `json:"thisCluster,omitempty" tf:"this_cluster,omitempty"`
}

type TargetClusterThisClusterInitParameters struct {
}

type TargetClusterThisClusterObservation struct {
}

type TargetClusterThisClusterParameters struct {
}

type ThisClusterInitParameters struct {
}

type ThisClusterObservation struct {
}

type ThisClusterParameters struct {
}

// KafkaConnectorSpec defines the desired state of KafkaConnector
type KafkaConnectorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KafkaConnectorParameters `json:"forProvider"`
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
	InitProvider KafkaConnectorInitParameters `json:"initProvider,omitempty"`
}

// KafkaConnectorStatus defines the observed state of KafkaConnector.
type KafkaConnectorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KafkaConnectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// KafkaConnector is the Schema for the KafkaConnectors API. Manages a connectors of a Kafka cluster within Yandex.Cloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type KafkaConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   KafkaConnectorSpec   `json:"spec"`
	Status KafkaConnectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaConnectorList contains a list of KafkaConnectors
type KafkaConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaConnector `json:"items"`
}

// Repository type metadata.
var (
	KafkaConnector_Kind             = "KafkaConnector"
	KafkaConnector_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KafkaConnector_Kind}.String()
	KafkaConnector_KindAPIVersion   = KafkaConnector_Kind + "." + CRDGroupVersion.String()
	KafkaConnector_GroupVersionKind = CRDGroupVersion.WithKind(KafkaConnector_Kind)
)

func init() {
	SchemeBuilder.Register(&KafkaConnector{}, &KafkaConnectorList{})
}
