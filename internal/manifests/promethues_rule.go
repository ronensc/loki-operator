package manifests

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// BuildPrometheusRule returns a list of k8s objects for Loki PrometheusRule
func BuildPrometheusRule(opts Options) ([]client.Object, error) {
	prometheusRule := NewPrometheusRule()

	return []client.Object{
		prometheusRule,
	}, nil
}

// NewPrometheusRule creates a prometheus rule
func NewPrometheusRule() *monitoringv1.PrometheusRule {
	return &monitoringv1.PrometheusRule{
		TypeMeta: metav1.TypeMeta{
			Kind:       monitoringv1.PrometheusRuleKind,
			APIVersion: monitoringv1.SchemeGroupVersion.String(),
		},

		ObjectMeta: metav1.ObjectMeta{
			Name: "myrule2",
			// Labels:    labels,
		},
		Spec: monitoringv1.PrometheusRuleSpec{
			Groups: []monitoringv1.RuleGroup{},
		},
	}
}
