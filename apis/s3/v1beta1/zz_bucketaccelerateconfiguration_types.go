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

type BucketAccelerateConfigurationInitParameters struct {

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Transfer acceleration state of the bucket. Valid values: Enabled, Suspended.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BucketAccelerateConfigurationObservation struct {

	// Name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// The bucket or bucket and expected_bucket_owner separated by a comma (,) if the latter is provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Transfer acceleration state of the bucket. Valid values: Enabled, Suspended.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BucketAccelerateConfigurationParameters struct {

	// Name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Account ID of the expected bucket owner.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Transfer acceleration state of the bucket. Valid values: Enabled, Suspended.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

// BucketAccelerateConfigurationSpec defines the desired state of BucketAccelerateConfiguration
type BucketAccelerateConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketAccelerateConfigurationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider BucketAccelerateConfigurationInitParameters `json:"initProvider,omitempty"`
}

// BucketAccelerateConfigurationStatus defines the observed state of BucketAccelerateConfiguration.
type BucketAccelerateConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketAccelerateConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketAccelerateConfiguration is the Schema for the BucketAccelerateConfigurations API. Provides an S3 bucket accelerate configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketAccelerateConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.status) || (has(self.initProvider) && has(self.initProvider.status))",message="spec.forProvider.status is a required parameter"
	Spec   BucketAccelerateConfigurationSpec   `json:"spec"`
	Status BucketAccelerateConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketAccelerateConfigurationList contains a list of BucketAccelerateConfigurations
type BucketAccelerateConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketAccelerateConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketAccelerateConfiguration_Kind             = "BucketAccelerateConfiguration"
	BucketAccelerateConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketAccelerateConfiguration_Kind}.String()
	BucketAccelerateConfiguration_KindAPIVersion   = BucketAccelerateConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketAccelerateConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketAccelerateConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketAccelerateConfiguration{}, &BucketAccelerateConfigurationList{})
}
