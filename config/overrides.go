package config

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

var (
	resourceGroup = map[string]string{

		"azurerm_redis_cache": "cache",
	}
	resourceKind = map[string]string{

		"azurerm_redis_cache": "RedisCache",
	}
)

// default api-group & kind configuration for all resources
func groupKindOverride(r *ujconfig.Resource) {
	if _, ok := resourceGroup[r.Name]; ok {
		r.ShortGroup = resourceGroup[r.Name]
	}

	if _, ok := resourceKind[r.Name]; ok {
		r.Kind = resourceKind[r.Name]
	}
}
