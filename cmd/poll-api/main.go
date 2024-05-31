package main

import (
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yunusemre12500/poll-api/cmd/poll-api/defaults"
	"github.com/yunusemre12500/poll-api/cmd/poll-api/flags"
	controllers "github.com/yunusemre12500/poll-api/internal/infrastructure/controllers/poll/v1"
	database "github.com/yunusemre12500/poll-api/internal/infrastructure/database"
	repositories "github.com/yunusemre12500/poll-api/internal/infrastructure/repositories/poll/v1"
	"github.com/yunusemre12500/poll-api/internal/infrastructure/server"
)

var (
	bindAddress    string
	bindPort       uint
	DatabaseURI    string
	logFormat      string
	LogPrettyPrint bool
)

func init() {
	flag.StringVar(&bindAddress, flags.BindAddress, defaults.BindAddress, "")
	flag.UintVar(&bindPort, flags.BindPort, defaults.BindPort, "")
	flag.StringVar(&DatabaseURI, flags.DatabaseURI, defaults.DatabaseURI, "")
	flag.StringVar(&logFormat, flags.LogFormat, defaults.LogFormat, "")
	flag.BoolVar(&LogPrettyPrint, flags.LogPrettyPrint, defaults.LogPrettyPrint, "")
}

func main() {
	flag.Parse()

	logger := logrus.New()

	if LogPrettyPrint && logFormat != "json" {
		panic("Pretty log output is only supported with JSON format.")
	}

	switch logFormat {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: LogPrettyPrint,
		})
	case "text":
	case "":
		logger.SetFormatter(&logrus.TextFormatter{
			DisableLevelTruncation: true,
		})
	default:
		panic(fmt.Sprintf("Unsupported log format: %s", logFormat))
	}

	client, err := database.Connect(DatabaseURI)

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Database: %s", err))
	}

	pollCollection := client.Database("public").Collection("polls")

	pollRepository := repositories.NewMongoPollRepository(*pollCollection)

	pollController := controllers.NewHTTPPollController(&pollRepository)

	address := fmt.Sprintf("%s:%d", bindAddress, bindPort)

	server := server.NewHTTPServer(logger, address)

	pollController.AddRoutes(server.Mux())

	server.Start()
}
