package resourcemanager

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/namespaced/resourcemanager/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/namespaced/resourcemanager"
)

// Configure adds configurations for resourcemanager group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_resourcemanager_cloud_iam_binding", func(r *ujconfig.Resource) {
		r.References["cloud_id"] = ujconfig.Reference{
			Type: "Cloud",
		}
	})
	p.AddResourceConfigurator("yandex_resourcemanager_cloud_iam_member", func(r *ujconfig.Resource) {
		r.References["cloud_id"] = ujconfig.Reference{
			Type: "Cloud",
		}
	})
	p.AddResourceConfigurator("yandex_resourcemanager_folder", func(r *ujconfig.Resource) {
		r.References["cloud_id"] = ujconfig.Reference{
			Type: "Cloud",
		}
	})
	p.AddResourceConfigurator("yandex_resourcemanager_folder_iam_binding", func(r *ujconfig.Resource) {
		r.References["folder_id"] = ujconfig.Reference{
			Type: "Folder",
		}
	})
	p.AddResourceConfigurator("yandex_resourcemanager_folder_iam_member", func(r *ujconfig.Resource) {
		r.References["folder_id"] = ujconfig.Reference{
			Type: "Folder",
		}
	})
	p.AddResourceConfigurator("yandex_resourcemanager_folder_iam_policy", func(r *ujconfig.Resource) {
		r.References["folder_id"] = ujconfig.Reference{
			Type: "Folder",
		}
	})
}
