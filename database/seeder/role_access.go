package seeder

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type RoleAccess struct {
	ID         string   `json:"id"`
	RoleID     string   `json:"role_id"`
	ModuleID   string   `json:"module_id"`
	TenantID   *string  `json:"tenant_id"`
	Permission []string `json:"permission"`
	Status     string   `json:"status"`
	CreatedBy  *string  `json:"created_by"`
}

func (s *Seeder) downRoleAccess(ctx context.Context) (err error) {
	query := `
	DELETE FROM role_access WHERE id = $1
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing role_access query: %v", err)
		return
	}
	defer stmt.Close()

	var roleAccess []RoleAccess
	err = s.getData("database/seeder/data/role-access.json", &roleAccess)
	if err != nil {
		log.Error().Msgf("Error getting role_access data: %v", err)
		return
	}

	for _, access := range roleAccess {
		_, err = stmt.ExecContext(ctx, access.ID)
		if err != nil {
			log.Error().Msgf("Error deleting role_access data: %v", err)
			return
		}
	}

	return
}

func (s *Seeder) SeedRoleAccess(ctx context.Context) (err error) {
	query := `
	INSERT INTO role_access (id, role_id, module_id, tenant_id, permission, status, created_by)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing role_access query: %v", err)
		return
	}
	defer stmt.Close()

	var roleAccess []RoleAccess
	err = s.getData("database/seeder/data/role-access.json", &roleAccess)
	if err != nil {
		log.Error().Msgf("Error getting role_access data: %v", err)
		return
	}

	for _, access := range roleAccess {
		var permissions []byte
		if permissions, err = json.Marshal(access.Permission); err != nil {
			return err
		}

		_, err = stmt.ExecContext(ctx, access.ID, access.RoleID, access.ModuleID, access.TenantID, permissions, access.Status, access.CreatedBy)
		if err != nil {
			log.Error().Msgf("Error inserting role_access data: %v", err)
			return
		}
	}

	return
}
