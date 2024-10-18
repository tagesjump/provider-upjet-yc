// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1"
	iam "github.com/tagesjump/provider-upjet-yc/config/iam"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ServiceAccount.
func (mg *ServiceAccount) ResolveReferences(ctx context.Context, c client.Reader) error {
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

// ResolveReferences of this ServiceAccountAPIKey.
func (mg *ServiceAccountAPIKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccountID")
	}
	mg.Spec.InitProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceAccountIAMBinding.
func (mg *ServiceAccountIAMBinding) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Members),
		Extract:       iam.ServiceAccountRefValue(),
		References:    mg.Spec.ForProvider.ServiceAccountRef,
		Selector:      mg.Spec.ForProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Members")
	}
	mg.Spec.ForProvider.Members = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.ServiceAccountRef = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Members),
		Extract:       iam.ServiceAccountRefValue(),
		References:    mg.Spec.InitProvider.ServiceAccountRef,
		Selector:      mg.Spec.InitProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Members")
	}
	mg.Spec.InitProvider.Members = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.ServiceAccountRef = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccountID")
	}
	mg.Spec.InitProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceAccountIAMMember.
func (mg *ServiceAccountIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Member),
		Extract:      iam.ServiceAccountRefValue(),
		Reference:    mg.Spec.ForProvider.ServiceAccountRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Member")
	}
	mg.Spec.ForProvider.Member = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Member),
		Extract:      iam.ServiceAccountRefValue(),
		Reference:    mg.Spec.InitProvider.ServiceAccountRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Member")
	}
	mg.Spec.InitProvider.Member = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccountID")
	}
	mg.Spec.InitProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceAccountIAMPolicy.
func (mg *ServiceAccountIAMPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccountID")
	}
	mg.Spec.InitProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceAccountKey.
func (mg *ServiceAccountKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccountID")
	}
	mg.Spec.InitProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceAccountStaticAccessKey.
func (mg *ServiceAccountStaticAccessKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &ServiceAccountList{},
			Managed: &ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccountID")
	}
	mg.Spec.InitProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountIDRef = rsp.ResolvedReference

	return nil
}
