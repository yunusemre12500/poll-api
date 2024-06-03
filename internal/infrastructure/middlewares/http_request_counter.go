package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yunusemre12500/poll-api/internal/infrastructure/metrics"
)

func NewHTTPRequestCounterMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer ctx.Next()

		metrics.HTTPRequestsTotal.With(prometheus.Labels{
			"method": ctx.Request.Method,
			"path":   ctx.Request.URL.Path,
			"status": fmt.Sprintf("%d", ctx.Writer.Status()),
		}).Inc()

	}
}
