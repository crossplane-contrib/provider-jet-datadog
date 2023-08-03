/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane-contrib/provider-jet-datadog/config/downtime"

	"github.com/upbound/upjet/pkg/config"
	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-jet-datadog/config/dashboard"
	"github.com/crossplane-contrib/provider-jet-datadog/config/monitor"
	"github.com/crossplane-contrib/provider-jet-datadog/config/role"
	"github.com/crossplane-contrib/provider-jet-datadog/config/synthetics"
)

const (
	resourcePrefix = "datadog"
	modulePath     = "github.com/crossplane-contrib/provider-jet-datadog"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	providerOpts := config.WithBasePackages(config.BasePackages{
		APIVersion: []string{
			// Default package for ProviderConfig APIs
			"apis/v1alpha1",
		},
		Controller: []string{
			// Default package for ProviderConfig controllers
			"internal/controller/providerconfig",
		},
	})

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata), providerOpts,
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		monitor.Configure,
		synthetics.Configure,
		downtime.Configure,
		role.Configure,
		dashboard.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
