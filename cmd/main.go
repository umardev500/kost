package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/application"
)

func init() {
	log.Logger = log.With().Caller().Logger()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	app := application.NewApplication()
	err := app.Start(ctx)
	if err != nil {
		log.Fatal().Msgf("Error starting application: %s", err)
	}
}
