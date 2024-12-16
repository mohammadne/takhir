package common

import (
	"fmt"

	"github.com/mohammadne/takhir/internal"
	"github.com/prometheus/client_golang/prometheus"
)

func RegisterMetric(name string, labels []string) *prometheus.CounterVec {
	vector := prometheus.NewCounterVec(prometheus.CounterOpts{
		Help:      fmt.Sprintf("counter vector for %s", name),
		Namespace: internal.Namespace,
		Subsystem: internal.System,
		Name:      name,
	}, labels)
	prometheus.MustRegister(vector)
	return vector
}
