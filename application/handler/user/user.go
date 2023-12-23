package user

import "github.com/umardev500/kost/domain"

type userHandler struct {
	uc domain.UserUsecase
}

func NewUserhandler(uc domain.UserUsecase) domain.UserHandler {
	return &userHandler{
		uc: uc,
	}
}
