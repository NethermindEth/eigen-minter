package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type Metrics struct {
	pushGatewayURL     string
	triggerCounter     prometheus.Counter
	pressButtonCounter *prometheus.CounterVec
}

func NewMetrics(pushGatewayURL string) *Metrics {
	return &Metrics{
		pushGatewayURL: pushGatewayURL,
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
	pusher := push.New(m.pushGatewayURL, "eigen_minter")
	pusher.Collector(m.triggerCounter)
	pusher.Collector(m.pressButtonCounter)
	return pusher.Push()
}
