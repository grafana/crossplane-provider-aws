/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *ReportDefinition) ResolveReferences( // ResolveReferences of this ReportDefinition.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.S3Bucket),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.S3BucketRef,
			Selector:     mg.Spec.ForProvider.S3BucketSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.S3Bucket")
	}
	mg.Spec.ForProvider.S3Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.S3BucketRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.S3Bucket),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.S3BucketRef,
			Selector:     mg.Spec.InitProvider.S3BucketSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.S3Bucket")
	}
	mg.Spec.InitProvider.S3Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.S3BucketRef = rsp.ResolvedReference

	return nil
}
