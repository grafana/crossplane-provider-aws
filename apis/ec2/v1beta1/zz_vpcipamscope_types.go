// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VPCIpamScopeInitParameters struct {

	// A description for the scope you're creating.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the IPAM for which you're creating this scope.
	// +crossplane:generate:reference:type=VPCIpam
	IpamID *string `json:"ipamId,omitempty" tf:"ipam_id,omitempty"`

	// Reference to a VPCIpam to populate ipamId.
	// +kubebuilder:validation:Optional
	IpamIDRef *v1.Reference `json:"ipamIdRef,omitempty" tf:"-"`

	// Selector for a VPCIpam to populate ipamId.
	// +kubebuilder:validation:Optional
	IpamIDSelector *v1.Selector `json:"ipamIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VPCIpamScopeObservation struct {

	// The Amazon Resource Name (ARN) of the scope.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A description for the scope you're creating.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the IPAM Scope.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the IPAM for which you're creating this scope.
	IpamArn *string `json:"ipamArn,omitempty" tf:"ipam_arn,omitempty"`

	// The ID of the IPAM for which you're creating this scope.
	IpamID *string `json:"ipamId,omitempty" tf:"ipam_id,omitempty"`

	// The type of the scope.
	IpamScopeType *string `json:"ipamScopeType,omitempty" tf:"ipam_scope_type,omitempty"`

	// Defines if the scope is the default scope or not.
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// The number of pools in the scope.
	PoolCount *float64 `json:"poolCount,omitempty" tf:"pool_count,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VPCIpamScopeParameters struct {

	// A description for the scope you're creating.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the IPAM for which you're creating this scope.
	// +crossplane:generate:reference:type=VPCIpam
	// +kubebuilder:validation:Optional
	IpamID *string `json:"ipamId,omitempty" tf:"ipam_id,omitempty"`

	// Reference to a VPCIpam to populate ipamId.
	// +kubebuilder:validation:Optional
	IpamIDRef *v1.Reference `json:"ipamIdRef,omitempty" tf:"-"`

	// Selector for a VPCIpam to populate ipamId.
	// +kubebuilder:validation:Optional
	IpamIDSelector *v1.Selector `json:"ipamIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VPCIpamScopeSpec defines the desired state of VPCIpamScope
type VPCIpamScopeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCIpamScopeParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VPCIpamScopeInitParameters `json:"initProvider,omitempty"`
}

// VPCIpamScopeStatus defines the observed state of VPCIpamScope.
type VPCIpamScopeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCIpamScopeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VPCIpamScope is the Schema for the VPCIpamScopes API. Creates a scope for AWS IPAM.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCIpamScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCIpamScopeSpec   `json:"spec"`
	Status            VPCIpamScopeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCIpamScopeList contains a list of VPCIpamScopes
type VPCIpamScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCIpamScope `json:"items"`
}

// Repository type metadata.
var (
	VPCIpamScope_Kind             = "VPCIpamScope"
	VPCIpamScope_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCIpamScope_Kind}.String()
	VPCIpamScope_KindAPIVersion   = VPCIpamScope_Kind + "." + CRDGroupVersion.String()
	VPCIpamScope_GroupVersionKind = CRDGroupVersion.WithKind(VPCIpamScope_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCIpamScope{}, &VPCIpamScopeList{})
}
