package seeder

import (
	"context"

	"github.com/rs/zerolog/log"
)

type Role struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	TenantID  *string `json:"tenant_id"`
	Status    string  `json:"status"`
	CreatedBy *string `json:"created_by"`
}

func (s *Seeder) downRoles(ctx context.Context) (err error) {
	query := `
	DELETE FROM roles WHERE id = $1
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing role query: %v", err)
		return
	}
	defer stmt.Close()

	var roles []Role
	err = s.getData("database/seeder/data/roles.json", &roles)
	if err != nil {
		log.Error().Msgf("Error getting role data: %v", err)
		return
	}

	for _, role := range roles {
		_, err = stmt.ExecContext(ctx, role.ID)
		if err != nil {
			log.Error().Msgf("Error deleting role data: %v", err)
			return
		}
	}

	return
}

func (s *Seeder) SeedRoles(ctx context.Context) (err error) {
	query := `
	INSERT INTO roles (id, name, tenant_id, status, created_by)
	VALUES ($1, $2, $3, $4, $5)
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing role query: %v", err)
		return
	}
	defer stmt.Close()

	var roles []Role
	err = s.getData("database/seeder/data/roles.json", &roles)
	if err != nil {
		log.Error().Msgf("Error getting role data: %v", err)
		return
	}

	for _, role := range roles {
		_, err = stmt.ExecContext(ctx, role.ID, role.Name, role.TenantID, role.Status, role.CreatedBy)
		if err != nil {
			log.Error().Msgf("Error inserting role data: %v", err)
			return
		}
	}

	return
}
