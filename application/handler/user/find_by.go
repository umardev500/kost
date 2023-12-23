package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/kost/application/handler"
	"github.com/umardev500/kost/constants"
	"github.com/umardev500/kost/domain/model"
	"github.com/umardev500/kost/utils"
)

func (uh *userHandler) FindByID(c *fiber.Ctx) (err error) {
	var id = c.Params("id")
	var queries = c.Queries()

	userID := uuid.MustParse(id)
	result, err := uh.uc.FindByID(c.Context(), model.UserFind{
		Params: model.UserFindParams{
			ID: &userID,
		},
		Filters: model.UserFilter{
			Status: constants.Status(queries["status"]),
		},
	})
	if err != nil {
		return utils.ErrorHandler(c, err)
	}

	return handler.FiberOK(c, fiber.StatusOK, "Find user by id's", result)
}
