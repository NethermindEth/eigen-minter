/*
Copyright 2022 Nethermind

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
package e2e

import (
	"io"
	"strings"
	"testing"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"gotest.tools/v3/assert"
)

type ExpectedMetric struct {
	Name   string
	Type   string
	Help   string
	Labels []string
}

var expectedMetrics = map[string]ExpectedMetric{
	"eigen_minter_trigger_total": {
		Name: "eigen_minter_trigger_total",
		Type: "counter",
		Help: "Total number of times the application was triggered",
	},
	"eigen_minter_press_button_total": {
		Name:   "eigen_minter_press_button_total",
		Type:   "counter",
		Help:   "Total number of pressButton calls",
		Labels: []string{"status"},
	},
}

func parseMetrics(reader io.Reader) (map[string]*dto.MetricFamily, error) {
	parser := expfmt.TextParser{}
	return parser.TextToMetricFamilies(reader)
}

func validateMetrics(t *testing.T, parsedMetrics map[string]*dto.MetricFamily, expectedMetrics map[string]ExpectedMetric) {
	t.Helper()
	t.Logf("Validating metrics")
	foundMetrics := make(map[string]bool)
	for name, expected := range expectedMetrics {
		metricFamily, found := parsedMetrics[name]
		if !found {
			t.Logf("Metric %s not found in the response\n", name)
			continue
		}
		foundMetrics[name] = true

		// Check the metric type
		assert.Equal(t, metricFamily.GetType().String(), strings.ToUpper(expected.Type), "Metric %s has incorrect type. Expected: %s, Got: %s\n", name, strings.ToUpper(expected.Type), metricFamily.GetType())

		// Check the help string
		assert.Equal(t, metricFamily.GetHelp(), expected.Help, "Metric %s has incorrect help string. Expected: %s, Got: %s\n", name, expected.Help, metricFamily.GetHelp())

		// Check labels
		for _, metric := range metricFamily.Metric {
			labelNames := make(map[string]bool)
			for _, label := range metric.Label {
				labelNames[label.GetName()] = true
			}
			for _, expectedLabel := range expected.Labels {
				assert.Equal(t, labelNames[expectedLabel], true, "Metric %s is missing expected label: %s\n", name, expectedLabel)
			}
		}
	}
}
