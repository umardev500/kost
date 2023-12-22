package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/umardev500/kost/config"
	"github.com/umardev500/kost/routes/api"
)

type Router struct {
	app      *fiber.App
	storage  fiber.Storage
	db       config.DB
	validate *validator.Validate
}

func NewRouter(
	a *fiber.App,
	s fiber.Storage,
	db config.DB,
	v *validator.Validate,
) *Router {
	return &Router{
		app:      a,
		storage:  s,
		db:       db,
		validate: v,
	}
}

func (r Router) Load() {

	// Middwlares
	r.app.Use(cors.New())

	api := api.NewApi(r.validate, r.storage, r.db)
	r.app.Route("api", api.Load)
}
