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

type PostgresqlUserInitParameters struct {

	// +crossplane:generate:reference:type=PostgresqlCluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a PostgresqlCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a PostgresqlCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// The maximum number of connections per user. (Default 50)
	ConnLimit *float64 `json:"connLimit,omitempty" tf:"conn_limit,omitempty"`

	// Inhibits deletion of the user. Can either be true, false or unspecified.
	DeletionProtection *string `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// List of the user's grants.
	Grants []*string `json:"grants,omitempty" tf:"grants,omitempty"`

	// User's ability to login.
	Login *bool `json:"login,omitempty" tf:"login,omitempty"`

	// The name of the user.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of permissions granted to the user. The structure is documented below.
	Permission []PostgresqlUserPermissionInitParameters `json:"permission,omitempty" tf:"permission,omitempty"`

	// Map of user settings. List of settings is documented below.
	// +mapType=granular
	Settings map[string]*string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type PostgresqlUserObservation struct {
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// The maximum number of connections per user. (Default 50)
	ConnLimit *float64 `json:"connLimit,omitempty" tf:"conn_limit,omitempty"`

	// Inhibits deletion of the user. Can either be true, false or unspecified.
	DeletionProtection *string `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// List of the user's grants.
	Grants []*string `json:"grants,omitempty" tf:"grants,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// User's ability to login.
	Login *bool `json:"login,omitempty" tf:"login,omitempty"`

	// The name of the user.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of permissions granted to the user. The structure is documented below.
	Permission []PostgresqlUserPermissionObservation `json:"permission,omitempty" tf:"permission,omitempty"`

	// Map of user settings. List of settings is documented below.
	// +mapType=granular
	Settings map[string]*string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type PostgresqlUserParameters struct {

	// +crossplane:generate:reference:type=PostgresqlCluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a PostgresqlCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a PostgresqlCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// The maximum number of connections per user. (Default 50)
	// +kubebuilder:validation:Optional
	ConnLimit *float64 `json:"connLimit,omitempty" tf:"conn_limit,omitempty"`

	// Inhibits deletion of the user. Can either be true, false or unspecified.
	// +kubebuilder:validation:Optional
	DeletionProtection *string `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// List of the user's grants.
	// +kubebuilder:validation:Optional
	Grants []*string `json:"grants,omitempty" tf:"grants,omitempty"`

	// User's ability to login.
	// +kubebuilder:validation:Optional
	Login *bool `json:"login,omitempty" tf:"login,omitempty"`

	// The name of the user.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The password of the user.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Set of permissions granted to the user. The structure is documented below.
	// +kubebuilder:validation:Optional
	Permission []PostgresqlUserPermissionParameters `json:"permission,omitempty" tf:"permission,omitempty"`

	// Map of user settings. List of settings is documented below.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Settings map[string]*string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type PostgresqlUserPermissionInitParameters struct {

	// The name of the database that the permission grants access to.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`
}

type PostgresqlUserPermissionObservation struct {

	// The name of the database that the permission grants access to.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`
}

type PostgresqlUserPermissionParameters struct {

	// The name of the database that the permission grants access to.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`
}

// PostgresqlUserSpec defines the desired state of PostgresqlUser
type PostgresqlUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PostgresqlUserParameters `json:"forProvider"`
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
	InitProvider PostgresqlUserInitParameters `json:"initProvider,omitempty"`
}

// PostgresqlUserStatus defines the observed state of PostgresqlUser.
type PostgresqlUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PostgresqlUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PostgresqlUser is the Schema for the PostgresqlUsers API. Manages a PostgreSQL user within Yandex.Cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type PostgresqlUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	Spec   PostgresqlUserSpec   `json:"spec"`
	Status PostgresqlUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlUserList contains a list of PostgresqlUsers
type PostgresqlUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlUser `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlUser_Kind             = "PostgresqlUser"
	PostgresqlUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PostgresqlUser_Kind}.String()
	PostgresqlUser_KindAPIVersion   = PostgresqlUser_Kind + "." + CRDGroupVersion.String()
	PostgresqlUser_GroupVersionKind = CRDGroupVersion.WithKind(PostgresqlUser_Kind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlUser{}, &PostgresqlUserList{})
}
