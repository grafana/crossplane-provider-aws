//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthorityObservation) DeepCopyInto(out *CertificateAuthorityObservation) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthorityObservation.
func (in *CertificateAuthorityObservation) DeepCopy() *CertificateAuthorityObservation {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthorityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateAuthorityParameters) DeepCopyInto(out *CertificateAuthorityParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateAuthorityParameters.
func (in *CertificateAuthorityParameters) DeepCopy() *CertificateAuthorityParameters {
	if in == nil {
		return nil
	}
	out := new(CertificateAuthorityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuth) DeepCopyInto(out *ClusterAuth) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuth.
func (in *ClusterAuth) DeepCopy() *ClusterAuth {
	if in == nil {
		return nil
	}
	out := new(ClusterAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAuth) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthList) DeepCopyInto(out *ClusterAuthList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterAuth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthList.
func (in *ClusterAuthList) DeepCopy() *ClusterAuthList {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAuthList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthObservation) DeepCopyInto(out *ClusterAuthObservation) {
	*out = *in
	if in.LastRefreshTime != nil {
		in, out := &in.LastRefreshTime, &out.LastRefreshTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthObservation.
func (in *ClusterAuthObservation) DeepCopy() *ClusterAuthObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthParameters) DeepCopyInto(out *ClusterAuthParameters) {
	*out = *in
	if in.ClusterNameRef != nil {
		in, out := &in.ClusterNameRef, &out.ClusterNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterNameSelector != nil {
		in, out := &in.ClusterNameSelector, &out.ClusterNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.RefreshPeriod != nil {
		in, out := &in.RefreshPeriod, &out.RefreshPeriod
		*out = new(metav1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthParameters.
func (in *ClusterAuthParameters) DeepCopy() *ClusterAuthParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthSpec) DeepCopyInto(out *ClusterAuthSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthSpec.
func (in *ClusterAuthSpec) DeepCopy() *ClusterAuthSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthStatus) DeepCopyInto(out *ClusterAuthStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthStatus.
func (in *ClusterAuthStatus) DeepCopy() *ClusterAuthStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservation) DeepCopyInto(out *ClusterObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CertificateAuthority != nil {
		in, out := &in.CertificateAuthority, &out.CertificateAuthority
		*out = make([]CertificateAuthorityObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KubernetesNetworkConfig != nil {
		in, out := &in.KubernetesNetworkConfig, &out.KubernetesNetworkConfig
		*out = make([]KubernetesNetworkConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlatformVersion != nil {
		in, out := &in.PlatformVersion, &out.PlatformVersion
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VPCConfig != nil {
		in, out := &in.VPCConfig, &out.VPCConfig
		*out = make([]VPCConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservation.
func (in *ClusterObservation) DeepCopy() *ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameters) DeepCopyInto(out *ClusterParameters) {
	*out = *in
	if in.EnabledClusterLogTypes != nil {
		in, out := &in.EnabledClusterLogTypes, &out.EnabledClusterLogTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.EncryptionConfig != nil {
		in, out := &in.EncryptionConfig, &out.EncryptionConfig
		*out = make([]EncryptionConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KubernetesNetworkConfig != nil {
		in, out := &in.KubernetesNetworkConfig, &out.KubernetesNetworkConfig
		*out = make([]KubernetesNetworkConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OutpostConfig != nil {
		in, out := &in.OutpostConfig, &out.OutpostConfig
		*out = make([]OutpostConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.RoleArnRef != nil {
		in, out := &in.RoleArnRef, &out.RoleArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleArnSelector != nil {
		in, out := &in.RoleArnSelector, &out.RoleArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VPCConfig != nil {
		in, out := &in.VPCConfig, &out.VPCConfig
		*out = make([]VPCConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameters.
func (in *ClusterParameters) DeepCopy() *ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlanePlacementObservation) DeepCopyInto(out *ControlPlanePlacementObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlanePlacementObservation.
func (in *ControlPlanePlacementObservation) DeepCopy() *ControlPlanePlacementObservation {
	if in == nil {
		return nil
	}
	out := new(ControlPlanePlacementObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlanePlacementParameters) DeepCopyInto(out *ControlPlanePlacementParameters) {
	*out = *in
	if in.GroupName != nil {
		in, out := &in.GroupName, &out.GroupName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlanePlacementParameters.
func (in *ControlPlanePlacementParameters) DeepCopy() *ControlPlanePlacementParameters {
	if in == nil {
		return nil
	}
	out := new(ControlPlanePlacementParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfigObservation) DeepCopyInto(out *EncryptionConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfigObservation.
func (in *EncryptionConfigObservation) DeepCopy() *EncryptionConfigObservation {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfigParameters) DeepCopyInto(out *EncryptionConfigParameters) {
	*out = *in
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = make([]ProviderParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfigParameters.
func (in *EncryptionConfigParameters) DeepCopy() *EncryptionConfigParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityObservation) DeepCopyInto(out *IdentityObservation) {
	*out = *in
	if in.Oidc != nil {
		in, out := &in.Oidc, &out.Oidc
		*out = make([]OidcObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityObservation.
func (in *IdentityObservation) DeepCopy() *IdentityObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityParameters) DeepCopyInto(out *IdentityParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityParameters.
func (in *IdentityParameters) DeepCopy() *IdentityParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesNetworkConfigObservation) DeepCopyInto(out *KubernetesNetworkConfigObservation) {
	*out = *in
	if in.ServiceIPv6Cidr != nil {
		in, out := &in.ServiceIPv6Cidr, &out.ServiceIPv6Cidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesNetworkConfigObservation.
func (in *KubernetesNetworkConfigObservation) DeepCopy() *KubernetesNetworkConfigObservation {
	if in == nil {
		return nil
	}
	out := new(KubernetesNetworkConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesNetworkConfigParameters) DeepCopyInto(out *KubernetesNetworkConfigParameters) {
	*out = *in
	if in.IPFamily != nil {
		in, out := &in.IPFamily, &out.IPFamily
		*out = new(string)
		**out = **in
	}
	if in.ServiceIPv4Cidr != nil {
		in, out := &in.ServiceIPv4Cidr, &out.ServiceIPv4Cidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesNetworkConfigParameters.
func (in *KubernetesNetworkConfigParameters) DeepCopy() *KubernetesNetworkConfigParameters {
	if in == nil {
		return nil
	}
	out := new(KubernetesNetworkConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcObservation) DeepCopyInto(out *OidcObservation) {
	*out = *in
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcObservation.
func (in *OidcObservation) DeepCopy() *OidcObservation {
	if in == nil {
		return nil
	}
	out := new(OidcObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcParameters) DeepCopyInto(out *OidcParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcParameters.
func (in *OidcParameters) DeepCopy() *OidcParameters {
	if in == nil {
		return nil
	}
	out := new(OidcParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutpostConfigObservation) DeepCopyInto(out *OutpostConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutpostConfigObservation.
func (in *OutpostConfigObservation) DeepCopy() *OutpostConfigObservation {
	if in == nil {
		return nil
	}
	out := new(OutpostConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutpostConfigParameters) DeepCopyInto(out *OutpostConfigParameters) {
	*out = *in
	if in.ControlPlaneInstanceType != nil {
		in, out := &in.ControlPlaneInstanceType, &out.ControlPlaneInstanceType
		*out = new(string)
		**out = **in
	}
	if in.ControlPlanePlacement != nil {
		in, out := &in.ControlPlanePlacement, &out.ControlPlanePlacement
		*out = make([]ControlPlanePlacementParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OutpostArns != nil {
		in, out := &in.OutpostArns, &out.OutpostArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutpostConfigParameters.
func (in *OutpostConfigParameters) DeepCopy() *OutpostConfigParameters {
	if in == nil {
		return nil
	}
	out := new(OutpostConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderObservation) DeepCopyInto(out *ProviderObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderObservation.
func (in *ProviderObservation) DeepCopy() *ProviderObservation {
	if in == nil {
		return nil
	}
	out := new(ProviderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderParameters) DeepCopyInto(out *ProviderParameters) {
	*out = *in
	if in.KeyArn != nil {
		in, out := &in.KeyArn, &out.KeyArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderParameters.
func (in *ProviderParameters) DeepCopy() *ProviderParameters {
	if in == nil {
		return nil
	}
	out := new(ProviderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCConfigObservation) DeepCopyInto(out *VPCConfigObservation) {
	*out = *in
	if in.ClusterSecurityGroupID != nil {
		in, out := &in.ClusterSecurityGroupID, &out.ClusterSecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCConfigObservation.
func (in *VPCConfigObservation) DeepCopy() *VPCConfigObservation {
	if in == nil {
		return nil
	}
	out := new(VPCConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCConfigParameters) DeepCopyInto(out *VPCConfigParameters) {
	*out = *in
	if in.EndpointPrivateAccess != nil {
		in, out := &in.EndpointPrivateAccess, &out.EndpointPrivateAccess
		*out = new(bool)
		**out = **in
	}
	if in.EndpointPublicAccess != nil {
		in, out := &in.EndpointPublicAccess, &out.EndpointPublicAccess
		*out = new(bool)
		**out = **in
	}
	if in.PublicAccessCidrs != nil {
		in, out := &in.PublicAccessCidrs, &out.PublicAccessCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityGroupIDRefs != nil {
		in, out := &in.SecurityGroupIDRefs, &out.SecurityGroupIDRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityGroupIDSelector != nil {
		in, out := &in.SecurityGroupIDSelector, &out.SecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIDRefs != nil {
		in, out := &in.SubnetIDRefs, &out.SubnetIDRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCConfigParameters.
func (in *VPCConfigParameters) DeepCopy() *VPCConfigParameters {
	if in == nil {
		return nil
	}
	out := new(VPCConfigParameters)
	in.DeepCopyInto(out)
	return out
}
