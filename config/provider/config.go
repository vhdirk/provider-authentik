package provider

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "provider"

var flowUUIDRef = config.Reference{
	Type:      "github.com/vhdirk/provider-authentik/apis/authentik/v1alpha1.Flow",
	Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)`,
}

// Configure configures the provider provider.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_provider_ldap", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LDAP"

		r.References["authorization_flow"] = flowUUIDRef
	})
	p.AddResourceConfigurator("authentik_provider_oauth2", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OAuth2"

		r.References["authorization_flow"] = flowUUIDRef
		r.References["property_mappings"] = config.Reference{
			Type:      "github.com/vhdirk/provider-authentik/apis/authentik/v1alpha1.ScopeMapping",
			Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
	p.AddResourceConfigurator("authentik_provider_proxy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Proxy"

		r.References["authorization_flow"] = flowUUIDRef
	})
	p.AddResourceConfigurator("authentik_provider_radius", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Radius"

		r.References["authorization_flow"] = flowUUIDRef
	})
	p.AddResourceConfigurator("authentik_provider_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SAML"

		r.References["authorization_flow"] = flowUUIDRef
	})
	p.AddResourceConfigurator("authentik_provider_scim", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SCIM"

		r.References["authorization_flow"] = flowUUIDRef
	})
}
