// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/tagesjump/provider-upjet-yc/apis/resourcemanager/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Dashboard.
func (mg *Dashboard) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Parametrization); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Parametrization[i3].Parameters); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Parametrization[i3].Parameters[i4].LabelValues); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Parametrization[i3].Parameters[i4].LabelValues[i5].FolderID),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.Parametrization[i3].Parameters[i4].LabelValues[i5].FolderIDRef,
					Selector:     mg.Spec.ForProvider.Parametrization[i3].Parameters[i4].LabelValues[i5].FolderIDSelector,
					To: reference.To{
						List:    &v1alpha1.FolderList{},
						Managed: &v1alpha1.Folder{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Parametrization[i3].Parameters[i4].LabelValues[i5].FolderID")
				}
				mg.Spec.ForProvider.Parametrization[i3].Parameters[i4].LabelValues[i5].FolderID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Parametrization[i3].Parameters[i4].LabelValues[i5].FolderIDRef = rsp.ResolvedReference

			}
		}
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

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Parametrization); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Parametrization[i3].Parameters); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Parametrization[i3].Parameters[i4].LabelValues); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Parametrization[i3].Parameters[i4].LabelValues[i5].FolderID),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.Parametrization[i3].Parameters[i4].LabelValues[i5].FolderIDRef,
					Selector:     mg.Spec.InitProvider.Parametrization[i3].Parameters[i4].LabelValues[i5].FolderIDSelector,
					To: reference.To{
						List:    &v1alpha1.FolderList{},
						Managed: &v1alpha1.Folder{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.Parametrization[i3].Parameters[i4].LabelValues[i5].FolderID")
				}
				mg.Spec.InitProvider.Parametrization[i3].Parameters[i4].LabelValues[i5].FolderID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.Parametrization[i3].Parameters[i4].LabelValues[i5].FolderIDRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}
