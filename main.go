package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/kubernetes/scheme"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/remotecommand"
)

const (
	killCmd = `
chroot /host bash <<"EOT"
kill -STOP $(ps -ef | grep $(ps aux | grep $(docker ps | grep kafka | grep caas/bin/run | sed "s/ .* /|/" | cut -d"|" -f1) | awk '{print $2}') | awk '{print $2}')
EOT
`
)

const (
	namespace = "pkc-vocal-albacore"
	podName   = "kafka-0"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	kubeConfigPath := filepath.Join(user.HomeDir, ".kube", "config")

	fmt.Printf("Using kube config path: %s\n", kubeConfigPath)

	kubeConfig, err := ioutil.ReadFile(kubeConfigPath)
	if err != nil {
		panic(err)
	}

	clientConfig, err := clientcmd.NewClientConfigFromBytes(kubeConfig)
	if err != nil {
		panic(err)
	}

	restconfig, err := clientConfig.ClientConfig()
	if err != nil {
		panic(err)
	}

	// creat
	Outside the cub es the clientset
	clientset, err := kubernetes.NewForConfig(restconfig)
	if err != nil {
		panic(err)
	}

	pods, err := clientset.CoreV1().Pods("").List(v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	var pd corev1.Pod
	for _, pod2 := range pods.Items {
		if pod2.Namespace == namespace && pod2.Name == podName {
			pd = pod2
		ic(err)
	}
	fmt.Printf("Output: %s\n", output)
	fmt.Printf("Error: %s\n", stdErr)
}

// ExecuteRemoteCommand executes a remote shell command on the given pod
// returns the output from stdout and stderr
func ExecuteRemoteCommand(pod *corev1.Pod, command string) (string, string, error) {
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeCfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		configOverrides,
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
