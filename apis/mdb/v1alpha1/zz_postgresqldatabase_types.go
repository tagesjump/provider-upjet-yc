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

type PostgresqlDatabaseExtensionInitParameters struct {

	// Name of the database extension. For more information on available extensions see the official documentation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Version of the extension.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type PostgresqlDatabaseExtensionObservation struct {

	// Name of the database extension. For more information on available extensions see the official documentation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Version of the extension.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type PostgresqlDatabaseExtensionParameters struct {

	// Name of the database extension. For more information on available extensions see the official documentation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Version of the extension.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type PostgresqlDatabaseInitParameters struct {

	// +crossplane:generate:reference:type=PostgresqlCluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a PostgresqlCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a PostgresqlCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Inhibits deletion of the database. Can either be true, false or unspecified.
	DeletionProtection *string `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Set of database extensions. The structure is documented below
	Extension []PostgresqlDatabaseExtensionInitParameters `json:"extension,omitempty" tf:"extension,omitempty"`

	// POSIX locale for string sorting order. Forbidden to change in an existing database.
	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	// POSIX locale for character classification. Forbidden to change in an existing database.
	LcType *string `json:"lcType,omitempty" tf:"lc_type,omitempty"`

	// The name of the database.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the user assigned as the owner of the database. Forbidden to change in an existing database.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Name of the template database.
	TemplateDB *string `json:"templateDb,omitempty" tf:"template_db,omitempty"`
}

type PostgresqlDatabaseObservation struct {
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Inhibits deletion of the database. Can either be true, false or unspecified.
	DeletionProtection *string `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Set of database extensions. The structure is documented below
	Extension []PostgresqlDatabaseExtensionObservation `json:"extension,omitempty" tf:"extension,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// POSIX locale for string sorting order. Forbidden to change in an existing database.
	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	// POSIX locale for character classification. Forbidden to change in an existing database.
	LcType *string `json:"lcType,omitempty" tf:"lc_type,omitempty"`

	// The name of the database.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the user assigned as the owner of the database. Forbidden to change in an existing database.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Name of the template database.
	TemplateDB *string `json:"templateDb,omitempty" tf:"template_db,omitempty"`
}

type PostgresqlDatabaseParameters struct {

	// +crossplane:generate:reference:type=PostgresqlCluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a PostgresqlCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a PostgresqlCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Inhibits deletion of the database. Can either be true, false or unspecified.
	// +kubebuilder:validation:Optional
	DeletionProtection *string `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Set of database extensions. The structure is documented below
	// +kubebuilder:validation:Optional
	Extension []PostgresqlDatabaseExtensionParameters `json:"extension,omitempty" tf:"extension,omitempty"`

	// POSIX locale for string sorting order. Forbidden to change in an existing database.
	// +kubebuilder:validation:Optional
	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	// POSIX locale for character classification. Forbidden to change in an existing database.
	// +kubebuilder:validation:Optional
	LcType *string `json:"lcType,omitempty" tf:"lc_type,omitempty"`

	// The name of the database.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Name of the user assigned as the owner of the database. Forbidden to change in an existing database.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Name of the template database.
	// +kubebuilder:validation:Optional
	TemplateDB *string `json:"templateDb,omitempty" tf:"template_db,omitempty"`
}

// PostgresqlDatabaseSpec defines the desired state of PostgresqlDatabase
type PostgresqlDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PostgresqlDatabaseParameters `json:"forProvider"`
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
	InitProvider PostgresqlDatabaseInitParameters `json:"initProvider,omitempty"`
}

// PostgresqlDatabaseStatus defines the observed state of PostgresqlDatabase.
type PostgresqlDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PostgresqlDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PostgresqlDatabase is the Schema for the PostgresqlDatabases API. Manages a PostgreSQL database within Yandex.Cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type PostgresqlDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.owner) || (has(self.initProvider) && has(self.initProvider.owner))",message="spec.forProvider.owner is a required parameter"
	Spec   PostgresqlDatabaseSpec   `json:"spec"`
	Status PostgresqlDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlDatabaseList contains a list of PostgresqlDatabases
type PostgresqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlDatabase `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlDatabase_Kind             = "PostgresqlDatabase"
	PostgresqlDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PostgresqlDatabase_Kind}.String()
	PostgresqlDatabase_KindAPIVersion   = PostgresqlDatabase_Kind + "." + CRDGroupVersion.String()
	PostgresqlDatabase_GroupVersionKind = CRDGroupVersion.WithKind(PostgresqlDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlDatabase{}, &PostgresqlDatabaseList{})
}
