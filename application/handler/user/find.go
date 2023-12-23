package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/kost/application/handler"
	"github.com/umardev500/kost/constants"
	"github.com/umardev500/kost/domain/model"
	"github.com/umardev500/kost/utils"
)

func (uh *userHandler) Find(c *fiber.Ctx) (err error) {
	// Paginations
	var pagNum int64 = 1
	var pageSize int64 = 10
	utils.Atoi(&pagNum, c.Query("page"))
	utils.Atoi(&pageSize, c.Query("size"))
	pagination := model.PaginationParams{
		PageNum:  pagNum,
		PageSize: pageSize,
	}

	// Filters
	var status = c.Query("status")
	filters := model.UserFilter{
		Status: constants.Status(status),
	}

	search := c.Query("search")
	params := model.UserFindParams{
		Search: &search,
	}

	find := model.UserFind{
		Pagination: pagination,
		Filters:    filters,
		Params:     params,
	}

	// Payloads
	payload, err := uh.uc.Find(c.Context(), find)
	if err != nil {
		return utils.ErrorHandler(c, err)
	}

	return handler.FiberOK(c, fiber.StatusOK, "Find users", payload)
}
