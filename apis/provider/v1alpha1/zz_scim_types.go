/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SCIMInitParameters struct {

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String)
	PropertyMappings []*string `json:"propertyMappings,omitempty" tf:"property_mappings,omitempty"`

	// (List of String)
	PropertyMappingsGroup []*string `json:"propertyMappingsGroup,omitempty" tf:"property_mappings_group,omitempty"`

	// (String)
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SCIMObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String)
	PropertyMappings []*string `json:"propertyMappings,omitempty" tf:"property_mappings,omitempty"`

	// (List of String)
	PropertyMappingsGroup []*string `json:"propertyMappingsGroup,omitempty" tf:"property_mappings_group,omitempty"`

	// (String)
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SCIMParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String)
	// +kubebuilder:validation:Optional
	PropertyMappings []*string `json:"propertyMappings,omitempty" tf:"property_mappings,omitempty"`

	// (List of String)
	// +kubebuilder:validation:Optional
	PropertyMappingsGroup []*string `json:"propertyMappingsGroup,omitempty" tf:"property_mappings_group,omitempty"`

	// (String, Sensitive)
	// +kubebuilder:validation:Optional
	TokenSecretRef v1.SecretKeySelector `json:"tokenSecretRef" tf:"-"`

	// (String)
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

// SCIMSpec defines the desired state of SCIM
type SCIMSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SCIMParameters `json:"forProvider"`
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
	InitProvider SCIMInitParameters `json:"initProvider,omitempty"`
}

// SCIMStatus defines the observed state of SCIM.
type SCIMStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SCIMObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SCIM is the Schema for the SCIMs API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type SCIM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tokenSecretRef)",message="spec.forProvider.tokenSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.url) || (has(self.initProvider) && has(self.initProvider.url))",message="spec.forProvider.url is a required parameter"
	Spec   SCIMSpec   `json:"spec"`
	Status SCIMStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SCIMList contains a list of SCIMs
type SCIMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SCIM `json:"items"`
}

// Repository type metadata.
var (
	SCIM_Kind             = "SCIM"
	SCIM_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SCIM_Kind}.String()
	SCIM_KindAPIVersion   = SCIM_Kind + "." + CRDGroupVersion.String()
	SCIM_GroupVersionKind = CRDGroupVersion.WithKind(SCIM_Kind)
)

func init() {
	SchemeBuilder.Register(&SCIM{}, &SCIMList{})
}