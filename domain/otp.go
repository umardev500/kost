package domain

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/kost/domain/model/otp"
)

type OtpHandler interface {
	SendOtp(*fiber.Ctx) error
}

type OtpUsecase interface {
	SendOtp(context.Context, otp.OtpParams) error
}
