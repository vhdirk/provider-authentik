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

// SigningKeyRef references a signing key by id
var CertificateKeyPairRef = config.Reference{
	TerraformName: "authentik_certificate_key_pair",
	Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
}
