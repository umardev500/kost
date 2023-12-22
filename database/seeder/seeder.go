package seeder

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/kost/config"
)

type Seeder struct {
	tx *config.Trx
}

// NewSeeder returns a new instance of the Seeder struct.
//
// It takes a pointer to a `config.Trx` object as a parameter.
// The function returns a pointer to a `Seeder` struct.
func NewSeeder(tx *config.Trx) *Seeder {
	return &Seeder{
		tx: tx,
	}
}

// getData reads data from a file and unmarshals it into a struct.
//
// It takes in the file path as a string and the target struct as an interface{}.
// It returns an error if there was an issue reading or unmarshalling the data.
func (s *Seeder) getData(filePath string, result interface{}) (err error) {
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	filePath = filepath.Join(dir, filePath)
	f, err := os.ReadFile(filePath)
	if err != nil {
		log.Error().Msgf("Error reading data: %v", err)
		return
	}

	err = json.Unmarshal(f, &result)
	if err != nil {
		log.Error().Msgf("Error unmarshalling data: %v", err)
		return
	}

	return
}

func (s *Seeder) Populate(ctx context.Context) (err error) {

	s.tx.WithTransaction(ctx, func(ctx context.Context) (err error) {
		err = s.SeedTenants(ctx)
		if err != nil {
			return
		}
		err = s.SeedUsers(ctx)
		if err != nil {
			return
		}
		err = s.SeedUserDetails(ctx)
		if err != nil {
			return
		}
		err = s.SeedRoles(ctx)
		if err != nil {
			return
		}

		err = s.SeedRoleUsers(ctx)
		if err != nil {
			return
		}

		err = s.SeedModules(ctx)
		if err != nil {
			return
		}

		err = s.SeedRoleAccess(ctx)
		if err != nil {
			return
		}

		err = s.SeedKosts(ctx)
		if err != nil {
			return
		}

		return
	})

	return nil
}

func (s *Seeder) Down(ctx context.Context) (err error) {
	s.tx.WithTransaction(ctx, func(ctx context.Context) (err error) {
		log.Info().Msg("Down kosts")
		err = s.downKosts(ctx)
		if err != nil {
			return
		}
		log.Info().Msg("Down role access")
		err = s.downRoleAccess(ctx)
		if err != nil {
			return
		}
		log.Info().Msg("Down modules")
		err = s.downModules(ctx)
		if err != nil {
			return
		}
		log.Info().Msg("Down role users")
		err = s.downRoleUsers(ctx)
		if err != nil {
			return
		}
		log.Info().Msg("Down user details")
		err = s.downUserDetails(ctx)
		if err != nil {
			return
		}
		log.Info().Msg("Down users")
		err = s.downUsers(ctx)
		if err != nil {
			return
		}
		log.Info().Msg("Down roles")
		err = s.downRoles(ctx)
		if err != nil {
			return
		}
		log.Info().Msg("Down tenants")
		err = s.downTenants(ctx)

		log.Info().Msg("---Down seeded---")

		return
	})

	return
}
