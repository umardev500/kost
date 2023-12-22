package otp

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/kost/domain"
)

type otpUsecase struct {
	storage fiber.Storage
}

func NewOtpUsecase(s fiber.Storage) domain.OtpUsecase {
	return &otpUsecase{
		storage: s,
	}
}
