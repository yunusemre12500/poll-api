package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

func CreateRegistry() *prometheus.Registry {
	registry := prometheus.NewRegistry()

	registry.MustRegister(
		HTTPRequestsTotal,
		HTTPResponseLatencyInMs,
	)

	return registry
}

var HTTPRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Total number of requests.",
}, []string{"method", "path", "status"})

var HTTPResponseLatencyInMs = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_response_latency_ms",
	Help:    "Response latency in milliseconds.",
	Buckets: []float64{0.1, 0.2, 0.25, 0.3, .35, 0.4, 0.45, 0.5},
}, []string{"method", "path"})
