package mdb

import (
	"fmt"
	"strings"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/tagesjump/provider-upjet-yc/config/iam"
	"github.com/tagesjump/provider-upjet-yc/config/vpc"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/mdb/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/mdb"
)

func usernames(attr map[string]interface{}) []string {
	usersInterface, ok := attr["user"]
	if !ok {
		return nil
	}
	users, ok := usersInterface.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, len(users))
	for i, userInterface := range users {
		user, ok := userInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if username, ok := user["name"].(string); ok {
			result[i] = username
		}
	}
	return result
}

func passwords(attr map[string]interface{}) []string {
	usersInterface, ok := attr["user"]
	if !ok {
		return nil
	}
	users, ok := usersInterface.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, len(users))
	for i, userInterface := range users {
		user, ok := userInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if password, ok := user["password"].(string); ok {
			result[i] = password
		}
	}
	return result
}

func databases(attr map[string]interface{}) []string {
	databasesInterface, ok := attr["database"]
	if !ok {
		return nil
	}
	databases, ok := databasesInterface.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, len(databases))
	for i, databaseInterface := range databases {
		database, ok := databaseInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if databaseName, ok := database["name"].(string); ok {
			result[i] = databaseName
		}
	}
	return result
}

func fqdns(attr map[string]interface{}) []string {
	hostsInterface, ok := attr["host"]
	if !ok {
		return nil
	}
	hosts, ok := hostsInterface.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, len(hosts))
	for i, hostInterface := range hosts {
		host, ok := hostInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if fqdn, ok := host["fqdn"].(string); ok {
			result[i] = fqdn
		}
	}
	return result
}

func hostnames(attr map[string]interface{}) []string {
	hostsInterface, ok := attr["host"]
	if !ok {
		return nil
	}
	hosts, ok := hostsInterface.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, len(hosts))
	for i, hostInterface := range hosts {
		host, ok := hostInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if fqdn, ok := host["name"].(string); ok {
			result[i] = fqdn
		}
	}
	return result
}

func postgresqlConnectionStrings(attr map[string]interface{}) map[string]string {
	connstrings := make(map[string]string)
	hosts := fqdns(attr)
	for _, db := range databases(attr) {
		ps := passwords(attr)
		for i, u := range usernames(attr) {
			password := ps[i]
			connstrings[fmt.Sprintf("connection-string.%s.%s", u, db)] =
				fmt.Sprintf(
					"host=%s port=6432 sslmode=verify-full dbname=%s user=%s target_session_attrs=read-write password=%s",
					strings.Join(hosts, ","), db, u, password)
		}
	}
	return connstrings
}

func mysqlConnectionStrings(attr map[string]interface{}) map[string]string {
	connstrings := make(map[string]string)
	for _, db := range databases(attr) {
		ps := passwords(attr)
		for i, u := range usernames(attr) {
			password := ps[i]
			for _, host := range fqdns(attr) {
				connstrings[fmt.Sprintf("connection-string.%s.%s.%s", u, db, host)] =
					fmt.Sprintf("mysql://%s:%s@%s/%s", u, password, host, db)
			}
		}
	}
	return connstrings
}

func mongodbConnectionStrings(attr map[string]interface{}) map[string]string {
	connstrings := make(map[string]string)
	hosts := hostnames(attr)
	for _, db := range databases(attr) {
		ps := passwords(attr)
		for i, u := range usernames(attr) {
			password := ps[i]
			connstrings[fmt.Sprintf("connection-string.%s.%s", u, db)] =
				fmt.Sprintf(
					"mongodb://%s:%s@%s/%s",
					u, password, strings.Join(hosts, ","), db)
		}
	}
	return connstrings
}

func elasticsearchConnectionStrings(attr map[string]interface{}) map[string]string {
	connstrings := make(map[string]string)
	for _, db := range databases(attr) {
		ps := passwords(attr)
		for i, u := range usernames(attr) {
			password := ps[i]
			for _, host := range fqdns(attr) {
				connstrings[fmt.Sprintf("connection-string.%s.%s.%s", u, db, host)] =
					fmt.Sprintf("https://%s:%s@%s:9200/%s",
						u, password, host, db)
			}

		}
	}
	return connstrings
}

func postgresqlConnDetails(attr map[string]interface{}) map[string][]byte {
	conn := make(map[string][]byte)
	for i, v := range fqdns(attr) {
		conn[fmt.Sprintf("attribute.host.%d.fqdn", i)] = []byte(v)
	}
	for i, v := range usernames(attr) {
		conn[fmt.Sprintf("attribute.user.%d.name", i)] = []byte(v)
	}
	for i, v := range databases(attr) {
		conn[fmt.Sprintf("attribute.database.%d.name", i)] = []byte(v)
	}
	for k, v := range postgresqlConnectionStrings(attr) {
		conn[k] = []byte(v)
	}

	return conn
}

func mysqlConnDetails(attr map[string]interface{}) map[string][]byte {
	conn := make(map[string][]byte)
	for i, v := range fqdns(attr) {
		conn[fmt.Sprintf("attribute.host.%d.fqdn", i)] = []byte(v)
	}
	for i, v := range usernames(attr) {
		conn[fmt.Sprintf("attribute.user.%d.name", i)] = []byte(v)
	}
	for i, v := range databases(attr) {
		conn[fmt.Sprintf("attribute.database.%d.name", i)] = []byte(v)
	}
	for k, v := range mysqlConnectionStrings(attr) {
		conn[k] = []byte(v)
	}

	return conn
}

func mongodbConnDetails(attr map[string]interface{}) map[string][]byte {
	conn := make(map[string][]byte)
	for i, v := range hostnames(attr) {
		conn[fmt.Sprintf("attribute.host.%d.name", i)] = []byte(v)
	}
	for i, v := range usernames(attr) {
		conn[fmt.Sprintf("attribute.user.%d.name", i)] = []byte(v)
	}
	for i, v := range databases(attr) {
		conn[fmt.Sprintf("attribute.database.%d.name", i)] = []byte(v)
	}
	for k, v := range mongodbConnectionStrings(attr) {
		conn[k] = []byte(v)
	}

	return conn
}

func elasticsearchConnDetails(attr map[string]interface{}) map[string][]byte {
	conn := make(map[string][]byte)
	for i, v := range hostnames(attr) {
		conn[fmt.Sprintf("attribute.host.%d.name", i)] = []byte(v)
	}
	for i, v := range usernames(attr) {
		conn[fmt.Sprintf("attribute.user.%d.name", i)] = []byte(v)
	}
	for i, v := range databases(attr) {
		conn[fmt.Sprintf("attribute.database.%d.name", i)] = []byte(v)
	}
	for i, v := range fqdns(attr) {
		conn[fmt.Sprintf("attribute.fqdn.%d.name", i)] = []byte(v)
	}
	for k, v := range elasticsearchConnectionStrings(attr) {
		conn[k] = []byte(v)
	}
	if clusterID, ok := attr["id"]; ok {
		if clusterIDString, ok := clusterID.(string); ok {
			conn["attribute.id"] = []byte(clusterIDString)
		}
	}
	return conn
}

func redisConnDetails(attr map[string]interface{}) map[string][]byte {
	conn := make(map[string][]byte)
	for i, v := range fqdns(attr) {
		conn[fmt.Sprintf("attribute.host.%d.fqdn", i)] = []byte(v)
	}
	for i, v := range databases(attr) {
		conn[fmt.Sprintf("attribute.database.%d.name", i)] = []byte(v)
	}

	return conn
}

// Configure adds configurations for mdb group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_mdb_postgresql_cluster", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return postgresqlConnDetails(attr), nil
		}
	})

	p.AddResourceConfigurator("yandex_mdb_postgresql_database", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			Type: "PostgresqlCluster",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_mdb_postgresql_user", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			Type: "PostgresqlCluster",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_mdb_mysql_cluster", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return mysqlConnDetails(attr), nil
		}
	})
	p.AddResourceConfigurator("yandex_mdb_mysql_database", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			Type: "MySQLCluster",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_mdb_mysql_user", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			Type: "MySQLCluster",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_mdb_redis_cluster", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return redisConnDetails(attr), nil
		}
	})
	p.AddResourceConfigurator("yandex_mdb_mongodb_cluster", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return mongodbConnDetails(attr), nil
		}
	})
	p.AddResourceConfigurator("yandex_mdb_elasticsearch_cluster", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["service_account_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return elasticsearchConnDetails(attr), nil
		}
	})

	p.AddResourceConfigurator("yandex_mdb_kafka_cluster", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["subnet_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
		// TODO: AdditionalConnDetails
	})
	p.AddResourceConfigurator("yandex_mdb_kafka_topic", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			Type: "KafkaCluster",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("yandex_mdb_kafka_user", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			Type: "KafkaCluster",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("yandex_mdb_kafka_connector", func(r *ujconfig.Resource) {
		r.References["cluster_id"] = ujconfig.Reference{
			Type: "KafkaCluster",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("yandex_mdb_clickhouse_cluster", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["service_account_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
		}
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
		// TODO: AdditionalConnDetails
	})

	p.AddResourceConfigurator("yandex_mdb_greenplum_cluster", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
		// TODO: AdditionalConnDetails
	})

	p.AddResourceConfigurator("yandex_mdb_sqlserver_cluster", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
		// TODO: AdditionalConnDetails
	})

	p.AddResourceConfigurator("yandex_mdb_opensearch_cluster", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["service_account_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
		}
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["config.opensearch.node_groups.subnet_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["config.dashboard.node_groups.subnet_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.UseAsync = true
		// TODO: AdditionalConnDetails
	})
}
