package base

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = ""

// Reference a flow
var FlowUUIDRef = config.Reference{
	Type:      "github.com/vhdirk/provider-authentik/apis/authentik/v1alpha1.Flow",
	Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)`,
}

// Configure configures the base provider.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_application", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Application"

		// // TODO: Support generic references
		// // https://github.com/crossplane/upjet/issues/95
		// r.References["protocol_provider"] = config.Reference{
		// 	Type:      "github.com/vhdirk/provider-authentik/apis/provider/v1alpha1.Proxy",
		// 	Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
		// }
	})
	p.AddResourceConfigurator("authentik_outpost", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Outpost"
	})
	p.AddResourceConfigurator("authentik_service_connection_kubernetes", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ServiceConnectionKubernetes"
	})
	p.AddResourceConfigurator("authentik_blueprint", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Blueprint"
	})
	p.AddResourceConfigurator("authentik_flow", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Flow"
	})
	p.AddResourceConfigurator("authentik_flow_stage_binding", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FlowStageBinding"
	})
	p.AddResourceConfigurator("authentik_scope_mapping", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ScopeMapping"
	})
	p.AddResourceConfigurator("authentik_certificate_key_pair", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CertificateKeyPair"
	})
	p.AddResourceConfigurator("authentik_tenant", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Tenant"
	})
	p.AddResourceConfigurator("authentik_event_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EventRule"
	})
	p.AddResourceConfigurator("authentik_event_transport", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EventTransport"
	})
}
