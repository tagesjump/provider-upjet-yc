// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CaptchaInitParameters struct {

	// (List of String)
	AllowedSites []*string `json:"allowedSites,omitempty" tf:"allowed_sites,omitempty"`

	// (String)
	ChallengeType *string `json:"challengeType,omitempty" tf:"challenge_type,omitempty"`

	// (String)
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	// (String)
	Complexity *string `json:"complexity,omitempty" tf:"complexity,omitempty"`

	// (Boolean)
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) (see below for nested schema)
	OverrideVariant []OverrideVariantInitParameters `json:"overrideVariant,omitempty" tf:"override_variant,omitempty"`

	// (String)
	PreCheckType *string `json:"preCheckType,omitempty" tf:"pre_check_type,omitempty"`

	// (Block List) (see below for nested schema)
	SecurityRule []SecurityRuleInitParameters `json:"securityRule,omitempty" tf:"security_rule,omitempty"`

	// (String)
	StyleJSON *string `json:"styleJson,omitempty" tf:"style_json,omitempty"`

	// (Boolean)
	TurnOffHostnameCheck *bool `json:"turnOffHostnameCheck,omitempty" tf:"turn_off_hostname_check,omitempty"`
}

type CaptchaObservation struct {

	// (List of String)
	AllowedSites []*string `json:"allowedSites,omitempty" tf:"allowed_sites,omitempty"`

	// (String)
	ChallengeType *string `json:"challengeType,omitempty" tf:"challenge_type,omitempty"`

	// (String)
	ClientKey *string `json:"clientKey,omitempty" tf:"client_key,omitempty"`

	// (String)
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	// (String)
	Complexity *string `json:"complexity,omitempty" tf:"complexity,omitempty"`

	// (String)
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (Boolean)
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// (String)
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) (see below for nested schema)
	OverrideVariant []OverrideVariantObservation `json:"overrideVariant,omitempty" tf:"override_variant,omitempty"`

	// (String)
	PreCheckType *string `json:"preCheckType,omitempty" tf:"pre_check_type,omitempty"`

	// (Block List) (see below for nested schema)
	SecurityRule []SecurityRuleObservation `json:"securityRule,omitempty" tf:"security_rule,omitempty"`

	// (String)
	StyleJSON *string `json:"styleJson,omitempty" tf:"style_json,omitempty"`

	// (Boolean)
	Suspend *bool `json:"suspend,omitempty" tf:"suspend,omitempty"`

	// (Boolean)
	TurnOffHostnameCheck *bool `json:"turnOffHostnameCheck,omitempty" tf:"turn_off_hostname_check,omitempty"`
}

type CaptchaParameters struct {

	// (List of String)
	// +kubebuilder:validation:Optional
	AllowedSites []*string `json:"allowedSites,omitempty" tf:"allowed_sites,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ChallengeType *string `json:"challengeType,omitempty" tf:"challenge_type,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Complexity *string `json:"complexity,omitempty" tf:"complexity,omitempty"`

	// (Boolean)
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	OverrideVariant []OverrideVariantParameters `json:"overrideVariant,omitempty" tf:"override_variant,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PreCheckType *string `json:"preCheckType,omitempty" tf:"pre_check_type,omitempty"`

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	SecurityRule []SecurityRuleParameters `json:"securityRule,omitempty" tf:"security_rule,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	StyleJSON *string `json:"styleJson,omitempty" tf:"style_json,omitempty"`

	// (Boolean)
	// +kubebuilder:validation:Optional
	TurnOffHostnameCheck *bool `json:"turnOffHostnameCheck,omitempty" tf:"turn_off_hostname_check,omitempty"`
}

type ConditionInitParameters struct {

	// (Block List) (see below for nested schema)
	Headers []HeadersInitParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	Host []HostInitParameters `json:"host,omitempty" tf:"host,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	SourceIP []SourceIPInitParameters `json:"sourceIp,omitempty" tf:"source_ip,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	URI []URIInitParameters `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ConditionObservation struct {

	// (Block List) (see below for nested schema)
	Headers []HeadersObservation `json:"headers,omitempty" tf:"headers,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	Host []HostObservation `json:"host,omitempty" tf:"host,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	SourceIP []SourceIPObservation `json:"sourceIp,omitempty" tf:"source_ip,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	URI []URIObservation `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ConditionParameters struct {

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Headers []HeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Host []HostParameters `json:"host,omitempty" tf:"host,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	SourceIP []SourceIPParameters `json:"sourceIp,omitempty" tf:"source_ip,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	URI []URIParameters `json:"uri,omitempty" tf:"uri,omitempty"`
}

type GeoIPMatchInitParameters struct {

	// (List of String)
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`
}

type GeoIPMatchObservation struct {

	// (List of String)
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`
}

type GeoIPMatchParameters struct {

	// (List of String)
	// +kubebuilder:validation:Optional
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`
}

type GeoIPNotMatchInitParameters struct {

	// (List of String)
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`
}

type GeoIPNotMatchObservation struct {

	// (List of String)
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`
}

type GeoIPNotMatchParameters struct {

	// (List of String)
	// +kubebuilder:validation:Optional
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`
}

type HeadersInitParameters struct {

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Value []ValueInitParameters `json:"value,omitempty" tf:"value,omitempty"`
}

type HeadersObservation struct {

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Value []ValueObservation `json:"value,omitempty" tf:"value,omitempty"`
}

type HeadersParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Value []ValueParameters `json:"value" tf:"value,omitempty"`
}

type HostInitParameters struct {

	// (Block List) (see below for nested schema)
	Hosts []HostsInitParameters `json:"hosts,omitempty" tf:"hosts,omitempty"`
}

type HostObservation struct {

	// (Block List) (see below for nested schema)
	Hosts []HostsObservation `json:"hosts,omitempty" tf:"hosts,omitempty"`
}

type HostParameters struct {

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Hosts []HostsParameters `json:"hosts,omitempty" tf:"hosts,omitempty"`
}

type HostsInitParameters struct {

	// (String)
	ExactMatch *string `json:"exactMatch,omitempty" tf:"exact_match,omitempty"`

	// (String)
	ExactNotMatch *string `json:"exactNotMatch,omitempty" tf:"exact_not_match,omitempty"`

	// (String)
	PireRegexMatch *string `json:"pireRegexMatch,omitempty" tf:"pire_regex_match,omitempty"`

	// (String)
	PireRegexNotMatch *string `json:"pireRegexNotMatch,omitempty" tf:"pire_regex_not_match,omitempty"`

	// (String)
	PrefixMatch *string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`

	// (String)
	PrefixNotMatch *string `json:"prefixNotMatch,omitempty" tf:"prefix_not_match,omitempty"`
}

type HostsObservation struct {

	// (String)
	ExactMatch *string `json:"exactMatch,omitempty" tf:"exact_match,omitempty"`

	// (String)
	ExactNotMatch *string `json:"exactNotMatch,omitempty" tf:"exact_not_match,omitempty"`

	// (String)
	PireRegexMatch *string `json:"pireRegexMatch,omitempty" tf:"pire_regex_match,omitempty"`

	// (String)
	PireRegexNotMatch *string `json:"pireRegexNotMatch,omitempty" tf:"pire_regex_not_match,omitempty"`

	// (String)
	PrefixMatch *string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`

	// (String)
	PrefixNotMatch *string `json:"prefixNotMatch,omitempty" tf:"prefix_not_match,omitempty"`
}

type HostsParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	ExactMatch *string `json:"exactMatch,omitempty" tf:"exact_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ExactNotMatch *string `json:"exactNotMatch,omitempty" tf:"exact_not_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PireRegexMatch *string `json:"pireRegexMatch,omitempty" tf:"pire_regex_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PireRegexNotMatch *string `json:"pireRegexNotMatch,omitempty" tf:"pire_regex_not_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PrefixMatch *string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PrefixNotMatch *string `json:"prefixNotMatch,omitempty" tf:"prefix_not_match,omitempty"`
}

type IPRangesMatchInitParameters struct {

	// (List of String)
	IPRanges []*string `json:"ipRanges,omitempty" tf:"ip_ranges,omitempty"`
}

type IPRangesMatchObservation struct {

	// (List of String)
	IPRanges []*string `json:"ipRanges,omitempty" tf:"ip_ranges,omitempty"`
}

type IPRangesMatchParameters struct {

	// (List of String)
	// +kubebuilder:validation:Optional
	IPRanges []*string `json:"ipRanges,omitempty" tf:"ip_ranges,omitempty"`
}

type IPRangesNotMatchInitParameters struct {

	// (List of String)
	IPRanges []*string `json:"ipRanges,omitempty" tf:"ip_ranges,omitempty"`
}

type IPRangesNotMatchObservation struct {

	// (List of String)
	IPRanges []*string `json:"ipRanges,omitempty" tf:"ip_ranges,omitempty"`
}

type IPRangesNotMatchParameters struct {

	// (List of String)
	// +kubebuilder:validation:Optional
	IPRanges []*string `json:"ipRanges,omitempty" tf:"ip_ranges,omitempty"`
}

type OverrideVariantInitParameters struct {

	// (String)
	ChallengeType *string `json:"challengeType,omitempty" tf:"challenge_type,omitempty"`

	// (String)
	Complexity *string `json:"complexity,omitempty" tf:"complexity,omitempty"`

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	PreCheckType *string `json:"preCheckType,omitempty" tf:"pre_check_type,omitempty"`

	// (String)
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type OverrideVariantObservation struct {

	// (String)
	ChallengeType *string `json:"challengeType,omitempty" tf:"challenge_type,omitempty"`

	// (String)
	Complexity *string `json:"complexity,omitempty" tf:"complexity,omitempty"`

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	PreCheckType *string `json:"preCheckType,omitempty" tf:"pre_check_type,omitempty"`

	// (String)
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type OverrideVariantParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	ChallengeType *string `json:"challengeType,omitempty" tf:"challenge_type,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Complexity *string `json:"complexity,omitempty" tf:"complexity,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PreCheckType *string `json:"preCheckType,omitempty" tf:"pre_check_type,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type PathInitParameters struct {

	// (String)
	ExactMatch *string `json:"exactMatch,omitempty" tf:"exact_match,omitempty"`

	// (String)
	ExactNotMatch *string `json:"exactNotMatch,omitempty" tf:"exact_not_match,omitempty"`

	// (String)
	PireRegexMatch *string `json:"pireRegexMatch,omitempty" tf:"pire_regex_match,omitempty"`

	// (String)
	PireRegexNotMatch *string `json:"pireRegexNotMatch,omitempty" tf:"pire_regex_not_match,omitempty"`

	// (String)
	PrefixMatch *string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`

	// (String)
	PrefixNotMatch *string `json:"prefixNotMatch,omitempty" tf:"prefix_not_match,omitempty"`
}

type PathObservation struct {

	// (String)
	ExactMatch *string `json:"exactMatch,omitempty" tf:"exact_match,omitempty"`

	// (String)
	ExactNotMatch *string `json:"exactNotMatch,omitempty" tf:"exact_not_match,omitempty"`

	// (String)
	PireRegexMatch *string `json:"pireRegexMatch,omitempty" tf:"pire_regex_match,omitempty"`

	// (String)
	PireRegexNotMatch *string `json:"pireRegexNotMatch,omitempty" tf:"pire_regex_not_match,omitempty"`

	// (String)
	PrefixMatch *string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`

	// (String)
	PrefixNotMatch *string `json:"prefixNotMatch,omitempty" tf:"prefix_not_match,omitempty"`
}

type PathParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	ExactMatch *string `json:"exactMatch,omitempty" tf:"exact_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ExactNotMatch *string `json:"exactNotMatch,omitempty" tf:"exact_not_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PireRegexMatch *string `json:"pireRegexMatch,omitempty" tf:"pire_regex_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PireRegexNotMatch *string `json:"pireRegexNotMatch,omitempty" tf:"pire_regex_not_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PrefixMatch *string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PrefixNotMatch *string `json:"prefixNotMatch,omitempty" tf:"prefix_not_match,omitempty"`
}

type QueriesInitParameters struct {

	// (String)
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Value []QueriesValueInitParameters `json:"value,omitempty" tf:"value,omitempty"`
}

type QueriesObservation struct {

	// (String)
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Value []QueriesValueObservation `json:"value,omitempty" tf:"value,omitempty"`
}

type QueriesParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Value []QueriesValueParameters `json:"value" tf:"value,omitempty"`
}

type QueriesValueInitParameters struct {

	// (String)
	ExactMatch *string `json:"exactMatch,omitempty" tf:"exact_match,omitempty"`

	// (String)
	ExactNotMatch *string `json:"exactNotMatch,omitempty" tf:"exact_not_match,omitempty"`

	// (String)
	PireRegexMatch *string `json:"pireRegexMatch,omitempty" tf:"pire_regex_match,omitempty"`

	// (String)
	PireRegexNotMatch *string `json:"pireRegexNotMatch,omitempty" tf:"pire_regex_not_match,omitempty"`

	// (String)
	PrefixMatch *string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`

	// (String)
	PrefixNotMatch *string `json:"prefixNotMatch,omitempty" tf:"prefix_not_match,omitempty"`
}

type QueriesValueObservation struct {

	// (String)
	ExactMatch *string `json:"exactMatch,omitempty" tf:"exact_match,omitempty"`

	// (String)
	ExactNotMatch *string `json:"exactNotMatch,omitempty" tf:"exact_not_match,omitempty"`

	// (String)
	PireRegexMatch *string `json:"pireRegexMatch,omitempty" tf:"pire_regex_match,omitempty"`

	// (String)
	PireRegexNotMatch *string `json:"pireRegexNotMatch,omitempty" tf:"pire_regex_not_match,omitempty"`

	// (String)
	PrefixMatch *string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`

	// (String)
	PrefixNotMatch *string `json:"prefixNotMatch,omitempty" tf:"prefix_not_match,omitempty"`
}

type QueriesValueParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	ExactMatch *string `json:"exactMatch,omitempty" tf:"exact_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ExactNotMatch *string `json:"exactNotMatch,omitempty" tf:"exact_not_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PireRegexMatch *string `json:"pireRegexMatch,omitempty" tf:"pire_regex_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PireRegexNotMatch *string `json:"pireRegexNotMatch,omitempty" tf:"pire_regex_not_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PrefixMatch *string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PrefixNotMatch *string `json:"prefixNotMatch,omitempty" tf:"prefix_not_match,omitempty"`
}

type SecurityRuleInitParameters struct {

	// (Block List, Max: 1) (see below for nested schema)
	Condition []ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	OverrideVariantUUID *string `json:"overrideVariantUuid,omitempty" tf:"override_variant_uuid,omitempty"`

	// (Number)
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type SecurityRuleObservation struct {

	// (Block List, Max: 1) (see below for nested schema)
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	OverrideVariantUUID *string `json:"overrideVariantUuid,omitempty" tf:"override_variant_uuid,omitempty"`

	// (Number)
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type SecurityRuleParameters struct {

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	OverrideVariantUUID *string `json:"overrideVariantUuid,omitempty" tf:"override_variant_uuid,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type SourceIPInitParameters struct {

	// (Block List, Max: 1) (see below for nested schema)
	GeoIPMatch []GeoIPMatchInitParameters `json:"geoIpMatch,omitempty" tf:"geo_ip_match,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	GeoIPNotMatch []GeoIPNotMatchInitParameters `json:"geoIpNotMatch,omitempty" tf:"geo_ip_not_match,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	IPRangesMatch []IPRangesMatchInitParameters `json:"ipRangesMatch,omitempty" tf:"ip_ranges_match,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	IPRangesNotMatch []IPRangesNotMatchInitParameters `json:"ipRangesNotMatch,omitempty" tf:"ip_ranges_not_match,omitempty"`
}

type SourceIPObservation struct {

	// (Block List, Max: 1) (see below for nested schema)
	GeoIPMatch []GeoIPMatchObservation `json:"geoIpMatch,omitempty" tf:"geo_ip_match,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	GeoIPNotMatch []GeoIPNotMatchObservation `json:"geoIpNotMatch,omitempty" tf:"geo_ip_not_match,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	IPRangesMatch []IPRangesMatchObservation `json:"ipRangesMatch,omitempty" tf:"ip_ranges_match,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	IPRangesNotMatch []IPRangesNotMatchObservation `json:"ipRangesNotMatch,omitempty" tf:"ip_ranges_not_match,omitempty"`
}

type SourceIPParameters struct {

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	GeoIPMatch []GeoIPMatchParameters `json:"geoIpMatch,omitempty" tf:"geo_ip_match,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	GeoIPNotMatch []GeoIPNotMatchParameters `json:"geoIpNotMatch,omitempty" tf:"geo_ip_not_match,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	IPRangesMatch []IPRangesMatchParameters `json:"ipRangesMatch,omitempty" tf:"ip_ranges_match,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	IPRangesNotMatch []IPRangesNotMatchParameters `json:"ipRangesNotMatch,omitempty" tf:"ip_ranges_not_match,omitempty"`
}

type URIInitParameters struct {

	// (Block List, Max: 1) (see below for nested schema)
	Path []PathInitParameters `json:"path,omitempty" tf:"path,omitempty"`

	// (Block List) (see below for nested schema)
	Queries []QueriesInitParameters `json:"queries,omitempty" tf:"queries,omitempty"`
}

type URIObservation struct {

	// (Block List, Max: 1) (see below for nested schema)
	Path []PathObservation `json:"path,omitempty" tf:"path,omitempty"`

	// (Block List) (see below for nested schema)
	Queries []QueriesObservation `json:"queries,omitempty" tf:"queries,omitempty"`
}

type URIParameters struct {

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Path []PathParameters `json:"path,omitempty" tf:"path,omitempty"`

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Queries []QueriesParameters `json:"queries,omitempty" tf:"queries,omitempty"`
}

type ValueInitParameters struct {

	// (String)
	ExactMatch *string `json:"exactMatch,omitempty" tf:"exact_match,omitempty"`

	// (String)
	ExactNotMatch *string `json:"exactNotMatch,omitempty" tf:"exact_not_match,omitempty"`

	// (String)
	PireRegexMatch *string `json:"pireRegexMatch,omitempty" tf:"pire_regex_match,omitempty"`

	// (String)
	PireRegexNotMatch *string `json:"pireRegexNotMatch,omitempty" tf:"pire_regex_not_match,omitempty"`

	// (String)
	PrefixMatch *string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`

	// (String)
	PrefixNotMatch *string `json:"prefixNotMatch,omitempty" tf:"prefix_not_match,omitempty"`
}

type ValueObservation struct {

	// (String)
	ExactMatch *string `json:"exactMatch,omitempty" tf:"exact_match,omitempty"`

	// (String)
	ExactNotMatch *string `json:"exactNotMatch,omitempty" tf:"exact_not_match,omitempty"`

	// (String)
	PireRegexMatch *string `json:"pireRegexMatch,omitempty" tf:"pire_regex_match,omitempty"`

	// (String)
	PireRegexNotMatch *string `json:"pireRegexNotMatch,omitempty" tf:"pire_regex_not_match,omitempty"`

	// (String)
	PrefixMatch *string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`

	// (String)
	PrefixNotMatch *string `json:"prefixNotMatch,omitempty" tf:"prefix_not_match,omitempty"`
}

type ValueParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	ExactMatch *string `json:"exactMatch,omitempty" tf:"exact_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ExactNotMatch *string `json:"exactNotMatch,omitempty" tf:"exact_not_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PireRegexMatch *string `json:"pireRegexMatch,omitempty" tf:"pire_regex_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PireRegexNotMatch *string `json:"pireRegexNotMatch,omitempty" tf:"pire_regex_not_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PrefixMatch *string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PrefixNotMatch *string `json:"prefixNotMatch,omitempty" tf:"prefix_not_match,omitempty"`
}

// CaptchaSpec defines the desired state of Captcha
type CaptchaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CaptchaParameters `json:"forProvider"`
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
	InitProvider CaptchaInitParameters `json:"initProvider,omitempty"`
}

// CaptchaStatus defines the observed state of Captcha.
type CaptchaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CaptchaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Captcha is the Schema for the Captchas API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Captcha struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CaptchaSpec   `json:"spec"`
	Status            CaptchaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CaptchaList contains a list of Captchas
type CaptchaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Captcha `json:"items"`
}

// Repository type metadata.
var (
	Captcha_Kind             = "Captcha"
	Captcha_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Captcha_Kind}.String()
	Captcha_KindAPIVersion   = Captcha_Kind + "." + CRDGroupVersion.String()
	Captcha_GroupVersionKind = CRDGroupVersion.WithKind(Captcha_Kind)
)

func init() {
	SchemeBuilder.Register(&Captcha{}, &CaptchaList{})
}
