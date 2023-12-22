package application

import (
	"context"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/config"
	"github.com/umardev500/kost/routes"
	"github.com/umardev500/kost/storage"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Start(ctx context.Context) (err error) {
	app := fiber.New()
	ch := make(chan error, 1)

	rDB := config.NewRedis()
	strg := storage.NewRedisStorage(rDB)
	db := config.NewPostgres(ctx)
	v := validator.New()
	router := routes.NewRouter(app, strg, db, v)
	router.Load()

	func() {
		port := os.Getenv("APP_PORT")
		msg := fmt.Sprintf("🚀 Server is running on port %s", port)
		log.Info().Msg(msg)

		addr := fmt.Sprintf(":%s", port)
		if err := app.Listen(addr); err != nil {
			ch <- err
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		app.Shutdown() // Shutdown the application
		fmt.Println()
		msg := "🚀 Application is shutting down gracefully... 😴"
		log.Info().Msg(msg)
	}

	return
}
