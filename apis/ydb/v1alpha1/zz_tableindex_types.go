// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TableIndexInitParameters struct {
	Columns []*string `json:"columns,omitempty" tf:"columns,omitempty"`

	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	Cover []*string `json:"cover,omitempty" tf:"cover,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TableID *string `json:"tableId,omitempty" tf:"table_id,omitempty"`

	TablePath *string `json:"tablePath,omitempty" tf:"table_path,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TableIndexObservation struct {
	Columns []*string `json:"columns,omitempty" tf:"columns,omitempty"`

	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	Cover []*string `json:"cover,omitempty" tf:"cover,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TableID *string `json:"tableId,omitempty" tf:"table_id,omitempty"`

	TablePath *string `json:"tablePath,omitempty" tf:"table_path,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TableIndexParameters struct {

	// +kubebuilder:validation:Optional
	Columns []*string `json:"columns,omitempty" tf:"columns,omitempty"`

	// +kubebuilder:validation:Optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// +kubebuilder:validation:Optional
	Cover []*string `json:"cover,omitempty" tf:"cover,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	TableID *string `json:"tableId,omitempty" tf:"table_id,omitempty"`

	// +kubebuilder:validation:Optional
	TablePath *string `json:"tablePath,omitempty" tf:"table_path,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// TableIndexSpec defines the desired state of TableIndex
type TableIndexSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableIndexParameters `json:"forProvider"`
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
	InitProvider TableIndexInitParameters `json:"initProvider,omitempty"`
}

// TableIndexStatus defines the observed state of TableIndex.
type TableIndexStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableIndexObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TableIndex is the Schema for the TableIndexs API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type TableIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.columns) || (has(self.initProvider) && has(self.initProvider.columns))",message="spec.forProvider.columns is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   TableIndexSpec   `json:"spec"`
	Status TableIndexStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableIndexList contains a list of TableIndexs
type TableIndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableIndex `json:"items"`
}

// Repository type metadata.
var (
	TableIndex_Kind             = "TableIndex"
	TableIndex_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableIndex_Kind}.String()
	TableIndex_KindAPIVersion   = TableIndex_Kind + "." + CRDGroupVersion.String()
	TableIndex_GroupVersionKind = CRDGroupVersion.WithKind(TableIndex_Kind)
)

func init() {
	SchemeBuilder.Register(&TableIndex{}, &TableIndexList{})
}
