package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HTTPServer struct {
	ctx context.Context
	mux *http.ServeMux
	srv *http.Server
}

func NewHTTPServer(address string) *HTTPServer {
	ctx := context.Background()

	mux := http.NewServeMux()

	srv := http.Server{
		Addr:    address,
		Handler: mux,
	}

	return &HTTPServer{
		ctx: ctx,
		mux: mux,
		srv: &srv,
	}
}

func (httpServer *HTTPServer) Start() {
	slog.Info("Starting HTTP server...")

	done := make(chan os.Signal, 1)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := httpServer.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start HTTP server: %s", err)
			os.Exit(1)
		}
	}()

	slog.Info("HTTP server started.")

	<-done

	ctx, cancel := context.WithTimeout(httpServer.ctx, 5*time.Second)

	slog.Info("Shutting down HTTP server...")

	defer cancel()

	if err := httpServer.srv.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown HTTP server gracefully: %v", err)
	}
}

func (httpServer *HTTPServer) Mux() *http.ServeMux {
	return httpServer.mux
}
