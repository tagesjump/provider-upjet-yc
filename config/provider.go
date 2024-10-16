/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"github.com/yandex-cloud/terraform-provider-yandex/yandex"
	yandex_framework "github.com/yandex-cloud/terraform-provider-yandex/yandex-framework/provider"

	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane/upjet/pkg/registry/reference"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/tagesjump/provider-upjet-yc/config/alb"
	"github.com/tagesjump/provider-upjet-yc/config/cdn"
	"github.com/tagesjump/provider-upjet-yc/config/cm"
	"github.com/tagesjump/provider-upjet-yc/config/compute"
	"github.com/tagesjump/provider-upjet-yc/config/container"
	"github.com/tagesjump/provider-upjet-yc/config/datatransfer"
	"github.com/tagesjump/provider-upjet-yc/config/dns"
	"github.com/tagesjump/provider-upjet-yc/config/function"
	"github.com/tagesjump/provider-upjet-yc/config/iam"
	"github.com/tagesjump/provider-upjet-yc/config/iot"
	"github.com/tagesjump/provider-upjet-yc/config/kms"
	"github.com/tagesjump/provider-upjet-yc/config/kubernetes"
	"github.com/tagesjump/provider-upjet-yc/config/lb"
	"github.com/tagesjump/provider-upjet-yc/config/loadtesting"
	"github.com/tagesjump/provider-upjet-yc/config/lockbox"
	"github.com/tagesjump/provider-upjet-yc/config/logging"
	"github.com/tagesjump/provider-upjet-yc/config/mdb"
	"github.com/tagesjump/provider-upjet-yc/config/message"
	"github.com/tagesjump/provider-upjet-yc/config/monitoring"
	"github.com/tagesjump/provider-upjet-yc/config/organizationmanager"
	"github.com/tagesjump/provider-upjet-yc/config/resourcemanager"
	"github.com/tagesjump/provider-upjet-yc/config/serverless"
	"github.com/tagesjump/provider-upjet-yc/config/storage"
	"github.com/tagesjump/provider-upjet-yc/config/vpc"
	"github.com/tagesjump/provider-upjet-yc/config/ydb"
)

const (
	resourcePrefix = "yandex-cloud"
	modulePath     = "github.com/tagesjump/provider-upjet-yc"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// workaround for the TF Google v2.41.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(generationProvider bool) (*ujconfig.Provider, error) {
	var providerInstance *schema.Provider
	var err error
	if generationProvider {
		providerInstance, err = getProviderSchema(providerSchema)
	} else {
		providerInstance = yandex.NewSDKProvider()
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("yandex-cloud.upjet.crossplane.io"),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(providerInstance),
		ujconfig.WithTerraformPluginSDKIncludeList(resourceList(TerraformPluginSDKExternalNameConfigs)),
		ujconfig.WithTerraformPluginFrameworkProvider(yandex_framework.NewFrameworkProvider()),
		ujconfig.WithTerraformPluginFrameworkIncludeList(resourceList(TerraformPluginFrameworkExternalNameConfigs)),
		ujconfig.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
	)

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
		logging.Configure,
		message.Configure,
		ydb.Configure,
		datatransfer.Configure,
		lockbox.Configure,
		monitoring.Configure,
		loadtesting.Configure,
		serverless.Configure,
		iot.Configure,
		function.Configure,
		cm.Configure,
		cdn.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// resourceList returns the list of resources that have external
// name configured in the specified table.
func resourceList(t map[string]ujconfig.ExternalName) []string {
	l := make([]string, len(t))
	i := 0
	for n := range t {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = n + "$"
		i++
	}
	return l
}
