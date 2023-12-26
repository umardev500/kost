package user

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/domain/model"
)

func (ur *userRepository) Delete(ctx context.Context, data model.ActionData) (aff int64, err error) {
	query := `
		DELETE FROM users WHERE id = $1
	`
	db := ur.tr.GetConn(ctx)
	result, err := db.ExecContext(ctx, query, data.ID)
	if err != nil {
		log.Error().Msgf("Error while deleting user")
		return
	}

	aff, err = result.RowsAffected()
	return
}
