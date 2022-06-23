package controllers

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	rsctv1alpha1 "github.com/mjturek/rsct-operator/api/v1alpha1"
)

// updateRSCTStatus will do something magical one day.
func (r *RSCTReconciler) updateRSCTStatus(ctx context.Context, rsct *rsctv1alpha1.RSCT, currentDaemonSet *appsv1.DaemonSet) error {
	return nil
}

func (r *RSCTReconciler) getRSCTDaemonSetPodList(ctx context.Context, rsct *rsctv1alpha1.RSCT) (*corev1.PodList, error){
        // List the pods that are part of the daemonset
        podList := &corev1.PodList{}
        listOpts := []client.ListOption{
                client.InNamespace(r.Config.Namespace),
                client.MatchingLabels(labelsForMemcached(r.Config.Name)),
        }
        if err = r.Client.List(ctx, podList, listOpts...); err != nil {
                return ctrl.Result{}, err
        }
	return podList
}
