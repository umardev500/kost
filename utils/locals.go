package utils

import "github.com/gofiber/fiber/v2"

func GetUserID(c *fiber.Ctx) (result *string) {
	id := c.Locals("id")
	if id == nil {
		return
	}

	idStr := id.(string)
	result = &idStr

	return
}
