/*
Copyright 2021 The Kubernetes Authors.

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

package metrics

import (
	"k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"
	"k8s.io/pod-security-admission/api"
)

const (
	namespace = "pod_security"
)

var SecurityEvaluation = metrics.NewCounterVec(
	&metrics.CounterOpts{
		Name:           "pod_security_evaluations_total",
		Help:           "Counter of pod security evaluations.",
		StabilityLevel: metrics.ALPHA,
	},
	[]string{"decision", "policy_level", "policy_version", "mode", "operation", "resource", "subresource"},
)

type Decision string
type Mode string

const (
	ModeAudit = "audit"
	ModeEnforce = "enforce"
	ModeWarn = "warn"
	DecisionAllow = "allow" // Policy evaluated, request allowed
	DecisionDeny = "deny" // Policy evaluated, request denied
	DecisionIgnore = "ignore" // Request not relevant, policy not evaluated
	DecisionExempt = "exempt" // Request exempt, policy not evaluated
	DecisionError = "error" // Error preventing evaluation, policy not evaluated
)

type EvaluationRecorder interface {
	// TODO: fill in args required to record https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/2579-psp-replacemenonitoring
	RecordEvaluation(decision Decision, policy api.LevelVersion, evalMode Mode, attrs api.Attributes)
}

type PrometheusRecorder struct {
}

func NewPrometheusRecorder() *PrometheusRecorder {
	legacyregistry.MustRegister(SecurityEvaluation)
	return &PrometheusRecorder{}
}

func (r PrometheusRecorder) RecordEvaluation(decision Decision, policy api.LevelVersion, evalMode Mode, attrs api.Attributes) {
	dec := string(decision)
	operation := string(attrs.GetOperation())
	resource := attrs.GetResource().String()
	subresource := attrs.GetSubresource()
	var version string
	if policy.Valid() {
		v, err := api.ParseVersion(policy.Version.String())
		if err != nil || v.Latest(){
			version = "latest"
		}else{
			if v.Older(){
				// TODO
			}
		}
		SecurityEvaluation.WithLabelValues(dec, string(policy.Level),
			version, string(evalMode), operation, resource, subresource).Inc()
	}

}
