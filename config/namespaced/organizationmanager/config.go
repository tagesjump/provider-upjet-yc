package organizationmanager

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/namespaced/organizationmanager/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/namespaced/organizationmanager"
)

// Cross-references to members have been disabled because
// they can interfere and it is impossible to make them per user.

// Configure adds configurations for organizationmanager group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_organizationmanager_group_iam_member", func(r *ujconfig.Resource) {
		r.References["group_id"] = ujconfig.Reference{
			Type: "Group",
		}
		// r.References["member"] = ujconfig.Reference{
		// 	Type:              "ServiceAccount",
		// 	Extractor:         fmt.Sprintf("%s.%s", iam.ConfigPath, iam.ServiceAccountRefValueFn),
		// 	RefFieldName:      "ServiceAccountRef",
		// 	SelectorFieldName: "ServiceAccountSelector",
		// }
	})
	p.AddResourceConfigurator("yandex_organizationmanager_group_membership", func(r *ujconfig.Resource) {
		r.References["group_id"] = ujconfig.Reference{
			Type: "Group",
		}
	})
	// p.AddResourceConfigurator("yandex_organizationmanager_organization_iam_binding", func (r *ujconfig.Resource) {
	// r.References["members"] = ujconfig.Reference{
	// 	Type:              "ServiceAccount",
	// 	Extractor:         fmt.Sprintf("%s.%s", iam.ConfigPath, iam.ServiceAccountRefValueFn),
	// 	RefFieldName:      "ServiceAccountRef",
	// 	SelectorFieldName: "ServiceAccountSelector",
	// }
	// })
	// p.AddResourceConfigurator("yandex_organizationmanager_organization_iam_member", func (r *ujconfig.Resource) {
	// r.References["member"] = ujconfig.Reference{
	// 	Type:              "ServiceAccount",
	// 	Extractor:         fmt.Sprintf("%s.%s", iam.ConfigPath, iam.ServiceAccountRefValueFn),
	// 	RefFieldName:      "ServiceAccountRef",
	// 	SelectorFieldName: "ServiceAccountSelector",
	// }
	// })
}
