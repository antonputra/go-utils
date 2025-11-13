package main

import (
	"fmt"
	"time"

	mon "github.com/antonputra/go-utils/monitoring"
	"github.com/antonputra/go-utils/util"
	"github.com/prometheus/client_golang/prometheus"
)

func work(m *mon.Metrics) {
	now := time.Now()
	fmt.Println("doing work...")
	time.Sleep(1 * time.Millisecond)
	m.Hist.WithLabelValues("redis").Observe(time.Since(now).Seconds())
}

func main() {
	reg := prometheus.NewRegistry()

	port := 8082
	appVersion := "v0.1.0"

	gaugeLabels := []string{"version"}
	counterLabels := []string{}
	histLabels := []string{"target"}

	m := mon.NewMetrics("client", gaugeLabels, counterLabels, histLabels, reg)
	mon.StartPrometheus(port, reg)
	m.Gauge.WithLabelValues(appVersion).Set(1)

	util.DoWork(work, 5, 20, m)
}
