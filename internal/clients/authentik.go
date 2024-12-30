/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/pkg/errors"
	"github.com/vhdirk/crossplane-provider-authentik/apis/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal authentik credentials as JSON"
	errExtractSecretKey     = "cannot extract from secret key when none specified"
	errGetCredentialsSecret = "cannot get credentials secret"
)

var requiredAuthentikConfigKeys = []string{
	"url",
	"token",
	"insecure",
}
var optionalAuthentikConfigKeys = []string{}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		creds, err := ExtractCredentials(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}

		// set provider configuration
		ps.Configuration = map[string]any{}
		// Iterate over the requiredAuthentikConfigKeys, they must be set
		for _, key := range requiredAuthentikConfigKeys {
			if value, ok := creds[key]; ok {
				if !ok {
					// Return an error if a required key is missing
					return ps, errors.Errorf("required Authentik configuration key '%s' is missing", key)
				}
				ps.Configuration[key] = value
			}
		}

		// Iterate over the optionalAuthentikConfigKeys, they can be set and do not have to be in the creds map
		for _, key := range optionalAuthentikConfigKeys {
			if value, ok := creds[key]; ok {
				ps.Configuration[key] = value
			}
		}

		return ps, nil
	}
}

// ExtractCredentials Function that extracts credentials from the secret provided to providerconfig
func ExtractCredentials(ctx context.Context, source xpv1.CredentialsSource, client client.Client, selector xpv1.CommonCredentialSelectors) (map[string]string, error) {
	creds := make(map[string]string)

	// first try to see if the secret contains a proper key-value map
	if selector.SecretRef == nil {
		return nil, errors.New(errExtractSecretKey)
	}
	secret := &corev1.Secret{}
	if err := client.Get(ctx, types.NamespacedName{Namespace: selector.SecretRef.Namespace, Name: selector.SecretRef.Name}, secret); err != nil {
		return nil, errors.Wrap(err, errGetCredentialsSecret)
	}
	if _, ok := secret.Data[selector.SecretRef.Key]; !ok {
		for k, v := range secret.Data {
			creds[k] = string(v)
		}
		return creds, nil
	}

	// if that fails, use Crossplane's way of extracting a JSON document
	rawData, err := resource.CommonCredentialExtractor(ctx, source, client, selector)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(rawData, &creds); err != nil {
		return nil, errors.Wrap(err, errUnmarshalCredentials)
	}

	return creds, nil
}
