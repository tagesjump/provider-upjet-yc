package config

import (
	"fmt"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/tagesjump/provider-upjet-yc/config/cluster/resourcemanager"
)

// TerraformPluginSDKExternalNameConfigs contains all external name configurations
// belonging to Terraform Plugin SDKv2 resources to be reconciled
// under the no-fork architecture for this provider.
var TerraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{
	"yandex_alb_backend_group": config.IdentifierFromProvider,
	"yandex_alb_http_router": config.IdentifierFromProvider,
	"yandex_alb_load_balancer": config.IdentifierFromProvider,
	"yandex_alb_target_group": config.IdentifierFromProvider,
	"yandex_alb_virtual_host": config.IdentifierFromProvider,

	"yandex_api_gateway": config.IdentifierFromProvider,

	"yandex_audit_trails_trail": config.IdentifierFromProvider,

	"yandex_backup_policy": config.IdentifierFromProvider,
	"yandex_backup_policy_bindings": config.IdentifierFromProvider,

	"yandex_cdn_origin_group": config.IdentifierFromProvider,
	"yandex_cdn_resource": config.IdentifierFromProvider,

	"yandex_cm_certificate": config.IdentifierFromProvider,

	"yandex_compute_disk": config.IdentifierFromProvider,
	"yandex_compute_image": config.IdentifierFromProvider,
	"yandex_compute_instance": config.IdentifierFromProvider,
	"yandex_compute_instance_group": config.IdentifierFromProvider,
	"yandex_compute_placement_group": config.IdentifierFromProvider,
	"yandex_compute_snapshot": config.IdentifierFromProvider,
	"yandex_compute_snapshot_schedule": config.IdentifierFromProvider,

	"yandex_container_registry_ip_permission": config.IdentifierFromProvider,
	"yandex_container_repository_lifecycle_policy": config.IdentifierFromProvider,

	"yandex_dataproc_cluster": config.IdentifierFromProvider,


	"yandex_datatransfer_transfer": config.IdentifierFromProvider,

	"yandex_dns_recordset": config.IdentifierFromProvider,
	"yandex_dns_zone": config.IdentifierFromProvider,

	"yandex_function": config.IdentifierFromProvider,
	"yandex_function_scaling_policy": config.IdentifierFromProvider,
	"yandex_function_trigger": config.IdentifierFromProvider,

	"yandex_iam_oauth_client_secret": config.IdentifierFromProvider,
	"yandex_iam_service_account_api_key": config.IdentifierFromProvider,
	"yandex_iam_service_account_iam_policy": config.IdentifierFromProvider,
	"yandex_iam_service_account_key": config.IdentifierFromProvider,
	"yandex_iam_service_account_static_access_key": config.IdentifierFromProvider,

	"yandex_iot_core_broker": config.IdentifierFromProvider,
	"yandex_iot_core_device": config.IdentifierFromProvider,
	"yandex_iot_core_registry": config.IdentifierFromProvider,

	"yandex_kms_secret_ciphertext": config.IdentifierFromProvider,

	"yandex_kubernetes_cluster": config.IdentifierFromProvider,
	"yandex_kubernetes_node_group": config.IdentifierFromProvider,

	"yandex_lb_network_load_balancer": config.IdentifierFromProvider,

	"yandex_loadtesting_agent": config.IdentifierFromProvider,

	"yandex_lockbox_secret": config.IdentifierFromProvider,
	"yandex_lockbox_secret_version": config.IdentifierFromProvider,
	"yandex_lockbox_secret_version_hashed": config.IdentifierFromProvider,

	"yandex_mdb_clickhouse_cluster": config.IdentifierFromProvider,
	"yandex_mdb_greenplum_cluster": config.IdentifierFromProvider,
	"yandex_mdb_kafka_cluster": config.IdentifierFromProvider,
	"yandex_mdb_kafka_connector": config.IdentifierFromProvider,
	"yandex_mdb_kafka_topic": config.IdentifierFromProvider,
	"yandex_mdb_kafka_user": config.IdentifierFromProvider,
	"yandex_mdb_mongodb_cluster": config.IdentifierFromProvider,
	"yandex_mdb_mysql_cluster": config.IdentifierFromProvider,
	"yandex_mdb_mysql_database": config.IdentifierFromProvider,
	"yandex_mdb_mysql_user": config.IdentifierFromProvider,
	"yandex_mdb_postgresql_cluster": config.IdentifierFromProvider,
	"yandex_mdb_postgresql_database": config.IdentifierFromProvider,
	"yandex_mdb_postgresql_user": config.IdentifierFromProvider,
	"yandex_mdb_redis_cluster": config.IdentifierFromProvider,
	"yandex_mdb_sqlserver_cluster": config.IdentifierFromProvider,

	"yandex_message_queue": config.IdentifierFromProvider,

	"yandex_monitoring_dashboard": config.IdentifierFromProvider,

	"yandex_organizationmanager_group_mapping": config.IdentifierFromProvider,
	"yandex_organizationmanager_group_mapping_item": config.IdentifierFromProvider,
	"yandex_organizationmanager_group_membership": config.IdentifierFromProvider,
	"yandex_organizationmanager_os_login_settings": config.IdentifierFromProvider,
	"yandex_organizationmanager_saml_federation": config.IdentifierFromProvider,
	"yandex_organizationmanager_saml_federation_user_account": config.IdentifierFromProvider,

	"yandex_resourcemanager_folder_iam_policy": config.IdentifierFromProvider,

	"yandex_serverless_container": config.IdentifierFromProvider,
	"yandex_serverless_eventrouter_connector": config.IdentifierFromProvider,
	"yandex_serverless_eventrouter_rule": config.IdentifierFromProvider,

	"yandex_smartcaptcha_captcha": config.IdentifierFromProvider,

	"yandex_storage_bucket": config.IdentifierFromProvider,
	"yandex_storage_object": config.IdentifierFromProvider,

	"yandex_sws_advanced_rate_limiter_profile": config.IdentifierFromProvider,
	"yandex_sws_security_profile": config.IdentifierFromProvider,
	"yandex_sws_waf_profile": config.IdentifierFromProvider,

	"yandex_vpc_address": config.IdentifierFromProvider,
	"yandex_vpc_default_security_group": config.IdentifierFromProvider,
	"yandex_vpc_gateway": config.IdentifierFromProvider,
	"yandex_vpc_network": config.IdentifierFromProvider,
	"yandex_vpc_private_endpoint": config.IdentifierFromProvider,
	"yandex_vpc_route_table": config.IdentifierFromProvider,
	"yandex_vpc_security_group": config.IdentifierFromProvider,
	"yandex_vpc_subnet": config.IdentifierFromProvider,

	"yandex_ydb_database_dedicated": config.IdentifierFromProvider,
	"yandex_ydb_database_serverless": config.IdentifierFromProvider,
	"yandex_ydb_table": config.IdentifierFromProvider,
	"yandex_ydb_table_changefeed": config.IdentifierFromProvider,
	"yandex_ydb_table_index": config.IdentifierFromProvider,
	"yandex_ydb_topic": config.IdentifierFromProvider,
}

// TerraformPluginFrameworkExternalNameConfigs contains all external
// name configurations belonging to Terraform Plugin Framework
// resources to be reconciled under the no-fork architecture for this
// provider.
var TerraformPluginFrameworkExternalNameConfigs = map[string]config.ExternalName{
	"yandex_airflow_cluster": config.IdentifierFromProvider,

	"yandex_billing_cloud_binding": config.IdentifierFromProvider,

	"yandex_cloud_desktops_desktop": config.IdentifierFromProvider,
	"yandex_cloud_desktops_desktop_group": config.IdentifierFromProvider,

	"yandex_cloudregistry_registry": config.IdentifierFromProvider,
	"yandex_cloudregistry_registry_iam_binding": config.IdentifierFromProvider,
	"yandex_cloudregistry_registry_ip_permission": config.IdentifierFromProvider,

	"yandex_cm_certificate_iam_binding": config.IdentifierFromProvider,
	"yandex_cm_certificate_iam_member": config.IdentifierFromProvider,

	"yandex_compute_disk_iam_binding": config.IdentifierFromProvider,
	"yandex_compute_disk_placement_group": config.IdentifierFromProvider,
	"yandex_compute_disk_placement_group_iam_binding": config.IdentifierFromProvider,
	"yandex_compute_filesystem": config.IdentifierFromProvider,
	"yandex_compute_filesystem_iam_binding": config.IdentifierFromProvider,
	"yandex_compute_gpu_cluster": config.IdentifierFromProvider,
	"yandex_compute_gpu_cluster_iam_binding": config.IdentifierFromProvider,
	"yandex_compute_image_iam_binding": config.IdentifierFromProvider,
	"yandex_compute_instance_iam_binding": config.IdentifierFromProvider,
	"yandex_compute_placement_group_iam_binding": config.IdentifierFromProvider,
	"yandex_compute_snapshot_iam_binding": config.IdentifierFromProvider,
	"yandex_compute_snapshot_schedule_iam_binding": config.IdentifierFromProvider,

	"yandex_connectionmanager_connection": config.IdentifierFromProvider,

	"yandex_container_registry": config.IdentifierFromProvider,
	"yandex_container_registry_iam_binding": config.IdentifierFromProvider,
	"yandex_container_repository": config.IdentifierFromProvider,
	"yandex_container_repository_iam_binding": config.IdentifierFromProvider,

	"yandex_datatransfer_endpoint": config.IdentifierFromProvider,

	"yandex_datasphere_community": config.IdentifierFromProvider,
	"yandex_datasphere_community_iam_binding": config.IdentifierFromProvider,
	"yandex_datasphere_project": config.IdentifierFromProvider,
	"yandex_datasphere_project_iam_binding": config.IdentifierFromProvider,

	"yandex_dns_zone_iam_binding": config.IdentifierFromProvider,

	"yandex_function_iam_binding": config.IdentifierFromProvider,

	"yandex_gitlab_instance": config.IdentifierFromProvider,

	"yandex_iam_oauth_client": config.IdentifierFromProvider,
	"yandex_iam_service_account": config.IdentifierFromProvider,
	"yandex_iam_service_account_iam_binding": config.IdentifierFromProvider,
	"yandex_iam_service_account_iam_member": config.IdentifierFromProvider,
	"yandex_iam_workload_identity_federated_credential": config.IdentifierFromProvider,
	"yandex_iam_workload_identity_oidc_federation": config.IdentifierFromProvider,
	"yandex_iam_workload_identity_oidc_federation_iam_binding": config.IdentifierFromProvider,

	"yandex_kms_asymmetric_encryption_key": config.IdentifierFromProvider,
	"yandex_kms_asymmetric_encryption_key_iam_binding": config.IdentifierFromProvider,
	"yandex_kms_asymmetric_encryption_key_iam_member": config.IdentifierFromProvider,
	"yandex_kms_asymmetric_signature_key": config.IdentifierFromProvider,
	"yandex_kms_asymmetric_signature_key_iam_binding": config.IdentifierFromProvider,
	"yandex_kms_asymmetric_signature_key_iam_member": config.IdentifierFromProvider,
	"yandex_kms_symmetric_key": config.IdentifierFromProvider,
	"yandex_kms_symmetric_key_iam_binding": config.IdentifierFromProvider,
	"yandex_kms_symmetric_key_iam_member": config.IdentifierFromProvider,

	"yandex_kubernetes_cluster_iam_binding": config.IdentifierFromProvider,
	"yandex_kubernetes_cluster_iam_member": config.IdentifierFromProvider,
	"yandex_kubernetes_marketplace_helm_release": config.IdentifierFromProvider,

	"yandex_lb_target_group": config.IdentifierFromProvider,

	"yandex_lockbox_secret_iam_binding": config.IdentifierFromProvider,
	"yandex_lockbox_secret_iam_member": config.IdentifierFromProvider,

	"yandex_logging_group": config.IdentifierFromProvider,

	"yandex_mdb_clickhouse_cluster_v2": config.IdentifierFromProvider,
	"yandex_mdb_clickhouse_database": config.IdentifierFromProvider,
	"yandex_mdb_clickhouse_user": config.IdentifierFromProvider,
	"yandex_mdb_greenplum_cluster_v2": config.IdentifierFromProvider,
	"yandex_mdb_greenplum_resource_group": config.IdentifierFromProvider,
	"yandex_mdb_greenplum_user": config.IdentifierFromProvider,
	"yandex_mdb_mongodb_database": config.IdentifierFromProvider,
	"yandex_mdb_mongodb_user": config.IdentifierFromProvider,
	"yandex_mdb_mysql_cluster_v2": config.IdentifierFromProvider,
	"yandex_mdb_opensearch_cluster": config.IdentifierFromProvider,
	"yandex_mdb_postgresql_cluster_v2": config.IdentifierFromProvider,
	"yandex_mdb_redis_cluster_v2": config.IdentifierFromProvider,
	"yandex_mdb_redis_user": config.IdentifierFromProvider,
	"yandex_mdb_sharded_postgresql_cluster": config.IdentifierFromProvider,
	"yandex_mdb_sharded_postgresql_database": config.IdentifierFromProvider,
	"yandex_mdb_sharded_postgresql_shard": config.IdentifierFromProvider,
	"yandex_mdb_sharded_postgresql_user": config.IdentifierFromProvider,

	"yandex_metastore_cluster": config.IdentifierFromProvider,

	"yandex_organizationmanager_group": config.IdentifierFromProvider,
	"yandex_organizationmanager_group_iam_member": config.IdentifierFromProvider,
	"yandex_organizationmanager_idp_application_oauth_application": config.IdentifierFromProvider,
	"yandex_organizationmanager_idp_application_oauth_application_assignment": config.IdentifierFromProvider,
	"yandex_organizationmanager_idp_application_saml_application": config.IdentifierFromProvider,
	"yandex_organizationmanager_idp_application_saml_application_assignment": config.IdentifierFromProvider,
	"yandex_organizationmanager_idp_application_saml_signature_certificate": config.IdentifierFromProvider,
	"yandex_organizationmanager_idp_user": config.IdentifierFromProvider,
	"yandex_organizationmanager_idp_userpool": config.IdentifierFromProvider,
	"yandex_organizationmanager_idp_userpool_domain": config.IdentifierFromProvider,
	"yandex_organizationmanager_mfa_enforcement": config.IdentifierFromProvider,
	"yandex_organizationmanager_mfa_enforcement_audience": config.IdentifierFromProvider,
	"yandex_organizationmanager_organization_iam_binding": config.IdentifierFromProvider,
	"yandex_organizationmanager_organization_iam_member": config.IdentifierFromProvider,
	"yandex_organizationmanager_user_ssh_key": config.IdentifierFromProvider,

	"yandex_resourcemanager_cloud": config.IdentifierFromProvider,
	"yandex_resourcemanager_cloud_iam_binding": config.IdentifierFromProvider,
	"yandex_resourcemanager_cloud_iam_member": config.IdentifierFromProvider,
	"yandex_resourcemanager_folder": config.IdentifierFromProvider,
	"yandex_resourcemanager_folder_iam_binding": config.IdentifierFromProvider,
	"yandex_resourcemanager_folder_iam_member": config.IdentifierFromProvider,

	"yandex_serverless_container_iam_binding": config.IdentifierFromProvider,
	"yandex_serverless_eventrouter_bus": config.IdentifierFromProvider,

	"yandex_spark_cluster": config.IdentifierFromProvider,

	"yandex_storage_bucket_grant": config.IdentifierFromProvider,
	"yandex_storage_bucket_iam_binding": config.IdentifierFromProvider,
	"yandex_storage_bucket_policy": config.IdentifierFromProvider,

	"yandex_trino_access_control": config.IdentifierFromProvider,
	"yandex_trino_catalog": config.IdentifierFromProvider,
	"yandex_trino_cluster": config.IdentifierFromProvider,

	"yandex_vpc_security_group_rule": config.IdentifierFromProvider,

	"yandex_ydb_database_iam_binding": config.IdentifierFromProvider,

	"yandex_yq_monitoring_connection": config.IdentifierFromProvider,
	"yandex_yq_object_storage_binding": config.IdentifierFromProvider,
	"yandex_yq_object_storage_connection": config.IdentifierFromProvider,
	"yandex_yq_ydb_connection": config.IdentifierFromProvider,
	"yandex_yq_yds_binding": config.IdentifierFromProvider,
	"yandex_yq_yds_connection": config.IdentifierFromProvider,

	"yandex_ytsaurus_cluster": config.IdentifierFromProvider,
}

// cliReconciledExternalNameConfigs contains all external name configurations
// belonging to Terraform resources to be reconciled under the CLI-based
// architecture for this provider.
var cliReconciledExternalNameConfigs = map[string]config.ExternalName{}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := TerraformPluginSDKExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		} else if e, ok := TerraformPluginFrameworkExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		} else if e, ok := cliReconciledExternalNameConfigs[r.Name]; ok {
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
