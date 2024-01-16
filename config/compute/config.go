package compute

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/tagesjump/provider-upjet-yc/config/iam"
	"github.com/tagesjump/provider-upjet-yc/config/vpc"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/compute/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/compute"
)

// Configure adds configurations for compute group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_compute_instance", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
		}
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["network_interface.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["network_interface.security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["boot_disk.initialize_params.image_id"] = ujconfig.Reference{
			Type: "Image",
		}
		r.References["boot_disk.disk_id"] = ujconfig.Reference{
			Type: "Disk",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("yandex_compute_instance_group", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
		}
		r.References["instance_template.network_interface.network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["instance_template.network_interface.subnet_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["instance_template.network_interface.security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["instance_template.boot_disk.initialize_params.image_id"] = ujconfig.Reference{
			Type: "Image",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("yandex_compute_snapshot", func(r *ujconfig.Resource) {
		r.References["source_disk_id"] = ujconfig.Reference{
			Type: "Disk",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("yandex_compute_snapshot_schedule", func(r *ujconfig.Resource) {
		r.References["disk_ids"] = ujconfig.Reference{
			Type: "Disk",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("yandex_compute_disk", func(r *ujconfig.Resource) {
		r.References["image_id"] = ujconfig.Reference{
			Type: "Image",
		}
		r.UseAsync = true
	})

}
