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

type DomainServiceAccessPolicyInitParameters struct {

	// The access rules you want to configure. These rules replace any existing rules. See the AWS documentation for details.
	AccessPolicy *string `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// The CloudSearch domain name the policy applies to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudsearch/v1beta1.Domain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Reference to a Domain in cloudsearch to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameRef *v1.Reference `json:"domainNameRef,omitempty" tf:"-"`

	// Selector for a Domain in cloudsearch to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameSelector *v1.Selector `json:"domainNameSelector,omitempty" tf:"-"`
}

type DomainServiceAccessPolicyObservation struct {

	// The access rules you want to configure. These rules replace any existing rules. See the AWS documentation for details.
	AccessPolicy *string `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// The CloudSearch domain name the policy applies to.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DomainServiceAccessPolicyParameters struct {

	// The access rules you want to configure. These rules replace any existing rules. See the AWS documentation for details.
	// +kubebuilder:validation:Optional
	AccessPolicy *string `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`

	// The CloudSearch domain name the policy applies to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudsearch/v1beta1.Domain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Reference to a Domain in cloudsearch to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameRef *v1.Reference `json:"domainNameRef,omitempty" tf:"-"`

	// Selector for a Domain in cloudsearch to populate domainName.
	// +kubebuilder:validation:Optional
	DomainNameSelector *v1.Selector `json:"domainNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DomainServiceAccessPolicySpec defines the desired state of DomainServiceAccessPolicy
type DomainServiceAccessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainServiceAccessPolicyParameters `json:"forProvider"`
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
	InitProvider DomainServiceAccessPolicyInitParameters `json:"initProvider,omitempty"`
}

// DomainServiceAccessPolicyStatus defines the observed state of DomainServiceAccessPolicy.
type DomainServiceAccessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainServiceAccessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DomainServiceAccessPolicy is the Schema for the DomainServiceAccessPolicys API. Provides an CloudSearch domain service access policy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DomainServiceAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessPolicy) || (has(self.initProvider) && has(self.initProvider.accessPolicy))",message="spec.forProvider.accessPolicy is a required parameter"
	Spec   DomainServiceAccessPolicySpec   `json:"spec"`
	Status DomainServiceAccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainServiceAccessPolicyList contains a list of DomainServiceAccessPolicys
type DomainServiceAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainServiceAccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	DomainServiceAccessPolicy_Kind             = "DomainServiceAccessPolicy"
	DomainServiceAccessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainServiceAccessPolicy_Kind}.String()
	DomainServiceAccessPolicy_KindAPIVersion   = DomainServiceAccessPolicy_Kind + "." + CRDGroupVersion.String()
	DomainServiceAccessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DomainServiceAccessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainServiceAccessPolicy{}, &DomainServiceAccessPolicyList{})
}
