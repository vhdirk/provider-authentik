/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/vhdirk/provider-authentik/apis/authentik/v1alpha1"
	v1alpha1directory "github.com/vhdirk/provider-authentik/apis/directory/v1alpha1"
	v1alpha1policy "github.com/vhdirk/provider-authentik/apis/policy/v1alpha1"
	v1alpha1propertymapping "github.com/vhdirk/provider-authentik/apis/propertymapping/v1alpha1"
	v1alpha1provider "github.com/vhdirk/provider-authentik/apis/provider/v1alpha1"
	v1alpha1stage "github.com/vhdirk/provider-authentik/apis/stage/v1alpha1"
	v1alpha1apis "github.com/vhdirk/provider-authentik/apis/v1alpha1"
	v1beta1 "github.com/vhdirk/provider-authentik/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1directory.SchemeBuilder.AddToScheme,
		v1alpha1policy.SchemeBuilder.AddToScheme,
		v1alpha1propertymapping.SchemeBuilder.AddToScheme,
		v1alpha1provider.SchemeBuilder.AddToScheme,
		v1alpha1stage.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}