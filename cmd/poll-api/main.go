package main

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/yunusemre12500/poll-api/cmd/poll-api/defaults"
	"github.com/yunusemre12500/poll-api/cmd/poll-api/flags"
	controllers "github.com/yunusemre12500/poll-api/internal/infrastructure/controllers/poll/v1"
	repositories "github.com/yunusemre12500/poll-api/internal/infrastructure/repositories/poll/v1"
	"github.com/yunusemre12500/poll-api/internal/infrastructure/server"
)

var (
	bindAddress string
	bindPort    uint
)

func init() {
	flag.StringVar(&bindAddress, flags.BindAddress, defaults.BindAddress, "")
	flag.UintVar(&bindPort, flags.BindPort, defaults.BindPort, "")
}

func main() {
	flag.Parse()

	address := fmt.Sprintf("%s:%d", bindAddress, bindPort)

	pollRepository := repositories.NewInMemoryPollRepository()

	pollController := controllers.NewHTTPPollController(&pollRepository)

	server := server.NewHTTPServer(address)

	pollController.AddRoutes(server.Mux())

	slog.Info("Starting server...")

	server.Start()
}
