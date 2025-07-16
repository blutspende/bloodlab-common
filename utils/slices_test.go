package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlices_ConvertBytes2Dto1D(t *testing.T) {
	// Arrange
	input := [][]byte{
		[]byte("first"),
		[]byte("second"),
		[]byte("third"),
	}
	// Act
	result := ConvertBytes2Dto1D(input)
	// Assert
	assert.Equal(t, []byte("first\nsecond\nthird"), result)
}
func TestSlices_ConvertBytes2Dto1D_WithEmpty(t *testing.T) {
	// Arrange
	input := [][]byte{
		[]byte("first"),
		[]byte(""),
		[]byte("third"),
	}
	// Act
	result := ConvertBytes2Dto1D(input)
	// Assert
	assert.Equal(t, []byte("first\n\nthird"), result)
}
func TestSlices_ConvertBytes2Dto1D_WithUTF8(t *testing.T) {
	// Arrange
	input := [][]byte{
		[]byte("first"),
		[]byte("űúőéáéí"),
		[]byte("third"),
	}
	// Act
	result := ConvertBytes2Dto1D(input)
	// Assert
	assert.Equal(t, []byte("first\nűúőéáéí\nthird"), result)
}

func TestSlices_ConvertBytes2Dto1DWithCheck(t *testing.T) {
	// Arrange
	input := [][]byte{
		[]byte("first"),
		[]byte("second"),
		[]byte("third"),
	}
	// Act
	result, err := ConvertBytes2Dto1DWithCheck(input)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, []byte("first\nsecond\nthird"), result)
}
func TestSlices_ConvertBytes2Dto1DWithCheck_Error(t *testing.T) {
	// Arrange
	input := [][]byte{
		[]byte("first"),
		[]byte("sec\nond"),
		[]byte("third"),
	}
	// Act
	result, err := ConvertBytes2Dto1DWithCheck(input)
	// Assert
	assert.Error(t, err)
	assert.Equal(t, []byte{}, result)
}

func TestSlices_ConvertBytes1Dto2D(t *testing.T) {
	// Arrange
	input := []byte("first\nsecond\nthird")
	// Act
	result := ConvertBytes1Dto2D(input)
	// Assert
	assert.Len(t, result, 3)
	assert.Equal(t, []byte("first"), result[0])
	assert.Equal(t, []byte("second"), result[1])
	assert.Equal(t, []byte("third"), result[2])
}
