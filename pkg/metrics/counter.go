package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type Counter interface {
	IncrementVector(values ...string)
}

type counter struct {
	vector *prometheus.CounterVec
}

func MustRegisterCounter(name, namespace, subsystem string, labels []string) Counter {
	vector := counterVector(name, namespace, subsystem, labels)
	prometheus.MustRegister(vector)
	return &counter{vector: vector}
}

func RegisterCounter(name, namespace, subsystem string, labels []string) (Counter, error) {
	vector := counterVector(name, namespace, subsystem, labels)
	if err := prometheus.Register(vector); err != nil {
		return nil, fmt.Errorf("error while registering counter vector: %v", err)
	}
	return &counter{vector: vector}, nil
}

func counterVector(name, namespace, subsystem string, labels []string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Help:      fmt.Sprintf("counter vector for %s", name),
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      name,
	}, labels)
}

func (c *counter) IncrementVector(values ...string) {
	c.vector.WithLabelValues(values...).Inc()
}
