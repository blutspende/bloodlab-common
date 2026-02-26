package pagination

type PaginatedQuery struct {
	PageSize  int    `form:"pageSize" json:"pageSize" minimum:"0" default:"25" example:"25"`
	Page      int    `form:"page" json:"page" minimum:"0" default:"0" example:"1"`
	Direction string `form:"direction" json:"direction" example:"ascending"`
	Sort      string `form:"sort" json:"sort" example:"code"`
}
type FilteredPaginatedQuery struct {
	PaginatedQuery
	SearchTerm *string `form:"search" json:"search" example:"PLASMA"`
}

type PaginatedResponse struct {
	PageSize    int `json:"pageSize" example:"25"`
	CurrentPage int `json:"currentPage" example:"1"`
	TotalCount  int `json:"totalCount" example:"40"`
	TotalPages  int `json:"totalPages" example:"2"`
}

// Helper functions

func (p PaginatedQuery) IsPaged() bool {
	return p.PageSize > 0
}
func (p PaginatedQuery) IsUnPaged() bool {
	return p.PageSize == 0
}

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
func NewPaginatedResponse(pageSize, currentPage, totalCount int) PaginatedResponse {
	return PaginatedResponse{
		PageSize:    pageSize,
		CurrentPage: currentPage,
		TotalCount:  totalCount,
		TotalPages:  TotalPages(totalCount, pageSize),
	}
}

// Standardization

const MaxSafeInt = 9007199254740991

var StandardPageSizes = []int{25, 50, 100}
var ValidPageSizes = []int{0, 25, 50, 100, MaxSafeInt}

func StandardisePaginatedQuery(page PaginatedQuery) PaginatedQuery {
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
