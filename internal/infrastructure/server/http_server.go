package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	pollController "github.com/yunusemre12500/poll-api/internal/infrastructure/controllers/poll/v1"
	"github.com/yunusemre12500/poll-api/internal/infrastructure/metrics"
	"github.com/yunusemre12500/poll-api/internal/infrastructure/middlewares"
)

type HTTPServer struct {
	engine *gin.Engine
}

func NewHTTPServer() *HTTPServer {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	engine.HandleMethodNotAllowed = true

	engine.SetTrustedProxies(nil)

	engine.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    "MethodNotAllowed",
			"message": "Method not allowed.",
		})
	})

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    "NotFound",
			"message": "Not found.",
		})
	})

	return &HTTPServer{
		engine: engine,
	}
}

func (server *HTTPServer) registerMetricsRoute() {
	registry := metrics.CreateRegistry()

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	server.engine.GET("/metrics", gin.WrapH(handler))
}

func (server *HTTPServer) registerMiddlewares() {
	server.engine.Use(
		middlewares.NewHTTPRequestCounterMiddleware(),
		middlewares.NewHTTPResponseLatencyCounterMiddleware(),
	)
}

func (server *HTTPServer) AddPollController(controller *pollController.HTTPPollController) {
	v1 := server.engine.Group("/v1")

	v1.POST("/polls", controller.Create)
	v1.GET("/polls/:id", controller.GetByID)
	v1.GET("/polls", controller.List)
}

func (server *HTTPServer) ListenAndServe(address string) error {
	server.registerMiddlewares()
	server.registerMetricsRoute()

	return server.engine.Run(address)
}
