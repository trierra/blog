package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	//HealthRegistry is a registry for metrics, showing the system health (errors, failures, etc)
	HealthRegistry *prometheus.Registry

	// ErrorsMetric metric shows total errors count
	// MUST contain a label with a "error" key
	ErrorsMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "errors_total",
			Help: "Total amount of errors",
		},
		[]string{"error"},
	)

	// ActionDeclinedMetric is a metric that counts declined actions.
	// MUST contain a labels with a "action" and "item" keys
	ActionDeclinedMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "action_declined",
			Help: "",
		},
		[]string{"action", "item"},
	)

	// ActionTakenMetric is a metric which counts actions that has been taken.
	// MUST contain a labels with a "action" and "item" keys
	ActionTakenMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "action_taken",
			Help: "",
		},
		[]string{"action", "item"},
	)
)

func init() {
	//create a registry
	HealthRegistry = prometheus.NewRegistry()

	//register health-related metrics into registry

	HealthRegistry.MustRegister(ErrorsMetric)
	HealthRegistry.MustRegister(ActionDeclinedMetric)
	HealthRegistry.MustRegister(ActionTakenMetric)
}

// ErrorInc calls prometheus counter which counts amount of errors,
// occurred on error, fatal and panic log levels
func ErrorInc() {
	ErrorsMetric.With(prometheus.Labels{"error": "error"}).Inc()
}

// ActionTakenMetricInc increases a number of taken actions on a particular volume
func ActionTakenMetricInc(itemName string, actionName string) {
	ActionTakenMetric.With(prometheus.Labels{"item": itemName, "action": actionName}).Inc()
}

// ActionDeclinedMetricInc  increases a number of declined actions on a particular volume
func ActionDeclinedMetricInc(itemName string, actionName string) {
	ActionDeclinedMetric.With(prometheus.Labels{"item": itemName, "action": actionName}).Inc()
}
