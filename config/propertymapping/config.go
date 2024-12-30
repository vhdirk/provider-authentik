package propertymapping

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "propertymapping"

// Configure configures the propertymapping provider.
func Configure(p *config.Provider) {
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_google_workspace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "GoogleWorkspace"
	})
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_ldap", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LDAP"
	})
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MicrosoftEntra"
	})
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_notification", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Notification"
	})
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_rac", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RAC"
	})
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SAML"
	})
	// deprectated
	p.AddResourceConfigurator("authentik_property_mapping_scim", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SCIM"
	})

	p.AddResourceConfigurator("authentik_property_mapping_provider_google_workspace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderGoogleWorkspace"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderMicrosoftEntra"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_rac", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderRAC"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_radius", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderRadius"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderSAML"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_scim", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderSCIM"
	})
	p.AddResourceConfigurator("authentik_property_mapping_provider_scope", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ProviderScope"
	})

	p.AddResourceConfigurator("authentik_property_mapping_source_kerberos", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceKerberos"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_ldap", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceLDAP"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_oauth", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceOAuth"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_plex", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourcePlex"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceSAML"
	})
	p.AddResourceConfigurator("authentik_property_mapping_source_scim", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceSCIM"
	})

}
