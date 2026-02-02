package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagination_TotalPages_Whole(t *testing.T) {
	// Arrange
	totalCount := 10
	pageSize := 5
	// Act
	result := TotalPages(totalCount, pageSize)
	// Assert
	assert.Equal(t, 2, result)
}
func TestPagination_TotalPages_Rounded(t *testing.T) {
	// Arrange
	totalCount := 10
	pageSize := 4
	// Act
	result := TotalPages(totalCount, pageSize)
	// Assert
	assert.Equal(t, 3, result)
}

func TestPagination_StandardisePagination_WrongPageSize(t *testing.T) {
	// Arrange
	input := Pagination{
		PageSize: 33,
		Page:     1,
	}
	// Act
	result := StandardisePagination(input)
	// Assert
	assert.Equal(t, 25, result.PageSize)
	assert.Equal(t, 1, result.Page)
}
func TestPagination_StandardisePagination_WrongPage(t *testing.T) {
	// Arrange
	input := Pagination{
		PageSize: 50,
		Page:     -1,
	}
	// Act
	result := StandardisePagination(input)
	// Assert
	assert.Equal(t, 50, result.PageSize)
	assert.Equal(t, 0, result.Page)
}
func TestPagination_StandardisePagination_BigPageSize(t *testing.T) {
	// Arrange
	input := Pagination{
		PageSize: 110,
		Page:     0,
	}
	// Act
	result := StandardisePagination(input)
	// Assert
	assert.Equal(t, MaxSafeInt, result.PageSize)
	assert.Equal(t, 0, result.Page)
}
func TestPagination_StandardisePagination_ZeroSize(t *testing.T) {
	// Arrange
	input := Pagination{
		PageSize: 0,
		Page:     0,
	}
	// Act
	result := StandardisePagination(input)
	// Assert
	assert.Equal(t, 0, result.PageSize)
	assert.Equal(t, 0, result.Page)
}
