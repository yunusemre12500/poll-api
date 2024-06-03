package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewMetricsMiddleware(handler http.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer ctx.Next()

		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
