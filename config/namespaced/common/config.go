package common

import (
	xpref "github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/resource"
)

const (
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/namespaced/common"
	// ExtractPublicKeyFuncPath resource ID extractor access key
	ExtractPublicKeyFuncPath = ConfigPath + ".ExtractAccessKey()"
	// ExtractSpecNameFuncPath  resource ID extractor func name
	ExtractSpecNameFuncPath = ConfigPath + ".ExtractSpecName()"
)

// ExtractAccessKey extracts the value of `spec.atProvider.accessKey`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractAccessKey() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		o, err := tr.GetObservation()
		if err != nil {
			return ""
		}
		if k := o["access_key"]; k != nil {
			return k.(string)
		}
		return ""
	}
}

// ExtractSpecName extracts the value of `spec.forProvider.name`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractSpecName() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		o, err := tr.GetParameters()
		if err != nil {
			return ""
		}
		if k := o["name"]; k != nil {
			return k.(string)
		}
		return ""
	}
}
