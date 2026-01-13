package message

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/tagesjump/provider-upjet-yc/config/namespaced/common"

	"github.com/tagesjump/provider-upjet-yc/config/namespaced/iam"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/namespaced/message/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/namespaced/message"
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
