package seeder

import (
	"context"

	"github.com/rs/zerolog/log"
)

type RoleUser struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	RoleID    string `json:"role_id"`
	Status    string `json:"status"`
	CreatedBy string `json:"created_by"`
}

func (s *Seeder) downRoleUsers(ctx context.Context) (err error) {
	query := `
	DELETE FROM role_users WHERE id = $1
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing role_user query: %v", err)
		return
	}
	defer stmt.Close()

	var roleUsers []RoleUser
	err = s.getData("database/seeder/data/role-users.json", &roleUsers)
	if err != nil {
		log.Error().Msgf("Error getting role_user data: %v", err)
		return
	}

	for _, roleUser := range roleUsers {
		_, err = stmt.ExecContext(ctx, roleUser.ID)
		if err != nil {
			log.Error().Msgf("Error deleting role_user data: %v", err)
			return
		}
	}

	return
}

func (s *Seeder) SeedRoleUsers(ctx context.Context) (err error) {
	query := `
	INSERT INTO role_users (id, user_id, role_id, status, created_by)
	VALUES ($1, $2, $3, $4, $5)
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing role_user query: %v", err)
		return
	}
	defer stmt.Close()

	var roleUsers []RoleUser
	err = s.getData("database/seeder/data/role-users.json", &roleUsers)
	if err != nil {
		log.Error().Msgf("Error getting role_user data: %v", err)
		return
	}

	for _, roleUser := range roleUsers {
		_, err = stmt.ExecContext(ctx, roleUser.ID, roleUser.UserID, roleUser.RoleID, roleUser.Status, roleUser.CreatedBy)
		if err != nil {
			log.Error().Msgf("Error inserting role_user data: %v", err)
			return
		}
	}

	return
}
