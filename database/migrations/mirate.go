package migrations

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func NewMigration(ctx context.Context) {
	start := time.Now()
	log.Info().Msgf("Starting migrations 🛠️...")

	db := config.NewPostgres(ctx)
	defer db.Close()
	defer db.DB.Close()
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatal().Msgf("Error creating postgres driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal().Msgf("Error creating migration instance: %v", err)
	}

	m.Up()

	log.Info().Msgf("Migrations completed \033[32m🎉 (\U000023F3 %s)\033[0m", time.Since(start))
}
