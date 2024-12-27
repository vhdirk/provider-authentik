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

type ProviderMicrosoftEntraInitParameters struct {
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ProviderMicrosoftEntraObservation struct {
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ProviderMicrosoftEntraParameters struct {

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// ProviderMicrosoftEntraSpec defines the desired state of ProviderMicrosoftEntra
type ProviderMicrosoftEntraSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProviderMicrosoftEntraParameters `json:"forProvider"`
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
	InitProvider ProviderMicrosoftEntraInitParameters `json:"initProvider,omitempty"`
}

// ProviderMicrosoftEntraStatus defines the observed state of ProviderMicrosoftEntra.
type ProviderMicrosoftEntraStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProviderMicrosoftEntraObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ProviderMicrosoftEntra is the Schema for the ProviderMicrosoftEntras API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type ProviderMicrosoftEntra struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.expression) || (has(self.initProvider) && has(self.initProvider.expression))",message="spec.forProvider.expression is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ProviderMicrosoftEntraSpec   `json:"spec"`
	Status ProviderMicrosoftEntraStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderMicrosoftEntraList contains a list of ProviderMicrosoftEntras
type ProviderMicrosoftEntraList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderMicrosoftEntra `json:"items"`
}

// Repository type metadata.
var (
	ProviderMicrosoftEntra_Kind             = "ProviderMicrosoftEntra"
	ProviderMicrosoftEntra_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProviderMicrosoftEntra_Kind}.String()
	ProviderMicrosoftEntra_KindAPIVersion   = ProviderMicrosoftEntra_Kind + "." + CRDGroupVersion.String()
	ProviderMicrosoftEntra_GroupVersionKind = CRDGroupVersion.WithKind(ProviderMicrosoftEntra_Kind)
)

func init() {
	SchemeBuilder.Register(&ProviderMicrosoftEntra{}, &ProviderMicrosoftEntraList{})
}
