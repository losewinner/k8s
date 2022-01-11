// This is a generated file. Do not edit directly.

module k8s.io/sample-apiserver

go 1.16

require (
	github.com/google/gofuzz v1.1.0
	github.com/spf13/cobra v1.2.1
	k8s.io/apimachinery v0.0.0
	k8s.io/apiserver v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/code-generator v0.0.0
	k8s.io/component-base v0.0.0
	k8s.io/kube-openapi v0.0.0-restframework-35b5ddc
	k8s.io/utils v0.0.0-20211208161948-7d6a63dca704
)

replace (
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/apiserver => ../apiserver
	k8s.io/client-go => ../client-go
	k8s.io/code-generator => ../code-generator
	k8s.io/component-base => ../component-base
	k8s.io/kube-openapi => github.com/austince/kube-openapi v0.0.0-restframework-35b5ddc
	k8s.io/sample-apiserver => ../sample-apiserver
)
