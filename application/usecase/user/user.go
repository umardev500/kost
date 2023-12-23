package usecase

import (
	"github.com/umardev500/kost/config"
	"github.com/umardev500/kost/domain"
)

type userUsecase struct {
	repo domain.UserRepository
	tr   *config.Trx
}

func NewUserUsecase(r domain.UserRepository, tr *config.Trx) domain.UserUsecase {
	return &userUsecase{
		repo: r,
		tr:   tr,
	}
}
