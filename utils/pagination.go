package utils

import (
	"github.com/umardev500/kost/domain/model"
)

func GetOffset(params model.PaginationParams) int64 {
	offset := (params.PageNum - 1) * params.PageSize
	return offset
}

// GetPagesTotal get page total
func GetPagesTotal(countTotal, pageSize int64) int64 {
	var pagesTotal int64 = countTotal / pageSize
	remainder := (countTotal % pageSize)
	if remainder != 0 {
		pagesTotal = 1
	}

	return pagesTotal
}
