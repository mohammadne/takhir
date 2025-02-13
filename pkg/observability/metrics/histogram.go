package metrics

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Histogram interface {
	ObserveResponseTime(duration time.Time, values ...string)
}

type histogram struct {
	vector *prometheus.HistogramVec
}

func MustRegisterHistogram(name, namespace, subsystem string, labels []string) Histogram {
	vector := histogramVector(name, namespace, subsystem, labels)
	prometheus.MustRegister(vector)
	return &histogram{vector: vector}
}

func RegisterHistogram(name, namespace, subsystem string, labels []string) (Histogram, error) {
	vector := histogramVector(name, namespace, subsystem, labels)
	if err := prometheus.Register(vector); err != nil {
		return nil, fmt.Errorf("error while registering histogram vector: %v", err)
	}
	return &histogram{vector: vector}, nil
}

func histogramVector(name, namespace, subsystem string, labels []string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Help:      fmt.Sprintf("histogram vector for %s", name),
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      name,
		Buckets:   []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5},
	}, labels)
}

func (c *histogram) ObserveResponseTime(start time.Time, values ...string) {
	c.vector.WithLabelValues(values...).Observe(time.Since(start).Seconds())
}
