/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"datadog_dashboard":               config.IdentifierFromProvider,
	"datadog_dashboard_json":          config.IdentifierFromProvider,
	"datadog_downtime":                config.IdentifierFromProvider,
	"datadog_monitor":                 config.IdentifierFromProvider,
	"datadog_monitor_json":            config.IdentifierFromProvider,
	"datadog_service_level_objective": config.IdentifierFromProvider,
	"datadog_role":                    config.IdentifierFromProvider,
	"datadog_synthetics_test":         config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
