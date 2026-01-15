package vpc

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/namespaced/vpc/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/namespaced/vpc"
)

// Configure adds configurations for vpc group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_vpc_subnet", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: "Network",
		}
	})
	p.AddResourceConfigurator("yandex_vpc_default_security_group", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: "Network",
		}
	})
	p.AddResourceConfigurator("yandex_vpc_security_group", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: "Network",
		}
	})
	p.AddResourceConfigurator("yandex_vpc_security_group_rule", func(r *ujconfig.Resource) {
		r.References["security_group_binding"] = ujconfig.Reference{
			Type: "SecurityGroup",
		}
	})
	p.AddResourceConfigurator("yandex_vpc_route_table", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: "Network",
		}
		r.References["static_route.gateway_id"] = ujconfig.Reference{
			Type: "Gateway",
		}
	})
}
