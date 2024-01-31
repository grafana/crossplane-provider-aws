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

type DNSEntryInitParameters struct {
}

type DNSEntryObservation struct {

	// The DNS name.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// The ID of the private hosted zone.
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`
}

type DNSEntryParameters struct {
}

type DNSOptionsInitParameters struct {

	// The DNS records created for the endpoint. Valid values are ipv4, dualstack, service-defined, and ipv6.
	DNSRecordIPType *string `json:"dnsRecordIpType,omitempty" tf:"dns_record_ip_type,omitempty"`

	// Indicates whether to enable private DNS only for inbound endpoints. This option is available only for services that support both gateway and interface endpoints. It routes traffic that originates from the VPC to the gateway endpoint and traffic that originates from on-premises to the interface endpoint. Default is false. Can only be specified if private_dns_enabled is true.
	PrivateDNSOnlyForInboundResolverEndpoint *bool `json:"privateDnsOnlyForInboundResolverEndpoint,omitempty" tf:"private_dns_only_for_inbound_resolver_endpoint,omitempty"`
}

type DNSOptionsObservation struct {

	// The DNS records created for the endpoint. Valid values are ipv4, dualstack, service-defined, and ipv6.
	DNSRecordIPType *string `json:"dnsRecordIpType,omitempty" tf:"dns_record_ip_type,omitempty"`

	// Indicates whether to enable private DNS only for inbound endpoints. This option is available only for services that support both gateway and interface endpoints. It routes traffic that originates from the VPC to the gateway endpoint and traffic that originates from on-premises to the interface endpoint. Default is false. Can only be specified if private_dns_enabled is true.
	PrivateDNSOnlyForInboundResolverEndpoint *bool `json:"privateDnsOnlyForInboundResolverEndpoint,omitempty" tf:"private_dns_only_for_inbound_resolver_endpoint,omitempty"`
}

type DNSOptionsParameters struct {

	// The DNS records created for the endpoint. Valid values are ipv4, dualstack, service-defined, and ipv6.
	// +kubebuilder:validation:Optional
	DNSRecordIPType *string `json:"dnsRecordIpType,omitempty" tf:"dns_record_ip_type,omitempty"`

	// Indicates whether to enable private DNS only for inbound endpoints. This option is available only for services that support both gateway and interface endpoints. It routes traffic that originates from the VPC to the gateway endpoint and traffic that originates from on-premises to the interface endpoint. Default is false. Can only be specified if private_dns_enabled is true.
	// +kubebuilder:validation:Optional
	PrivateDNSOnlyForInboundResolverEndpoint *bool `json:"privateDnsOnlyForInboundResolverEndpoint,omitempty" tf:"private_dns_only_for_inbound_resolver_endpoint,omitempty"`
}

type VPCEndpointInitParameters_2 struct {

	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// The DNS options for the endpoint. See dns_options below.
	DNSOptions []DNSOptionsInitParameters `json:"dnsOptions,omitempty" tf:"dns_options,omitempty"`

	// The IP address type for the endpoint. Valid values are ipv4, dualstack, and ipv6.
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// A policy to attach to the endpoint that controls access to the service. This is a JSON formatted string. Defaults to full access. All Gateway and some Interface endpoints support policies - see the relevant AWS documentation for more details.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type Interface. Most users will want this enabled to allow services within the VPC to automatically use the endpoint.
	// Defaults to false.
	PrivateDNSEnabled *bool `json:"privateDnsEnabled,omitempty" tf:"private_dns_enabled,omitempty"`

	// The service name. For AWS services the service name is usually in the form com.amazonaws.<region>.<service> (the SageMaker Notebook service is an exception to this rule, the service name is in the form aws.sagemaker.<region>.notebook).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCEndpointService
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("service_name",true)
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Reference to a VPCEndpointService in ec2 to populate serviceName.
	// +kubebuilder:validation:Optional
	ServiceNameRef *v1.Reference `json:"serviceNameRef,omitempty" tf:"-"`

	// Selector for a VPCEndpointService in ec2 to populate serviceName.
	// +kubebuilder:validation:Optional
	ServiceNameSelector *v1.Selector `json:"serviceNameSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VPC endpoint type, Gateway, GatewayLoadBalancer, or Interface. Defaults to Gateway.
	VPCEndpointType *string `json:"vpcEndpointType,omitempty" tf:"vpc_endpoint_type,omitempty"`

	// The ID of the VPC in which the endpoint will be used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type VPCEndpointObservation_2 struct {

	// The Amazon Resource Name (ARN) of the VPC endpoint.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type Gateway.
	CidrBlocks []*string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`

	// The DNS entries for the VPC Endpoint. Applicable for endpoints of type Interface. DNS blocks are documented below.
	DNSEntry []DNSEntryObservation `json:"dnsEntry,omitempty" tf:"dns_entry,omitempty"`

	// The DNS options for the endpoint. See dns_options below.
	DNSOptions []DNSOptionsObservation `json:"dnsOptions,omitempty" tf:"dns_options,omitempty"`

	// The ID of the VPC endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address type for the endpoint. Valid values are ipv4, dualstack, and ipv6.
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type Interface.
	// +listType=set
	NetworkInterfaceIds []*string `json:"networkInterfaceIds,omitempty" tf:"network_interface_ids,omitempty"`

	// The ID of the AWS account that owns the VPC endpoint.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// A policy to attach to the endpoint that controls access to the service. This is a JSON formatted string. Defaults to full access. All Gateway and some Interface endpoints support policies - see the relevant AWS documentation for more details.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The prefix list ID of the exposed AWS service. Applicable for endpoints of type Gateway.
	PrefixListID *string `json:"prefixListId,omitempty" tf:"prefix_list_id,omitempty"`

	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type Interface. Most users will want this enabled to allow services within the VPC to automatically use the endpoint.
	// Defaults to false.
	PrivateDNSEnabled *bool `json:"privateDnsEnabled,omitempty" tf:"private_dns_enabled,omitempty"`

	// Whether or not the VPC Endpoint is being managed by its service - true or false.
	RequesterManaged *bool `json:"requesterManaged,omitempty" tf:"requester_managed,omitempty"`

	// One or more route table IDs. Applicable for endpoints of type Gateway.
	// +listType=set
	RouteTableIds []*string `json:"routeTableIds,omitempty" tf:"route_table_ids,omitempty"`

	// The ID of one or more security groups to associate with the network interface. Applicable for endpoints of type Interface.
	// If no security groups are specified, the VPC's default security group is associated with the endpoint.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The service name. For AWS services the service name is usually in the form com.amazonaws.<region>.<service> (the SageMaker Notebook service is an exception to this rule, the service name is in the form aws.sagemaker.<region>.notebook).
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The state of the VPC endpoint.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type GatewayLoadBalancer and Interface. Interface type endpoints cannot function without being assigned to a subnet.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The VPC endpoint type, Gateway, GatewayLoadBalancer, or Interface. Defaults to Gateway.
	VPCEndpointType *string `json:"vpcEndpointType,omitempty" tf:"vpc_endpoint_type,omitempty"`

	// The ID of the VPC in which the endpoint will be used.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCEndpointParameters_2 struct {

	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	// +kubebuilder:validation:Optional
	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`

	// The DNS options for the endpoint. See dns_options below.
	// +kubebuilder:validation:Optional
	DNSOptions []DNSOptionsParameters `json:"dnsOptions,omitempty" tf:"dns_options,omitempty"`

	// The IP address type for the endpoint. Valid values are ipv4, dualstack, and ipv6.
	// +kubebuilder:validation:Optional
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// A policy to attach to the endpoint that controls access to the service. This is a JSON formatted string. Defaults to full access. All Gateway and some Interface endpoints support policies - see the relevant AWS documentation for more details.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type Interface. Most users will want this enabled to allow services within the VPC to automatically use the endpoint.
	// Defaults to false.
	// +kubebuilder:validation:Optional
	PrivateDNSEnabled *bool `json:"privateDnsEnabled,omitempty" tf:"private_dns_enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The service name. For AWS services the service name is usually in the form com.amazonaws.<region>.<service> (the SageMaker Notebook service is an exception to this rule, the service name is in the form aws.sagemaker.<region>.notebook).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCEndpointService
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("service_name",true)
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Reference to a VPCEndpointService in ec2 to populate serviceName.
	// +kubebuilder:validation:Optional
	ServiceNameRef *v1.Reference `json:"serviceNameRef,omitempty" tf:"-"`

	// Selector for a VPCEndpointService in ec2 to populate serviceName.
	// +kubebuilder:validation:Optional
	ServiceNameSelector *v1.Selector `json:"serviceNameSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VPC endpoint type, Gateway, GatewayLoadBalancer, or Interface. Defaults to Gateway.
	// +kubebuilder:validation:Optional
	VPCEndpointType *string `json:"vpcEndpointType,omitempty" tf:"vpc_endpoint_type,omitempty"`

	// The ID of the VPC in which the endpoint will be used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// VPCEndpointSpec defines the desired state of VPCEndpoint
type VPCEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCEndpointParameters_2 `json:"forProvider"`
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
	InitProvider VPCEndpointInitParameters_2 `json:"initProvider,omitempty"`
}

// VPCEndpointStatus defines the observed state of VPCEndpoint.
type VPCEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCEndpointObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VPCEndpoint is the Schema for the VPCEndpoints API. Provides a VPC Endpoint resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCEndpointSpec   `json:"spec"`
	Status            VPCEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointList contains a list of VPCEndpoints
type VPCEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpoint `json:"items"`
}

// Repository type metadata.
var (
	VPCEndpoint_Kind             = "VPCEndpoint"
	VPCEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCEndpoint_Kind}.String()
	VPCEndpoint_KindAPIVersion   = VPCEndpoint_Kind + "." + CRDGroupVersion.String()
	VPCEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(VPCEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCEndpoint{}, &VPCEndpointList{})
}
