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

type EndpointInitParameters struct {

	// The Amazon Kinesis data stream configuration.
	KinesisStreamConfig []KinesisStreamConfigInitParameters `json:"kinesisStreamConfig,omitempty" tf:"kinesis_stream_config,omitempty"`

	// The type of data stream where real-time log data is sent. The only valid value is Kinesis.
	StreamType *string `json:"streamType,omitempty" tf:"stream_type,omitempty"`
}

type EndpointObservation struct {

	// The Amazon Kinesis data stream configuration.
	KinesisStreamConfig []KinesisStreamConfigObservation `json:"kinesisStreamConfig,omitempty" tf:"kinesis_stream_config,omitempty"`

	// The type of data stream where real-time log data is sent. The only valid value is Kinesis.
	StreamType *string `json:"streamType,omitempty" tf:"stream_type,omitempty"`
}

type EndpointParameters struct {

	// The Amazon Kinesis data stream configuration.
	// +kubebuilder:validation:Optional
	KinesisStreamConfig []KinesisStreamConfigParameters `json:"kinesisStreamConfig" tf:"kinesis_stream_config,omitempty"`

	// The type of data stream where real-time log data is sent. The only valid value is Kinesis.
	// +kubebuilder:validation:Optional
	StreamType *string `json:"streamType" tf:"stream_type,omitempty"`
}

type KinesisStreamConfigInitParameters struct {

	// The ARN of an IAM role that CloudFront can use to send real-time log data to the Kinesis data stream.
	// See the AWS documentation for more information.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The ARN of the Kinesis data stream.
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

type KinesisStreamConfigObservation struct {

	// The ARN of an IAM role that CloudFront can use to send real-time log data to the Kinesis data stream.
	// See the AWS documentation for more information.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The ARN of the Kinesis data stream.
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`
}

type KinesisStreamConfigParameters struct {

	// The ARN of an IAM role that CloudFront can use to send real-time log data to the Kinesis data stream.
	// See the AWS documentation for more information.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The ARN of the Kinesis data stream.
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

type RealtimeLogConfigInitParameters struct {

	// The Amazon Kinesis data streams where real-time log data is sent.
	Endpoint []EndpointInitParameters `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The fields that are included in each real-time log record. See the AWS documentation for supported values.
	// +listType=set
	Fields []*string `json:"fields,omitempty" tf:"fields,omitempty"`

	// The unique name to identify this real-time log configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between 1 and 100, inclusive.
	SamplingRate *float64 `json:"samplingRate,omitempty" tf:"sampling_rate,omitempty"`
}

type RealtimeLogConfigObservation struct {

	// The ARN (Amazon Resource Name) of the CloudFront real-time log configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Amazon Kinesis data streams where real-time log data is sent.
	Endpoint []EndpointObservation `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The fields that are included in each real-time log record. See the AWS documentation for supported values.
	// +listType=set
	Fields []*string `json:"fields,omitempty" tf:"fields,omitempty"`

	// The ID of the CloudFront real-time log configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique name to identify this real-time log configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between 1 and 100, inclusive.
	SamplingRate *float64 `json:"samplingRate,omitempty" tf:"sampling_rate,omitempty"`
}

type RealtimeLogConfigParameters struct {

	// The Amazon Kinesis data streams where real-time log data is sent.
	// +kubebuilder:validation:Optional
	Endpoint []EndpointParameters `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The fields that are included in each real-time log record. See the AWS documentation for supported values.
	// +kubebuilder:validation:Optional
	// +listType=set
	Fields []*string `json:"fields,omitempty" tf:"fields,omitempty"`

	// The unique name to identify this real-time log configuration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between 1 and 100, inclusive.
	// +kubebuilder:validation:Optional
	SamplingRate *float64 `json:"samplingRate,omitempty" tf:"sampling_rate,omitempty"`
}

// RealtimeLogConfigSpec defines the desired state of RealtimeLogConfig
type RealtimeLogConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RealtimeLogConfigParameters `json:"forProvider"`
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
	InitProvider RealtimeLogConfigInitParameters `json:"initProvider,omitempty"`
}

// RealtimeLogConfigStatus defines the observed state of RealtimeLogConfig.
type RealtimeLogConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RealtimeLogConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RealtimeLogConfig is the Schema for the RealtimeLogConfigs API. Provides a CloudFront real-time log configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RealtimeLogConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endpoint) || (has(self.initProvider) && has(self.initProvider.endpoint))",message="spec.forProvider.endpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fields) || (has(self.initProvider) && has(self.initProvider.fields))",message="spec.forProvider.fields is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.samplingRate) || (has(self.initProvider) && has(self.initProvider.samplingRate))",message="spec.forProvider.samplingRate is a required parameter"
	Spec   RealtimeLogConfigSpec   `json:"spec"`
	Status RealtimeLogConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RealtimeLogConfigList contains a list of RealtimeLogConfigs
type RealtimeLogConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RealtimeLogConfig `json:"items"`
}

// Repository type metadata.
var (
	RealtimeLogConfig_Kind             = "RealtimeLogConfig"
	RealtimeLogConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RealtimeLogConfig_Kind}.String()
	RealtimeLogConfig_KindAPIVersion   = RealtimeLogConfig_Kind + "." + CRDGroupVersion.String()
	RealtimeLogConfig_GroupVersionKind = CRDGroupVersion.WithKind(RealtimeLogConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&RealtimeLogConfig{}, &RealtimeLogConfigList{})
}
