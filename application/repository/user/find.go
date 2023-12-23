package user

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/domain/model"
)

func (ur *userRepository) Find(ctx context.Context, find model.UserFind) (result model.UserFindAllData, err error) {
	whereClause := " WHERE 1 = 1"
	// var params = find.Params
	var filters = find.Filters
	var args []interface{} = []interface{}{}

	// Filter by status
	if filters.Status != "" {
		whereClause += fmt.Sprintf(" AND u.status = $%d", len(args)+1)
		args = append(args, filters.Status)
	}

	query := fmt.Sprintf(`
		SELECT u.* FROM users u
		%s
	`, whereClause)

	db := ur.tr.GetConn(ctx)
	rows, err := db.QueryxContext(ctx, query, args...)
	if err != nil {
		log.Error().Msgf("failed for user query context")
		return
	}

	var users []model.User

	for rows.Next() {
		var each model.User
		if err = rows.StructScan(&each); err != nil {
			log.Error().Msgf("failed for struct scanning")
			return
		}

		users = append(users, each)
	}

	queryCount := fmt.Sprintf(`SELECT COUNT(*) FROM users u %s`, whereClause)
	var total int64
	err = db.QueryRowxContext(ctx, queryCount, args...).Scan(&total)
	if err != nil {
		log.Error().Msgf("failed to select count users")
		return
	}

	result = model.UserFindAllData{
		Users: users,
		Total: total,
	}

	return
}
