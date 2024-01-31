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

type SAMLProviderInitParameters struct {

	// An XML document generated by an identity provider that supports SAML 2.0.
	SAMLMetadataDocument *string `json:"samlMetadataDocument,omitempty" tf:"saml_metadata_document,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SAMLProviderObservation struct {

	// The ARN assigned by AWS for this provider.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An XML document generated by an identity provider that supports SAML 2.0.
	SAMLMetadataDocument *string `json:"samlMetadataDocument,omitempty" tf:"saml_metadata_document,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The expiration date and time for the SAML provider in RFC1123 format, e.g., Mon, 02 Jan 2006 15:04:05 MST.
	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until,omitempty"`
}

type SAMLProviderParameters struct {

	// An XML document generated by an identity provider that supports SAML 2.0.
	// +kubebuilder:validation:Optional
	SAMLMetadataDocument *string `json:"samlMetadataDocument,omitempty" tf:"saml_metadata_document,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SAMLProviderSpec defines the desired state of SAMLProvider
type SAMLProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SAMLProviderParameters `json:"forProvider"`
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
	InitProvider SAMLProviderInitParameters `json:"initProvider,omitempty"`
}

// SAMLProviderStatus defines the observed state of SAMLProvider.
type SAMLProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SAMLProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SAMLProvider is the Schema for the SAMLProviders API. Provides an IAM SAML provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SAMLProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.samlMetadataDocument) || (has(self.initProvider) && has(self.initProvider.samlMetadataDocument))",message="spec.forProvider.samlMetadataDocument is a required parameter"
	Spec   SAMLProviderSpec   `json:"spec"`
	Status SAMLProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SAMLProviderList contains a list of SAMLProviders
type SAMLProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SAMLProvider `json:"items"`
}

// Repository type metadata.
var (
	SAMLProvider_Kind             = "SAMLProvider"
	SAMLProvider_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SAMLProvider_Kind}.String()
	SAMLProvider_KindAPIVersion   = SAMLProvider_Kind + "." + CRDGroupVersion.String()
	SAMLProvider_GroupVersionKind = CRDGroupVersion.WithKind(SAMLProvider_Kind)
)

func init() {
	SchemeBuilder.Register(&SAMLProvider{}, &SAMLProviderList{})
}
