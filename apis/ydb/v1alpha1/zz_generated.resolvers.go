// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha12 "github.com/tagesjump/provider-upjet-yc/apis/iam/v1alpha1"
	v1alpha1 "github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1"
	v1alpha11 "github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1"
	iam "github.com/tagesjump/provider-upjet-yc/config/iam"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DatabaseDedicated.
func (mg *DatabaseDedicated) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
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
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetIdsRefs,
		Selector:      mg.Spec.ForProvider.SubnetIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SubnetList{},
			Managed: &v1alpha11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIdsRefs = mrsp.ResolvedReferences

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
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NetworkID")
	}
	mg.Spec.InitProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.SubnetIdsRefs,
		Selector:      mg.Spec.InitProvider.SubnetIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SubnetList{},
			Managed: &v1alpha11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetIds")
	}
	mg.Spec.InitProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SubnetIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this DatabaseIAMBinding.
func (mg *DatabaseIAMBinding) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatabaseID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatabaseIDRef,
		Selector:     mg.Spec.ForProvider.DatabaseIDSelector,
		To: reference.To{
			List:    &DatabaseServerlessList{},
			Managed: &DatabaseServerless{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatabaseID")
	}
	mg.Spec.ForProvider.DatabaseID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Members),
		Extract:       iam.ServiceAccountRefValue(),
		References:    mg.Spec.ForProvider.ServiceAccountRef,
		Selector:      mg.Spec.ForProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &v1alpha12.ServiceAccountList{},
			Managed: &v1alpha12.ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Members")
	}
	mg.Spec.ForProvider.Members = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.ServiceAccountRef = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DatabaseID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.DatabaseIDRef,
		Selector:     mg.Spec.InitProvider.DatabaseIDSelector,
		To: reference.To{
			List:    &DatabaseServerlessList{},
			Managed: &DatabaseServerless{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DatabaseID")
	}
	mg.Spec.InitProvider.DatabaseID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DatabaseIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Members),
		Extract:       iam.ServiceAccountRefValue(),
		References:    mg.Spec.InitProvider.ServiceAccountRef,
		Selector:      mg.Spec.InitProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &v1alpha12.ServiceAccountList{},
			Managed: &v1alpha12.ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Members")
	}
	mg.Spec.InitProvider.Members = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.ServiceAccountRef = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this DatabaseServerless.
func (mg *DatabaseServerless) ResolveReferences(ctx context.Context, c client.Reader) error {
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
