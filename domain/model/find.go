package model

type PaginationParams struct {
	PageNum  int64 `json:"page"`
	PageSize int64 `json:"size"`
}
