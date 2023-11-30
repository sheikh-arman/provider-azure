/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/sheikh-arman/provider-azure/config/base"
	"github.com/sheikh-arman/provider-azure/config/cache"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "azure"
	modulePath     = "github.com/sheikh-arman/provider-azure"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("azure.kubedb.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	// API group overrides from Terraform import statements
	for _, r := range pc.Resources {
		groupKindOverride(r)
	}
	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		base.Configure,
		cache.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
