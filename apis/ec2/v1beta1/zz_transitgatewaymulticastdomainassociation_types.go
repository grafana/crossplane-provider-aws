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

type TransitGatewayMulticastDomainAssociationInitParameters struct {

	// The ID of the subnet to associate with the transit gateway multicast domain.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The ID of the transit gateway attachment.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGatewayVPCAttachment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// Reference to a TransitGatewayVPCAttachment in ec2 to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDRef *v1.Reference `json:"transitGatewayAttachmentIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayVPCAttachment in ec2 to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDSelector *v1.Selector `json:"transitGatewayAttachmentIdSelector,omitempty" tf:"-"`

	// The ID of the transit gateway multicast domain.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGatewayMulticastDomain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	TransitGatewayMulticastDomainID *string `json:"transitGatewayMulticastDomainId,omitempty" tf:"transit_gateway_multicast_domain_id,omitempty"`

	// Reference to a TransitGatewayMulticastDomain in ec2 to populate transitGatewayMulticastDomainId.
	// +kubebuilder:validation:Optional
	TransitGatewayMulticastDomainIDRef *v1.Reference `json:"transitGatewayMulticastDomainIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayMulticastDomain in ec2 to populate transitGatewayMulticastDomainId.
	// +kubebuilder:validation:Optional
	TransitGatewayMulticastDomainIDSelector *v1.Selector `json:"transitGatewayMulticastDomainIdSelector,omitempty" tf:"-"`
}

type TransitGatewayMulticastDomainAssociationObservation struct {

	// EC2 Transit Gateway Multicast Domain Association identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the subnet to associate with the transit gateway multicast domain.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The ID of the transit gateway attachment.
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainID *string `json:"transitGatewayMulticastDomainId,omitempty" tf:"transit_gateway_multicast_domain_id,omitempty"`
}

type TransitGatewayMulticastDomainAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the subnet to associate with the transit gateway multicast domain.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The ID of the transit gateway attachment.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGatewayVPCAttachment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// Reference to a TransitGatewayVPCAttachment in ec2 to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDRef *v1.Reference `json:"transitGatewayAttachmentIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayVPCAttachment in ec2 to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDSelector *v1.Selector `json:"transitGatewayAttachmentIdSelector,omitempty" tf:"-"`

	// The ID of the transit gateway multicast domain.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGatewayMulticastDomain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TransitGatewayMulticastDomainID *string `json:"transitGatewayMulticastDomainId,omitempty" tf:"transit_gateway_multicast_domain_id,omitempty"`

	// Reference to a TransitGatewayMulticastDomain in ec2 to populate transitGatewayMulticastDomainId.
	// +kubebuilder:validation:Optional
	TransitGatewayMulticastDomainIDRef *v1.Reference `json:"transitGatewayMulticastDomainIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayMulticastDomain in ec2 to populate transitGatewayMulticastDomainId.
	// +kubebuilder:validation:Optional
	TransitGatewayMulticastDomainIDSelector *v1.Selector `json:"transitGatewayMulticastDomainIdSelector,omitempty" tf:"-"`
}

// TransitGatewayMulticastDomainAssociationSpec defines the desired state of TransitGatewayMulticastDomainAssociation
type TransitGatewayMulticastDomainAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayMulticastDomainAssociationParameters `json:"forProvider"`
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
	InitProvider TransitGatewayMulticastDomainAssociationInitParameters `json:"initProvider,omitempty"`
}

// TransitGatewayMulticastDomainAssociationStatus defines the observed state of TransitGatewayMulticastDomainAssociation.
type TransitGatewayMulticastDomainAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayMulticastDomainAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TransitGatewayMulticastDomainAssociation is the Schema for the TransitGatewayMulticastDomainAssociations API. Manages an EC2 Transit Gateway Multicast Domain Association
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayMulticastDomainAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayMulticastDomainAssociationSpec   `json:"spec"`
	Status            TransitGatewayMulticastDomainAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayMulticastDomainAssociationList contains a list of TransitGatewayMulticastDomainAssociations
type TransitGatewayMulticastDomainAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayMulticastDomainAssociation `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayMulticastDomainAssociation_Kind             = "TransitGatewayMulticastDomainAssociation"
	TransitGatewayMulticastDomainAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayMulticastDomainAssociation_Kind}.String()
	TransitGatewayMulticastDomainAssociation_KindAPIVersion   = TransitGatewayMulticastDomainAssociation_Kind + "." + CRDGroupVersion.String()
	TransitGatewayMulticastDomainAssociation_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayMulticastDomainAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayMulticastDomainAssociation{}, &TransitGatewayMulticastDomainAssociationList{})
}
