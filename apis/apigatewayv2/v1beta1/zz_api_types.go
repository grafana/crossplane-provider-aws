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

type APIInitParameters struct {

	// An API key selection expression.
	// Valid values: $context.authorizer.usageIdentifierKey, $request.header.x-api-key. Defaults to $request.header.x-api-key.
	// Applicable for WebSocket APIs.
	APIKeySelectionExpression *string `json:"apiKeySelectionExpression,omitempty" tf:"api_key_selection_expression,omitempty"`

	// An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Cross-origin resource sharing (CORS) configuration. Applicable for HTTP APIs.
	CorsConfiguration []CorsConfigurationInitParameters `json:"corsConfiguration,omitempty" tf:"cors_configuration,omitempty"`

	// Part of quick create. Specifies any credentials required for the integration. Applicable for HTTP APIs.
	CredentialsArn *string `json:"credentialsArn,omitempty" tf:"credentials_arn,omitempty"`

	// Description of the API. Must be less than or equal to 1024 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether clients can invoke the API by using the default execute-api endpoint.
	// By default, clients can invoke the API with the default {api_id}.execute-api.{region}.amazonaws.com endpoint.
	// To require that clients use a custom domain name to invoke the API, disable the default endpoint.
	DisableExecuteAPIEndpoint *bool `json:"disableExecuteApiEndpoint,omitempty" tf:"disable_execute_api_endpoint,omitempty"`

	// Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to false. Applicable for HTTP APIs.
	FailOnWarnings *bool `json:"failOnWarnings,omitempty" tf:"fail_on_warnings,omitempty"`

	// Name of the API. Must be less than or equal to 128 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// API protocol. Valid values: HTTP, WEBSOCKET.
	ProtocolType *string `json:"protocolType,omitempty" tf:"protocol_type,omitempty"`

	// Part of quick create. Specifies any route key. Applicable for HTTP APIs.
	RouteKey *string `json:"routeKey,omitempty" tf:"route_key,omitempty"`

	// The route selection expression for the API.
	// Defaults to $request.method $request.path.
	RouteSelectionExpression *string `json:"routeSelectionExpression,omitempty" tf:"route_selection_expression,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
	// For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
	// The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Applicable for HTTP APIs.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Version identifier for the API. Must be between 1 and 64 characters in length.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type APIObservation struct {

	// URI of the API, of the form https://{api-id}.execute-api.{region}.amazonaws.com for HTTP APIs and wss://{api-id}.execute-api.{region}.amazonaws.com for WebSocket APIs.
	APIEndpoint *string `json:"apiEndpoint,omitempty" tf:"api_endpoint,omitempty"`

	// An API key selection expression.
	// Valid values: $context.authorizer.usageIdentifierKey, $request.header.x-api-key. Defaults to $request.header.x-api-key.
	// Applicable for WebSocket APIs.
	APIKeySelectionExpression *string `json:"apiKeySelectionExpression,omitempty" tf:"api_key_selection_expression,omitempty"`

	// ARN of the API.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Cross-origin resource sharing (CORS) configuration. Applicable for HTTP APIs.
	CorsConfiguration []CorsConfigurationObservation `json:"corsConfiguration,omitempty" tf:"cors_configuration,omitempty"`

	// Part of quick create. Specifies any credentials required for the integration. Applicable for HTTP APIs.
	CredentialsArn *string `json:"credentialsArn,omitempty" tf:"credentials_arn,omitempty"`

	// Description of the API. Must be less than or equal to 1024 characters in length.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether clients can invoke the API by using the default execute-api endpoint.
	// By default, clients can invoke the API with the default {api_id}.execute-api.{region}.amazonaws.com endpoint.
	// To require that clients use a custom domain name to invoke the API, disable the default endpoint.
	DisableExecuteAPIEndpoint *bool `json:"disableExecuteApiEndpoint,omitempty" tf:"disable_execute_api_endpoint,omitempty"`

	// ARN prefix to be used in an aws_lambda_permission's source_arn attribute
	// or in an aws_iam_policy to authorize access to the @connections API.
	// See the Amazon API Gateway Developer Guide for details.
	ExecutionArn *string `json:"executionArn,omitempty" tf:"execution_arn,omitempty"`

	// Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to false. Applicable for HTTP APIs.
	FailOnWarnings *bool `json:"failOnWarnings,omitempty" tf:"fail_on_warnings,omitempty"`

	// API identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the API. Must be less than or equal to 128 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// API protocol. Valid values: HTTP, WEBSOCKET.
	ProtocolType *string `json:"protocolType,omitempty" tf:"protocol_type,omitempty"`

	// Part of quick create. Specifies any route key. Applicable for HTTP APIs.
	RouteKey *string `json:"routeKey,omitempty" tf:"route_key,omitempty"`

	// The route selection expression for the API.
	// Defaults to $request.method $request.path.
	RouteSelectionExpression *string `json:"routeSelectionExpression,omitempty" tf:"route_selection_expression,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
	// For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
	// The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Applicable for HTTP APIs.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Version identifier for the API. Must be between 1 and 64 characters in length.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type APIParameters struct {

	// An API key selection expression.
	// Valid values: $context.authorizer.usageIdentifierKey, $request.header.x-api-key. Defaults to $request.header.x-api-key.
	// Applicable for WebSocket APIs.
	// +kubebuilder:validation:Optional
	APIKeySelectionExpression *string `json:"apiKeySelectionExpression,omitempty" tf:"api_key_selection_expression,omitempty"`

	// An OpenAPI specification that defines the set of routes and integrations to create as part of the HTTP APIs. Supported only for HTTP APIs.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Cross-origin resource sharing (CORS) configuration. Applicable for HTTP APIs.
	// +kubebuilder:validation:Optional
	CorsConfiguration []CorsConfigurationParameters `json:"corsConfiguration,omitempty" tf:"cors_configuration,omitempty"`

	// Part of quick create. Specifies any credentials required for the integration. Applicable for HTTP APIs.
	// +kubebuilder:validation:Optional
	CredentialsArn *string `json:"credentialsArn,omitempty" tf:"credentials_arn,omitempty"`

	// Description of the API. Must be less than or equal to 1024 characters in length.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether clients can invoke the API by using the default execute-api endpoint.
	// By default, clients can invoke the API with the default {api_id}.execute-api.{region}.amazonaws.com endpoint.
	// To require that clients use a custom domain name to invoke the API, disable the default endpoint.
	// +kubebuilder:validation:Optional
	DisableExecuteAPIEndpoint *bool `json:"disableExecuteApiEndpoint,omitempty" tf:"disable_execute_api_endpoint,omitempty"`

	// Whether warnings should return an error while API Gateway is creating or updating the resource using an OpenAPI specification. Defaults to false. Applicable for HTTP APIs.
	// +kubebuilder:validation:Optional
	FailOnWarnings *bool `json:"failOnWarnings,omitempty" tf:"fail_on_warnings,omitempty"`

	// Name of the API. Must be less than or equal to 128 characters in length.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// API protocol. Valid values: HTTP, WEBSOCKET.
	// +kubebuilder:validation:Optional
	ProtocolType *string `json:"protocolType,omitempty" tf:"protocol_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Part of quick create. Specifies any route key. Applicable for HTTP APIs.
	// +kubebuilder:validation:Optional
	RouteKey *string `json:"routeKey,omitempty" tf:"route_key,omitempty"`

	// The route selection expression for the API.
	// Defaults to $request.method $request.path.
	// +kubebuilder:validation:Optional
	RouteSelectionExpression *string `json:"routeSelectionExpression,omitempty" tf:"route_selection_expression,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes.
	// For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN.
	// The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Applicable for HTTP APIs.
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// Version identifier for the API. Must be between 1 and 64 characters in length.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type CorsConfigurationInitParameters struct {

	// Whether credentials are included in the CORS request.
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// Set of allowed HTTP headers.
	AllowHeaders []*string `json:"allowHeaders,omitempty" tf:"allow_headers,omitempty"`

	// Set of allowed HTTP methods.
	AllowMethods []*string `json:"allowMethods,omitempty" tf:"allow_methods,omitempty"`

	// Set of allowed origins.
	AllowOrigins []*string `json:"allowOrigins,omitempty" tf:"allow_origins,omitempty"`

	// Set of exposed HTTP headers.
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Number of seconds that the browser should cache preflight request results.
	MaxAge *int64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`
}

type CorsConfigurationObservation struct {

	// Whether credentials are included in the CORS request.
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// Set of allowed HTTP headers.
	AllowHeaders []*string `json:"allowHeaders,omitempty" tf:"allow_headers,omitempty"`

	// Set of allowed HTTP methods.
	AllowMethods []*string `json:"allowMethods,omitempty" tf:"allow_methods,omitempty"`

	// Set of allowed origins.
	AllowOrigins []*string `json:"allowOrigins,omitempty" tf:"allow_origins,omitempty"`

	// Set of exposed HTTP headers.
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Number of seconds that the browser should cache preflight request results.
	MaxAge *int64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`
}

type CorsConfigurationParameters struct {

	// Whether credentials are included in the CORS request.
	// +kubebuilder:validation:Optional
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// Set of allowed HTTP headers.
	// +kubebuilder:validation:Optional
	AllowHeaders []*string `json:"allowHeaders,omitempty" tf:"allow_headers,omitempty"`

	// Set of allowed HTTP methods.
	// +kubebuilder:validation:Optional
	AllowMethods []*string `json:"allowMethods,omitempty" tf:"allow_methods,omitempty"`

	// Set of allowed origins.
	// +kubebuilder:validation:Optional
	AllowOrigins []*string `json:"allowOrigins,omitempty" tf:"allow_origins,omitempty"`

	// Set of exposed HTTP headers.
	// +kubebuilder:validation:Optional
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Number of seconds that the browser should cache preflight request results.
	// +kubebuilder:validation:Optional
	MaxAge *int64 `json:"maxAge,omitempty" tf:"max_age,omitempty"`
}

// APISpec defines the desired state of API
type APISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIParameters `json:"forProvider"`
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
	InitProvider APIInitParameters `json:"initProvider,omitempty"`
}

// APIStatus defines the observed state of API.
type APIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// API is the Schema for the APIs API. Manages an Amazon API Gateway Version 2 API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type API struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocolType) || (has(self.initProvider) && has(self.initProvider.protocolType))",message="spec.forProvider.protocolType is a required parameter"
	Spec   APISpec   `json:"spec"`
	Status APIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIList contains a list of APIs
type APIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []API `json:"items"`
}

// Repository type metadata.
var (
	API_Kind             = "API"
	API_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: API_Kind}.String()
	API_KindAPIVersion   = API_Kind + "." + CRDGroupVersion.String()
	API_GroupVersionKind = CRDGroupVersion.WithKind(API_Kind)
)

func init() {
	SchemeBuilder.Register(&API{}, &APIList{})
}
