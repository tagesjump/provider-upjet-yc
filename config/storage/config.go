/*
Copyright 2021 The Crossplane Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storage

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/tagesjump/provider-upjet-yc/config/common"
	"github.com/tagesjump/provider-upjet-yc/config/iam"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/storage/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/storage"
)

// Configure adds configurations for storage group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_storage_bucket", func(r *ujconfig.Resource) {
		r.UseAsync = true
		r.References["access_key"] = ujconfig.Reference{
			Type:      fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccountStaticAccessKey"),
			Extractor: common.ExtractPublicKeyFuncPath,
		}
		r.LateInitializer = ujconfig.LateInitializer{
			IgnoredFields: []string{"anonymous_access_flags", "anonymous_access_flags.config_read", "anonymous_access_flags.list", "anonymous_access_flags.read"},
		}
	})
	p.AddResourceConfigurator("yandex_storage_object", func(r *ujconfig.Resource) {
		r.UseAsync = true
		r.References["access_key"] = ujconfig.Reference{
			Type:      fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccountStaticAccessKey"),
			Extractor: common.ExtractPublicKeyFuncPath,
		}
		r.References["bucket"] = ujconfig.Reference{
			Type: "Bucket",
		}
	})
}
