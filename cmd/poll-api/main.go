package main

import (
	"flag"
	"fmt"

	"github.com/yunusemre12500/poll-api/cmd/poll-api/defaults"
	"github.com/yunusemre12500/poll-api/cmd/poll-api/flags"
	pollController "github.com/yunusemre12500/poll-api/internal/infrastructure/controllers/poll/v1"
	"github.com/yunusemre12500/poll-api/internal/infrastructure/database"
	pollRepository "github.com/yunusemre12500/poll-api/internal/infrastructure/repositories/poll/v1"
	server "github.com/yunusemre12500/poll-api/internal/infrastructure/server"
	pollService "github.com/yunusemre12500/poll-api/internal/infrastructure/services/poll/v1"
)

var (
	databaseUrl         string
	httpListenerAddress string
	httpListenerPort    uint
)

func init() {
	flag.StringVar(&databaseUrl, flags.DatabaseURL, defaults.DatabaseURL, "URL to connect to database")
	flag.StringVar(&httpListenerAddress, flags.HTTPListenerAddress, defaults.HTTPListenerAddress, "IP address to listen on HTTP server")
	flag.UintVar(&httpListenerPort, flags.HTTPListenerPort, defaults.HTTPListenerPort, "Port number to listen on HTTP server")
}

func main() {
	flag.Parse()

	database, err := database.Connect(databaseUrl)

	if err != nil {
		panic("Failed to connect to database")
	}

	pollCollection := database.Collection("polls")
	pollRepository := pollRepository.NewMongoPollRepository(pollCollection)
	pollService := pollService.NewHTTPPollService(pollRepository)
	pollController := pollController.NewHTTPPollController(pollService)

	server := server.NewHTTPServer()

	server.AddPollController(pollController)

	address := fmt.Sprintf("%s:%d", httpListenerAddress, httpListenerPort)

	server.ListenAndServe(address)
}
