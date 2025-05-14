package encoding

import (
	"github.com/blutspende/bloodlab-common/enums/encoding"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncoding_ConvertFromEncodingToUtf8(t *testing.T) {
	// Arrange
	input := []byte("űúőéóüöáßäüöë")
	enc := encoding.UTF8
	// Act
	result, err := ConvertFromEncodingToUtf8(input, enc)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "űúőéóüöáßäüöë", result)
}

func TestEncoding_ConvertFromEncodingToUtf8Win1252(t *testing.T) {
	// Arrange
	input := []byte{0xDF, 0xE4, 0xFC, 0xF6, 0xEB}
	enc := encoding.Windows1252
	// Act
	result, err := ConvertFromEncodingToUtf8(input, enc)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "ßäüöë", result)
}

func TestEncoding_ConvertFromUtf8ToEncodingWin1252(t *testing.T) {
	// Arrange
	input := "ßäüöë"
	enc := encoding.Windows1252
	// Act
	result, err := ConvertFromUtf8ToEncoding(input, enc)
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
	_, err := ConvertFromEncodingToUtf8(input, encoding.Encoding(enc))
	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, "invalid_encoding: invalid encoding", err.Error())
}
