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

type InstanceStateInitParameters struct {

	// Whether to request a forced stop when state is stopped. Otherwise (i.e., state is running), ignored. When an instance is forced to stop, it does not flush file system caches or file system metadata, and you must subsequently perform file system check and repair. Not recommended for Windows instances. Defaults to false.
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// ID of the instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in ec2 to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in ec2 to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// - State of the instance. Valid values are stopped, running.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type InstanceStateObservation struct {

	// Whether to request a forced stop when state is stopped. Otherwise (i.e., state is running), ignored. When an instance is forced to stop, it does not flush file system caches or file system metadata, and you must subsequently perform file system check and repair. Not recommended for Windows instances. Defaults to false.
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// ID of the instance (matches instance_id).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// - State of the instance. Valid values are stopped, running.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type InstanceStateParameters struct {

	// Whether to request a forced stop when state is stopped. Otherwise (i.e., state is running), ignored. When an instance is forced to stop, it does not flush file system caches or file system metadata, and you must subsequently perform file system check and repair. Not recommended for Windows instances. Defaults to false.
	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// ID of the instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in ec2 to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in ec2 to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// - State of the instance. Valid values are stopped, running.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

// InstanceStateSpec defines the desired state of InstanceState
type InstanceStateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceStateParameters `json:"forProvider"`
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
	InitProvider InstanceStateInitParameters `json:"initProvider,omitempty"`
}

// InstanceStateStatus defines the observed state of InstanceState.
type InstanceStateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceStateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// InstanceState is the Schema for the InstanceStates API. Provides an EC2 instance state resource. This allows managing an instance power state.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InstanceState struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.state) || (has(self.initProvider) && has(self.initProvider.state))",message="spec.forProvider.state is a required parameter"
	Spec   InstanceStateSpec   `json:"spec"`
	Status InstanceStateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceStateList contains a list of InstanceStates
type InstanceStateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceState `json:"items"`
}

// Repository type metadata.
var (
	InstanceState_Kind             = "InstanceState"
	InstanceState_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceState_Kind}.String()
	InstanceState_KindAPIVersion   = InstanceState_Kind + "." + CRDGroupVersion.String()
	InstanceState_GroupVersionKind = CRDGroupVersion.WithKind(InstanceState_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceState{}, &InstanceStateList{})
}
