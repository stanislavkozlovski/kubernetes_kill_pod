package main

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	corev1 "k8s.io/api/core/v1"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"k8s.io/client-go/kubernetes/scheme"

	"k8s.io/client-go/tools/remotecommand"
)

func main() {
	fmt.Println("a")

	config, err := clientcmd.BuildConfigFromFlags("localhost:8080", "/Users/stanislav/.kube/config")
	if err != nil {
		panic(err)
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	pods, err := clientset.CoreV1().Pods("").List(v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	var pd corev1.Pod
	for _, pod2 := range pods.Items {
		fmt.Println(pod2.Namespace)
		fmt.Println(pod2.Name)
		if pod2.Namespace == "pkc-busy-seagull" && pod2.Name == "kafka-1" {
			pd = pod2
		}

	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	fmt.Printf("executing remote command on %s/%s\n", pd.Namespace, pd.Name)
	output, stdErr, err := ExecuteRemoteCommand(&pd, "ls -l")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output %s stdErr %s\n", output, stdErr)
}

// ExecuteRemoteCommand executes a remote shell command on the given pod
// returns the output from stdout and stderr
func ExecuteRemoteCommand(pod *corev1.Pod, command string) (string, string, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	loadingRules.ExplicitPath = "/Users/stanislav/.kube/config"
	configOverrides := &clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: "localhost:8080"}}
	kubeCfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		configOverrides ,
	)
	restCfg, err := kubeCfg.ClientConfig()
	if err != nil {
		return "", "", err
	}
	coreClient, err := corev1client.NewForConfig(restCfg)
	if err != nil {
		return "", "", err
	}

	outBuf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}
	request := coreClient.RESTClient().
		Post().
		Namespace(pod.Namespace).
		Resource("pods").
		Name(pod.Name).
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Command: []string{"/bin/sh", "-c", command},
			Stdin:   false,
			Stdout:  true,
			Stderr:  true,
		}, scheme.ParameterCodec)
	exec, err := remotecommand.NewSPDYExecutor(restCfg, "POST", request.URL())
	err = exec.Stream(remotecommand.StreamOptions{
		Stdout: outBuf,
		Stderr: errBuf,
	})
	if err != nil {
		return "", "", errors.Wrapf(err, "Failed executing command %s on %v/%v", command, pod.Namespace, pod.Name)
	}

	return outBuf.String(), errBuf.String(), nil
}
