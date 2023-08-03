/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	dashboard "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/dashboard/dashboard"
	dashboardjson "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/dashboard/dashboardjson"
	role "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/datadog/role"
	downtime "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/downtime/downtime"
	json "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/monitor/json"
	monitor "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/monitor/monitor"
	servicelevelobjective "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/monitor/servicelevelobjective"
	providerconfig "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/providerconfig"
	test "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/synthetics/test"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dashboard.Setup,
		dashboardjson.Setup,
		role.Setup,
		downtime.Setup,
		json.Setup,
		monitor.Setup,
		servicelevelobjective.Setup,
		providerconfig.Setup,
		test.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
