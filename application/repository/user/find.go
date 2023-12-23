package user

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/domain/model"
	"github.com/umardev500/kost/utils"
)

func (ur *userRepository) Find(ctx context.Context, find model.UserFind) (result model.UserFindAllData, err error) {
	var params = find.Params
	var filters = find.Filters
	var args []interface{} = []interface{}{}

	// Extended where clause
	whereClause := " WHERE 1 = 1"

	// Search
	var search *string = params.Search
	if search != nil {
		whereClause += fmt.Sprintf(" AND (username LIKE $%d OR email LIKE $%d)", len(args)+1, len(args)+2)
		args = append(args, fmt.Sprintf("%%%s%%", *search))
		args = append(args, fmt.Sprintf("%%%s%%", *search))
	}

	// Filters
	if filters.Status != "" {
		whereClause += fmt.Sprintf(" AND u.status = $%d", len(args)+1)
		args = append(args, filters.Status)
	}

	// Pagination
	offset := utils.GetOffset(find.Pagination)
	paging := fmt.Sprintf(` LIMIT $%d OFFSET $%d`, len(args)+1, len(args)+2)
	args = append(args, find.Pagination.PageSize)
	args = append(args, offset)

	query := fmt.Sprintf(`
		SELECT u.* FROM users u
		%s
		%s
	`, whereClause, paging)

	db := ur.tr.GetConn(ctx)
	rows, err := db.QueryxContext(ctx, query, args...)
	if err != nil {
		log.Error().Msgf("failed for user query context")
		return
	}

	// Scans
	var users []model.User
	for rows.Next() {
		var each model.User
		if err = rows.StructScan(&each); err != nil {
			log.Error().Msgf("failed for struct scanning")
			return
		}

		users = append(users, each)
	}

	newArgs := args[:len(args)-2] // exclude paging
	queryCount := fmt.Sprintf(`SELECT COUNT(*) FROM users u %s`, whereClause)
	var total int64
	err = db.QueryRowxContext(ctx, queryCount, newArgs...).Scan(&total)
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
