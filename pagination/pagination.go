package pagination

type Pagination struct {
	PageSize  int    `form:"pageSize" json:"pageSize" example:"25"`
	Page      int    `form:"page" json:"page" example:"1"`
	Direction string `form:"direction" json:"direction" example:"ascending"`
	Sort      string `form:"sort" json:"sort" example:"code"`
}

type PaginationDTO struct {
	PageSize   int `json:"pageSize" example:"25"`
	Page       int `json:"currentPage" example:"1"`
	TotalCount int `json:"totalCount" example:"40"`
	TotalPages int `json:"totalPages" example:"2"`
}

type FilteredPaginatedQuery struct {
	Pagination
	SearchTerm *string `form:"search" json:"search" example:"PLASMA"`
}

// Helper functions

func TotalPages(totalCount, pageSize int) int {
	totalPages := 1
	if totalCount > 0 && pageSize > 0 {
		totalPages = totalCount / pageSize
		if totalCount%pageSize != 0 {
			totalPages++
		}
	}
	return totalPages
}

func NewPaginationDTO(pageSize, page, totalCount int) PaginationDTO {
	return PaginationDTO{
		PageSize:   pageSize,
		Page:       page,
		TotalCount: totalCount,
		TotalPages: TotalPages(totalCount, pageSize),
	}
}

// Standardization

const MaxSafeInt = 9007199254740991

var StandardPageSizes = []int{25, 50, 100}
var ValidPageSizes = []int{0, 25, 50, 100, MaxSafeInt}

func StandardisePagination(page Pagination) Pagination {
	if page.Page < 0 {
		page.Page = 0
	}
	if page.PageSize < 0 {
		page.PageSize = 0
	} else if page.PageSize > 100 {
		page.PageSize = MaxSafeInt
	} else if page.PageSize != 0 && page.PageSize != 25 && page.PageSize != 50 && page.PageSize != 100 {
		page.PageSize = 25
	}
	return page
}
