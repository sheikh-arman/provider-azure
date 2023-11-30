/*
Copyright 2022 Upbound Inc.
*/

package cache

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures redis group
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("azurerm_redis_cache", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "redis_version")
	})
}
