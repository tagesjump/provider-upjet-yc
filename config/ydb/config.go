package ydb

import (
	"fmt"

	"github.com/tagesjump/provider-upjet-yc/config/iam"
	"github.com/tagesjump/provider-upjet-yc/config/vpc"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/ydb/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/ydb"
)

// Configure adds configurations for ydb group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_ydb_database_dedicated", func(r *ujconfig.Resource) {
		r.References["subnet_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
	})
	p.AddResourceConfigurator("yandex_ydb_database_iam_binding", func(r *ujconfig.Resource) {
		r.References["database_id"] = ujconfig.Reference{
			Type: "DatabaseServerless",
		}
		r.References["members"] = ujconfig.Reference{
			Type:              fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
			Extractor:         fmt.Sprintf("%s.%s", iam.ConfigPath, iam.ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})
}
