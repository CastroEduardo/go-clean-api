package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/pkg/metrics"
)

func Prometheus() gin.HandlerFunc {

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := float64(time.Since(start).Milliseconds()) // mantiene milisegundos
		metrics.HttpDuration.WithLabelValues(c.FullPath(), c.Request.Method, strconv.Itoa(c.Writer.Status())).
			Observe(duration)
	}

	// return func(c *gin.Context) {
	// 	start := time.Now()
	// 	path := c.FullPath()
	// 	method := c.Request.Method
	// 	c.Next()
	// 	status := c.Writer.Status()
	// 	metrics.HttpDuration.WithLabelValues(path, method, strconv.Itoa(status)).
	// 		Observe(float64(time.Since(start) / time.Millisecond))

	// }
}

// func PrometheusMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		start := time.Now()
// 		c.Next()
// 		duration := float64(time.Since(start).Milliseconds()) // mantiene milisegundos
// 		HttpDuration.WithLabelValues(c.FullPath(), c.Request.Method, strconv.Itoa(c.Writer.Status())).
// 			Observe(duration)
// 	}
// }
