package loadtesting

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/tagesjump/provider-upjet-yc/config/namespaced/compute"
	"github.com/tagesjump/provider-upjet-yc/config/namespaced/vpc"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/namespaced/loadtesting/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/namespaced/loadtesting"
)

// Configure adds configurations for loadtesting group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_loadtesting_agent", func(r *ujconfig.Resource) {
		r.References["compute_instance.service_account_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["compute_instance.network_interface.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["compute_instance.boot_disk.disk_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", compute.ApisPackagePath, "Disk"),
		}
		r.UseAsync = true
	})

}
