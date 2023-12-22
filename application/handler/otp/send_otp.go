package otp

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/kost/domain/model/otp"
	"github.com/umardev500/kost/utils"
)

func (oh *otpHandler) SendOtp(c *fiber.Ctx) (err error) {
	var payload otp.OtpParams
	if err = c.BodyParser(&payload); err != nil {
		return utils.ErrorHandler(c, err, fiber.StatusBadRequest)
	}

	// Validate
	if err := oh.validate.Struct(payload); err != nil {
		return utils.ErrorHandler(c, err, fiber.StatusUnprocessableEntity)
	}

	params := otp.OtpParams{
		Dur: 5 * time.Minute,
		Value: otp.OtpValue{
			Name: payload.Value.Name,
		},
		Email:    payload.Email,
		Location: "templates/email/otp.html",
		Subject:  payload.Subject,
	}
	oh.uc.SendOtp(c.Context(), params)

	return c.JSON("OK")
}
