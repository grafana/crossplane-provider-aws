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

type ContentTypeProfileConfigInitParameters struct {

	// Object that contains an attribute items that contains the list of configurations for a field-level encryption content type-profile. See Content Type Profile.
	ContentTypeProfiles []ContentTypeProfilesInitParameters `json:"contentTypeProfiles,omitempty" tf:"content_type_profiles,omitempty"`

	// specifies what to do when an unknown content type is provided for the profile. If true, content is forwarded without being encrypted when the content type is unknown. If false (the default), an error is returned when the content type is unknown.
	ForwardWhenContentTypeIsUnknown *bool `json:"forwardWhenContentTypeIsUnknown,omitempty" tf:"forward_when_content_type_is_unknown,omitempty"`
}

type ContentTypeProfileConfigObservation struct {

	// Object that contains an attribute items that contains the list of configurations for a field-level encryption content type-profile. See Content Type Profile.
	ContentTypeProfiles []ContentTypeProfilesObservation `json:"contentTypeProfiles,omitempty" tf:"content_type_profiles,omitempty"`

	// specifies what to do when an unknown content type is provided for the profile. If true, content is forwarded without being encrypted when the content type is unknown. If false (the default), an error is returned when the content type is unknown.
	ForwardWhenContentTypeIsUnknown *bool `json:"forwardWhenContentTypeIsUnknown,omitempty" tf:"forward_when_content_type_is_unknown,omitempty"`
}

type ContentTypeProfileConfigParameters struct {

	// Object that contains an attribute items that contains the list of configurations for a field-level encryption content type-profile. See Content Type Profile.
	// +kubebuilder:validation:Optional
	ContentTypeProfiles []ContentTypeProfilesParameters `json:"contentTypeProfiles" tf:"content_type_profiles,omitempty"`

	// specifies what to do when an unknown content type is provided for the profile. If true, content is forwarded without being encrypted when the content type is unknown. If false (the default), an error is returned when the content type is unknown.
	// +kubebuilder:validation:Optional
	ForwardWhenContentTypeIsUnknown *bool `json:"forwardWhenContentTypeIsUnknown" tf:"forward_when_content_type_is_unknown,omitempty"`
}

type ContentTypeProfilesInitParameters struct {
	Items []ContentTypeProfilesItemsInitParameters `json:"items,omitempty" tf:"items,omitempty"`
}

type ContentTypeProfilesItemsInitParameters struct {

	// he content type for a field-level encryption content type-profile mapping. Valid value is application/x-www-form-urlencoded.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// The format for a field-level encryption content type-profile mapping. Valid value is URLEncoded.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The profile ID for a field-level encryption content type-profile mapping.
	ProfileID *string `json:"profileId,omitempty" tf:"profile_id,omitempty"`
}

type ContentTypeProfilesItemsObservation struct {

	// he content type for a field-level encryption content type-profile mapping. Valid value is application/x-www-form-urlencoded.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// The format for a field-level encryption content type-profile mapping. Valid value is URLEncoded.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The profile ID for a field-level encryption content type-profile mapping.
	ProfileID *string `json:"profileId,omitempty" tf:"profile_id,omitempty"`
}

type ContentTypeProfilesItemsParameters struct {

	// he content type for a field-level encryption content type-profile mapping. Valid value is application/x-www-form-urlencoded.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// The format for a field-level encryption content type-profile mapping. Valid value is URLEncoded.
	// +kubebuilder:validation:Optional
	Format *string `json:"format" tf:"format,omitempty"`

	// The profile ID for a field-level encryption content type-profile mapping.
	// +kubebuilder:validation:Optional
	ProfileID *string `json:"profileId,omitempty" tf:"profile_id,omitempty"`
}

type ContentTypeProfilesObservation struct {
	Items []ContentTypeProfilesItemsObservation `json:"items,omitempty" tf:"items,omitempty"`
}

type ContentTypeProfilesParameters struct {

	// +kubebuilder:validation:Optional
	Items []ContentTypeProfilesItemsParameters `json:"items" tf:"items,omitempty"`
}

type FieldLevelEncryptionConfigInitParameters struct {

	// An optional comment about the Field Level Encryption Config.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Content Type Profile Config specifies when to forward content if a content type isn't recognized and profiles to use as by default in a request if a query argument doesn't specify a profile to use.
	ContentTypeProfileConfig []ContentTypeProfileConfigInitParameters `json:"contentTypeProfileConfig,omitempty" tf:"content_type_profile_config,omitempty"`

	// Query Arg Profile Config that specifies when to forward content if a profile isn't found and the profile that can be provided as a query argument in a request.
	QueryArgProfileConfig []QueryArgProfileConfigInitParameters `json:"queryArgProfileConfig,omitempty" tf:"query_arg_profile_config,omitempty"`
}

type FieldLevelEncryptionConfigObservation struct {

	// Internal value used by CloudFront to allow future updates to the Field Level Encryption Config.
	CallerReference *string `json:"callerReference,omitempty" tf:"caller_reference,omitempty"`

	// An optional comment about the Field Level Encryption Config.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Content Type Profile Config specifies when to forward content if a content type isn't recognized and profiles to use as by default in a request if a query argument doesn't specify a profile to use.
	ContentTypeProfileConfig []ContentTypeProfileConfigObservation `json:"contentTypeProfileConfig,omitempty" tf:"content_type_profile_config,omitempty"`

	// The current version of the Field Level Encryption Config. For example: E2QWRUHAPOMQZL.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// The identifier for the Field Level Encryption Config. For example: K3D5EWEUDCCXON.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Query Arg Profile Config that specifies when to forward content if a profile isn't found and the profile that can be provided as a query argument in a request.
	QueryArgProfileConfig []QueryArgProfileConfigObservation `json:"queryArgProfileConfig,omitempty" tf:"query_arg_profile_config,omitempty"`
}

type FieldLevelEncryptionConfigParameters struct {

	// An optional comment about the Field Level Encryption Config.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Content Type Profile Config specifies when to forward content if a content type isn't recognized and profiles to use as by default in a request if a query argument doesn't specify a profile to use.
	// +kubebuilder:validation:Optional
	ContentTypeProfileConfig []ContentTypeProfileConfigParameters `json:"contentTypeProfileConfig,omitempty" tf:"content_type_profile_config,omitempty"`

	// Query Arg Profile Config that specifies when to forward content if a profile isn't found and the profile that can be provided as a query argument in a request.
	// +kubebuilder:validation:Optional
	QueryArgProfileConfig []QueryArgProfileConfigParameters `json:"queryArgProfileConfig,omitempty" tf:"query_arg_profile_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type QueryArgProfileConfigInitParameters struct {

	// Flag to set if you want a request to be forwarded to the origin even if the profile specified by the field-level encryption query argument, fle-profile, is unknown.
	ForwardWhenQueryArgProfileIsUnknown *bool `json:"forwardWhenQueryArgProfileIsUnknown,omitempty" tf:"forward_when_query_arg_profile_is_unknown,omitempty"`

	// Object that contains an attribute items that contains the list ofrofiles specified for query argument-profile mapping for field-level encryption. see Query Arg Profile.
	QueryArgProfiles []QueryArgProfilesInitParameters `json:"queryArgProfiles,omitempty" tf:"query_arg_profiles,omitempty"`
}

type QueryArgProfileConfigObservation struct {

	// Flag to set if you want a request to be forwarded to the origin even if the profile specified by the field-level encryption query argument, fle-profile, is unknown.
	ForwardWhenQueryArgProfileIsUnknown *bool `json:"forwardWhenQueryArgProfileIsUnknown,omitempty" tf:"forward_when_query_arg_profile_is_unknown,omitempty"`

	// Object that contains an attribute items that contains the list ofrofiles specified for query argument-profile mapping for field-level encryption. see Query Arg Profile.
	QueryArgProfiles []QueryArgProfilesObservation `json:"queryArgProfiles,omitempty" tf:"query_arg_profiles,omitempty"`
}

type QueryArgProfileConfigParameters struct {

	// Flag to set if you want a request to be forwarded to the origin even if the profile specified by the field-level encryption query argument, fle-profile, is unknown.
	// +kubebuilder:validation:Optional
	ForwardWhenQueryArgProfileIsUnknown *bool `json:"forwardWhenQueryArgProfileIsUnknown" tf:"forward_when_query_arg_profile_is_unknown,omitempty"`

	// Object that contains an attribute items that contains the list ofrofiles specified for query argument-profile mapping for field-level encryption. see Query Arg Profile.
	// +kubebuilder:validation:Optional
	QueryArgProfiles []QueryArgProfilesParameters `json:"queryArgProfiles,omitempty" tf:"query_arg_profiles,omitempty"`
}

type QueryArgProfilesInitParameters struct {
	Items []QueryArgProfilesItemsInitParameters `json:"items,omitempty" tf:"items,omitempty"`
}

type QueryArgProfilesItemsInitParameters struct {

	// Query argument for field-level encryption query argument-profile mapping.
	QueryArg *string `json:"queryArg,omitempty" tf:"query_arg,omitempty"`
}

type QueryArgProfilesItemsObservation struct {

	// The profile ID for a field-level encryption content type-profile mapping.
	ProfileID *string `json:"profileId,omitempty" tf:"profile_id,omitempty"`

	// Query argument for field-level encryption query argument-profile mapping.
	QueryArg *string `json:"queryArg,omitempty" tf:"query_arg,omitempty"`
}

type QueryArgProfilesItemsParameters struct {

	// The profile ID for a field-level encryption content type-profile mapping.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudfront/v1beta1.FieldLevelEncryptionProfile
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ProfileID *string `json:"profileId,omitempty" tf:"profile_id,omitempty"`

	// Reference to a FieldLevelEncryptionProfile in cloudfront to populate profileId.
	// +kubebuilder:validation:Optional
	ProfileIDRef *v1.Reference `json:"profileIdRef,omitempty" tf:"-"`

	// Selector for a FieldLevelEncryptionProfile in cloudfront to populate profileId.
	// +kubebuilder:validation:Optional
	ProfileIDSelector *v1.Selector `json:"profileIdSelector,omitempty" tf:"-"`

	// Query argument for field-level encryption query argument-profile mapping.
	// +kubebuilder:validation:Optional
	QueryArg *string `json:"queryArg" tf:"query_arg,omitempty"`
}

type QueryArgProfilesObservation struct {
	Items []QueryArgProfilesItemsObservation `json:"items,omitempty" tf:"items,omitempty"`
}

type QueryArgProfilesParameters struct {

	// +kubebuilder:validation:Optional
	Items []QueryArgProfilesItemsParameters `json:"items,omitempty" tf:"items,omitempty"`
}

// FieldLevelEncryptionConfigSpec defines the desired state of FieldLevelEncryptionConfig
type FieldLevelEncryptionConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FieldLevelEncryptionConfigParameters `json:"forProvider"`
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
	InitProvider FieldLevelEncryptionConfigInitParameters `json:"initProvider,omitempty"`
}

// FieldLevelEncryptionConfigStatus defines the observed state of FieldLevelEncryptionConfig.
type FieldLevelEncryptionConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FieldLevelEncryptionConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FieldLevelEncryptionConfig is the Schema for the FieldLevelEncryptionConfigs API. Provides a CloudFront Field-level Encryption Config resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FieldLevelEncryptionConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.contentTypeProfileConfig) || (has(self.initProvider) && has(self.initProvider.contentTypeProfileConfig))",message="spec.forProvider.contentTypeProfileConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.queryArgProfileConfig) || (has(self.initProvider) && has(self.initProvider.queryArgProfileConfig))",message="spec.forProvider.queryArgProfileConfig is a required parameter"
	Spec   FieldLevelEncryptionConfigSpec   `json:"spec"`
	Status FieldLevelEncryptionConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FieldLevelEncryptionConfigList contains a list of FieldLevelEncryptionConfigs
type FieldLevelEncryptionConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FieldLevelEncryptionConfig `json:"items"`
}

// Repository type metadata.
var (
	FieldLevelEncryptionConfig_Kind             = "FieldLevelEncryptionConfig"
	FieldLevelEncryptionConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FieldLevelEncryptionConfig_Kind}.String()
	FieldLevelEncryptionConfig_KindAPIVersion   = FieldLevelEncryptionConfig_Kind + "." + CRDGroupVersion.String()
	FieldLevelEncryptionConfig_GroupVersionKind = CRDGroupVersion.WithKind(FieldLevelEncryptionConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&FieldLevelEncryptionConfig{}, &FieldLevelEncryptionConfigList{})
}
