package message

import (
	"fmt"

	"github.com/tagesjump/provider-upjet-yc/config/common"
	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/tagesjump/provider-upjet-yc/config/iam"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/message/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/message"
)

// Configure adds configurations for yandex message queue group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_message_queue", func(r *ujconfig.Resource) {
		r.References["access_key"] = ujconfig.Reference{
			Type:      fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccountStaticAccessKey"),
			Extractor: common.ExtractPublicKeyFuncPath,
		}
		r.UseAsync = true
	})
}
