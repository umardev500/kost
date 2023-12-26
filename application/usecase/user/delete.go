package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/domain/model"
	"github.com/umardev500/kost/utils"
)

func (uuc *userUsecase) Delete(ctx context.Context, data model.ActionData) (err error) {
	aff, err := uuc.repo.Delete(ctx, data)
	if err != nil {
		newErr := utils.NewError()
		log.Error().Msgf("Error while deleting user: id=%s %v", newErr.ID, err)

		return
	}

	err = utils.ErrAffected(aff, err)
	if err != nil {
		newErr := utils.NewError()
		log.Warn().Msgf("Error while deleting user: id=%s %v", newErr.ID, err)
	}

	return
}
