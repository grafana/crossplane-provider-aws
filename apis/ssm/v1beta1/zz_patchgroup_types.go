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

type PatchGroupInitParameters struct {

	// The ID of the patch baseline to register the patch group with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ssm/v1beta1.PatchBaseline
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	BaselineID *string `json:"baselineId,omitempty" tf:"baseline_id,omitempty"`

	// Reference to a PatchBaseline in ssm to populate baselineId.
	// +kubebuilder:validation:Optional
	BaselineIDRef *v1.Reference `json:"baselineIdRef,omitempty" tf:"-"`

	// Selector for a PatchBaseline in ssm to populate baselineId.
	// +kubebuilder:validation:Optional
	BaselineIDSelector *v1.Selector `json:"baselineIdSelector,omitempty" tf:"-"`

	// The name of the patch group that should be registered with the patch baseline.
	PatchGroup *string `json:"patchGroup,omitempty" tf:"patch_group,omitempty"`
}

type PatchGroupObservation struct {

	// The ID of the patch baseline to register the patch group with.
	BaselineID *string `json:"baselineId,omitempty" tf:"baseline_id,omitempty"`

	// The name of the patch group and ID of the patch baseline separated by a comma (,).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the patch group that should be registered with the patch baseline.
	PatchGroup *string `json:"patchGroup,omitempty" tf:"patch_group,omitempty"`
}

type PatchGroupParameters struct {

	// The ID of the patch baseline to register the patch group with.
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

	// The name of the patch group that should be registered with the patch baseline.
	// +kubebuilder:validation:Optional
	PatchGroup *string `json:"patchGroup,omitempty" tf:"patch_group,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// PatchGroupSpec defines the desired state of PatchGroup
type PatchGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PatchGroupParameters `json:"forProvider"`
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
	InitProvider PatchGroupInitParameters `json:"initProvider,omitempty"`
}

// PatchGroupStatus defines the observed state of PatchGroup.
type PatchGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PatchGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// PatchGroup is the Schema for the PatchGroups API. Provides an SSM Patch Group resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PatchGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.patchGroup) || (has(self.initProvider) && has(self.initProvider.patchGroup))",message="spec.forProvider.patchGroup is a required parameter"
	Spec   PatchGroupSpec   `json:"spec"`
	Status PatchGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PatchGroupList contains a list of PatchGroups
type PatchGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PatchGroup `json:"items"`
}

// Repository type metadata.
var (
	PatchGroup_Kind             = "PatchGroup"
	PatchGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PatchGroup_Kind}.String()
	PatchGroup_KindAPIVersion   = PatchGroup_Kind + "." + CRDGroupVersion.String()
	PatchGroup_GroupVersionKind = CRDGroupVersion.WithKind(PatchGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&PatchGroup{}, &PatchGroupList{})
}
