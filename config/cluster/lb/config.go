package lb

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/tagesjump/provider-upjet-yc/config/cluster/vpc"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/cluster/lb/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/cluster/lb"
)

// Configure adds configurations for lb group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_lb_target_group", func(r *ujconfig.Resource) {
		r.References["target.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
	})
	p.AddResourceConfigurator("yandex_lb_network_load_balancer", func(r *ujconfig.Resource) {
		r.References["attached_target_group.target_group_id"] = ujconfig.Reference{
			Type: "TargetGroup",
		}
		r.References["listener.internal_address_spec.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
	})
}
