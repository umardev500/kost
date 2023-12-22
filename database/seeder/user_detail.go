package seeder

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/utils/date"
)

type UserDetails struct {
	ID        string   `json:"id"`
	UserID    string   `json:"user_id"`
	FullName  string   `json:"full_name"`
	DOB       date.Dob `json:"dob"`
	Gender    string   `json:"gender"`
	Phone     string   `json:"phone"`
	Avatar    string   `json:"avatar"`
	CreatedBy string   `json:"created_by"`
}

func (s *Seeder) downUserDetails(ctx context.Context) (err error) {
	query := `
	DELETE FROM user_details WHERE id = $1
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing user_details query: %v", err)
		return
	}
	defer stmt.Close()

	var details []UserDetails
	err = s.getData("database/seeder/data/user-details.json", &details)
	if err != nil {
		log.Error().Msgf("Error getting user_details data: %v", err)
		return
	}

	for _, detail := range details {
		_, err = stmt.ExecContext(ctx, detail.ID)
		if err != nil {
			log.Error().Msgf("Error deleting user_details data: %v", err)
			return
		}
	}

	return
}

func (s *Seeder) SeedUserDetails(ctx context.Context) (err error) {
	query := `
	INSERT INTO user_details (id, user_id, full_name, dob, gender, phone, avatar, created_by)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing user_details query: %v", err)
		return
	}
	defer stmt.Close()

	var details []UserDetails
	err = s.getData("database/seeder/data/user-details.json", &details)
	if err != nil {
		log.Error().Msgf("Error getting user_details data: %v", err)
		return
	}

	for _, detail := range details {
		_, err = stmt.ExecContext(ctx, detail.ID, detail.UserID, detail.FullName, detail.DOB.String(), detail.Gender, detail.Phone, detail.Avatar, detail.CreatedBy)
		if err != nil {
			log.Error().Msgf("Error inserting user_details data: %v", err)
			return
		}
	}

	return
}
