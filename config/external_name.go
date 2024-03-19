package config

import (
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/tagesjump/provider-upjet-yc/config/resourcemanager"
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
	"yandex_iam_service_account_static_access_key":            config.IdentifierFromProvider,
	"yandex_iam_service_account_iam_member":                   config.NameAsIdentifier,
	"yandex_organizationmanager_group":                        config.IdentifierFromProvider,
	"yandex_organizationmanager_group_iam_member":             config.IdentifierFromProvider,
	"yandex_organizationmanager_group_membership":             config.IdentifierFromProvider,
	"yandex_organizationmanager_organization_iam_binding":     config.IdentifierFromProvider,
	"yandex_organizationmanager_organization_iam_member":      config.IdentifierFromProvider,
	"yandex_organizationmanager_os_login_settings":            config.IdentifierFromProvider,
	"yandex_organizationmanager_user_ssh_key":                 config.IdentifierFromProvider,
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
	"yandex_mdb_clickhouse_cluster":                           config.IdentifierFromProvider,
	"yandex_mdb_elasticsearch_cluster":                        config.IdentifierFromProvider,
	"yandex_mdb_greenplum_cluster":                            config.IdentifierFromProvider,
	"yandex_mdb_kafka_cluster":                                config.IdentifierFromProvider,
	"yandex_mdb_kafka_connector":                              config.IdentifierFromProvider,
	"yandex_mdb_kafka_topic":                                  config.IdentifierFromProvider,
	"yandex_mdb_kafka_user":                                   config.IdentifierFromProvider,
	"yandex_mdb_mongodb_cluster":                              config.IdentifierFromProvider,
	"yandex_mdb_mongodb_user":                                 config.IdentifierFromProvider,
	"yandex_mdb_mongodb_database":                             config.IdentifierFromProvider,
	"yandex_mdb_mysql_cluster":                                config.IdentifierFromProvider,
	"yandex_mdb_mysql_database":                               config.IdentifierFromProvider,
	"yandex_mdb_mysql_user":                                   config.IdentifierFromProvider,
	"yandex_mdb_postgresql_cluster":                           config.IdentifierFromProvider,
	"yandex_mdb_postgresql_database":                          config.IdentifierFromProvider,
	"yandex_mdb_postgresql_user":                              config.IdentifierFromProvider,
	"yandex_mdb_redis_cluster":                                config.IdentifierFromProvider,
	"yandex_mdb_sqlserver_cluster":                            config.IdentifierFromProvider,
	"yandex_compute_disk":                                     config.IdentifierFromProvider,
	"yandex_compute_disk_placement_group":                     config.IdentifierFromProvider,
	"yandex_compute_filesystem":                               config.IdentifierFromProvider,
	"yandex_compute_gpu_cluster":                              config.IdentifierFromProvider,
	"yandex_compute_image":                                    config.IdentifierFromProvider,
	"yandex_compute_instance":                                 config.IdentifierFromProvider,
	"yandex_compute_instance_group":                           config.IdentifierFromProvider,
	"yandex_compute_instance_migrate":                         config.IdentifierFromProvider,
	"yandex_compute_placement_group":                          config.IdentifierFromProvider,
	"yandex_compute_snapshot":                                 config.IdentifierFromProvider,
	"yandex_compute_snapshot_schedule":                        config.IdentifierFromProvider,
	"yandex_container_registry":                               config.IdentifierFromProvider,
	"yandex_container_registry_iam_binding":                   config.IdentifierFromProvider,
	"yandex_container_registry_ip_permission":                 config.IdentifierFromProvider,
	"yandex_container_repository":                             config.IdentifierFromProvider,
	"yandex_container_repository_iam_binding":                 config.IdentifierFromProvider,
	"yandex_container_repository_lifecycle_policy":            config.IdentifierFromProvider,
	"yandex_alb_backend_group":                                config.IdentifierFromProvider,
	"yandex_alb_http_router":                                  config.IdentifierFromProvider,
	"yandex_alb_load_balancer":                                config.IdentifierFromProvider,
	"yandex_alb_target_group":                                 config.IdentifierFromProvider,
	"yandex_alb_virtual_host":                                 config.IdentifierFromProvider,
	"yandex_storage_bucket":                                   config.IdentifierFromProvider,
	"yandex_storage_object":                                   config.IdentifierFromProvider,
	"yandex_lb_network_load_balancer":                         config.IdentifierFromProvider,
	"yandex_lb_target_group":                                  config.IdentifierFromProvider,
	"yandex_dns_recordset":                                    config.IdentifierFromProvider,
	"yandex_dns_zone":                                         config.IdentifierFromProvider,
	"yandex_dns_zone_iam_binding":                             config.IdentifierFromProvider,
	"yandex_logging_group":                                    config.IdentifierFromProvider,
	"yandex_message_queue":                                    config.IdentifierFromProvider,
	"yandex_ydb_database_dedicated":                           config.IdentifierFromProvider,
	"yandex_ydb_database_iam_binding":                         config.IdentifierFromProvider,
	"yandex_ydb_database_serverless":                          config.IdentifierFromProvider,
	"yandex_ydb_table":                                        config.IdentifierFromProvider,
	"yandex_ydb_table_changefeed":                             config.IdentifierFromProvider,
	"yandex_ydb_table_index":                                  config.IdentifierFromProvider,
	"yandex_ydb_topic":                                        config.IdentifierFromProvider,
	"yandex_datatransfer_endpoint":                            config.IdentifierFromProvider,
	"yandex_datatransfer_transfer":                            config.IdentifierFromProvider,
	"yandex_lockbox_secret":                                   config.IdentifierFromProvider,
	"yandex_lockbox_secret_iam_binding":                       config.IdentifierFromProvider,
	"yandex_lockbox_secret_version":                           config.IdentifierFromProvider,
	"yandex_monitoring_dashboard":                             config.IdentifierFromProvider,
	"yandex_loadtesting_agent":                                config.IdentifierFromProvider,
	"yandex_mdb_opensearch_cluster":                           config.IdentifierFromProvider,
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
