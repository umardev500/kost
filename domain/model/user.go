package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/kost/constants"
)

type User struct {
	ID         uuid.UUID  `json:"id"`
	TenantID   *string    `json:"tenant_id,omitempty"`
	Email      string     `json:"email"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  *uuid.UUID `json:"created_by,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	UpdatedBy  *string    `json:"updated_by,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	DeletedBy  *string    `json:"deleted_by,omitempty"`
	DocVersion int        `json:"doc_version"`
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
	Filters UserFilter
}
