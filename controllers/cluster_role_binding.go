/*
Copyright 2021.

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
	"fmt"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	operatorv1alpha1 "github.com/mjturek/rsct-operator/api/v1alpha1"
)

// ensureRSCTClusterRoleBinding ensures that the privleged cluster role binding exists.
func (r *RSCTReconciler) ensureRSCTClusterRoleBinding(ctx context.Context, rsct *operatorv1alpha1.RSCT, sa *corev1.ServiceAccount) (bool, error) {
	nsName := types.NamespacedName{Namespace: rsct.Namespace, Name: rsct.Name}

	desired := desiredRSCTClusterRoleBinding(nsName, sa)

	if err := controllerutil.SetControllerReference(rsct, desired, r.Scheme); err != nil {
		return false, fmt.Errorf("failed to set the controller reference for cluster role binding: %w", err)
	}

	exist, err := r.existingRSCTClusterRoleBinding(ctx, nsName)
	if err != nil {
		return false, err
	}

	if !exist {
		if err := r.createRSCTClusterRoleBinding(ctx, desired); err != nil {
			return false, err
		}
	}

	return true, nil
}

// existingRSCTClusterRoleBinding checks if the RSCT cluster role binding exists.
func (r *RSCTReconciler) existingRSCTClusterRoleBinding(ctx context.Context, nsName types.NamespacedName) (bool, error) {
	crb := &rbacv1.ClusterRoleBinding{}
	if err := r.Client.Get(ctx, nsName, crb); err != nil {
		if errors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// desiredRSCTClusterRoleBinding returns the desired serivce account resource.
func desiredRSCTClusterRoleBinding(nsName types.NamespacedName, sa *corev1.ServiceAccount) *rbacv1.ClusterRoleBinding {
	// TODO(mjturek): Adjust RoleRef name for non-openshift cases
	return &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: nsName.Namespace,
			Name:      nsName.Name,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "system:openshift:scc:privileged",
		},
		Subjects: []rbacv1.Subject{
			rbacv1.Subject{
				Kind:      "ServiceAccount",
				Name:      sa.Name,
				Namespace: sa.Namespace,
			},
		},
	}
}

// createRSCTServiceAccount creates the given service account using the reconciler's client.
func (r *RSCTReconciler) createRSCTClusterRoleBinding(ctx context.Context, crb *rbacv1.ClusterRoleBinding) error {
	if err := r.Client.Create(ctx, crb); err != nil {
		return fmt.Errorf("failed to create RSCT cluster role binding %s/%s: %w", crb.Namespace, crb.Name, err)
	}

	return nil
}
