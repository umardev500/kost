package user

import (
	"context"
	"fmt"

	"github.com/umardev500/kost/domain/model"
)

func (ur *userRepository) FindByID(ctx context.Context, find model.UserFind) (result *model.User, err error) {
	whereClause := " WHERE 1 = 1"

	var params = find.Params
	var filters = find.Filters
	var args []interface{} = []interface{}{}

	if params.ID != nil {
		whereClause += fmt.Sprintf(" AND u.id = $%d", len(args)+1)
		args = append(args, params.ID)
	}

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
	var user model.User
	err = db.QueryRowxContext(ctx, query, args...).StructScan(&user)
	if err != nil {
		return
	}

	result = &user
	return
}
