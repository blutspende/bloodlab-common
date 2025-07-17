package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testString = "test"

func TestTypes_StringToPointer(t *testing.T) {
	// Arrange
	input := "test"
	// Act
	result := StringToPointer(input)
	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, "test", *result)
}
func TestTypes_StringToPointer_Literal(t *testing.T) {
	// Act
	result := StringToPointer("test")
	// Assert
	assert.Equal(t, "test", *result)
}
func TestTypes_StringToPointer_Constant(t *testing.T) {
	// Act
	result := StringToPointer(testString)
	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, "test", *result)
}
func TestTypes_StringToPointer_Empty(t *testing.T) {
	// Act
	result := StringToPointer("")
	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, "", *result)
}

func TestTypes_StringToPointerWithNil(t *testing.T) {
	// Arrange
	input := "test"
	// Act
	result := StringToPointerWithNil(input)
	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, "test", *result)
}
func TestTypes_StringToPointerWithNil_Literal(t *testing.T) {
	// Act
	result := StringToPointerWithNil("test")
	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, "test", *result)
}
func TestTypes_StringToPointerWithNil_Constant(t *testing.T) {
	// Act
	result := StringToPointerWithNil(testString)
	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, "test", *result)
}
func TestTypes_StringToPointerWithNil_Empty(t *testing.T) {
	// Act
	result := StringToPointerWithNil("")
	// Assert
	assert.Nil(t, result)
}
