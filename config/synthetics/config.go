package synthetics

import (
	tjconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the monitor group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("datadog_synthetics_test", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
	})
}
