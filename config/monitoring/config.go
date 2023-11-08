package monitoring

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/tagesjump/provider-upjet-yc/config/resourcemanager"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/monitoring/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/monitoring"
)

// Configure adds configurations for monitoring group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_monitoring_dashboard", func(r *ujconfig.Resource) {
		r.References["parametrization.parameters.label_values.folder_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", resourcemanager.ApisPackagePath, "Folder"),
		}
	})
}
