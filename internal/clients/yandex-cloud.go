/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/tagesjump/provider-upjet-yc/apis/v1beta1"
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

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
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
		return ps, errors.Wrap(configureNoForkYandexCloudClient(ctx, &ps, *tfProvider), "failed to configure the Terraform AzureAD provider meta")
	}
}

func configureNoForkYandexCloudClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	// Please be aware that this implementation relies on the schema.Provider
	// parameter `p` being a non-pointer. This is because normally
	// the Terraform plugin SDK normally configures the provider
	// only once and using a pointer argument here will cause
	// race conditions between resources referring to different
	// ProviderConfigs.
	diag := p.Configure(context.WithoutCancel(ctx), &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
	return nil
}
