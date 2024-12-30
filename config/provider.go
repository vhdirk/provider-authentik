/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"github.com/vhdirk/crossplane-provider-authentik/config/base"
	"github.com/vhdirk/crossplane-provider-authentik/config/directory"
	"github.com/vhdirk/crossplane-provider-authentik/config/policy"
	"github.com/vhdirk/crossplane-provider-authentik/config/propertymapping"
	"github.com/vhdirk/crossplane-provider-authentik/config/provider"
	"github.com/vhdirk/crossplane-provider-authentik/config/stage"

	authentikProvider "goauthentik.io/terraform-provider-authentik/provider"
)

const (
	resourcePrefix = "authentik"
	modulePath     = "github.com/vhdirk/crossplane-provider-authentik"
	rootGroup      = "authentik.crossplane.io"
)

// these will be set by the goreleaser configuration
// to appropriate values for the compiled binary
var version string = "dev"

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// workaround for the TF Azure v3.57.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(generationProvider bool) (*ujconfig.Provider, error) {
	var p *schema.Provider
	var err error
	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p = authentikProvider.Provider(version, false)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup(rootGroup),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginSDKIncludeList(ExternalNameConfigured()),
		ujconfig.WithTerraformPluginFrameworkIncludeList([]string{}), // For future resources
		ujconfig.WithTerraformProvider(p),
		// ujconfig.WithIncludeList(ExternalNameConfigured()),
		// ujconfig.WithTerraformPluginSDKIncludeList([]string{}),
		// ujconfig.WithTerraformPluginFrameworkIncludeList([]string{}), // For future resources
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		base.Configure,
		directory.Configure,
		policy.Configure,
		propertymapping.Configure,
		provider.Configure,
		stage.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}
