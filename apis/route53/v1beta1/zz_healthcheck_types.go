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

type HealthCheckInitParameters struct {

	// The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
	ChildHealthThreshold *int64 `json:"childHealthThreshold,omitempty" tf:"child_health_threshold,omitempty"`

	// For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
	ChildHealthchecks []*string `json:"childHealthchecks,omitempty" tf:"child_healthchecks,omitempty"`

	// The CloudWatchRegion that the CloudWatch alarm was created in.
	CloudwatchAlarmRegion *string `json:"cloudwatchAlarmRegion,omitempty" tf:"cloudwatch_alarm_region,omitempty"`

	// A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// A boolean value that indicates whether Route53 should send the fqdn to the endpoint when performing the health check. This defaults to AWS' defaults: when the type is "HTTPS" enable_sni defaults to true, when type is anything else enable_sni defaults to false.
	EnableSni *bool `json:"enableSni,omitempty" tf:"enable_sni,omitempty"`

	// The number of consecutive health checks that an endpoint must pass or fail.
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`

	// The fully qualified domain name of the endpoint to be checked. If a value is set for ip_address, the value set for fqdn will be passed in the Host header.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The IP address of the endpoint to be checked.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are Healthy , Unhealthy and LastKnownStatus.
	InsufficientDataHealthStatus *string `json:"insufficientDataHealthStatus,omitempty" tf:"insufficient_data_health_status,omitempty"`

	// A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
	InvertHealthcheck *bool `json:"invertHealthcheck,omitempty" tf:"invert_healthcheck,omitempty"`

	// A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
	MeasureLatency *bool `json:"measureLatency,omitempty" tf:"measure_latency,omitempty"`

	// The port of the endpoint to be checked.
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// This is a reference name used in Caller Reference
	// (helpful for identifying single health_check set amongst others)
	ReferenceName *string `json:"referenceName,omitempty" tf:"reference_name,omitempty"`

	// A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
	RequestInterval *int64 `json:"requestInterval,omitempty" tf:"request_interval,omitempty"`

	// The path that you want Amazon Route 53 to request when performing health checks.
	ResourcePath *string `json:"resourcePath,omitempty" tf:"resource_path,omitempty"`

	// The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is RECOVERY_CONTROL
	RoutingControlArn *string `json:"routingControlArn,omitempty" tf:"routing_control_arn,omitempty"`

	// String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with HTTP_STR_MATCH and HTTPS_STR_MATCH.
	SearchString *string `json:"searchString,omitempty" tf:"search_string,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The protocol to use when performing health checks. Valid values are HTTP, HTTPS, HTTP_STR_MATCH, HTTPS_STR_MATCH, TCP, CALCULATED, CLOUDWATCH_METRIC and RECOVERY_CONTROL.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthCheckObservation struct {

	// The Amazon Resource Name (ARN) of the Health Check.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
	ChildHealthThreshold *int64 `json:"childHealthThreshold,omitempty" tf:"child_health_threshold,omitempty"`

	// For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
	ChildHealthchecks []*string `json:"childHealthchecks,omitempty" tf:"child_healthchecks,omitempty"`

	// The name of the CloudWatch alarm.
	CloudwatchAlarmName *string `json:"cloudwatchAlarmName,omitempty" tf:"cloudwatch_alarm_name,omitempty"`

	// The CloudWatchRegion that the CloudWatch alarm was created in.
	CloudwatchAlarmRegion *string `json:"cloudwatchAlarmRegion,omitempty" tf:"cloudwatch_alarm_region,omitempty"`

	// A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// A boolean value that indicates whether Route53 should send the fqdn to the endpoint when performing the health check. This defaults to AWS' defaults: when the type is "HTTPS" enable_sni defaults to true, when type is anything else enable_sni defaults to false.
	EnableSni *bool `json:"enableSni,omitempty" tf:"enable_sni,omitempty"`

	// The number of consecutive health checks that an endpoint must pass or fail.
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`

	// The fully qualified domain name of the endpoint to be checked. If a value is set for ip_address, the value set for fqdn will be passed in the Host header.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The id of the health check
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address of the endpoint to be checked.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are Healthy , Unhealthy and LastKnownStatus.
	InsufficientDataHealthStatus *string `json:"insufficientDataHealthStatus,omitempty" tf:"insufficient_data_health_status,omitempty"`

	// A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
	InvertHealthcheck *bool `json:"invertHealthcheck,omitempty" tf:"invert_healthcheck,omitempty"`

	// A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
	MeasureLatency *bool `json:"measureLatency,omitempty" tf:"measure_latency,omitempty"`

	// The port of the endpoint to be checked.
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// This is a reference name used in Caller Reference
	// (helpful for identifying single health_check set amongst others)
	ReferenceName *string `json:"referenceName,omitempty" tf:"reference_name,omitempty"`

	// A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
	RequestInterval *int64 `json:"requestInterval,omitempty" tf:"request_interval,omitempty"`

	// The path that you want Amazon Route 53 to request when performing health checks.
	ResourcePath *string `json:"resourcePath,omitempty" tf:"resource_path,omitempty"`

	// The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is RECOVERY_CONTROL
	RoutingControlArn *string `json:"routingControlArn,omitempty" tf:"routing_control_arn,omitempty"`

	// String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with HTTP_STR_MATCH and HTTPS_STR_MATCH.
	SearchString *string `json:"searchString,omitempty" tf:"search_string,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The protocol to use when performing health checks. Valid values are HTTP, HTTPS, HTTP_STR_MATCH, HTTPS_STR_MATCH, TCP, CALCULATED, CLOUDWATCH_METRIC and RECOVERY_CONTROL.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthCheckParameters struct {

	// The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
	// +kubebuilder:validation:Optional
	ChildHealthThreshold *int64 `json:"childHealthThreshold,omitempty" tf:"child_health_threshold,omitempty"`

	// For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
	// +kubebuilder:validation:Optional
	ChildHealthchecks []*string `json:"childHealthchecks,omitempty" tf:"child_healthchecks,omitempty"`

	// The name of the CloudWatch alarm.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatch/v1beta1.MetricAlarm
	// +kubebuilder:validation:Optional
	CloudwatchAlarmName *string `json:"cloudwatchAlarmName,omitempty" tf:"cloudwatch_alarm_name,omitempty"`

	// Reference to a MetricAlarm in cloudwatch to populate cloudwatchAlarmName.
	// +kubebuilder:validation:Optional
	CloudwatchAlarmNameRef *v1.Reference `json:"cloudwatchAlarmNameRef,omitempty" tf:"-"`

	// Selector for a MetricAlarm in cloudwatch to populate cloudwatchAlarmName.
	// +kubebuilder:validation:Optional
	CloudwatchAlarmNameSelector *v1.Selector `json:"cloudwatchAlarmNameSelector,omitempty" tf:"-"`

	// The CloudWatchRegion that the CloudWatch alarm was created in.
	// +kubebuilder:validation:Optional
	CloudwatchAlarmRegion *string `json:"cloudwatchAlarmRegion,omitempty" tf:"cloudwatch_alarm_region,omitempty"`

	// A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// A boolean value that indicates whether Route53 should send the fqdn to the endpoint when performing the health check. This defaults to AWS' defaults: when the type is "HTTPS" enable_sni defaults to true, when type is anything else enable_sni defaults to false.
	// +kubebuilder:validation:Optional
	EnableSni *bool `json:"enableSni,omitempty" tf:"enable_sni,omitempty"`

	// The number of consecutive health checks that an endpoint must pass or fail.
	// +kubebuilder:validation:Optional
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`

	// The fully qualified domain name of the endpoint to be checked. If a value is set for ip_address, the value set for fqdn will be passed in the Host header.
	// +kubebuilder:validation:Optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The IP address of the endpoint to be checked.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are Healthy , Unhealthy and LastKnownStatus.
	// +kubebuilder:validation:Optional
	InsufficientDataHealthStatus *string `json:"insufficientDataHealthStatus,omitempty" tf:"insufficient_data_health_status,omitempty"`

	// A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
	// +kubebuilder:validation:Optional
	InvertHealthcheck *bool `json:"invertHealthcheck,omitempty" tf:"invert_healthcheck,omitempty"`

	// A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
	// +kubebuilder:validation:Optional
	MeasureLatency *bool `json:"measureLatency,omitempty" tf:"measure_latency,omitempty"`

	// The port of the endpoint to be checked.
	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// This is a reference name used in Caller Reference
	// (helpful for identifying single health_check set amongst others)
	// +kubebuilder:validation:Optional
	ReferenceName *string `json:"referenceName,omitempty" tf:"reference_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
	// +kubebuilder:validation:Optional
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`

	// The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
	// +kubebuilder:validation:Optional
	RequestInterval *int64 `json:"requestInterval,omitempty" tf:"request_interval,omitempty"`

	// The path that you want Amazon Route 53 to request when performing health checks.
	// +kubebuilder:validation:Optional
	ResourcePath *string `json:"resourcePath,omitempty" tf:"resource_path,omitempty"`

	// The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is RECOVERY_CONTROL
	// +kubebuilder:validation:Optional
	RoutingControlArn *string `json:"routingControlArn,omitempty" tf:"routing_control_arn,omitempty"`

	// String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with HTTP_STR_MATCH and HTTPS_STR_MATCH.
	// +kubebuilder:validation:Optional
	SearchString *string `json:"searchString,omitempty" tf:"search_string,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The protocol to use when performing health checks. Valid values are HTTP, HTTPS, HTTP_STR_MATCH, HTTPS_STR_MATCH, TCP, CALCULATED, CLOUDWATCH_METRIC and RECOVERY_CONTROL.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// HealthCheckSpec defines the desired state of HealthCheck
type HealthCheckSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HealthCheckParameters `json:"forProvider"`
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
	InitProvider HealthCheckInitParameters `json:"initProvider,omitempty"`
}

// HealthCheckStatus defines the observed state of HealthCheck.
type HealthCheckStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HealthCheckObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HealthCheck is the Schema for the HealthChecks API. Provides a Route53 health check.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type HealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   HealthCheckSpec   `json:"spec"`
	Status HealthCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthCheckList contains a list of HealthChecks
type HealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthCheck `json:"items"`
}

// Repository type metadata.
var (
	HealthCheck_Kind             = "HealthCheck"
	HealthCheck_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HealthCheck_Kind}.String()
	HealthCheck_KindAPIVersion   = HealthCheck_Kind + "." + CRDGroupVersion.String()
	HealthCheck_GroupVersionKind = CRDGroupVersion.WithKind(HealthCheck_Kind)
)

func init() {
	SchemeBuilder.Register(&HealthCheck{}, &HealthCheckList{})
}
