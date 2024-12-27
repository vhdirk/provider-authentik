package propertymapping

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "propertymapping"

// Configure configures the propertymapping provider.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_property_mapping_google_workspace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "GoogleWorkspace"
	})
	p.AddResourceConfigurator("authentik_property_mapping_ldap", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LDAP"
	})
	p.AddResourceConfigurator("authentik_property_mapping_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MicrosoftEntra"
	})
	p.AddResourceConfigurator("authentik_property_mapping_notification", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Notification"
	})
	p.AddResourceConfigurator("authentik_property_mapping_rac", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RAC"
	})
	p.AddResourceConfigurator("authentik_property_mapping_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SAML"
	})
	p.AddResourceConfigurator("authentik_property_mapping_scim", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SCIM"
	})
}
