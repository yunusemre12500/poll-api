package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yunusemre12500/poll-api/internal/infrastructure/metrics"
)

func NewHTTPResponseLatencyCounterMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		started := time.Now()

		defer ctx.Next()

		elapsed := time.Until(started)

		metrics.HTTPResponseLatencyInMs.With(prometheus.Labels{
			"method": ctx.Request.Method,
			"path":   ctx.Request.URL.Path,
		}).Observe(float64(elapsed.Milliseconds()))
	}
}
