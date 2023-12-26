package user

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/kost/application/handler"
	"github.com/umardev500/kost/domain/model"
	"github.com/umardev500/kost/utils"
)

func (uh *userHandler) Delete(c *fiber.Ctx) (err error) {
	var id = c.Params("id")

	uuidID := uuid.MustParse(id)
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	acData := model.ActionData{
		ID:     &uuidID,
		UserID: nil,
	}
	err = uh.uc.Delete(ctx, acData)
	if err != nil {
		return utils.ErrorHandler(c, err)
	}

	return handler.FiberOK(c, fiber.StatusOK, "Delete user success", nil)
}
