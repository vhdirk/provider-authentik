package base

import "github.com/crossplane/upjet/pkg/config"

// FlowRef references a flow by uuid
var FlowRef = config.Reference{
	TerraformName: "authentik_flow",
	Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("uuid",true)`,
}

// UserRef references a user by id
var UserRef = config.Reference{
	TerraformName: "authentik_user",
	Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
}

// CertificateKeyPairRef references a signing key by id
var CertificateKeyPairRef = config.Reference{
	TerraformName: "authentik_certificate_key_pair",
	Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
}

// PropertyMappingRACRef references a signing key by id
var PropertyMappingRACRef = config.Reference{
	TerraformName: "authentik_property_mapping_rac",
	Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
}

// PropertyMappingProviderScopeRef references a signing key by id
var PropertyMappingProviderScopeRef = config.Reference{
	TerraformName: "authentik_property_mapping_provider_scope",
	Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
}

// PropertyMappingProviderRACRef references a signing key by id
var PropertyMappingProviderRACRef = config.Reference{
	TerraformName: "authentik_property_mapping_provider_rac",
	Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
}
