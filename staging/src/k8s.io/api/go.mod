// This is a generated file. Do not edit directly.

module k8s.io/api

go 1.16

require (
	github.com/gogo/protobuf v1.3.2
	github.com/stretchr/testify v1.7.0
	k8s.io/apimachinery v0.0.0
)

replace (
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/kube-openapi => github.com/austince/kube-openapi v0.0.0-restframework-35b5ddc
)
