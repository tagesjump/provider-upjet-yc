package kms

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/tagesjump/provider-upjet-yc/config/cluster/iam"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/cluster/kms/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/cluster/kms"
)

// Configure adds configurations for kms group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_kms_asymmetric_encryption_key_iam_binding", func(r *ujconfig.Resource) {
		r.References["asymmetric_encryption_key_id"] = ujconfig.Reference{
			Type: "AsymmetricEncryptionKey",
		}
		r.References["members"] = ujconfig.Reference{
			Type:              fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
			Extractor:         fmt.Sprintf("%s.%s", iam.ConfigPath, iam.ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})
	p.AddResourceConfigurator("yandex_kms_asymmetric_signature_key_iam_binding", func(r *ujconfig.Resource) {
		r.References["asymmetric_signature_key_id"] = ujconfig.Reference{
			Type: "AsymmetricSignatureKey",
		}
		r.References["members"] = ujconfig.Reference{
			Type:              fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
			Extractor:         fmt.Sprintf("%s.%s", iam.ConfigPath, iam.ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})
	p.AddResourceConfigurator("yandex_kms_secret_ciphertext", func(r *ujconfig.Resource) {
		r.References["key_id"] = ujconfig.Reference{
			Type: "SymmetricKey",
		}
	})
	p.AddResourceConfigurator("yandex_kms_symmetric_key_iam_binding", func(r *ujconfig.Resource) {
		r.References["symmetric_key_id"] = ujconfig.Reference{
			Type: "SymmetricKey",
		}
		r.References["members"] = ujconfig.Reference{
			Type:              fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
			Extractor:         fmt.Sprintf("%s.%s", iam.ConfigPath, iam.ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})
}
