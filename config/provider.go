/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/tagesjump/provider-upjet-yc/config/alb"
	"github.com/tagesjump/provider-upjet-yc/config/compute"
	"github.com/tagesjump/provider-upjet-yc/config/container"
	"github.com/tagesjump/provider-upjet-yc/config/dns"
	"github.com/tagesjump/provider-upjet-yc/config/iam"
	"github.com/tagesjump/provider-upjet-yc/config/kms"
	"github.com/tagesjump/provider-upjet-yc/config/kubernetes"
	"github.com/tagesjump/provider-upjet-yc/config/lb"
	"github.com/tagesjump/provider-upjet-yc/config/mdb"
	"github.com/tagesjump/provider-upjet-yc/config/organizationmanager"
	"github.com/tagesjump/provider-upjet-yc/config/resourcemanager"
	"github.com/tagesjump/provider-upjet-yc/config/storage"
	"github.com/tagesjump/provider-upjet-yc/config/vpc"
)

const (
	resourcePrefix = "yandex-cloud"
	modulePath     = "github.com/tagesjump/provider-upjet-yc"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("yandex-cloud.upjet.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		iam.Configure,
		organizationmanager.Configure,
		resourcemanager.Configure,
		kms.Configure,
		vpc.Configure,
		kubernetes.Configure,
		mdb.Configure,
		compute.Configure,
		container.Configure,
		alb.Configure,
		storage.Configure,
		lb.Configure,
		dns.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
