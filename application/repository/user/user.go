package user

import (
	"github.com/umardev500/kost/config"
	"github.com/umardev500/kost/domain"
)

type userRepository struct {
	tr *config.Trx
}

func NewUserRepository(tr *config.Trx) domain.UserRepository {
	return &userRepository{
		tr: tr,
	}
}
