package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/kost/domain/model"
)

func FiberOK(c *fiber.Ctx, code int, message string, data *model.Payload) error {
	payload := model.OK{
		Code:    code,
		Success: true,
		Message: message,
		Payload: data,
	}
	return c.JSON(payload)
}

func FiberErr(c *fiber.Ctx, code int, message string, detail interface{}) error {
	payload := model.Err{
		Code:    code,
		Success: false,
		Message: message,
		Detail:  detail,
	}
	return c.JSON(payload)
}
