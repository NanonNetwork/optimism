package metrics

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gogo/status"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/codes"
)

const (
	RPCServerSubsystem = "rpc_server"
	RPCClientSubsystem = "rpc_client"
	DAClientSubsystem  = "da_client"
)

type RPCMetricer interface {
	RecordRPCServerRequest(method string) func()
	RecordRPCClientRequest(method string) func(err error)
	RecordRPCClientResponse(method string, err error)
	RecordDAClientRequest(method string) func(err error)
	RecordDAClientResponse(method string, err error)
}

// RPCMetrics tracks all the RPC metrics for the op-service RPC.
type RPCMetrics struct {
	RPCServerRequestsTotal          *prometheus.CounterVec
	RPCServerRequestDurationSeconds *prometheus.HistogramVec
	RPCClientRequestsTotal          *prometheus.CounterVec
	RPCClientRequestDurationSeconds *prometheus.HistogramVec
	RPCClientResponsesTotal         *prometheus.CounterVec
	DAClientRequestsTotal           *prometheus.CounterVec
	DAClientRequestDurationSeconds  *prometheus.HistogramVec
	DAClientResponsesTotal          *prometheus.CounterVec
}

// MakeRPCMetrics creates a new RPCMetrics instance with the given process name, and
// namespace for the service.
func MakeRPCMetrics(ns string, factory Factory) RPCMetrics {
	return RPCMetrics{
		RPCServerRequestsTotal: factory.NewCounterVec(prometheus.CounterOpts{
			Namespace: ns,
			Subsystem: RPCServerSubsystem,
			Name:      "requests_total",
			Help:      "Total requests to the RPC server",
		}, []string{
			"method",
		}),
		RPCServerRequestDurationSeconds: factory.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: ns,
			Subsystem: RPCServerSubsystem,
			Name:      "request_duration_seconds",
			Buckets:   []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
			Help:      "Histogram of RPC server request durations",
		}, []string{
			"method",
		}),
		RPCClientRequestsTotal: factory.NewCounterVec(prometheus.CounterOpts{
			Namespace: ns,
			Subsystem: RPCClientSubsystem,
			Name:      "requests_total",
			Help:      "Total RPC requests initiated by the opnode's RPC client",
		}, []string{
			"method",
		}),
		RPCClientRequestDurationSeconds: factory.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: ns,
			Subsystem: RPCClientSubsystem,
			Name:      "request_duration_seconds",
			Buckets:   []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
			Help:      "Histogram of RPC client request durations",
		}, []string{
			"method",
		}),
		RPCClientResponsesTotal: factory.NewCounterVec(prometheus.CounterOpts{
			Namespace: ns,
			Subsystem: RPCClientSubsystem,
			Name:      "responses_total",
			Help:      "Total RPC request responses received by the opnode's RPC client",
		}, []string{
			"method",
			"error",
		}),
		DAClientRequestsTotal: factory.NewCounterVec(prometheus.CounterOpts{
			Namespace: ns,
			Subsystem: DAClientSubsystem,
			Name:      "requests_total",
			Help:      "Total DA requests initiated by the opnode's DA client",
		}, []string{
			"method",
		}),
		DAClientRequestDurationSeconds: factory.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: ns,
			Subsystem: DAClientSubsystem,
			Name:      "request_duration_seconds",
			Buckets:   []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
			Help:      "Histogram of DA client request durations",
		}, []string{
			"method",
		}),
		DAClientResponsesTotal: factory.NewCounterVec(prometheus.CounterOpts{
			Namespace: ns,
			Subsystem: DAClientSubsystem,
			Name:      "responses_total",
			Help:      "Total DA request responses received by the opnode's DA client",
		}, []string{
			"method",
			"error",
		}),
	}
}

// RecordRPCServerRequest is a helper method to record an incoming RPC
// call to the opnode's RPC server. It bumps the requests metric,
// and tracks how long it takes to serve a response.
func (m *RPCMetrics) RecordRPCServerRequest(method string) func() {
	m.RPCServerRequestsTotal.WithLabelValues(method).Inc()
	timer := prometheus.NewTimer(m.RPCServerRequestDurationSeconds.WithLabelValues(method))
	return func() {
		timer.ObserveDuration()
	}
}

// RecordRPCClientRequest is a helper method to record an RPC client
// request. It bumps the requests metric, tracks the response
// duration, and records the response's error code.
func (m *RPCMetrics) RecordRPCClientRequest(method string) func(err error) {
	m.RPCClientRequestsTotal.WithLabelValues(method).Inc()
	timer := prometheus.NewTimer(m.RPCClientRequestDurationSeconds.WithLabelValues(method))
	return func(err error) {
		m.RecordRPCClientResponse(method, err)
		timer.ObserveDuration()
	}
}

// RecordRPCClientResponse records an RPC response. It will
// convert the passed-in error into something metrics friendly.
// Nil errors get converted into <nil>, RPC errors are converted
// into rpc_<error code>, HTTP errors are converted into
// http_<status code>, and everything else is converted into
// <unknown>.
func (m *RPCMetrics) RecordRPCClientResponse(method string, err error) {
	var errStr string
	var rpcErr rpc.Error
	var httpErr rpc.HTTPError
	if err == nil {
		errStr = "<nil>"
	} else if errors.As(err, &rpcErr) {
		errStr = fmt.Sprintf("rpc_%d", rpcErr.ErrorCode())
	} else if errors.As(err, &httpErr) {
		errStr = fmt.Sprintf("http_%d", httpErr.StatusCode)
	} else if errors.Is(err, ethereum.NotFound) {
		errStr = "<not found>"
	} else {
		errStr = "<unknown>"
	}
	m.RPCClientResponsesTotal.WithLabelValues(method, errStr).Inc()
}

// RecordDAClientRequest is a helper method to record an DA client
// request. It bumps the requests metric, tracks the response
// duration, and records the response's error code.
func (m *RPCMetrics) RecordDAClientRequest(method string) func(err error) {
	m.DAClientRequestsTotal.WithLabelValues(method).Inc()
	timer := prometheus.NewTimer(m.DAClientRequestDurationSeconds.WithLabelValues(method))
	return func(err error) {
		m.RecordDAClientResponse(method, err)
		timer.ObserveDuration()
	}
}

// RecordDAClientResponse records an DA response, converting errors into metrics-friendly format.
// Nil errors are converted into <nil>, DA errors into grpc_<error code>, and everything else into <unknown>.
func (m *RPCMetrics) RecordDAClientResponse(method string, err error) {
	var errStr string
	// Handle DA client errors using gogo/status package
	if err == nil {
		errStr = "<nil>"
	} else {
		// Use gogo/status to get error status for DA errors
		st, ok := status.FromError(err)
		if ok {
			// Convert DA error status into a formatted string
			errStr = fmt.Sprintf("grpc_%s", codes.Code(st.Code()))
		} else {
			errStr = "<unknown>"
		}
	}

	// Increment the metric for DA client responses
	m.DAClientResponsesTotal.WithLabelValues(method, errStr).Inc()
}

type NoopRPCMetrics struct{}

func (n *NoopRPCMetrics) RecordRPCServerRequest(method string) func() {
	return func() {}
}

func (n *NoopRPCMetrics) RecordRPCClientRequest(method string) func(err error) {
	return func(err error) {}
}

func (n *NoopRPCMetrics) RecordRPCClientResponse(method string, err error) {
}

func (n *NoopRPCMetrics) RecordDAClientRequest(method string) func(err error) {
	return func(err error) {}
}

func (n *NoopRPCMetrics) RecordDAClientResponse(method string, err error) {
}

var _ RPCMetricer = (*NoopRPCMetrics)(nil)
