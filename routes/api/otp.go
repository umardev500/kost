package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/ekost/application/handler/otp"
	otpUc "github.com/umardev500/ekost/application/usecase/otp"
)

func (a *Api) Otp(router fiber.Router) {
	ou := otpUc.NewOtpUsecase(a.storage)
	oh := otp.NewOtpHandler(ou, a.validate)

	router.Post("/send-otp", oh.SendOtp)
}
