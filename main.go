package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	app "github.com/alelince/metaheuristics-playground/internal/app"
	utils "github.com/alelince/metaheuristics-playground/internal/utils"
)

func main() {
	utils.InitLogger()
	log.Info().Msg("--- Starting Metaheuristics Playground ---")

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Signal(syscall.SIGTERM),
		os.Signal(syscall.SIGINT),
	)
	defer stop()

	var host string
	var port int
	parseFlags(&host, &port)

	app := app.NewApplication(fmt.Sprintf("%s:%d", host, port))
	if err := app.Start(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to start the application")
		os.Exit(1)
	}
}

func parseFlags(hostPtr *string, portPtr *int) {
	var help bool

	flag.StringVar(hostPtr, "host", "127.0.0.1", "REST API hostname or IP address")
	flag.IntVar(portPtr, "port", 8080, "REST API port")
	flag.BoolVar(&help, "help", false, "Show help")

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}
}
