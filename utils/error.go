package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/kost/domain/model"
)

type CustomError struct {
	ID         uuid.UUID
	Message    interface{}
	TimeStamp  time.Time
	StatusCode int
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

func NewError() CustomError {
	return CustomError{
		ID:        uuid.New(),
		TimeStamp: time.Now(),
	}
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	id := uuid.New()
	payload := model.Err{
		ID:      &id,
		Code:    500,
		Success: false,
		Message: fiber.ErrInternalServerError.Message,
		Detail:  nil,
	}

	if customErr, ok := err.(CustomError); ok {
		payload.ID = &customErr.ID
		payload.Code = customErr.StatusCode
		payload.Message = customErr.Error()
		return c.JSON(payload)
	}

	payload.Message = fiber.ErrInternalServerError.Message
	return c.JSON(payload)
}
