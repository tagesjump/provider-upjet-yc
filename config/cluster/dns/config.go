package dns

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/tagesjump/provider-upjet-yc/config/cluster/vpc"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/cluster/dns/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/cluster/dns"
)

// Configure adds configurations for dns group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_dns_zone", func(r *ujconfig.Resource) {
		r.References["private_networks"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
	})
	p.AddResourceConfigurator("yandex_dns_recordset", func(r *ujconfig.Resource) {
		r.References["zone_id"] = ujconfig.Reference{
			Type: "Zone",
		}
	})
}
