/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"github.com/yandex-cloud/terraform-provider-yandex/yandex"
	yandex_framework "github.com/yandex-cloud/terraform-provider-yandex/yandex-framework/provider"

	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane/upjet/v2/pkg/registry/reference"
	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	albCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/alb"
	cdnCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/cdn"
	cmCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/cm"
	computeCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/compute"
	containerCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/container"
	datatransferCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/datatransfer"
	dnsCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/dns"
	functionCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/function"
	iamCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/iam"
	iotCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/iot"
	kmsCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/kms"
	kubernetesCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/kubernetes"
	lbCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/lb"
	loadtestingCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/loadtesting"
	lockboxCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/lockbox"
	loggingCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/logging"
	mdbCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/mdb"
	messageCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/message"
	monitoringCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/monitoring"
	organizationmanagerCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/organizationmanager"
	resourcemanagerCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/resourcemanager"
	serverlessCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/serverless"
	storageCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/storage"
	vpcCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/vpc"
	ydbCluster "github.com/tagesjump/provider-upjet-yc/config/cluster/ydb"

	albNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/alb"
	cdnNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/cdn"
	cmNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/cm"
	computeNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/compute"
	containerNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/container"
	datatransferNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/datatransfer"
	dnsNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/dns"
	functionNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/function"
	iamNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/iam"
	iotNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/iot"
	kmsNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/kms"
	kubernetesNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/kubernetes"
	lbNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/lb"
	loadtestingNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/loadtesting"
	lockboxNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/lockbox"
	loggingNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/logging"
	mdbNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/mdb"
	messageNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/message"
	monitoringNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/monitoring"
	organizationmanagerNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/organizationmanager"
	resourcemanagerNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/resourcemanager"
	serverlessNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/serverless"
	storageNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/storage"
	vpcNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/vpc"
	ydbNamespaced "github.com/tagesjump/provider-upjet-yc/config/namespaced/ydb"
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
		iamCluster.Configure,
		organizationmanagerCluster.Configure,
		resourcemanagerCluster.Configure,
		kmsCluster.Configure,
		vpcCluster.Configure,
		kubernetesCluster.Configure,
		mdbCluster.Configure,
		computeCluster.Configure,
		containerCluster.Configure,
		albCluster.Configure,
		storageCluster.Configure,
		lbCluster.Configure,
		dnsCluster.Configure,
		loggingCluster.Configure,
		messageCluster.Configure,
		ydbCluster.Configure,
		datatransferCluster.Configure,
		lockboxCluster.Configure,
		monitoringCluster.Configure,
		loadtestingCluster.Configure,
		serverlessCluster.Configure,
		iotCluster.Configure,
		functionCluster.Configure,
		cmCluster.Configure,
		cdnCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// GetProviderNamespaced returns namespaced provider configuration
func GetProviderNamespaced(generationProvider bool) (*ujconfig.Provider, error) {
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
		ujconfig.WithRootGroup("yandex-cloud.upjet.m.crossplane.io"),
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
		iamNamespaced.Configure,
		organizationmanagerNamespaced.Configure,
		resourcemanagerNamespaced.Configure,
		kmsNamespaced.Configure,
		vpcNamespaced.Configure,
		kubernetesNamespaced.Configure,
		mdbNamespaced.Configure,
		computeNamespaced.Configure,
		containerNamespaced.Configure,
		albNamespaced.Configure,
		storageNamespaced.Configure,
		lbNamespaced.Configure,
		dnsNamespaced.Configure,
		loggingNamespaced.Configure,
		messageNamespaced.Configure,
		ydbNamespaced.Configure,
		datatransferNamespaced.Configure,
		lockboxNamespaced.Configure,
		monitoringNamespaced.Configure,
		loadtestingNamespaced.Configure,
		serverlessNamespaced.Configure,
		iotNamespaced.Configure,
		functionNamespaced.Configure,
		cmNamespaced.Configure,
		cdnNamespaced.Configure,
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
