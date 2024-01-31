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

type ResourceAssociationInitParameters struct {

	// Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Amazon Resource Name (ARN) of the RAM Resource Share.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ram/v1beta1.ResourceShare
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	ResourceShareArn *string `json:"resourceShareArn,omitempty" tf:"resource_share_arn,omitempty"`

	// Reference to a ResourceShare in ram to populate resourceShareArn.
	// +kubebuilder:validation:Optional
	ResourceShareArnRef *v1.Reference `json:"resourceShareArnRef,omitempty" tf:"-"`

	// Selector for a ResourceShare in ram to populate resourceShareArn.
	// +kubebuilder:validation:Optional
	ResourceShareArnSelector *v1.Selector `json:"resourceShareArnSelector,omitempty" tf:"-"`
}

type ResourceAssociationObservation struct {

	// The Amazon Resource Name (ARN) of the resource share.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Amazon Resource Name (ARN) of the RAM Resource Share.
	ResourceShareArn *string `json:"resourceShareArn,omitempty" tf:"resource_share_arn,omitempty"`
}

type ResourceAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Amazon Resource Name (ARN) of the RAM Resource Share.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ram/v1beta1.ResourceShare
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ResourceShareArn *string `json:"resourceShareArn,omitempty" tf:"resource_share_arn,omitempty"`

	// Reference to a ResourceShare in ram to populate resourceShareArn.
	// +kubebuilder:validation:Optional
	ResourceShareArnRef *v1.Reference `json:"resourceShareArnRef,omitempty" tf:"-"`

	// Selector for a ResourceShare in ram to populate resourceShareArn.
	// +kubebuilder:validation:Optional
	ResourceShareArnSelector *v1.Selector `json:"resourceShareArnSelector,omitempty" tf:"-"`
}

// ResourceAssociationSpec defines the desired state of ResourceAssociation
type ResourceAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceAssociationParameters `json:"forProvider"`
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
	InitProvider ResourceAssociationInitParameters `json:"initProvider,omitempty"`
}

// ResourceAssociationStatus defines the observed state of ResourceAssociation.
type ResourceAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ResourceAssociation is the Schema for the ResourceAssociations API. Manages a Resource Access Manager (RAM) Resource Association.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResourceAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceArn) || (has(self.initProvider) && has(self.initProvider.resourceArn))",message="spec.forProvider.resourceArn is a required parameter"
	Spec   ResourceAssociationSpec   `json:"spec"`
	Status ResourceAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceAssociationList contains a list of ResourceAssociations
type ResourceAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceAssociation `json:"items"`
}

// Repository type metadata.
var (
	ResourceAssociation_Kind             = "ResourceAssociation"
	ResourceAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceAssociation_Kind}.String()
	ResourceAssociation_KindAPIVersion   = ResourceAssociation_Kind + "." + CRDGroupVersion.String()
	ResourceAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ResourceAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceAssociation{}, &ResourceAssociationList{})
}
