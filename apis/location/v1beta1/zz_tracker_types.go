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

type TrackerInitParameters struct {

	// The optional description for the tracker resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The position filtering method of the tracker resource. Valid values: TimeBased, DistanceBased, AccuracyBased. Default: TimeBased.
	PositionFiltering *string `json:"positionFiltering,omitempty" tf:"position_filtering,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TrackerObservation struct {

	// The timestamp for when the tracker resource was created in ISO 8601 format.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The optional description for the tracker resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The position filtering method of the tracker resource. Valid values: TimeBased, DistanceBased, AccuracyBased. Default: TimeBased.
	PositionFiltering *string `json:"positionFiltering,omitempty" tf:"position_filtering,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The Amazon Resource Name (ARN) for the tracker resource. Used when you need to specify a resource across all AWS.
	TrackerArn *string `json:"trackerArn,omitempty" tf:"tracker_arn,omitempty"`

	// The timestamp for when the tracker resource was last updated in ISO 8601 format.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type TrackerParameters struct {

	// The optional description for the tracker resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The position filtering method of the tracker resource. Valid values: TimeBased, DistanceBased, AccuracyBased. Default: TimeBased.
	// +kubebuilder:validation:Optional
	PositionFiltering *string `json:"positionFiltering,omitempty" tf:"position_filtering,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TrackerSpec defines the desired state of Tracker
type TrackerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrackerParameters `json:"forProvider"`
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
	InitProvider TrackerInitParameters `json:"initProvider,omitempty"`
}

// TrackerStatus defines the observed state of Tracker.
type TrackerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrackerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Tracker is the Schema for the Trackers API. Provides a Location Service Tracker.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Tracker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrackerSpec   `json:"spec"`
	Status            TrackerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrackerList contains a list of Trackers
type TrackerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tracker `json:"items"`
}

// Repository type metadata.
var (
	Tracker_Kind             = "Tracker"
	Tracker_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Tracker_Kind}.String()
	Tracker_KindAPIVersion   = Tracker_Kind + "." + CRDGroupVersion.String()
	Tracker_GroupVersionKind = CRDGroupVersion.WithKind(Tracker_Kind)
)

func init() {
	SchemeBuilder.Register(&Tracker{}, &TrackerList{})
}
