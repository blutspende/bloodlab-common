package encoding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncoding_ConvertFromEncodingToUtf8(t *testing.T) {
	// Arrange
	input := []byte("űúőéóüöáßäüöë")
	enc := UTF8
	// Act
	result, err := ConvertFromEncodingToUTF8(input, enc)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "űúőéóüöáßäüöë", result)
}

func TestEncoding_ConvertFromEncodingToUtf8Win1252(t *testing.T) {
	// Arrange
	input := []byte{0xDF, 0xE4, 0xFC, 0xF6, 0xEB}
	enc := Windows1252
	// Act
	result, err := ConvertFromEncodingToUTF8(input, enc)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "ßäüöë", result)
}

func TestEncoding_ConvertFromUtf8ToEncodingWin1252(t *testing.T) {
	// Arrange
	input := "ßäüöë"
	enc := Windows1252
	// Act
	result, err := ConvertFromUTF8ToEncoding(input, enc)
	// Assert
	assert.Nil(t, err)
	expected := []byte{0xDF, 0xE4, 0xFC, 0xF6, 0xEB}
	assert.Equal(t, expected, result)
}

func TestEncoding_ConvertFromEncodingToUtf8InvalidEncoding(t *testing.T) {
	// Arrange
	input := []byte("invalid encoding")
	enc := "invalid_encoding"
	// Act
	_, err := ConvertFromEncodingToUTF8(input, Encoding(enc))
	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, "invalid_encoding: invalid encoding", err.Error())
}
