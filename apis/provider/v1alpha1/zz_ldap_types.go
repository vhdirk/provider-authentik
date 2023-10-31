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

type LDAPInitParameters struct {

	// (String)
	BaseDn *string `json:"baseDn,omitempty" tf:"base_dn,omitempty"`

	// (String)
	BindFlow *string `json:"bindFlow,omitempty" tf:"bind_flow,omitempty"`

	// (String) Defaults to direct.
	// Defaults to `direct`.
	BindMode *string `json:"bindMode,omitempty" tf:"bind_mode,omitempty"`

	// (String)
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// (Number) Defaults to 4000.
	// Defaults to `4000`.
	GIDStartNumber *float64 `json:"gidStartNumber,omitempty" tf:"gid_start_number,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	MfaSupport *bool `json:"mfaSupport,omitempty" tf:"mfa_support,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	SearchGroup *string `json:"searchGroup,omitempty" tf:"search_group,omitempty"`

	// (String) Defaults to direct.
	// Defaults to `direct`.
	SearchMode *string `json:"searchMode,omitempty" tf:"search_mode,omitempty"`

	// (String)
	TLSServerName *string `json:"tlsServerName,omitempty" tf:"tls_server_name,omitempty"`

	// (Number) Defaults to 2000.
	// Defaults to `2000`.
	UIDStartNumber *float64 `json:"uidStartNumber,omitempty" tf:"uid_start_number,omitempty"`
}

type LDAPObservation struct {

	// (String)
	BaseDn *string `json:"baseDn,omitempty" tf:"base_dn,omitempty"`

	// (String)
	BindFlow *string `json:"bindFlow,omitempty" tf:"bind_flow,omitempty"`

	// (String) Defaults to direct.
	// Defaults to `direct`.
	BindMode *string `json:"bindMode,omitempty" tf:"bind_mode,omitempty"`

	// (String)
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// (Number) Defaults to 4000.
	// Defaults to `4000`.
	GIDStartNumber *float64 `json:"gidStartNumber,omitempty" tf:"gid_start_number,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	MfaSupport *bool `json:"mfaSupport,omitempty" tf:"mfa_support,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	SearchGroup *string `json:"searchGroup,omitempty" tf:"search_group,omitempty"`

	// (String) Defaults to direct.
	// Defaults to `direct`.
	SearchMode *string `json:"searchMode,omitempty" tf:"search_mode,omitempty"`

	// (String)
	TLSServerName *string `json:"tlsServerName,omitempty" tf:"tls_server_name,omitempty"`

	// (Number) Defaults to 2000.
	// Defaults to `2000`.
	UIDStartNumber *float64 `json:"uidStartNumber,omitempty" tf:"uid_start_number,omitempty"`
}

type LDAPParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	BaseDn *string `json:"baseDn,omitempty" tf:"base_dn,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	BindFlow *string `json:"bindFlow,omitempty" tf:"bind_flow,omitempty"`

	// (String) Defaults to direct.
	// Defaults to `direct`.
	// +kubebuilder:validation:Optional
	BindMode *string `json:"bindMode,omitempty" tf:"bind_mode,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// (Number) Defaults to 4000.
	// Defaults to `4000`.
	// +kubebuilder:validation:Optional
	GIDStartNumber *float64 `json:"gidStartNumber,omitempty" tf:"gid_start_number,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	MfaSupport *bool `json:"mfaSupport,omitempty" tf:"mfa_support,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	SearchGroup *string `json:"searchGroup,omitempty" tf:"search_group,omitempty"`

	// (String) Defaults to direct.
	// Defaults to `direct`.
	// +kubebuilder:validation:Optional
	SearchMode *string `json:"searchMode,omitempty" tf:"search_mode,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	TLSServerName *string `json:"tlsServerName,omitempty" tf:"tls_server_name,omitempty"`

	// (Number) Defaults to 2000.
	// Defaults to `2000`.
	// +kubebuilder:validation:Optional
	UIDStartNumber *float64 `json:"uidStartNumber,omitempty" tf:"uid_start_number,omitempty"`
}

// LDAPSpec defines the desired state of LDAP
type LDAPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LDAPParameters `json:"forProvider"`
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
	InitProvider LDAPInitParameters `json:"initProvider,omitempty"`
}

// LDAPStatus defines the observed state of LDAP.
type LDAPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LDAPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LDAP is the Schema for the LDAPs API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type LDAP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.baseDn) || (has(self.initProvider) && has(self.initProvider.baseDn))",message="spec.forProvider.baseDn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bindFlow) || (has(self.initProvider) && has(self.initProvider.bindFlow))",message="spec.forProvider.bindFlow is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   LDAPSpec   `json:"spec"`
	Status LDAPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LDAPList contains a list of LDAPs
type LDAPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LDAP `json:"items"`
}

// Repository type metadata.
var (
	LDAP_Kind             = "LDAP"
	LDAP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LDAP_Kind}.String()
	LDAP_KindAPIVersion   = LDAP_Kind + "." + CRDGroupVersion.String()
	LDAP_GroupVersionKind = CRDGroupVersion.WithKind(LDAP_Kind)
)

func init() {
	SchemeBuilder.Register(&LDAP{}, &LDAPList{})
}