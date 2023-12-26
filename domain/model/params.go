package model

import "github.com/google/uuid"

type ActionData struct {
	ID     *uuid.UUID `json:"id"`
	UserID *uuid.UUID `json:"user_id"`
}
