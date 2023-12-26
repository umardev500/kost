package usecase

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/domain/model"
	"github.com/umardev500/kost/utils"
)

func (uuc *userUsecase) Delete(ctx context.Context, data model.ActionData) (err error) {
	aff, err := uuc.repo.Delete(ctx, data)
	if err != nil {
		newErr := utils.NewError()
		log.Error().Msgf("Error while deleting user: id=%s %v", newErr.ID, err)

		return newErr
	}

	err = utils.ErrAffected(aff, err)
	if err != nil {
		newErr := utils.NewError()
		newErr.StatusCode = fiber.StatusOK
		newErr.Message = "Delete not affected"
		log.Warn().Msgf("Error while deleting user: id=%s %v", newErr.ID, err)
		return newErr
	}

	return
}
