package cm

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/cm/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/cm"
)

// Configure adds configurations for cm group.
func Configure(p *ujconfig.Provider) {}
