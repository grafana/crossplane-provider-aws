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

type RuleConfigInitParameters struct {

	// Logical negation of the rule.
	Inverted *bool `json:"inverted,omitempty" tf:"inverted,omitempty"`

	// Number of controls that must be set when you specify an ATLEAST type rule.
	Threshold *int64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// Rule type. Valid values are ATLEAST, AND, and OR.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RuleConfigObservation struct {

	// Logical negation of the rule.
	Inverted *bool `json:"inverted,omitempty" tf:"inverted,omitempty"`

	// Number of controls that must be set when you specify an ATLEAST type rule.
	Threshold *int64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// Rule type. Valid values are ATLEAST, AND, and OR.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RuleConfigParameters struct {

	// Logical negation of the rule.
	// +kubebuilder:validation:Optional
	Inverted *bool `json:"inverted" tf:"inverted,omitempty"`

	// Number of controls that must be set when you specify an ATLEAST type rule.
	// +kubebuilder:validation:Optional
	Threshold *int64 `json:"threshold" tf:"threshold,omitempty"`

	// Rule type. Valid values are ATLEAST, AND, and OR.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type SafetyRuleInitParameters struct {

	// Gating controls for the new gating rule. That is, routing controls that are evaluated by the rule configuration that you specify.
	GatingControls []*string `json:"gatingControls,omitempty" tf:"gating_controls,omitempty"`

	// Name describing the safety rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration block for safety rule criteria. See below.
	RuleConfig []RuleConfigInitParameters `json:"ruleConfig,omitempty" tf:"rule_config,omitempty"`

	// Routing controls that can only be set or unset if the specified rule_config evaluates to true for the specified gating_controls.
	TargetControls []*string `json:"targetControls,omitempty" tf:"target_controls,omitempty"`

	// Evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail.
	WaitPeriodMs *int64 `json:"waitPeriodMs,omitempty" tf:"wait_period_ms,omitempty"`
}

type SafetyRuleObservation struct {

	// ARN of the safety rule.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Routing controls that are part of transactions that are evaluated to determine if a request to change a routing control state is allowed.
	AssertedControls []*string `json:"assertedControls,omitempty" tf:"asserted_controls,omitempty"`

	// ARN of the control panel in which this safety rule will reside.
	ControlPanelArn *string `json:"controlPanelArn,omitempty" tf:"control_panel_arn,omitempty"`

	// Gating controls for the new gating rule. That is, routing controls that are evaluated by the rule configuration that you specify.
	GatingControls []*string `json:"gatingControls,omitempty" tf:"gating_controls,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name describing the safety rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration block for safety rule criteria. See below.
	RuleConfig []RuleConfigObservation `json:"ruleConfig,omitempty" tf:"rule_config,omitempty"`

	// Status of the safety rule. PENDING when it is being created/updated, PENDING_DELETION when it is being deleted, and DEPLOYED otherwise.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Routing controls that can only be set or unset if the specified rule_config evaluates to true for the specified gating_controls.
	TargetControls []*string `json:"targetControls,omitempty" tf:"target_controls,omitempty"`

	// Evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail.
	WaitPeriodMs *int64 `json:"waitPeriodMs,omitempty" tf:"wait_period_ms,omitempty"`
}

type SafetyRuleParameters struct {

	// Routing controls that are part of transactions that are evaluated to determine if a request to change a routing control state is allowed.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53recoverycontrolconfig/v1beta1.RoutingControl
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	AssertedControls []*string `json:"assertedControls,omitempty" tf:"asserted_controls,omitempty"`

	// References to RoutingControl in route53recoverycontrolconfig to populate assertedControls.
	// +kubebuilder:validation:Optional
	AssertedControlsRefs []v1.Reference `json:"assertedControlsRefs,omitempty" tf:"-"`

	// Selector for a list of RoutingControl in route53recoverycontrolconfig to populate assertedControls.
	// +kubebuilder:validation:Optional
	AssertedControlsSelector *v1.Selector `json:"assertedControlsSelector,omitempty" tf:"-"`

	// ARN of the control panel in which this safety rule will reside.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53recoverycontrolconfig/v1beta1.ControlPanel
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	ControlPanelArn *string `json:"controlPanelArn,omitempty" tf:"control_panel_arn,omitempty"`

	// Reference to a ControlPanel in route53recoverycontrolconfig to populate controlPanelArn.
	// +kubebuilder:validation:Optional
	ControlPanelArnRef *v1.Reference `json:"controlPanelArnRef,omitempty" tf:"-"`

	// Selector for a ControlPanel in route53recoverycontrolconfig to populate controlPanelArn.
	// +kubebuilder:validation:Optional
	ControlPanelArnSelector *v1.Selector `json:"controlPanelArnSelector,omitempty" tf:"-"`

	// Gating controls for the new gating rule. That is, routing controls that are evaluated by the rule configuration that you specify.
	// +kubebuilder:validation:Optional
	GatingControls []*string `json:"gatingControls,omitempty" tf:"gating_controls,omitempty"`

	// Name describing the safety rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block for safety rule criteria. See below.
	// +kubebuilder:validation:Optional
	RuleConfig []RuleConfigParameters `json:"ruleConfig,omitempty" tf:"rule_config,omitempty"`

	// Routing controls that can only be set or unset if the specified rule_config evaluates to true for the specified gating_controls.
	// +kubebuilder:validation:Optional
	TargetControls []*string `json:"targetControls,omitempty" tf:"target_controls,omitempty"`

	// Evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail.
	// +kubebuilder:validation:Optional
	WaitPeriodMs *int64 `json:"waitPeriodMs,omitempty" tf:"wait_period_ms,omitempty"`
}

// SafetyRuleSpec defines the desired state of SafetyRule
type SafetyRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SafetyRuleParameters `json:"forProvider"`
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
	InitProvider SafetyRuleInitParameters `json:"initProvider,omitempty"`
}

// SafetyRuleStatus defines the observed state of SafetyRule.
type SafetyRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SafetyRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SafetyRule is the Schema for the SafetyRules API. Provides an AWS Route 53 Recovery Control Config Safety Rule
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SafetyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ruleConfig) || (has(self.initProvider) && has(self.initProvider.ruleConfig))",message="spec.forProvider.ruleConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.waitPeriodMs) || (has(self.initProvider) && has(self.initProvider.waitPeriodMs))",message="spec.forProvider.waitPeriodMs is a required parameter"
	Spec   SafetyRuleSpec   `json:"spec"`
	Status SafetyRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SafetyRuleList contains a list of SafetyRules
type SafetyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SafetyRule `json:"items"`
}

// Repository type metadata.
var (
	SafetyRule_Kind             = "SafetyRule"
	SafetyRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SafetyRule_Kind}.String()
	SafetyRule_KindAPIVersion   = SafetyRule_Kind + "." + CRDGroupVersion.String()
	SafetyRule_GroupVersionKind = CRDGroupVersion.WithKind(SafetyRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SafetyRule{}, &SafetyRuleList{})
}
