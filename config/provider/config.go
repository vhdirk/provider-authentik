package provider

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/vhdirk/crossplane-provider-authentik/config/base"
)

const shortGroup = "provider"

// Configure configures the provider provider.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_provider_google_workspace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "GoogleWorkspace"

		r.References["authorization_flow"] = base.FlowRef
	})
	p.AddResourceConfigurator("authentik_provider_ldap", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LDAP"

		r.References["authorization_flow"] = base.FlowRef
	})
	p.AddResourceConfigurator("authentik_provider_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MicrosoftEntra"

		r.References["authorization_flow"] = base.FlowRef
	})

	p.AddResourceConfigurator("authentik_provider_oauth2", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OAuth2"

		r.References["authorization_flow"] = base.FlowRef
		r.References["authentication_flow"] = base.FlowRef
		r.References["invalidation_flow"] = base.FlowRef
		r.References["signing_key"] = base.CertificateKeyPairRef
		r.References["encryption_key"] = base.CertificateKeyPairRef

		// r.References["property_mappings"] = base Prop
	})
	p.AddResourceConfigurator("authentik_provider_proxy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Proxy"

		r.References["authorization_flow"] = base.FlowRef
		r.References["authentication_flow"] = base.FlowRef
		r.References["invalidation_flow"] = base.FlowRef

	})
	p.AddResourceConfigurator("authentik_provider_rac", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RAC"

		r.References["authentication_flow"] = base.FlowRef
		r.References["authorization_flow"] = base.FlowRef

	})
	p.AddResourceConfigurator("authentik_provider_radius", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Radius"

		r.References["authorization_flow"] = base.FlowRef
	})
	p.AddResourceConfigurator("authentik_provider_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SAML"

		r.References["authorization_flow"] = base.FlowRef
	})
	p.AddResourceConfigurator("authentik_provider_scim", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SCIM"

		r.References["authorization_flow"] = base.FlowRef
	})
}
