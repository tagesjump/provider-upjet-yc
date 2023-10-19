/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/tagesjump/provider-upjet-yc/apis/v1beta1"
)

const (
	folderID              = "folder_id"
	cloudID               = "cloud_id"
	endpoint              = "endpoint"
	serviceAccountKeyFile = "service_account_key_file"
	regionID              = "region_id"
	zoneID                = "zone"
	token                 = "token"
	allowInsecure         = "insecure"
	plainText             = "plaintext"
	maxRetries            = "max_retries"
	storageEndpoint       = "storage_endpoint"
	storageAccessKey      = "storage_access_key"
	storageSecretKey      = "storage_secret_key"
	ymqEndpoint           = "ymq_endpoint"
	ymqAccessKey          = "ymq_access_key"
	ymqSecretKey          = "ymq_secret_key"
	sharedCredentialsFile = "shared_credentials_file"
	profile               = "profile"
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
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

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

		// The default folder ID where resources will be placed
		if value, ok := creds[folderID]; ok {
			ps.Configuration[folderID] = value
		}

		// ID of Yandex.Cloud tenant
		if value, ok := creds[cloudID]; ok {
			ps.Configuration[cloudID] = value
		}

		// The API endpoint for Yandex.Cloud SDK client
		if value, ok := creds[endpoint]; ok {
			ps.Configuration[endpoint] = value
		}

		// The region where operations will take place
		if value, ok := creds[regionID]; ok {
			ps.Configuration[regionID] = value
		}

		// The zone where operations will take place. Examples: are ru-central1-a, ru-central2-c, etc
		if value, ok := creds[zoneID]; ok {
			ps.Configuration[zoneID] = value
		}

		// The access token for API operations
		if value, ok := creds[token]; ok {
			ps.Configuration[token] = value
		}

		// Explicitly allow the provider to perform "insecure" SSL requests, default value is `false
		if value, ok := creds[allowInsecure]; ok {
			ps.Configuration[allowInsecure] = value
		}

		// Disable use of TLS. Default value is `false`.
		if value, ok := creds[plainText]; ok {
			ps.Configuration[plainText] = value
		}

		// The maximum number of times an API request is being executed. If the API request still fails, an error is thrown.
		if value, ok := creds[maxRetries]; ok {
			ps.Configuration[maxRetries] = value
		}

		// Yandex.Cloud storage service endpoint
		if value, ok := creds[storageEndpoint]; ok {
			ps.Configuration[storageEndpoint] = value
		}

		// Yandex.Cloud storage service access key. Used when a storage data/resource doesn't have an access key explicitly specified
		if value, ok := creds[storageAccessKey]; ok {
			ps.Configuration[storageAccessKey] = value
		}

		// Yandex.Cloud storage service secret key. Used when a storage data/resource doesn't have a secret key explicitly specified
		if value, ok := creds[storageSecretKey]; ok {
			ps.Configuration[storageSecretKey] = value
		}

		// Yandex.Cloud Message Queue service endpoint
		if value, ok := creds[ymqEndpoint]; ok {
			ps.Configuration[ymqEndpoint] = value
		}

		// Yandex.Cloud Message Queue service access key
		if value, ok := creds[ymqAccessKey]; ok {
			ps.Configuration[ymqAccessKey] = value
		}

		// Yandex.Cloud Message Queue service secret key
		if value, ok := creds[ymqSecretKey]; ok {
			ps.Configuration[ymqSecretKey] = value
		}

		// Path to shared credentials file
		if value, ok := creds[sharedCredentialsFile]; ok {
			ps.Configuration[sharedCredentialsFile] = value
		}

		// Profile to use in the shared credentials file. Default value is `default`.
		if value, ok := creds[profile]; ok {
			ps.Configuration[profile] = value
		}

		return ps, nil
	}
}
