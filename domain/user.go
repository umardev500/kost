package domain

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/kost/domain/model"
)

type UserHandler interface {
	Delete(*fiber.Ctx) error
	Find(*fiber.Ctx) error
	FindByID(*fiber.Ctx) error
}

type UserUsecase interface {
	FindByID(context.Context, model.UserFind) (*model.Payload, error)
	Find(context.Context, model.UserFind) (*model.Payload, error)
	Delete(context.Context, model.ActionData) error
}

type UserRepository interface {
	Delete(context.Context, model.ActionData) (int64, error)
	Find(context.Context, model.UserFind) (model.UserFindAllData, error)
	FindByID(context.Context, model.UserFind) (*model.User, error)
}
