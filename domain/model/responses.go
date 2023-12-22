package model

import "github.com/google/uuid"

type Payload struct {
	Data       interface{} `json:"data,omitempty" swaggertype:"object"`
	PageNum    int64       `json:"page_num,omitempty"`
	PageSize   int64       `json:"page_size,omitempty"`
	PagesTotal int64       `json:"pages_total,omitempty"`
}

type OK struct {
	Code    int      `json:"code"`
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Payload *Payload `json:"payload,omitempty" swaggerignore:"true"`
}

type Err struct {
	ID      *uuid.UUID  `json:"id,omitempty"`
	Code    int         `json:"code"`
	Success bool        `json:"success" example:"false"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail,omitempty" swaggerignore:"true"`
}
