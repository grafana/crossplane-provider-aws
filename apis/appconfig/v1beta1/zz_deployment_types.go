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

type DeploymentInitParameters struct {

	// Description of the deployment. Can be at most 1024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Environment ID. Must be between 4 and 7 characters in length.
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DeploymentObservation struct {

	// Application ID. Must be between 4 and 7 characters in length.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// ARN of the AppConfig Deployment.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Configuration profile ID. Must be between 4 and 7 characters in length.
	ConfigurationProfileID *string `json:"configurationProfileId,omitempty" tf:"configuration_profile_id,omitempty"`

	// Configuration version to deploy. Can be at most 1024 characters.
	ConfigurationVersion *string `json:"configurationVersion,omitempty" tf:"configuration_version,omitempty"`

	// Deployment number.
	DeploymentNumber *int64 `json:"deploymentNumber,omitempty" tf:"deployment_number,omitempty"`

	// Deployment strategy ID or name of a predefined deployment strategy. See Predefined Deployment Strategies for more details.
	DeploymentStrategyID *string `json:"deploymentStrategyId,omitempty" tf:"deployment_strategy_id,omitempty"`

	// Description of the deployment. Can be at most 1024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Environment ID. Must be between 4 and 7 characters in length.
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	// AppConfig application ID, environment ID, and deployment number separated by a slash (/).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// State of the deployment.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DeploymentParameters struct {

	// Application ID. Must be between 4 and 7 characters in length.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appconfig/v1beta1.Application
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Reference to a Application in appconfig to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDRef *v1.Reference `json:"applicationIdRef,omitempty" tf:"-"`

	// Selector for a Application in appconfig to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// Configuration profile ID. Must be between 4 and 7 characters in length.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appconfig/v1beta1.ConfigurationProfile
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("configuration_profile_id",true)
	// +kubebuilder:validation:Optional
	ConfigurationProfileID *string `json:"configurationProfileId,omitempty" tf:"configuration_profile_id,omitempty"`

	// Reference to a ConfigurationProfile in appconfig to populate configurationProfileId.
	// +kubebuilder:validation:Optional
	ConfigurationProfileIDRef *v1.Reference `json:"configurationProfileIdRef,omitempty" tf:"-"`

	// Selector for a ConfigurationProfile in appconfig to populate configurationProfileId.
	// +kubebuilder:validation:Optional
	ConfigurationProfileIDSelector *v1.Selector `json:"configurationProfileIdSelector,omitempty" tf:"-"`

	// Configuration version to deploy. Can be at most 1024 characters.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appconfig/v1beta1.HostedConfigurationVersion
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("version_number",true)
	// +kubebuilder:validation:Optional
	ConfigurationVersion *string `json:"configurationVersion,omitempty" tf:"configuration_version,omitempty"`

	// Reference to a HostedConfigurationVersion in appconfig to populate configurationVersion.
	// +kubebuilder:validation:Optional
	ConfigurationVersionRef *v1.Reference `json:"configurationVersionRef,omitempty" tf:"-"`

	// Selector for a HostedConfigurationVersion in appconfig to populate configurationVersion.
	// +kubebuilder:validation:Optional
	ConfigurationVersionSelector *v1.Selector `json:"configurationVersionSelector,omitempty" tf:"-"`

	// Deployment strategy ID or name of a predefined deployment strategy. See Predefined Deployment Strategies for more details.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appconfig/v1beta1.DeploymentStrategy
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DeploymentStrategyID *string `json:"deploymentStrategyId,omitempty" tf:"deployment_strategy_id,omitempty"`

	// Reference to a DeploymentStrategy in appconfig to populate deploymentStrategyId.
	// +kubebuilder:validation:Optional
	DeploymentStrategyIDRef *v1.Reference `json:"deploymentStrategyIdRef,omitempty" tf:"-"`

	// Selector for a DeploymentStrategy in appconfig to populate deploymentStrategyId.
	// +kubebuilder:validation:Optional
	DeploymentStrategyIDSelector *v1.Selector `json:"deploymentStrategyIdSelector,omitempty" tf:"-"`

	// Description of the deployment. Can be at most 1024 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Environment ID. Must be between 4 and 7 characters in length.
	// +kubebuilder:validation:Optional
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DeploymentSpec defines the desired state of Deployment
type DeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeploymentParameters `json:"forProvider"`
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
	InitProvider DeploymentInitParameters `json:"initProvider,omitempty"`
}

// DeploymentStatus defines the observed state of Deployment.
type DeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Deployment is the Schema for the Deployments API. Provides an AppConfig Deployment resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.environmentId) || (has(self.initProvider) && has(self.initProvider.environmentId))",message="spec.forProvider.environmentId is a required parameter"
	Spec   DeploymentSpec   `json:"spec"`
	Status DeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentList contains a list of Deployments
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Deployment `json:"items"`
}

// Repository type metadata.
var (
	Deployment_Kind             = "Deployment"
	Deployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Deployment_Kind}.String()
	Deployment_KindAPIVersion   = Deployment_Kind + "." + CRDGroupVersion.String()
	Deployment_GroupVersionKind = CRDGroupVersion.WithKind(Deployment_Kind)
)

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}
