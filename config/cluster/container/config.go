package container

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/tagesjump/provider-upjet-yc/config/cluster/iam"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/cluster/container/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/cluster/container"
)

// Configure adds configurations for container group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_container_registry_iam_binding", func(r *ujconfig.Resource) {
		r.References["registry_id"] = ujconfig.Reference{
			Type: "Registry",
		}
		r.References["members"] = ujconfig.Reference{
			Type:              fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
			Extractor:         fmt.Sprintf("%s.%s", iam.ConfigPath, iam.ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})

	p.AddResourceConfigurator("yandex_container_registry_ip_permission", func(r *ujconfig.Resource) {
		r.References["registry_id"] = ujconfig.Reference{
			Type: "Registry",
		}
	})

	p.AddResourceConfigurator("yandex_container_repository_iam_binding", func(r *ujconfig.Resource) {
		r.References["repository_id"] = ujconfig.Reference{
			Type: "Repository",
		}
		r.References["members"] = ujconfig.Reference{
			Type:              fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
			Extractor:         fmt.Sprintf("%s.%s", iam.ConfigPath, iam.ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})

	p.AddResourceConfigurator("yandex_container_repository_lifecycle_policy", func(r *ujconfig.Resource) {
		r.References["repository_id"] = ujconfig.Reference{
			Type: "Repository",
		}
	})
}
