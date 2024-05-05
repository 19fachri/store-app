package utils

import (
	"math"

	"github.com/19fachri/store-app/internal/store_server/servicemodels"
)

type PaginationUtil struct {
	offset int
	limit  int
}

func NewPagination(pageNo int, perPage int) PaginationUtil {
	return PaginationUtil{limit: perPage, offset: (pageNo - 1) * perPage}
}

func (pagination PaginationUtil) GetOffset() int {
	return pagination.offset
}

func (pagination PaginationUtil) GetLimit() int {
	return pagination.limit
}

func (pagination PaginationUtil) GetPageNo() int {
	return pagination.GetOffset() / pagination.GetLimit()
}

func (pagination PaginationUtil) GeneratePaginatedData(data interface{}, totalData int) servicemodels.Pagination {
	totalPage := uint(1)
	nextPage := false

	if pagination.GetLimit() > 0 {
		totalPage = uint(math.Ceil(float64(totalData) / float64(pagination.GetLimit())))
		nextPage = totalData > (pagination.GetLimit() + pagination.GetOffset())
	}

	return servicemodels.Pagination{
		Data:      data,
		TotalPage: totalPage,
		NextPage:  nextPage,
	}
}
