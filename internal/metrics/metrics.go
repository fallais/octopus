package metrics

import "github.com/prometheus/client_golang/prometheus"

// Metrics
var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "octopus",
			Subsystem: "octopus",
			Name:      "request_count",
			Help:      "Count of the requests.",
		},
		[]string{
			"method",
		},
	)
)

func init() {
	prometheus.MustRegister(requestCount)
}

// IncrementRequestCount increment the count
func IncrementRequestCount(method string) {
	requestCount.WithLabelValues(method).Inc()
}
