package utils

import (
	"github.com/google/uuid"
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

func TestTypes_UUIDToNullUUID(t *testing.T) {
	// Arrange
	input := uuid.MustParse("bfb292af-b806-4cbb-870b-b1f07e148df9")
	// Act
	result := UUIDToNullUUID(input)
	// Assert
	assert.True(t, result.Valid)
	assert.Equal(t, input, result.UUID)
}
func TestTypes_UUIDToNullUUID_Nil(t *testing.T) {
	// Arrange
	input := uuid.Nil
	// Act
	result := UUIDToNullUUID(input)
	// Assert
	assert.False(t, result.Valid)
	assert.Equal(t, input, result.UUID)
}
func TestTypes_UUIDToNullUUID_Empty(t *testing.T) {
	// Arrange
	input := uuid.UUID{}
	// Act
	result := UUIDToNullUUID(input)
	// Assert
	assert.False(t, result.Valid)
	assert.Equal(t, input, result.UUID)
}
