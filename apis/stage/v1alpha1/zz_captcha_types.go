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

type CaptchaInitParameters struct {

	// (String) Defaults to https://www.recaptcha.net/recaptcha/api/siteverify.
	// Defaults to `https://www.recaptcha.net/recaptcha/api/siteverify`.
	APIURL *string `json:"apiUrl,omitempty" tf:"api_url,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	ErrorOnInvalidScore *bool `json:"errorOnInvalidScore,omitempty" tf:"error_on_invalid_score,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	Interactive *bool `json:"interactive,omitempty" tf:"interactive,omitempty"`

	// (String) Defaults to https://www.recaptcha.net/recaptcha/api.js.
	// Defaults to `https://www.recaptcha.net/recaptcha/api.js`.
	JsURL *string `json:"jsUrl,omitempty" tf:"js_url,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive)
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// (String)
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// (Number) Defaults to 0.5.
	// Defaults to `0.5`.
	ScoreMaxThreshold *float64 `json:"scoreMaxThreshold,omitempty" tf:"score_max_threshold,omitempty"`

	// (Number) Defaults to 1.
	// Defaults to `1`.
	ScoreMinThreshold *float64 `json:"scoreMinThreshold,omitempty" tf:"score_min_threshold,omitempty"`
}

type CaptchaObservation struct {

	// (String) Defaults to https://www.recaptcha.net/recaptcha/api/siteverify.
	// Defaults to `https://www.recaptcha.net/recaptcha/api/siteverify`.
	APIURL *string `json:"apiUrl,omitempty" tf:"api_url,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	ErrorOnInvalidScore *bool `json:"errorOnInvalidScore,omitempty" tf:"error_on_invalid_score,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	Interactive *bool `json:"interactive,omitempty" tf:"interactive,omitempty"`

	// (String) Defaults to https://www.recaptcha.net/recaptcha/api.js.
	// Defaults to `https://www.recaptcha.net/recaptcha/api.js`.
	JsURL *string `json:"jsUrl,omitempty" tf:"js_url,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String)
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// (Number) Defaults to 0.5.
	// Defaults to `0.5`.
	ScoreMaxThreshold *float64 `json:"scoreMaxThreshold,omitempty" tf:"score_max_threshold,omitempty"`

	// (Number) Defaults to 1.
	// Defaults to `1`.
	ScoreMinThreshold *float64 `json:"scoreMinThreshold,omitempty" tf:"score_min_threshold,omitempty"`
}

type CaptchaParameters struct {

	// (String) Defaults to https://www.recaptcha.net/recaptcha/api/siteverify.
	// Defaults to `https://www.recaptcha.net/recaptcha/api/siteverify`.
	// +kubebuilder:validation:Optional
	APIURL *string `json:"apiUrl,omitempty" tf:"api_url,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	ErrorOnInvalidScore *bool `json:"errorOnInvalidScore,omitempty" tf:"error_on_invalid_score,omitempty"`

	// (Boolean) Defaults to false.
	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	Interactive *bool `json:"interactive,omitempty" tf:"interactive,omitempty"`

	// (String) Defaults to https://www.recaptcha.net/recaptcha/api.js.
	// Defaults to `https://www.recaptcha.net/recaptcha/api.js`.
	// +kubebuilder:validation:Optional
	JsURL *string `json:"jsUrl,omitempty" tf:"js_url,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive)
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// (String)
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// (Number) Defaults to 0.5.
	// Defaults to `0.5`.
	// +kubebuilder:validation:Optional
	ScoreMaxThreshold *float64 `json:"scoreMaxThreshold,omitempty" tf:"score_max_threshold,omitempty"`

	// (Number) Defaults to 1.
	// Defaults to `1`.
	// +kubebuilder:validation:Optional
	ScoreMinThreshold *float64 `json:"scoreMinThreshold,omitempty" tf:"score_min_threshold,omitempty"`
}

// CaptchaSpec defines the desired state of Captcha
type CaptchaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CaptchaParameters `json:"forProvider"`
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
	InitProvider CaptchaInitParameters `json:"initProvider,omitempty"`
}

// CaptchaStatus defines the observed state of Captcha.
type CaptchaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CaptchaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Captcha is the Schema for the Captchas API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type Captcha struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privateKeySecretRef)",message="spec.forProvider.privateKeySecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publicKey) || (has(self.initProvider) && has(self.initProvider.publicKey))",message="spec.forProvider.publicKey is a required parameter"
	Spec   CaptchaSpec   `json:"spec"`
	Status CaptchaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CaptchaList contains a list of Captchas
type CaptchaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Captcha `json:"items"`
}

// Repository type metadata.
var (
	Captcha_Kind             = "Captcha"
	Captcha_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Captcha_Kind}.String()
	Captcha_KindAPIVersion   = Captcha_Kind + "." + CRDGroupVersion.String()
	Captcha_GroupVersionKind = CRDGroupVersion.WithKind(Captcha_Kind)
)

func init() {
	SchemeBuilder.Register(&Captcha{}, &CaptchaList{})
}
