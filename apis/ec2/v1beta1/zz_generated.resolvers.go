/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/kms/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this EIP.
func (mg *EIP) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Instance),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceRef,
		Selector:     mg.Spec.ForProvider.InstanceSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instance")
	}
	mg.Spec.ForProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkInterface),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkInterfaceRef,
		Selector:     mg.Spec.ForProvider.NetworkInterfaceSelector,
		To: reference.To{
			List:    &NetworkInterfaceList{},
			Managed: &NetworkInterface{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkInterface")
	}
	mg.Spec.ForProvider.NetworkInterface = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkInterfaceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EBSBlockDevice); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EBSBlockDevice[i3].KMSKeyID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EBSBlockDevice[i3].KMSKeyIDRef,
			Selector:     mg.Spec.ForProvider.EBSBlockDevice[i3].KMSKeyIDSelector,
			To: reference.To{
				List:    &v1beta1.KeyList{},
				Managed: &v1beta1.Key{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EBSBlockDevice[i3].KMSKeyID")
		}
		mg.Spec.ForProvider.EBSBlockDevice[i3].KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EBSBlockDevice[i3].KMSKeyIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkInterface); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkInterface[i3].NetworkInterfaceID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NetworkInterface[i3].NetworkInterfaceIDRef,
			Selector:     mg.Spec.ForProvider.NetworkInterface[i3].NetworkInterfaceIDSelector,
			To: reference.To{
				List:    &NetworkInterfaceList{},
				Managed: &NetworkInterface{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkInterface[i3].NetworkInterfaceID")
		}
		mg.Spec.ForProvider.NetworkInterface[i3].NetworkInterfaceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NetworkInterface[i3].NetworkInterfaceIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.RootBlockDevice); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RootBlockDevice[i3].KMSKeyID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RootBlockDevice[i3].KMSKeyIDRef,
			Selector:     mg.Spec.ForProvider.RootBlockDevice[i3].KMSKeyIDSelector,
			To: reference.To{
				List:    &v1beta1.KeyList{},
				Managed: &v1beta1.Key{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.RootBlockDevice[i3].KMSKeyID")
		}
		mg.Spec.ForProvider.RootBlockDevice[i3].KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.RootBlockDevice[i3].KMSKeyIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VPCSecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.VPCSecurityGroupIDSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCSecurityGroupIds")
	}
	mg.Spec.ForProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCSecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this NetworkInterface.
func (mg *NetworkInterface) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroups")
	}
	mg.Spec.ForProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityGroup.
func (mg *SecurityGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &VPCList{},
			Managed: &VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityGroupRule.
func (mg *SecurityGroupRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.PrefixListIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.PrefixListIDRefs,
		Selector:      mg.Spec.ForProvider.PrefixListIDSelector,
		To: reference.To{
			List:    &ManagedPrefixListList{},
			Managed: &ManagedPrefixList{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrefixListIds")
	}
	mg.Spec.ForProvider.PrefixListIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.PrefixListIDRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupID")
	}
	mg.Spec.ForProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecurityGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceSecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SourceSecurityGroupIDRef,
		Selector:     mg.Spec.ForProvider.SourceSecurityGroupIDSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceSecurityGroupID")
	}
	mg.Spec.ForProvider.SourceSecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceSecurityGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Subnet.
func (mg *Subnet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &VPCList{},
			Managed: &VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VPCEndpoint.
func (mg *VPCEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceName),
		Extract:      resource.ExtractParamPath("service_name", true),
		Reference:    mg.Spec.ForProvider.ServiceNameRef,
		Selector:     mg.Spec.ForProvider.ServiceNameSelector,
		To: reference.To{
			List:    &VPCEndpointServiceList{},
			Managed: &VPCEndpointService{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceName")
	}
	mg.Spec.ForProvider.ServiceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &VPCList{},
			Managed: &VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VPCEndpointConnectionNotification.
func (mg *VPCEndpointConnectionNotification) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCEndpointServiceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VPCEndpointServiceIDRef,
		Selector:     mg.Spec.ForProvider.VPCEndpointServiceIDSelector,
		To: reference.To{
			List:    &VPCEndpointServiceList{},
			Managed: &VPCEndpointService{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCEndpointServiceID")
	}
	mg.Spec.ForProvider.VPCEndpointServiceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCEndpointServiceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VPCEndpointRouteTableAssociation.
func (mg *VPCEndpointRouteTableAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCEndpointID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VPCEndpointIDRef,
		Selector:     mg.Spec.ForProvider.VPCEndpointIDSelector,
		To: reference.To{
			List:    &VPCEndpointList{},
			Managed: &VPCEndpoint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCEndpointID")
	}
	mg.Spec.ForProvider.VPCEndpointID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCEndpointIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VPCEndpointServiceAllowedPrincipal.
func (mg *VPCEndpointServiceAllowedPrincipal) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCEndpointServiceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VPCEndpointServiceIDRef,
		Selector:     mg.Spec.ForProvider.VPCEndpointServiceIDSelector,
		To: reference.To{
			List:    &VPCEndpointServiceList{},
			Managed: &VPCEndpointService{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCEndpointServiceID")
	}
	mg.Spec.ForProvider.VPCEndpointServiceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCEndpointServiceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VPCEndpointSubnetAssociation.
func (mg *VPCEndpointSubnetAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCEndpointID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VPCEndpointIDRef,
		Selector:     mg.Spec.ForProvider.VPCEndpointIDSelector,
		To: reference.To{
			List:    &VPCEndpointList{},
			Managed: &VPCEndpoint{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCEndpointID")
	}
	mg.Spec.ForProvider.VPCEndpointID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCEndpointIDRef = rsp.ResolvedReference

	return nil
}
