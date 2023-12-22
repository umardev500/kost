package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/kost/config"
)

type Api struct {
	validate *validator.Validate
	storage  fiber.Storage
	db       config.DB
}

func NewApi(
	v *validator.Validate,
	s fiber.Storage,
	db config.DB,
) *Api {
	return &Api{
		validate: v,
		storage:  s,
		db:       db,
	}
}

func (a *Api) Load(r fiber.Router) {
	r.Route("otp", a.Otp)
}
