package otp

import (
	"github.com/go-playground/validator/v10"
	"github.com/umardev500/kost/domain"
)

type otpHandler struct {
	uc       domain.OtpUsecase
	validate *validator.Validate
}

func NewOtpHandler(uc domain.OtpUsecase, v *validator.Validate) domain.OtpHandler {
	return &otpHandler{
		uc:       uc,
		validate: v,
	}
}
