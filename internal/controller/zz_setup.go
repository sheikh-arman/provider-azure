// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	providerregistration "github.com/sheikh-arman/provider-azure/internal/controller/azure/providerregistration"
	resourcegroup "github.com/sheikh-arman/provider-azure/internal/controller/azure/resourcegroup"
	subscription "github.com/sheikh-arman/provider-azure/internal/controller/azure/subscription"
	rediscache "github.com/sheikh-arman/provider-azure/internal/controller/cache/rediscache"
	providerconfig "github.com/sheikh-arman/provider-azure/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerregistration.Setup,
		resourcegroup.Setup,
		subscription.Setup,
		rediscache.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
