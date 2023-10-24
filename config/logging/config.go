package logging

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/logging/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/logging"
)

// Configure adds configurations for logging group.
func Configure(p *ujconfig.Provider) {}
