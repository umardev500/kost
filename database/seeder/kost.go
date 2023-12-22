package seeder

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type Kost struct {
	ID            string   `json:"id"`
	TenantID      string   `json:"tenant_id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Photos        []string `json:"photos"`
	Videos        []string `json:"videos"`
	ProvinceID    string   `json:"province_id"`
	CityID        string   `json:"city_id"`
	SubdistrictID string   `json:"subdistrict_id"`
	VillageID     string   `json:"village_id"`
	ZipCode       string   `json:"zip_code"`
	Status        string   `json:"status"`
	CreatedBy     string   `json:"created_by"`
}

func (s *Seeder) downKosts(ctx context.Context) (err error) {
	query := `
	DELETE FROM kosts WHERE id = $1
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing kosts query: %v", err)
		return
	}
	defer stmt.Close()

	var kosts []Kost
	err = s.getData("database/seeder/data/kosts.json", &kosts)
	if err != nil {
		log.Error().Msgf("Error getting kosts data: %v", err)
		return
	}

	for _, kost := range kosts {
		_, err = stmt.ExecContext(ctx, kost.ID)
		if err != nil {
			log.Error().Msgf("Error deleting kosts data: %v", err)
			return
		}
	}

	return
}

func (s *Seeder) SeedKosts(ctx context.Context) (err error) {
	query := `
	INSERT INTO kosts (id, tenant_id, name, description, photos, videos, province_id, city_id, subdistrict_id, village_id, zip_code, status, created_by)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing kosts query: %v", err)
		return
	}
	defer stmt.Close()

	var kosts []Kost
	err = s.getData("database/seeder/data/kosts.json", &kosts)
	if err != nil {
		log.Error().Msgf("Error getting kosts data: %v", err)
		return
	}

	for _, kost := range kosts {
		var photos []byte
		if photos, err = json.Marshal(kost.Photos); err != nil {
			log.Error().Msgf("Error marshalling photos JSON: %v", err)
			return
		}

		var videos []byte
		if videos, err = json.Marshal(kost.Videos); err != nil {
			log.Error().Msgf("Error marshalling videos JSON: %v", err)
			return
		}

		_, err = stmt.ExecContext(ctx, kost.ID, kost.TenantID, kost.Name, kost.Description, &photos, &videos, kost.ProvinceID, kost.CityID, kost.SubdistrictID, kost.VillageID, kost.ZipCode, kost.Status, kost.CreatedBy)
		if err != nil {
			log.Error().Msgf("Error inserting kosts data: %v", err)
			return
		}
	}

	return
}
