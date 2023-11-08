package datatransfer

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/tagesjump/provider-upjet-yc/config/common"
	"github.com/tagesjump/provider-upjet-yc/config/mdb"
	"github.com/tagesjump/provider-upjet-yc/config/vpc"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/datatransfer/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/datatransfer"
)

// Configure adds configurations for datatransfer group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_datatransfer_endpoint", func(r *ujconfig.Resource) {
		r.References["settings.postgres_target.connection.mdb_cluster_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "PostgresqlCluster"),
		}
		r.References["settings.mysql_target.connection.mdb_cluster_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MySQLCluster"),
		}
		r.References["settings.mongo_target.connection.connection_options.mdb_cluster_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MongodbCluster"),
		}
		r.References["settings.postgres_source.connection.mdb_cluster_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "PostgresqlCluster"),
		}
		r.References["settings.mysql_source.connection.mdb_cluster_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MySQLCluster"),
		}
		r.References["settings.mongo_source.connection.connection_options.mdb_cluster_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MongodbCluster"),
		}
		r.References["settings.mysql_target.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.postgres_target.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.mongo_target.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.clickhouse_target.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.mysql_source.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.postgres_source.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.mongo_source.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.clickhouse_source.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.mysql_source.connection.on_premise.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.mysql_source.user"] = ujconfig.Reference{
			Type:      fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MySQLUser"),
			Extractor: common.ExtractSpecNameFuncPath,
		}
		r.References["settings.mysql_source.database"] = ujconfig.Reference{
			Type:      fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MySQLDatabase"),
			Extractor: common.ExtractSpecNameFuncPath,
		}
		r.References["settings.mysql_source.security_groups"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["settings.mysql_target.connection.on_premise.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.mysql_target.user"] = ujconfig.Reference{
			Type:      fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MySQLUser"),
			Extractor: common.ExtractSpecNameFuncPath,
		}
		r.References["settings.mysql_target.database"] = ujconfig.Reference{
			Type:      fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MySQLDatabase"),
			Extractor: common.ExtractSpecNameFuncPath,
		}
		r.References["settings.mysql_target.security_groups"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["settings.postgres_source.connection.on_premise.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.postgres_source.user"] = ujconfig.Reference{
			Type:      fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "PostgresqlUser"),
			Extractor: common.ExtractSpecNameFuncPath,
		}
		r.References["settings.postgres_source.database"] = ujconfig.Reference{
			Type:      fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "PostgresqlDatabase"),
			Extractor: common.ExtractSpecNameFuncPath,
		}
		r.References["settings.postgres_source.security_groups"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["settings.postgres_target.connection.on_premise.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.postgres_target.user"] = ujconfig.Reference{
			Type:      fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "PostgresqlUser"),
			Extractor: common.ExtractSpecNameFuncPath,
		}
		r.References["settings.postgres_target.database"] = ujconfig.Reference{
			Type:      fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "PostgresqlDatabase"),
			Extractor: common.ExtractSpecNameFuncPath,
		}
		r.References["settings.postgres_target.security_groups"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["settings.mongo_source.connection.connection_options.on_premise.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.clickhouse_source.connection.connection_options.on_premise.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
	})
	p.AddResourceConfigurator("yandex_datatransfer_transfer", func(r *ujconfig.Resource) {
		r.References["source_id"] = ujconfig.Reference{
			Type: "Endpoint",
		}
		r.References["target_id"] = ujconfig.Reference{
			Type: "Endpoint",
		}
		r.UseAsync = true
	})
}
