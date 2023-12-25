package seeder

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/constants"
)

type User struct {
	ID        string           `json:"id"`
	TenantID  *string          `json:"tenant_id"`
	Email     string           `json:"email"`
	Username  string           `json:"username"`
	Password  string           `json:"password"`
	CreatedBy *string          `json:"created_by"`
	Status    constants.Status `json:"status"`
}

func (s *Seeder) downUsers(ctx context.Context) (err error) {
	query := `
	DELETE FROM users WHERE id = $1
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing user query: %v", err)
		return
	}
	defer stmt.Close()

	var users []User
	err = s.getData("database/seeder/data/users.json", &users)
	if err != nil {
		log.Error().Msgf("Error getting user data: %v", err)
		return
	}

	for _, user := range users {
		_, err = stmt.ExecContext(ctx, user.ID)
		if err != nil {
			log.Error().Msgf("Error deleting user data: %v", err)
			return
		}
	}

	return
}

func (s *Seeder) SeedUsers(ctx context.Context) (err error) {
	query := `
	INSERT INTO users (id, tenant_id, email, username, password, status, created_by)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	db := s.tx.GetConn(ctx)
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		log.Error().Msgf("Error preparing user query: %v", err)
		return
	}
	defer stmt.Close()

	var users []User
	err = s.getData("database/seeder/data/users.json", &users)
	if err != nil {
		log.Error().Msgf("Error getting user data: %v", err)
		return
	}

	// Append dummy data
	for i := 0; i <= 9_000; i++ {
		userID := uuid.New().String()
		users = append(users, User{
			ID:        uuid.New().String(),
			Email:     faker.Email(),
			TenantID:  nil,
			Username:  fmt.Sprintf("%s%d", faker.Name(), i),
			Password:  fmt.Sprintf("password%d", rand.Int()),
			Status:    constants.StatusActive,
			CreatedBy: &userID,
		})
	}

	fmt.Println(len(users))
	// End of appending dummy data

	for _, user := range users {
		_, err = stmt.ExecContext(ctx, user.ID, user.TenantID, user.Email, user.Username, user.Password, user.Status, user.CreatedBy)
		if err != nil {
			log.Error().Msgf("Error inserting user data: %v", err)
			return
		}
	}

	return
}
