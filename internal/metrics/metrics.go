package metrics

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type Metrics struct {
	pushGatewayURL     string
	pusher             *push.Pusher
	triggerCounter     prometheus.Counter
	pressButtonCounter *prometheus.CounterVec
}

func NewMetrics(pushGatewayURL string) (*Metrics, error) {
	metrics := &Metrics{
		pushGatewayURL: pushGatewayURL,
		pusher:         push.New(pushGatewayURL, "eigen_minter"),
		triggerCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "eigen_minter_trigger_total",
			Help: "Total number of times the application was triggered",
		}),
		pressButtonCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "eigen_minter_press_button_total",
				Help: "Total number of pressButton calls",
			},
			[]string{"status"},
		),
	}
	if err := metrics.Init(); err != nil {
		return nil, err
	}
	return metrics, nil
}

func (m *Metrics) Init() error {
	resp, err := http.Get(m.pushGatewayURL + "/metrics")
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch metrics from pushgateway: %s", resp.Status)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	bodyString := string(body)
	initializers := []struct {
		regex   *regexp.Regexp
		counter prometheus.Counter
	}{
		{
			regex:   regexp.MustCompile(`eigen_minter_press_button_total{.*\bstatus="failure"} (\d+)`),
			counter: m.pressButtonCounter.WithLabelValues("failure"),
		},
		{
			regex:   regexp.MustCompile(`eigen_minter_press_button_total{.*\bstatus="success"} (\d+)`),
			counter: m.pressButtonCounter.WithLabelValues("success"),
		},
		{
			regex:   regexp.MustCompile(`eigen_minter_trigger_total{.*} (\d+)`),
			counter: m.triggerCounter,
		},
	}
	for _, initializer := range initializers {
		matches := initializer.regex.FindStringSubmatch(bodyString)
		if len(matches) > 0 {
			value, err := strconv.ParseFloat(matches[1], 64)
			if err != nil {
				return err
			}
			initializer.counter.Add(value)
		}
	}

	return nil
}

func (m *Metrics) RecordTrigger() {
	m.triggerCounter.Inc()
}

func (m *Metrics) RecordPressButtonSuccess() {
	m.pressButtonCounter.WithLabelValues("success").Inc()
}

func (m *Metrics) RecordPressButtonFailure() {
	m.pressButtonCounter.WithLabelValues("failure").Inc()
}

func (m *Metrics) Push() error {
	m.pusher.Collector(m.triggerCounter)
	m.pusher.Collector(m.pressButtonCounter)
	return m.pusher.Push()
}
