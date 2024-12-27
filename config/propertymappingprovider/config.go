package propertymappingprovider

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "propertymappingprovider"

// Configure configures the propertymapping provider.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_property_mapping_provider_google_workspace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "GoogleWorkspace"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MicrosoftEntra"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_rac", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RAC"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_radius", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Radius"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SAML"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_scim", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SCIM"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_scope", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Scope"
	})
}
