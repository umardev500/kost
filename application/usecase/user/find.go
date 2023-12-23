package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/domain/model"
	"github.com/umardev500/kost/utils"
)

func (uuc *userUsecase) Find(ctx context.Context, find model.UserFind) (payload *model.Payload, err error) {
	data, err := uuc.repo.Find(ctx, find)
	if err != nil {
		id := uuid.New()
		log.Error().Msgf("failed to find users. id=%s err=%v", id, err)
		return
	}

	payload = &model.Payload{
		PagesTotal: utils.GetPagesTotal(data.Total, find.Pagination.PageSize),
		Data:       data.Users,
		PageNum:    find.Pagination.PageNum,
		PageSize:   find.Pagination.PageSize,
	}

	return
}
