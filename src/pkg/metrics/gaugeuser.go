package metrics

import "github.com/prometheus/client_golang/prometheus"

var ActiveUsers = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "active_users",
		Help: "Current number of active users",
	},
)
