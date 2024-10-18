// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Address.
func (mg *Address) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DefaultSecurityGroup.
func (mg *DefaultSecurityGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.NetworkIDRef,
		Selector:     mg.Spec.InitProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NetworkID")
	}
	mg.Spec.InitProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Gateway.
func (mg *Gateway) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Network.
func (mg *Network) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PrivateEndpoint.
func (mg *PrivateEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EndpointAddress); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EndpointAddress[i3].SubnetID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.EndpointAddress[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.EndpointAddress[i3].SubnetIDSelector,
			To: reference.To{
				List:    &SubnetList{},
				Managed: &Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EndpointAddress[i3].SubnetID")
		}
		mg.Spec.ForProvider.EndpointAddress[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EndpointAddress[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.EndpointAddress); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EndpointAddress[i3].SubnetID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.EndpointAddress[i3].SubnetIDRef,
			Selector:     mg.Spec.InitProvider.EndpointAddress[i3].SubnetIDSelector,
			To: reference.To{
				List:    &SubnetList{},
				Managed: &Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EndpointAddress[i3].SubnetID")
		}
		mg.Spec.InitProvider.EndpointAddress[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EndpointAddress[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.NetworkIDRef,
		Selector:     mg.Spec.InitProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NetworkID")
	}
	mg.Spec.InitProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RouteTable.
func (mg *RouteTable) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StaticRoute); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StaticRoute[i3].GatewayID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.StaticRoute[i3].GatewayIDRef,
			Selector:     mg.Spec.ForProvider.StaticRoute[i3].GatewayIDSelector,
			To: reference.To{
				List:    &GatewayList{},
				Managed: &Gateway{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StaticRoute[i3].GatewayID")
		}
		mg.Spec.ForProvider.StaticRoute[i3].GatewayID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.StaticRoute[i3].GatewayIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.NetworkIDRef,
		Selector:     mg.Spec.InitProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NetworkID")
	}
	mg.Spec.InitProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.StaticRoute); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StaticRoute[i3].GatewayID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.StaticRoute[i3].GatewayIDRef,
			Selector:     mg.Spec.InitProvider.StaticRoute[i3].GatewayIDSelector,
			To: reference.To{
				List:    &GatewayList{},
				Managed: &Gateway{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.StaticRoute[i3].GatewayID")
		}
		mg.Spec.InitProvider.StaticRoute[i3].GatewayID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.StaticRoute[i3].GatewayIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this SecurityGroup.
func (mg *SecurityGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.NetworkIDRef,
		Selector:     mg.Spec.InitProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NetworkID")
	}
	mg.Spec.InitProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityGroupRule.
func (mg *SecurityGroupRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroupBinding),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SecurityGroupBindingRef,
		Selector:     mg.Spec.ForProvider.SecurityGroupBindingSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupBinding")
	}
	mg.Spec.ForProvider.SecurityGroupBinding = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecurityGroupBindingRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SecurityGroupBinding),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SecurityGroupBindingRef,
		Selector:     mg.Spec.InitProvider.SecurityGroupBindingSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroupBinding")
	}
	mg.Spec.InitProvider.SecurityGroupBinding = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SecurityGroupBindingRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Subnet.
func (mg *Subnet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.NetworkIDRef,
		Selector:     mg.Spec.InitProvider.NetworkIDSelector,
		To: reference.To{
			List:    &NetworkList{},
			Managed: &Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NetworkID")
	}
	mg.Spec.InitProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkIDRef = rsp.ResolvedReference

	return nil
}
