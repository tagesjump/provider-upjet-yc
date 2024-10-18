// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1alpha13 "github.com/tagesjump/provider-upjet-yc/apis/iam/v1alpha1"
	v1alpha12 "github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1"
	v1alpha1 "github.com/tagesjump/provider-upjet-yc/apis/storage/v1alpha1"
	v1alpha11 "github.com/tagesjump/provider-upjet-yc/apis/vpc/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", false),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &v1alpha1.BucketList{},
			Managed: &v1alpha1.Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ClusterConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ClusterConfig[i3].SubclusterSpec); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterConfig[i3].SubclusterSpec[i4].SubnetID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.ClusterConfig[i3].SubclusterSpec[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.ClusterConfig[i3].SubclusterSpec[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha11.SubnetList{},
					Managed: &v1alpha11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ClusterConfig[i3].SubclusterSpec[i4].SubnetID")
			}
			mg.Spec.ForProvider.ClusterConfig[i3].SubclusterSpec[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ClusterConfig[i3].SubclusterSpec[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha12.FolderList{},
			Managed: &v1alpha12.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &v1alpha13.ServiceAccountList{},
			Managed: &v1alpha13.ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", false),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &v1alpha1.BucketList{},
			Managed: &v1alpha1.Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.ClusterConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.ClusterConfig[i3].SubclusterSpec); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterConfig[i3].SubclusterSpec[i4].SubnetID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.ClusterConfig[i3].SubclusterSpec[i4].SubnetIDRef,
				Selector:     mg.Spec.InitProvider.ClusterConfig[i3].SubclusterSpec[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha11.SubnetList{},
					Managed: &v1alpha11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.ClusterConfig[i3].SubclusterSpec[i4].SubnetID")
			}
			mg.Spec.InitProvider.ClusterConfig[i3].SubclusterSpec[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.ClusterConfig[i3].SubclusterSpec[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha12.FolderList{},
			Managed: &v1alpha12.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccountID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &v1alpha13.ServiceAccountList{},
			Managed: &v1alpha13.ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccountID")
	}
	mg.Spec.InitProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountIDRef = rsp.ResolvedReference

	return nil
}