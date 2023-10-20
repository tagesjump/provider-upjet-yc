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

type SAMLFederationInitParameters struct {

	// Add new users automatically on successful authentication. The user will get the resource-manager.clouds.member role automatically, but you need to grant other roles to them. If the value is false, users who aren't added to the cloud can't log in, even if they have authenticated on your server.
	AutoCreateAccountOnLogin *bool `json:"autoCreateAccountOnLogin,omitempty" tf:"auto_create_account_on_login,omitempty"`

	// Use case-insensitive name ids.
	CaseInsensitiveNameIds *bool `json:"caseInsensitiveNameIds,omitempty" tf:"case_insensitive_name_ids,omitempty"`

	// The lifetime of a Browser cookie in seconds. If the cookie is still valid, the management console authenticates the user immediately and redirects them to the home page. The default value is 8h.
	CookieMaxAge *string `json:"cookieMaxAge,omitempty" tf:"cookie_max_age,omitempty"`

	// The description of the SAML Federation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the IdP server to be used for authentication. The IdP server also responds to IAM with this ID after the user authenticates.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// A set of key/value label pairs assigned to the SAML Federation.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the SAML Federation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The organization to attach this SAML Federation to.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Federation security settings, structure is documented below.
	SecuritySettings []SecuritySettingsInitParameters `json:"securitySettings,omitempty" tf:"security_settings,omitempty"`

	// Single sign-on endpoint binding type. Most Identity Providers support the POST binding type. SAML Binding is a mapping of a SAML protocol message onto standard messaging formats and/or communications protocols.
	SsoBinding *string `json:"ssoBinding,omitempty" tf:"sso_binding,omitempty"`

	// Single sign-on endpoint URL. Specify the link to the IdP login page here.
	SsoURL *string `json:"ssoUrl,omitempty" tf:"sso_url,omitempty"`
}

type SAMLFederationObservation struct {

	// Add new users automatically on successful authentication. The user will get the resource-manager.clouds.member role automatically, but you need to grant other roles to them. If the value is false, users who aren't added to the cloud can't log in, even if they have authenticated on your server.
	AutoCreateAccountOnLogin *bool `json:"autoCreateAccountOnLogin,omitempty" tf:"auto_create_account_on_login,omitempty"`

	// Use case-insensitive name ids.
	CaseInsensitiveNameIds *bool `json:"caseInsensitiveNameIds,omitempty" tf:"case_insensitive_name_ids,omitempty"`

	// The lifetime of a Browser cookie in seconds. If the cookie is still valid, the management console authenticates the user immediately and redirects them to the home page. The default value is 8h.
	CookieMaxAge *string `json:"cookieMaxAge,omitempty" tf:"cookie_max_age,omitempty"`

	// (Computed) The SAML Federation creation timestamp.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The description of the SAML Federation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the IdP server to be used for authentication. The IdP server also responds to IAM with this ID after the user authenticates.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// A set of key/value label pairs assigned to the SAML Federation.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the SAML Federation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The organization to attach this SAML Federation to.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Federation security settings, structure is documented below.
	SecuritySettings []SecuritySettingsObservation `json:"securitySettings,omitempty" tf:"security_settings,omitempty"`

	// Single sign-on endpoint binding type. Most Identity Providers support the POST binding type. SAML Binding is a mapping of a SAML protocol message onto standard messaging formats and/or communications protocols.
	SsoBinding *string `json:"ssoBinding,omitempty" tf:"sso_binding,omitempty"`

	// Single sign-on endpoint URL. Specify the link to the IdP login page here.
	SsoURL *string `json:"ssoUrl,omitempty" tf:"sso_url,omitempty"`
}

type SAMLFederationParameters struct {

	// Add new users automatically on successful authentication. The user will get the resource-manager.clouds.member role automatically, but you need to grant other roles to them. If the value is false, users who aren't added to the cloud can't log in, even if they have authenticated on your server.
	// +kubebuilder:validation:Optional
	AutoCreateAccountOnLogin *bool `json:"autoCreateAccountOnLogin,omitempty" tf:"auto_create_account_on_login,omitempty"`

	// Use case-insensitive name ids.
	// +kubebuilder:validation:Optional
	CaseInsensitiveNameIds *bool `json:"caseInsensitiveNameIds,omitempty" tf:"case_insensitive_name_ids,omitempty"`

	// The lifetime of a Browser cookie in seconds. If the cookie is still valid, the management console authenticates the user immediately and redirects them to the home page. The default value is 8h.
	// +kubebuilder:validation:Optional
	CookieMaxAge *string `json:"cookieMaxAge,omitempty" tf:"cookie_max_age,omitempty"`

	// The description of the SAML Federation.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the IdP server to be used for authentication. The IdP server also responds to IAM with this ID after the user authenticates.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// A set of key/value label pairs assigned to the SAML Federation.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the SAML Federation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The organization to attach this SAML Federation to.
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Federation security settings, structure is documented below.
	// +kubebuilder:validation:Optional
	SecuritySettings []SecuritySettingsParameters `json:"securitySettings,omitempty" tf:"security_settings,omitempty"`

	// Single sign-on endpoint binding type. Most Identity Providers support the POST binding type. SAML Binding is a mapping of a SAML protocol message onto standard messaging formats and/or communications protocols.
	// +kubebuilder:validation:Optional
	SsoBinding *string `json:"ssoBinding,omitempty" tf:"sso_binding,omitempty"`

	// Single sign-on endpoint URL. Specify the link to the IdP login page here.
	// +kubebuilder:validation:Optional
	SsoURL *string `json:"ssoUrl,omitempty" tf:"sso_url,omitempty"`
}

type SecuritySettingsInitParameters struct {

	// Enable encrypted assertions.
	EncryptedAssertions *bool `json:"encryptedAssertions,omitempty" tf:"encrypted_assertions,omitempty"`
}

type SecuritySettingsObservation struct {

	// Enable encrypted assertions.
	EncryptedAssertions *bool `json:"encryptedAssertions,omitempty" tf:"encrypted_assertions,omitempty"`
}

type SecuritySettingsParameters struct {

	// Enable encrypted assertions.
	// +kubebuilder:validation:Optional
	EncryptedAssertions *bool `json:"encryptedAssertions" tf:"encrypted_assertions,omitempty"`
}

// SAMLFederationSpec defines the desired state of SAMLFederation
type SAMLFederationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SAMLFederationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SAMLFederationInitParameters `json:"initProvider,omitempty"`
}

// SAMLFederationStatus defines the observed state of SAMLFederation.
type SAMLFederationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SAMLFederationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SAMLFederation is the Schema for the SAMLFederations API. Allows management of a single SAML Federation within an existing Yandex.Cloud Organization.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud.upjet}
type SAMLFederation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.issuer) || (has(self.initProvider) && has(self.initProvider.issuer))",message="spec.forProvider.issuer is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.organizationId) || (has(self.initProvider) && has(self.initProvider.organizationId))",message="spec.forProvider.organizationId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ssoBinding) || (has(self.initProvider) && has(self.initProvider.ssoBinding))",message="spec.forProvider.ssoBinding is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ssoUrl) || (has(self.initProvider) && has(self.initProvider.ssoUrl))",message="spec.forProvider.ssoUrl is a required parameter"
	Spec   SAMLFederationSpec   `json:"spec"`
	Status SAMLFederationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SAMLFederationList contains a list of SAMLFederations
type SAMLFederationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SAMLFederation `json:"items"`
}

// Repository type metadata.
var (
	SAMLFederation_Kind             = "SAMLFederation"
	SAMLFederation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SAMLFederation_Kind}.String()
	SAMLFederation_KindAPIVersion   = SAMLFederation_Kind + "." + CRDGroupVersion.String()
	SAMLFederation_GroupVersionKind = CRDGroupVersion.WithKind(SAMLFederation_Kind)
)

func init() {
	SchemeBuilder.Register(&SAMLFederation{}, &SAMLFederationList{})
}
