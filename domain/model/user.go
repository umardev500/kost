package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/kost/constants"
)

type User struct {
	ID         uuid.UUID        `json:"id" db:"id"`
	TenantID   *string          `json:"tenant_id,omitempty" db:"tenant_id"`
	Email      string           `json:"email" db:"email"`
	Username   string           `json:"username" db:"username"`
	Password   string           `json:"password" db:"password"`
	Status     constants.Status `json:"status" db:"status"`
	CreatedAt  time.Time        `json:"created_at" db:"created_at"`
	CreatedBy  *uuid.UUID       `json:"created_by,omitempty" db:"created_by"`
	UpdatedAt  *time.Time       `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy  *string          `json:"updated_by,omitempty" db:"updated_by"`
	DeletedAt  *time.Time       `json:"deleted_at,omitempty" db:"deleted_at"`
	DeletedBy  *string          `json:"deleted_by,omitempty" db:"deleted_by"`
	DocVersion int              `json:"doc_version" db:"doc_version"`
}

type UserFindParams struct {
	ID       *uuid.UUID `json:"id"`
	Username *string    `json:"username"`
	Search   *string    `json:"search"`
}

type UserFilter struct {
	Status constants.Status `json:"status" validate:"omitempty,oneof=active inactive"`
}

type UserFind struct {
	Params  UserFindParams `json:"params"`
	Filters UserFilter     `json:"filters"`
}
