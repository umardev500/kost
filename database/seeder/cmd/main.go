package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/config"
	"github.com/umardev500/kost/database/seeder"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	db := config.NewPostgres(context.Background())
	defer db.Close()
	trx := config.NewTransaction(db)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	caller(trx, ctx)
}

func caller(tx *config.Trx, ctx context.Context) (err error) {
	seeder := seeder.NewSeeder(tx)

	var options = []string{"Populate", "Down"}
	for i, val := range options {
		fmt.Printf("%d. %s\n", i+1, val)
	}

	var input string
	fmt.Print("Select option: ")
	fmt.Scanln(&input)

	inputNum, err := strconv.Atoi(input)
	if err != nil {
		log.Error().Msg("Enter the right number")

		caller(tx, ctx)
		return nil
	}

	if inputNum == 1 {
		err = seeder.Populate(ctx)
		if err != nil {
			return
		}
	} else if inputNum == 2 {
		err = seeder.Down(ctx)
		if err != nil {
			return
		}
	} else {
		log.Info().Msg("Please select one of list")
		caller(tx, ctx)
		return nil
	}

	return
}
