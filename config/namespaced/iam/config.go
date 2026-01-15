package iam

import (
	"encoding/json"
	"fmt"

	"github.com/crossplane/crossplane-runtime/v2/pkg/meta"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/pkg/errors"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/namespaced/iam/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/namespaced/iam"
	// ServiceAccountRefValueFn is the name of resolver.
	ServiceAccountRefValueFn = "ServiceAccountRefValue()"
)

func serviceAccountKey(attr map[string]interface{}) ([]byte, error) {
	for _, attribute := range []string{"id", "service_account_id", "created_at", "key_algorithm", "public_key", "private_key"} {
		if _, ok := attr[attribute]; !ok {
			return nil, errors.New(fmt.Sprintf("The attribute %s is missing", attribute))
		}
	}
	result := map[string]string{
		"id":                 attr["id"].(string),
		"service_account_id": attr["service_account_id"].(string),
		"created_at":         attr["created_at"].(string),
		"key_algorithm":      attr["key_algorithm"].(string),
		"public_key":         attr["public_key"].(string),
		"private_key":        attr["private_key"].(string),
	}
	encoded, err := json.Marshal(result)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot marshal service account key")
	}
	return encoded, nil
}

func serviceAccountStaticKey(attr map[string]interface{}) (map[string][]byte, error) {
	if _, ok := attr["access_key"]; !ok {
		return nil, nil
	}
	if _, ok := attr["id"]; !ok {
		return nil, nil
	}
	return map[string][]byte{
		"attribute.access_key": []byte(attr["access_key"].(string)),
		"attribute.id":         []byte(attr["id"].(string)),
	}, nil
}

// Configure adds configurations for iam group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_iam_service_account_key", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			Type: "ServiceAccount",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			serviceAccountKey, err := serviceAccountKey(attr)
			return map[string][]byte{
				"service_account_key": serviceAccountKey,
			}, err
		}
	})
	p.AddResourceConfigurator("yandex_iam_service_account_static_access_key", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			Type: "ServiceAccount",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = serviceAccountStaticKey
	})
	p.AddResourceConfigurator("yandex_iam_service_account_api_key", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			Type: "ServiceAccount",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = serviceAccountStaticKey
	})
	p.AddResourceConfigurator("yandex_iam_service_account_iam_member", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			Type: "ServiceAccount",
		}
		r.References["member"] = ujconfig.Reference{
			Type:              "ServiceAccount",
			Extractor:         fmt.Sprintf("%s.%s", ConfigPath, ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})

	p.AddResourceConfigurator("yandex_iam_service_account_iam_policy", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			Type: "ServiceAccount",
		}
		r.References["members"] = ujconfig.Reference{
			Type:              "ServiceAccount",
			Extractor:         fmt.Sprintf("%s.%s", ConfigPath, ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})

	p.AddResourceConfigurator("yandex_iam_service_account_iam_binding", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			Type: "ServiceAccount",
		}
		r.References["members"] = ujconfig.Reference{
			Type:              "ServiceAccount",
			Extractor:         fmt.Sprintf("%s.%s", ConfigPath, ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})
}

// ServiceAccountRefValue returns an extractor that returns templated value with service account id of ServiceAccount.
func ServiceAccountRefValue() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		if externalName := meta.GetExternalName(mg); externalName != "" {
			return fmt.Sprintf("serviceAccount:%s", externalName)
		}
		return ""
	}
}
