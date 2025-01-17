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
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func checkPushgatewayMetrics(t *testing.T, expectedMetrics map[string]ExpectedMetric, port int) {
	t.Helper()
	response, err := getRequest(t, fmt.Sprintf("http://localhost:%d/metrics", port), 1*time.Second)
	require.NoError(t, err, "request to pushgateway should succeed")
	require.Equal(t, http.StatusOK, response.StatusCode, "pushgateway should be up, but it is not")

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err, "should read response body")
	defer response.Body.Close()

	reader := strings.NewReader(string(body))
	metrics, err := parseMetrics(reader)
	require.NoError(t, err, "metrics should be parsed successfully")

	validateMetrics(t, metrics, expectedMetrics)
}
