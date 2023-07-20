/*
Copyright 2023 YashPimple.

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

package controllers

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	apiv4alpha1 "github.com/YashPimple/Operator-sdk/api/v4alpha1"
)

// ScalarReconciler reconciles a Scalar object
type ScalarReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=api.scaleroperator.io,resources=scalars,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=api.scaleroperator.io,resources=scalars/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=api.scaleroperator.io,resources=scalars/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Scalar object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ScalarReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	//logicx
	log.Log.Info("Reconcile called")
	scalar := &apiv4alpha1.Scalar{}
	err := r.Get(ctx, req.NamespacedName, scalar)
	if err != nil {
		return ctrl.Result{}, err
	}

	startTime := scalar.Spec.Start
	endTime := scalar.Spec.End
	replicas := scalar.Spec.Replicas

	currentHour := time.Now().UTC().Hour()

	if currentHour >= startTime && currentHour <= endTime {
		if err := scaleDeployment(scalar, r, ctx, replicas); err != nil {
			return ctrl.Result{}, err
		}

		return ctrl.Result{}, nil
	}

	return ctrl.Result{RequeueAfter: time.Duration(30 * time.Second)}, nil
}

func scaleDeployment(scalar *apiv4alpha1.Scalar, r *ScalarReconciler, ctx context.Context, replicas int32) error {
	for _, deploy := range scalar.Spec.Deployments {
		deployment := &v1.Deployment{}
		err := r.Get(ctx, types.NamespacedName{
			Namespace: deploy.Namespace,
			Name:      deploy.Name,
		},
			deployment,
		)
		if err != nil {
			return err
		}

		if deployment.Spec.Replicas != &replicas {
			deployment.Spec.Replicas = &replicas
			err := r.Update(ctx, deployment)
			if err != nil {
				scalar.Status.Status = apiv4alpha1.FAILED
				return err
			}
			scalar.Status.Status = apiv4alpha1.SUCCESS
			err = r.Status().Update(ctx, scalar)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ScalarReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv4alpha1.Scalar{}).
		Complete(r)
}
