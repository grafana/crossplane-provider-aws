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

type DestinationPortRangeInitParameters struct {

	// Starting port of the range
	FromPort *int64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Ending port of the range
	ToPort *int64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type DestinationPortRangeObservation struct {

	// Starting port of the range
	FromPort *int64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Ending port of the range
	ToPort *int64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type DestinationPortRangeParameters struct {

	// Starting port of the range
	// +kubebuilder:validation:Optional
	FromPort *int64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Ending port of the range
	// +kubebuilder:validation:Optional
	ToPort *int64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type SourcePortRangeInitParameters struct {

	// Starting port of the range
	FromPort *int64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Ending port of the range
	ToPort *int64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type SourcePortRangeObservation struct {

	// Starting port of the range
	FromPort *int64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Ending port of the range
	ToPort *int64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type SourcePortRangeParameters struct {

	// Starting port of the range
	// +kubebuilder:validation:Optional
	FromPort *int64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// Ending port of the range
	// +kubebuilder:validation:Optional
	ToPort *int64 `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type TrafficMirrorFilterRuleInitParameters struct {

	// Description of the traffic mirror filter rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty" tf:"destination_cidr_block,omitempty"`

	// Destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	DestinationPortRange []DestinationPortRangeInitParameters `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`

	// Protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see Protocol Numbers on the Internet Assigned Numbers Authority (IANA) website.
	Protocol *int64 `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Action to take (accept | reject) on the filtered traffic. Valid values are accept and reject
	RuleAction *string `json:"ruleAction,omitempty" tf:"rule_action,omitempty"`

	// Number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber *int64 `json:"ruleNumber,omitempty" tf:"rule_number,omitempty"`

	// Source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock *string `json:"sourceCidrBlock,omitempty" tf:"source_cidr_block,omitempty"`

	// Source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	SourcePortRange []SourcePortRangeInitParameters `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`

	// Direction of traffic to be captured. Valid values are ingress and egress
	TrafficDirection *string `json:"trafficDirection,omitempty" tf:"traffic_direction,omitempty"`
}

type TrafficMirrorFilterRuleObservation struct {

	// ARN of the traffic mirror filter rule.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the traffic mirror filter rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty" tf:"destination_cidr_block,omitempty"`

	// Destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	DestinationPortRange []DestinationPortRangeObservation `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`

	// Name of the traffic mirror filter rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see Protocol Numbers on the Internet Assigned Numbers Authority (IANA) website.
	Protocol *int64 `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Action to take (accept | reject) on the filtered traffic. Valid values are accept and reject
	RuleAction *string `json:"ruleAction,omitempty" tf:"rule_action,omitempty"`

	// Number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber *int64 `json:"ruleNumber,omitempty" tf:"rule_number,omitempty"`

	// Source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock *string `json:"sourceCidrBlock,omitempty" tf:"source_cidr_block,omitempty"`

	// Source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	SourcePortRange []SourcePortRangeObservation `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`

	// Direction of traffic to be captured. Valid values are ingress and egress
	TrafficDirection *string `json:"trafficDirection,omitempty" tf:"traffic_direction,omitempty"`

	// ID of the traffic mirror filter to which this rule should be added
	TrafficMirrorFilterID *string `json:"trafficMirrorFilterId,omitempty" tf:"traffic_mirror_filter_id,omitempty"`
}

type TrafficMirrorFilterRuleParameters struct {

	// Description of the traffic mirror filter rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Destination CIDR block to assign to the Traffic Mirror rule.
	// +kubebuilder:validation:Optional
	DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty" tf:"destination_cidr_block,omitempty"`

	// Destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	// +kubebuilder:validation:Optional
	DestinationPortRange []DestinationPortRangeParameters `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`

	// Protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see Protocol Numbers on the Internet Assigned Numbers Authority (IANA) website.
	// +kubebuilder:validation:Optional
	Protocol *int64 `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Action to take (accept | reject) on the filtered traffic. Valid values are accept and reject
	// +kubebuilder:validation:Optional
	RuleAction *string `json:"ruleAction,omitempty" tf:"rule_action,omitempty"`

	// Number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	// +kubebuilder:validation:Optional
	RuleNumber *int64 `json:"ruleNumber,omitempty" tf:"rule_number,omitempty"`

	// Source CIDR block to assign to the Traffic Mirror rule.
	// +kubebuilder:validation:Optional
	SourceCidrBlock *string `json:"sourceCidrBlock,omitempty" tf:"source_cidr_block,omitempty"`

	// Source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	// +kubebuilder:validation:Optional
	SourcePortRange []SourcePortRangeParameters `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`

	// Direction of traffic to be captured. Valid values are ingress and egress
	// +kubebuilder:validation:Optional
	TrafficDirection *string `json:"trafficDirection,omitempty" tf:"traffic_direction,omitempty"`

	// ID of the traffic mirror filter to which this rule should be added
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TrafficMirrorFilter
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TrafficMirrorFilterID *string `json:"trafficMirrorFilterId,omitempty" tf:"traffic_mirror_filter_id,omitempty"`

	// Reference to a TrafficMirrorFilter in ec2 to populate trafficMirrorFilterId.
	// +kubebuilder:validation:Optional
	TrafficMirrorFilterIDRef *v1.Reference `json:"trafficMirrorFilterIdRef,omitempty" tf:"-"`

	// Selector for a TrafficMirrorFilter in ec2 to populate trafficMirrorFilterId.
	// +kubebuilder:validation:Optional
	TrafficMirrorFilterIDSelector *v1.Selector `json:"trafficMirrorFilterIdSelector,omitempty" tf:"-"`
}

// TrafficMirrorFilterRuleSpec defines the desired state of TrafficMirrorFilterRule
type TrafficMirrorFilterRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficMirrorFilterRuleParameters `json:"forProvider"`
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
	InitProvider TrafficMirrorFilterRuleInitParameters `json:"initProvider,omitempty"`
}

// TrafficMirrorFilterRuleStatus defines the observed state of TrafficMirrorFilterRule.
type TrafficMirrorFilterRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficMirrorFilterRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficMirrorFilterRule is the Schema for the TrafficMirrorFilterRules API. Provides an Traffic mirror filter rule
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TrafficMirrorFilterRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destinationCidrBlock) || (has(self.initProvider) && has(self.initProvider.destinationCidrBlock))",message="spec.forProvider.destinationCidrBlock is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ruleAction) || (has(self.initProvider) && has(self.initProvider.ruleAction))",message="spec.forProvider.ruleAction is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ruleNumber) || (has(self.initProvider) && has(self.initProvider.ruleNumber))",message="spec.forProvider.ruleNumber is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourceCidrBlock) || (has(self.initProvider) && has(self.initProvider.sourceCidrBlock))",message="spec.forProvider.sourceCidrBlock is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.trafficDirection) || (has(self.initProvider) && has(self.initProvider.trafficDirection))",message="spec.forProvider.trafficDirection is a required parameter"
	Spec   TrafficMirrorFilterRuleSpec   `json:"spec"`
	Status TrafficMirrorFilterRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficMirrorFilterRuleList contains a list of TrafficMirrorFilterRules
type TrafficMirrorFilterRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficMirrorFilterRule `json:"items"`
}

// Repository type metadata.
var (
	TrafficMirrorFilterRule_Kind             = "TrafficMirrorFilterRule"
	TrafficMirrorFilterRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficMirrorFilterRule_Kind}.String()
	TrafficMirrorFilterRule_KindAPIVersion   = TrafficMirrorFilterRule_Kind + "." + CRDGroupVersion.String()
	TrafficMirrorFilterRule_GroupVersionKind = CRDGroupVersion.WithKind(TrafficMirrorFilterRule_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficMirrorFilterRule{}, &TrafficMirrorFilterRuleList{})
}
