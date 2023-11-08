package dns

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/tagesjump/provider-upjet-yc/config/vpc"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/dns/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/dns"
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
