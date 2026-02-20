package common

import (
	"fmt"
	"strings"

	xpref "github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane/upjet/v2/pkg/resource"
)

const (
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/cluster/common"
	// ExtractPublicKeyFuncPath resource ID extractor access key
	ExtractPublicKeyFuncPath = ConfigPath + ".ExtractAccessKey()"
	// ExtractSpecNameFuncPath resource ID extractor func name
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

func MustLookup(root map[string]*schema.Schema, path ...string) *schema.Schema {
	current := root

	for i, p := range path {
		s, ok := current[p]
		if !ok {
			panic(fmt.Sprintf("schema path not found: %s", strings.Join(path, ".")))
		}

		if i == len(path)-1 {
			return s
		}

		res, ok := s.Elem.(*schema.Resource)
		if !ok {
			panic(fmt.Sprintf("schema path is not a resource: %s", strings.Join(path[:i+1], ".")))
		}

		current = res.Schema
	}

	panic("unreachable")
}
