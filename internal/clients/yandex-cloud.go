/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/crossplane/upjet/v2/pkg/terraform"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterv1beta1 "github.com/tagesjump/provider-upjet-yc/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/tagesjump/provider-upjet-yc/apis/namespaced/v1beta1"
)

const (
	// The default folder ID where resources will be placed
	folderID = "folder_id"
	// ID of Yandex.Cloud tenant
	cloudID = "cloud_id"
	// The API endpoint for Yandex.Cloud SDK client
	endpoint = "endpoint"
	// Either the path to or the contents of a Service Account key file in JSON format
	serviceAccountKeyFile = "service_account_key_file"
	// The region where operations will take place
	regionID = "region_id"
	// The zone where operations will take place. Examples: are ru-central1-a, ru-central2-c, etc
	zoneID = "zone"
	// The access token for API operations
	token = "token"
	// Explicitly allow the provider to perform "insecure" SSL requests, default value is `false
	allowInsecure = "insecure"
	// Disable use of TLS. Default value is `false`.
	plainText = "plaintext"
	// The maximum number of times an API request is being executed. If the API request still fails, an error is thrown.
	maxRetries = "max_retries"
	// Yandex.Cloud storage service endpoint
	storageEndpoint = "storage_endpoint"
	// Yandex.Cloud storage service access key. Used when a storage data/resource doesn't have an access key explicitly specified
	storageAccessKey = "storage_access_key"
	// Yandex.Cloud storage service secret key. Used when a storage data/resource doesn't have a secret key explicitly specified
	storageSecretKey = "storage_secret_key"
	// Yandex.Cloud Message Queue service endpoint
	ymqEndpoint = "ymq_endpoint"
	// Yandex.Cloud Message Queue service access key
	ymqAccessKey = "ymq_access_key"
	// Yandex.Cloud Message Queue service secret key
	ymqSecretKey = "ymq_secret_key"
	// Path to shared credentials file
	sharedCredentialsFile = "shared_credentials_file"
	// Profile to use in the shared credentials file. Default value is `default`.
	profile = "profile"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal yandex-cloud.upjet credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot resolve provider config")
		}

		data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// set provider configuration
		ps.Configuration = map[string]any{}

		// Either the path to or the contents of a Service Account key file in JSON format
		ps.Configuration[serviceAccountKeyFile] = string(data)

		for _, setting := range []string{
			folderID, cloudID, endpoint, regionID, zoneID, token,
			allowInsecure, plainText, maxRetries, storageEndpoint, storageAccessKey,
			storageSecretKey, ymqEndpoint, ymqAccessKey, ymqSecretKey, sharedCredentialsFile, profile,
		} {
			if value, ok := creds[setting]; ok {
				ps.Configuration[setting] = value
			}
		}
		return ps, errors.Wrap(configureNoForkYandexCloudClient(ctx, &ps, *tfProvider), "failed to configure the Terraform YandexCloud provider meta")
	}
}

func configureNoForkYandexCloudClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	// Please be aware that this implementation relies on the schema.Provider
	// parameter `p` being a non-pointer. This is because normally
	// the Terraform plugin SDK normally configures the provider
	// only once and using a pointer argument here will cause
	// race conditions between resources referring to different
	// ProviderConfigs.
	diag := p.Configure(ctx, &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
	return nil
}

func toSharedPCSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged:
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed resource")
	}
}

func resolveLegacy(ctx context.Context, client client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc)
}

func resolveModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, "unknown GVK for ProviderConfig")
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		// This indicates a programming error, types are not properly generated
		return nil, errors.New(" is not an Object")
	}

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	pcu := &namespacedv1beta1.ProviderConfigUsage{}
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		return nil, errors.New("unknown provider config type")
	}
	t := resource.NewProviderConfigUsageTracker(crClient, pcu)
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}
