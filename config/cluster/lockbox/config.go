package lockbox

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/tagesjump/provider-upjet-yc/config/cluster/iam"
	"github.com/tagesjump/provider-upjet-yc/config/cluster/kms"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/cluster/lockbox/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/cluster/lockbox"
)

// Configure adds configurations for lockbox group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_lockbox_secret", func(r *ujconfig.Resource) {
		r.References["kms_key_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", kms.ApisPackagePath, "SymmetricKey"),
		}
	})

	p.AddResourceConfigurator("yandex_lockbox_secret_iam_binding", func(r *ujconfig.Resource) {
		r.References["secret_id"] = ujconfig.Reference{
			Type: "Secret",
		}
		r.References["members"] = ujconfig.Reference{
			Type:              fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
			Extractor:         fmt.Sprintf("%s.%s", iam.ConfigPath, iam.ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})

	p.AddResourceConfigurator("yandex_lockbox_secret_version", func(r *ujconfig.Resource) {
		r.References["secret_id"] = ujconfig.Reference{
			Type: "Secret",
		}
	})

}
