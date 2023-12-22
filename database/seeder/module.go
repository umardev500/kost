package seeder

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type Module struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Features  []string `json:"features"`
	Status    string   `json:"status"`
	Level     string   `json:"level"`
	CreatedBy *string  `json:"created_by"`
}

func (s *Seeder) downModules(ctx context.Context) (err error) {
	query := `
	DELETE FROM modules WHERE id = $1
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing module query: %v", err)
		return
	}
	defer stmt.Close()

	var modules []Module
	err = s.getData("database/seeder/data/modules.json", &modules)
	if err != nil {
		log.Error().Msgf("Error getting module data: %v", err)
		return
	}

	for _, module := range modules {
		_, err = stmt.ExecContext(ctx, module.ID)
		if err != nil {
			log.Error().Msgf("Error deleting module data: %v", err)
			return
		}
	}

	return
}

func (s *Seeder) SeedModules(ctx context.Context) (err error) {
	query := `
	INSERT INTO modules (id, name, features, status, level, created_by)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing module query: %v", err)
		return
	}
	defer stmt.Close()

	var modules []Module
	err = s.getData("database/seeder/data/modules.json", &modules)
	if err != nil {
		log.Error().Msgf("Error getting module data: %v", err)
		return
	}

	for _, module := range modules {
		var features []byte
		if features, err = json.Marshal(module.Features); err != nil {
			return err
		}

		_, err = stmt.ExecContext(ctx, module.ID, module.Name, features, module.Status, module.Level, module.CreatedBy)
		if err != nil {
			log.Error().Msgf("Error inserting module data: %v", err)
			return
		}
	}

	return
}
