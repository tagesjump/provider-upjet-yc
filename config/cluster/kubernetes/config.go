package kubernetes

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/tagesjump/provider-upjet-yc/config/cluster/kms"
	"github.com/tagesjump/provider-upjet-yc/config/cluster/vpc"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/cluster/kubernetes/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/cluster/kubernetes"
)

// Configure adds configurations for vpc group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_kubernetes_cluster", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["master.security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["node_service_account_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["service_account_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["master.regional.location.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["master.master_location.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["master.zonal.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.LateInitializer = ujconfig.LateInitializer{
			IgnoredFields: []string{"master.regional.location"},
		}
		r.References["kms_provider.key_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", kms.ApisPackagePath, "SymmetricKey"),
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("yandex_kubernetes_node_group", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			Type: "Cluster",
		}
		r.References["allocation_policy.location.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["instance_template.network_interface.subnet_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["instance_template.network_interface.security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
	})
}
