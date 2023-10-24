package alb

import (
	"fmt"

	"github.com/tagesjump/provider-upjet-yc/config/vpc"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/tagesjump/provider-upjet-yc/apis/kms/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/tagesjump/provider-upjet-yc/config/kms"
)

// Configure adds configurations for alb group.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("yandex_alb_load_balancer", func(r *ujconfig.Resource) {
		r.References["network_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["security_group_ids"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["allocation_policy.location.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["listener.endpoint.address.internal_ipv4_address.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["listener.http.handler.http_router_id"] = ujconfig.Reference{
			Type: "HTTPRouter",
		}
		r.References["sni_handler.handler.stream_handler.backend_group_id"] = ujconfig.Reference{
			Type: "BackendGroup",
		}
		r.References["sni_handler.handler.http_handler.http_router_id"] = ujconfig.Reference{
			Type: "HTTPRouter",
		}
		r.References["default_handler.handler.stream_handler.backend_group_id"] = ujconfig.Reference{
			Type: "BackendGroup",
		}
		r.References["default_handler.handler.http_handler.http_router_id"] = ujconfig.Reference{
			Type: "HTTPRouter",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_alb_backend_group", func(r *ujconfig.Resource) {
		r.References["http_backend.target_group_ids"] = ujconfig.Reference{
			Type: "TargetGroup",
		}
		r.References["stream_backend.target_group_ids"] = ujconfig.Reference{
			Type: "TargetGroup",
		}
		r.References["grpc_backend.target_group_ids"] = ujconfig.Reference{
			Type: "TargetGroup",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_alb_http_router", func(r *ujconfig.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_alb_target_group", func(r *ujconfig.Resource) {
		r.References["target.subnet_id"] = ujconfig.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_alb_virtual_host", func(r *ujconfig.Resource) {
		r.References["http_router_id"] = ujconfig.Reference{
			Type: "HTTPRouter",
		}
		r.References["route.http_route.http_route_action.backend_group_id"] = ujconfig.Reference{
			Type: "BackendGroup",
		}
		r.References["route.grpc_route.grpc_route_action.backend_group_id"] = ujconfig.Reference{
			Type: "BackendGroup",
		}
		r.UseAsync = true
	})
}
