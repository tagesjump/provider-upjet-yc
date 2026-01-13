package serverless

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/namespaced/serverless/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/namespaced/serverless"
)

// Configure adds configurations for serverless group.
func Configure(p *ujconfig.Provider) {}
