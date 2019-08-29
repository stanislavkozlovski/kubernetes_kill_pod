module github.com/confluentinc/cc-system-tests

require (
	github.com/pkg/errors v0.8.0
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/apiserver v0.0.0
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/helm v2.10.0-rc.2.0.20190801165946-b39cf0a5096b+incompatible
	k8s.io/kubernetes v1.11.4 // indirect
)

replace (
	// all the k8s mods are set to the 1.15.1 versions
	k8s.io/api => k8s.io/api v0.0.0-20190718183219-b59d8169aab5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190718185103-d1ef975d28ce
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190718184206-a1aa83af71a7
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190313123343-44a48934c135
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190718183610-8e956561bbf5
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190718190308-f8e43aa19282
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20190718190146-f7b0473036f9
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190612205613-18da4a14b22b
	k8s.io/component-base => k8s.io/component-base v0.0.0-20190718183727-0ececfbe9772
	k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190531030430-6117653b35f1
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20190718190424-bef8d46b95de
	k8s.io/klog => k8s.io/klog v0.1.0 // indirect
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190718184434-a064d4d1ed7a
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20190718190030-ea930fedc880
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20181114233023-0317810137be
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20190718185641-5233cb7cb41e
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20190718185913-d5429d807831
	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20190718185757-9b45f80d5747
	k8s.io/kubernetes => k8s.io/kubernetes v1.15.1
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20190718190548-039b99e58dbd
	k8s.io/metrics => k8s.io/metrics v0.0.0-20190718185242-1e1642704fe6
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20190718184639-baafa86838c0
	k8s.io/utils => k8s.io/utils v0.0.0-20190712204705-3dccf664f023
)
