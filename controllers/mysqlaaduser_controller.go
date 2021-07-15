// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	azurev1alpha2 "github.com/Azure/azure-service-operator/api/v1alpha2"
)

// MySQLAADUserReconciler reconciles a MySQLAADUser object
type MySQLAADUserReconciler struct {
	Reconciler *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=mysqlaadusers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources={mysqlaadusers/status,mysqlaadusers/finalizers},verbs=get;update;patch

func (r *MySQLAADUserReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req, &azurev1alpha2.MySQLAADUser{})
}

// SetupWithManager runs reconcile loop with manager
func (r *MySQLAADUserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha2.MySQLAADUser{}).
		Complete(r)
}
