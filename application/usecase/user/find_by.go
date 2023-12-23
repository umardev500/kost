package usecase

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/domain/model"
	"github.com/umardev500/kost/utils"
)

func (uuc *userUsecase) FindByID(ctx context.Context, find model.UserFind) (result *model.Payload, err error) {
	data, err := uuc.repo.FindByID(ctx, find)
	if err != nil {
		if err == sql.ErrNoRows {
			newErr := utils.NewError()
			newErr.StatusCode = fiber.StatusNotFound
			newErr.Message = fiber.ErrNotFound.Message
			log.Info().Msgf("User data not found. id=%s", newErr.ID)

			return nil, newErr
		}

		newErr := utils.NewError()
		log.Error().Msgf("Error while finding user: id=%s %v", newErr.ID, err)

		return
	}

	result = &model.Payload{
		Data: data,
	}

	return
}
