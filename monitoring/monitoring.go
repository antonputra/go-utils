package monitoring

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics holds a set of Prometheus metrics for the application.
type Metrics struct {
	Gauge   prometheus.GaugeVec
	Hist    *prometheus.HistogramVec
	Counter prometheus.CounterVec
}

// NewMetrics initializes and registers a new set of Prometheus metrics.
func NewMetrics(appName string, gaugeLabels []string, counterLabels []string, histLabels []string, reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		Gauge: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: appName,
			Name:      "info",
			Help:      "Metadata about the app.",
		}, gaugeLabels),
		Counter: *prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: appName,
			Name:      "count",
			Help:      "Number of operations.",
		}, counterLabels),
		// Native Prometheus Histograms
		Hist: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace:                      appName,
			Name:                           "operation_duration_seconds",
			Help:                           "Duration of the operation.",
			NativeHistogramBucketFactor:    1.02,
			NativeHistogramMaxBucketNumber: 500,
		}, histLabels),
	}
	reg.MustRegister(m.Gauge, m.Counter, m.Hist)

	return m
}

// StartPrometheus starts an HTTP server on the specified port to expose Prometheus metrics.
func StartPrometheus(port int, reg *prometheus.Registry) {
	pMux := http.NewServeMux()
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	pMux.Handle("/metrics", promHandler)

	metricsPort := fmt.Sprintf(":%d", port)
	go func() {
		slog.Info("starting the Prometheus server", "port", port)
		log.Fatal(http.ListenAndServe(metricsPort, pMux))
	}()
}
