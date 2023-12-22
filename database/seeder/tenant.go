package seeder

import (
	"context"

	"github.com/rs/zerolog/log"
)

type Tenant struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	IDType        string `json:"id_type"`
	IDNumber      string `json:"id_no"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Subdomain     string `json:"subdomain"`
	Address       string `json:"address"`
	ProvinceID    string `json:"province_id"`
	CityID        string `json:"city_id"`
	SubdistrictID string `json:"subdistrict_id"`
	VillageID     string `json:"village_id"`
	ZipCode       string `json:"zip_code"`
	Status        string `json:"status"`
	CreatedBy     string `json:"created_by"`
}

func (s *Seeder) downTenants(ctx context.Context) (err error) {
	query := `
	DELETE FROM tenants WHERE id = $1
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing tenants query: %v", err)
		return
	}
	defer stmt.Close()

	var tenants []Tenant
	err = s.getData("database/seeder/data/tenants.json", &tenants)
	if err != nil {
		log.Error().Msgf("Error getting tenants data: %v", err)
		return
	}

	for _, tenant := range tenants {
		_, err = stmt.ExecContext(ctx, tenant.ID)
		if err != nil {
			log.Error().Msgf("Error deleting tenants data: %v", err)
			return
		}
	}

	return
}

func (s *Seeder) SeedTenants(ctx context.Context) (err error) {
	query := `
	INSERT INTO tenants (id, name, id_type, id_no, email, phone, subdomain, address, province_id, city_id, subdistrict_id, village_id, zip_code, status, created_by)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing creds data: %v", err)
		return
	}
	defer stmt.Close()

	var tenants []Tenant
	err = s.getData("database/seeder/data/tenants.json", &tenants)
	if err != nil {
		log.Error().Msgf("Error getting tenant data: %v", err)
		return
	}

	for _, tenant := range tenants {
		_, err = stmt.ExecContext(
			ctx,
			tenant.ID,
			tenant.Name,
			tenant.IDType,
			tenant.IDNumber,
			tenant.Email,
			tenant.Phone,
			tenant.Subdomain,
			tenant.Address,
			tenant.ProvinceID,
			tenant.CityID,
			tenant.SubdistrictID,
			tenant.VillageID,
			tenant.ZipCode,
			tenant.Status,
			tenant.CreatedBy,
		)

		if err != nil {
			log.Error().Msgf("Error inserting tenant data: %v", err)
			return
		}
	}

	return
}
