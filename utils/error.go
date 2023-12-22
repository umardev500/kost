package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/ekost/domain/model/errs"
	"github.com/umardev500/ekost/domain/model/responses"
)

func ErrorHandler(c *fiber.Ctx, err error, code int) error {
	payload := responses.Err{
		ID:      nil,
		Code:    code,
		Success: false,
		Message: "",
		Detail:  nil,
	}

	if customErr, ok := err.(errs.CustomError); ok {
		payload.ID = &customErr.ID
		payload.Code = customErr.StatusCode
		payload.Message = customErr.Error()
		return c.JSON(payload)
	}

	payload.Code = 500
	payload.Message = fiber.ErrInternalServerError.Message
	return c.JSON(payload)
}
