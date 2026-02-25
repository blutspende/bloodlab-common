package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlices_JoinByteSlicesWithLF(t *testing.T) {
	// Arrange
	input := [][]byte{
		[]byte("first"),
		[]byte("second"),
		[]byte("third"),
	}
	// Act
	result := JoinByteSlicesWithLF(input)
	// Assert
	assert.Equal(t, []byte("first\nsecond\nthird"), result)
}
func TestSlices_JoinByteSlicesWithLF_WithEmpty(t *testing.T) {
	// Arrange
	input := [][]byte{
		[]byte("first"),
		[]byte(""),
		[]byte("third"),
	}
	// Act
	result := JoinByteSlicesWithLF(input)
	// Assert
	assert.Equal(t, []byte("first\n\nthird"), result)
}
func TestSlices_JoinByteSlicesWithLF_WithUTF8(t *testing.T) {
	// Arrange
	input := [][]byte{
		[]byte("first"),
		[]byte("űúőéáéí"),
		[]byte("third"),
	}
	// Act
	result := JoinByteSlicesWithLF(input)
	// Assert
	assert.Equal(t, []byte("first\nűúőéáéí\nthird"), result)
}

func TestSlices_JoinSingleLineByteSlicesWithLF(t *testing.T) {
	// Arrange
	input := [][]byte{
		[]byte("first"),
		[]byte("second"),
		[]byte("third"),
	}
	// Act
	result, err := JoinSingleLineByteSlicesWithLF(input)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, []byte("first\nsecond\nthird"), result)
}
func TestSlices_JoinSingleLineByteSlicesWithLF_Error(t *testing.T) {
	// Arrange
	input := [][]byte{
		[]byte("first"),
		[]byte("sec\nond"),
		[]byte("third"),
	}
	// Act
	result, err := JoinSingleLineByteSlicesWithLF(input)
	// Assert
	assert.ErrorIs(t, err, ErrSliceContainsLF)
	assert.Equal(t, []byte{}, result)
}

func TestSlices_SplitByteSliceByLF(t *testing.T) {
	// Arrange
	input := []byte("first\nsecond\nthird")
	// Act
	result := SplitByteSliceByLF(input)
	// Assert
	assert.Len(t, result, 3)
	assert.Equal(t, []byte("first"), result[0])
	assert.Equal(t, []byte("second"), result[1])
	assert.Equal(t, []byte("third"), result[2])
}
