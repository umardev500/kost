package domain

import (
	"context"

	"github.com/umardev500/kost/domain/model"
)

type UserRepository interface {
	FindBy(context.Context, model.UserFind) (*model.User, error)
}
