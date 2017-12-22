package replicator

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	copiedCtr = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "stream_replicator_copied_msgs",
		Help: "How many messages were copied",
	}, []string{"name", "worker"})

	failedCtr = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "stream_replicator_failed_msgs",
		Help: "How many messages failed to copy to the remote server",
	}, []string{"name", "worker"})

	ackFailedCtr = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "stream_replicator_acks_failed",
		Help: "How many times ack'ing a message failed",
	}, []string{"name", "worker"})

	processTime = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "stream_replicator_processing_time",
		Help: "How long it took to process messages",
	}, []string{"name", "worker"})

	reconnectCtr = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "stream_replicator_connection_reconnections",
		Help: "Number of times the connector reconnected to the middleware",
	}, []string{"name", "worker"})

	closedCtr = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "stream_replicator_connection_closed",
		Help: "Number of times the connection was closed",
	}, []string{"name", "worker"})

	errorCtr = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "stream_replicator_connection_errors",
		Help: "Number of times the connection encountered an error",
	}, []string{"name", "worker"})
)

func init() {
	prometheus.MustRegister(copiedCtr)
	prometheus.MustRegister(failedCtr)
	prometheus.MustRegister(ackFailedCtr)
	prometheus.MustRegister(processTime)
	prometheus.MustRegister(reconnectCtr)
	prometheus.MustRegister(closedCtr)
	prometheus.MustRegister(errorCtr)
}
