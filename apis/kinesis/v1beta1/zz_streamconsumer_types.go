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

type StreamConsumerInitParameters struct {

	// Name of the stream consumer.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// –  Amazon Resource Name (ARN) of the data stream the consumer is registered with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kinesis/v1beta1.Stream
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`

	// Reference to a Stream in kinesis to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnRef *v1.Reference `json:"streamArnRef,omitempty" tf:"-"`

	// Selector for a Stream in kinesis to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnSelector *v1.Selector `json:"streamArnSelector,omitempty" tf:"-"`
}

type StreamConsumerObservation struct {

	// Amazon Resource Name (ARN) of the stream consumer.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Approximate timestamp in RFC3339 format of when the stream consumer was created.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// Amazon Resource Name (ARN) of the stream consumer.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the stream consumer.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// –  Amazon Resource Name (ARN) of the data stream the consumer is registered with.
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`
}

type StreamConsumerParameters struct {

	// Name of the stream consumer.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// –  Amazon Resource Name (ARN) of the data stream the consumer is registered with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kinesis/v1beta1.Stream
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`

	// Reference to a Stream in kinesis to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnRef *v1.Reference `json:"streamArnRef,omitempty" tf:"-"`

	// Selector for a Stream in kinesis to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnSelector *v1.Selector `json:"streamArnSelector,omitempty" tf:"-"`
}

// StreamConsumerSpec defines the desired state of StreamConsumer
type StreamConsumerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StreamConsumerParameters `json:"forProvider"`
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
	InitProvider StreamConsumerInitParameters `json:"initProvider,omitempty"`
}

// StreamConsumerStatus defines the observed state of StreamConsumer.
type StreamConsumerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StreamConsumerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// StreamConsumer is the Schema for the StreamConsumers API. Manages a Kinesis Stream Consumer.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StreamConsumer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   StreamConsumerSpec   `json:"spec"`
	Status StreamConsumerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamConsumerList contains a list of StreamConsumers
type StreamConsumerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StreamConsumer `json:"items"`
}

// Repository type metadata.
var (
	StreamConsumer_Kind             = "StreamConsumer"
	StreamConsumer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StreamConsumer_Kind}.String()
	StreamConsumer_KindAPIVersion   = StreamConsumer_Kind + "." + CRDGroupVersion.String()
	StreamConsumer_GroupVersionKind = CRDGroupVersion.WithKind(StreamConsumer_Kind)
)

func init() {
	SchemeBuilder.Register(&StreamConsumer{}, &StreamConsumerList{})
}
