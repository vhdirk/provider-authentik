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

type PromptInitParameters struct {

	// (List of String)
	Fields []*string `json:"fields,omitempty" tf:"fields,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String)
	ValidationPolicies []*string `json:"validationPolicies,omitempty" tf:"validation_policies,omitempty"`
}

type PromptObservation struct {

	// (List of String)
	Fields []*string `json:"fields,omitempty" tf:"fields,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String)
	ValidationPolicies []*string `json:"validationPolicies,omitempty" tf:"validation_policies,omitempty"`
}

type PromptParameters struct {

	// (List of String)
	// +kubebuilder:validation:Optional
	Fields []*string `json:"fields,omitempty" tf:"fields,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (List of String)
	// +kubebuilder:validation:Optional
	ValidationPolicies []*string `json:"validationPolicies,omitempty" tf:"validation_policies,omitempty"`
}

// PromptSpec defines the desired state of Prompt
type PromptSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PromptParameters `json:"forProvider"`
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
	InitProvider PromptInitParameters `json:"initProvider,omitempty"`
}

// PromptStatus defines the observed state of Prompt.
type PromptStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PromptObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Prompt is the Schema for the Prompts API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type Prompt struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fields) || (has(self.initProvider) && has(self.initProvider.fields))",message="spec.forProvider.fields is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   PromptSpec   `json:"spec"`
	Status PromptStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PromptList contains a list of Prompts
type PromptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Prompt `json:"items"`
}

// Repository type metadata.
var (
	Prompt_Kind             = "Prompt"
	Prompt_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Prompt_Kind}.String()
	Prompt_KindAPIVersion   = Prompt_Kind + "." + CRDGroupVersion.String()
	Prompt_GroupVersionKind = CRDGroupVersion.WithKind(Prompt_Kind)
)

func init() {
	SchemeBuilder.Register(&Prompt{}, &PromptList{})
}
