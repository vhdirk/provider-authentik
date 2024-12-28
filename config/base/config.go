package base

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const shortGroup = ""

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

		// provider_id in authentik is just an integer reference
		// since it uses Django model inheritance, we don't need type checking
		r.TerraformResource.Schema["protocol_provider"] = &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		}

		// We'll resolve the reference at runtime since we can't type-check
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"protocol_provider"},
		}

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

		r.TerraformResource.Schema["slug"] = &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    false,
			Description: "Unique identifier for the flow.",
		}

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
