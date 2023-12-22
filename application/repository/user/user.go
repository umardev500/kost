package user

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/config"
	"github.com/umardev500/kost/domain"
	"github.com/umardev500/kost/domain/model"
)

type userRepository struct {
	tr *config.Trx
}

func NewUserRepository(tr *config.Trx) domain.UserRepository {
	return &userRepository{
		tr: tr,
	}
}

func (ur *userRepository) FindBy(ctx context.Context, find model.UserFind) (result *model.User, err error) {
	query := `
	SELECT u.* FROM users u
	`
	whereClause := " WHERE 1 = 1"

	var params = find.Params
	var filters = find.Filters
	var args []interface{} = []interface{}{}

	// Find by username
	if params.Username != nil {
		whereClause += fmt.Sprintf(" AND u.username = $%d", len(args)+1)
		args = append(args, params.Username)
	}

	if params.ID != nil {
		whereClause += fmt.Sprintf(" AND u.id = $%d", len(args)+1)
		args = append(args, params.ID)
	}

	// Filter by status
	if filters.Status != "" {
		whereClause += fmt.Sprintf(" AND u.status = $%d", len(args)+1)
		args = append(args, filters.Status)
	}

	db := ur.tr.GetConn(ctx)
	err = db.QueryRowxContext(ctx, query, args...).StructScan(&result)
	if err != nil {
		log.Error().Msgf("Error while finding user: %v", err)
		return
	}

	return
}
