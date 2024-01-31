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

type DefaultPatchBaselineInitParameters struct {

	// ID of the patch baseline.
	// Can be an ID or an ARN.
	// When specifying an AWS-provided patch baseline, must be the ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ssm/v1beta1.PatchBaseline
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	BaselineID *string `json:"baselineId,omitempty" tf:"baseline_id,omitempty"`

	// Reference to a PatchBaseline in ssm to populate baselineId.
	// +kubebuilder:validation:Optional
	BaselineIDRef *v1.Reference `json:"baselineIdRef,omitempty" tf:"-"`

	// Selector for a PatchBaseline in ssm to populate baselineId.
	// +kubebuilder:validation:Optional
	BaselineIDSelector *v1.Selector `json:"baselineIdSelector,omitempty" tf:"-"`

	// The operating system the patch baseline applies to.
	// Valid values are
	// AMAZON_LINUX,
	// AMAZON_LINUX_2,
	// AMAZON_LINUX_2022,
	// CENTOS,
	// DEBIAN,
	// MACOS,
	// ORACLE_LINUX,
	// RASPBIAN,
	// REDHAT_ENTERPRISE_LINUX,
	// ROCKY_LINUX,
	// SUSE,
	// UBUNTU, and
	// WINDOWS.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ssm/v1beta1.PatchBaseline
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("operating_system",false)
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// Reference to a PatchBaseline in ssm to populate operatingSystem.
	// +kubebuilder:validation:Optional
	OperatingSystemRef *v1.Reference `json:"operatingSystemRef,omitempty" tf:"-"`

	// Selector for a PatchBaseline in ssm to populate operatingSystem.
	// +kubebuilder:validation:Optional
	OperatingSystemSelector *v1.Selector `json:"operatingSystemSelector,omitempty" tf:"-"`
}

type DefaultPatchBaselineObservation struct {

	// ID of the patch baseline.
	// Can be an ID or an ARN.
	// When specifying an AWS-provided patch baseline, must be the ARN.
	BaselineID *string `json:"baselineId,omitempty" tf:"baseline_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The operating system the patch baseline applies to.
	// Valid values are
	// AMAZON_LINUX,
	// AMAZON_LINUX_2,
	// AMAZON_LINUX_2022,
	// CENTOS,
	// DEBIAN,
	// MACOS,
	// ORACLE_LINUX,
	// RASPBIAN,
	// REDHAT_ENTERPRISE_LINUX,
	// ROCKY_LINUX,
	// SUSE,
	// UBUNTU, and
	// WINDOWS.
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`
}

type DefaultPatchBaselineParameters struct {

	// ID of the patch baseline.
	// Can be an ID or an ARN.
	// When specifying an AWS-provided patch baseline, must be the ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ssm/v1beta1.PatchBaseline
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	BaselineID *string `json:"baselineId,omitempty" tf:"baseline_id,omitempty"`

	// Reference to a PatchBaseline in ssm to populate baselineId.
	// +kubebuilder:validation:Optional
	BaselineIDRef *v1.Reference `json:"baselineIdRef,omitempty" tf:"-"`

	// Selector for a PatchBaseline in ssm to populate baselineId.
	// +kubebuilder:validation:Optional
	BaselineIDSelector *v1.Selector `json:"baselineIdSelector,omitempty" tf:"-"`

	// The operating system the patch baseline applies to.
	// Valid values are
	// AMAZON_LINUX,
	// AMAZON_LINUX_2,
	// AMAZON_LINUX_2022,
	// CENTOS,
	// DEBIAN,
	// MACOS,
	// ORACLE_LINUX,
	// RASPBIAN,
	// REDHAT_ENTERPRISE_LINUX,
	// ROCKY_LINUX,
	// SUSE,
	// UBUNTU, and
	// WINDOWS.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ssm/v1beta1.PatchBaseline
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("operating_system",false)
	// +kubebuilder:validation:Optional
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system,omitempty"`

	// Reference to a PatchBaseline in ssm to populate operatingSystem.
	// +kubebuilder:validation:Optional
	OperatingSystemRef *v1.Reference `json:"operatingSystemRef,omitempty" tf:"-"`

	// Selector for a PatchBaseline in ssm to populate operatingSystem.
	// +kubebuilder:validation:Optional
	OperatingSystemSelector *v1.Selector `json:"operatingSystemSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DefaultPatchBaselineSpec defines the desired state of DefaultPatchBaseline
type DefaultPatchBaselineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultPatchBaselineParameters `json:"forProvider"`
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
	InitProvider DefaultPatchBaselineInitParameters `json:"initProvider,omitempty"`
}

// DefaultPatchBaselineStatus defines the observed state of DefaultPatchBaseline.
type DefaultPatchBaselineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultPatchBaselineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DefaultPatchBaseline is the Schema for the DefaultPatchBaselines API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DefaultPatchBaseline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultPatchBaselineSpec   `json:"spec"`
	Status            DefaultPatchBaselineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultPatchBaselineList contains a list of DefaultPatchBaselines
type DefaultPatchBaselineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultPatchBaseline `json:"items"`
}

// Repository type metadata.
var (
	DefaultPatchBaseline_Kind             = "DefaultPatchBaseline"
	DefaultPatchBaseline_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultPatchBaseline_Kind}.String()
	DefaultPatchBaseline_KindAPIVersion   = DefaultPatchBaseline_Kind + "." + CRDGroupVersion.String()
	DefaultPatchBaseline_GroupVersionKind = CRDGroupVersion.WithKind(DefaultPatchBaseline_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultPatchBaseline{}, &DefaultPatchBaselineList{})
}
