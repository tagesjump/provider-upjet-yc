/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"fmt"

	"github.com/tagesjump/provider-upjet-yc/config/resourcemanager"
	"github.com/upbound/upjet/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"yandex_iam_service_account":                              config.NameAsIdentifier,
	"yandex_iam_service_account_key":                          config.NameAsIdentifier,
	"yandex_iam_service_account_api_key":                      config.NameAsIdentifier,
	"yandex_iam_service_account_iam_policy":                   config.NameAsIdentifier,
	"yandex_iam_service_account_iam_binding":                  config.NameAsIdentifier,
	"yandex_iam_service_account_static_access_key":            config.NameAsIdentifier,
	"yandex_iam_service_account_iam_member":                   config.NameAsIdentifier,
	"yandex_organizationmanager_group":                        config.IdentifierFromProvider,
	"yandex_organizationmanager_group_iam_member":             config.IdentifierFromProvider,
	"yandex_organizationmanager_group_membership":             config.IdentifierFromProvider,
	"yandex_organizationmanager_organization_iam_binding":     config.IdentifierFromProvider,
	"yandex_organizationmanager_organization_iam_member":      config.IdentifierFromProvider,
	"yandex_organizationmanager_saml_federation":              config.IdentifierFromProvider,
	"yandex_organizationmanager_saml_federation_user_account": config.IdentifierFromProvider,
	"yandex_resourcemanager_cloud":                            config.IdentifierFromProvider,
	"yandex_resourcemanager_cloud_iam_binding":                config.IdentifierFromProvider,
	"yandex_resourcemanager_cloud_iam_member":                 config.IdentifierFromProvider,
	"yandex_resourcemanager_folder":                           config.IdentifierFromProvider,
	"yandex_resourcemanager_folder_iam_binding":               config.IdentifierFromProvider,
	"yandex_resourcemanager_folder_iam_member":                config.IdentifierFromProvider,
	"yandex_resourcemanager_folder_iam_policy":                config.IdentifierFromProvider,
	"yandex_kms_asymmetric_encryption_key":                    config.IdentifierFromProvider,
	"yandex_kms_asymmetric_encryption_key_iam_binding":        config.IdentifierFromProvider,
	"yandex_kms_asymmetric_signature_key":                     config.IdentifierFromProvider,
	"yandex_kms_asymmetric_signature_key_iam_binding":         config.IdentifierFromProvider,
	"yandex_kms_secret_ciphertext":                            config.IdentifierFromProvider,
	"yandex_kms_symmetric_key":                                config.IdentifierFromProvider,
	"yandex_kms_symmetric_key_iam_binding":                    config.IdentifierFromProvider,
	"yandex_vpc_network":                                      config.IdentifierFromProvider,
	"yandex_vpc_route_table":                                  config.IdentifierFromProvider,
	"yandex_vpc_gateway":                                      config.IdentifierFromProvider,
	"yandex_vpc_subnet":                                       config.IdentifierFromProvider,
	"yandex_vpc_default_security_group":                       config.IdentifierFromProvider,
	"yandex_vpc_security_group":                               config.IdentifierFromProvider,
	"yandex_vpc_security_group_rule":                          config.IdentifierFromProvider,
	"yandex_vpc_address":                                      config.IdentifierFromProvider,
	"yandex_kubernetes_cluster":                               config.IdentifierFromProvider,
	"yandex_kubernetes_node_group":                            config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
		if (r.ShortGroup != "resourcemanager" && r.ShortGroup != "organizationmanager") ||
			r.Name == "yandex_resourcemanager_folder_iam_member" ||
			r.Name == "yandex_resourcemanager_folder_iam_binding" ||
			r.Name == "yandex_resourcemanager_folder_iam_policy" {
			r.References["folder_id"] = config.Reference{
				Type: fmt.Sprintf("%s.%s", resourcemanager.ApisPackagePath, "Folder"),
			}
		} else {
			r.References["folder_id"] = config.Reference{
				Type: "Folder",
			}
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
