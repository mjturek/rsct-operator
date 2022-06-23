package controllers

import (
	"bytes"
	"context"
	"fmt"

	rsctv1alpha1 "github.com/mjturek/rsct-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
)

const rmcDomainStatusCmd = "rmcdomainstatus -s ctrmc -a IP"

// updateRSCTStatus will do something magical one day.
func (r *RSCTReconciler) updateRSCTStatus(ctx context.Context, rsct *rsctv1alpha1.RSCT, currentDaemonSet *appsv1.DaemonSet) error {
	podList := r.getRSCTDaemonSetPodList(contxt.TODO())
	cmd := strings.Split(rmcDomainStatusCmd, " ")

        for _, pod := range podList.Items() {

		 execOnPod(cmd, )
        }
	return nil
}

func getKubeConfig() clientcmd.ClientConfig {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
}

func execOnPod(cmd []string, podName types.NamespacedName, kubeClient kubernetes.Interface) (string, error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	req := kubeClient.CoreV1().RESTClient().
		Post().
		Name(podName.Name).
		Namespace(podName.Namespace).
		SubResource("exec")

	req.VersionedParams(&corev1.PodExecOptions{
		Command: cmd,
		Stdout:  true,
		Stderr:  true,
	}, scheme.ParameterCodec)

	clientConfig, err := getKubeConfig().ClientConfig()
	if err != nil {
		return "", fmt.Errorf("failed to get client config: %v", err)
	}

	exec, err := remotecommand.NewSPDYExecutor(clientConfig, "POST", req.URL())
	if err != nil {
		return "", fmt.Errorf("failed to create Executor: %v", err)
	}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdout: &stdout,
		Stderr: &stderr,
		Tty:    false,
	})
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v", err)
	}
	if stderr.Len() > 0 {
		return "", fmt.Errorf("stderr: %v", stderr.String())
	}
	return stdout.String(), nil
}
