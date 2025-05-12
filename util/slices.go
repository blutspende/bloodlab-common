package util

import "fmt"

func SliceContains[V comparable](search V, data []V) bool {
	for _, value := range data {
		if value == search {
			return true
		}
	}
	return false
}

func ConvertBytes2Dto1D(twoDim [][]byte) ([]byte, error) {
	result := []byte{}
	for i, row := range twoDim {
		if SliceContains('\u0000', row) {
			return []byte{}, fmt.Errorf("message contains invalid characters")
		}
		if i > 0 {
			result = append(result, 0)
		}
		result = append(result, row...)
	}
	return result, nil
}

func ConvertBytes1Dto2D(oneDim []byte) [][]byte {
	result := [][]byte{}
	startByte := 0
	for i := 0; i < len(oneDim); i++ {
		if oneDim[i] == 0 {
			result = append(result, oneDim[startByte:i])
			startByte = i + 1
		}
	}
	if startByte < len(oneDim) {
		result = append(result, oneDim[startByte:])
	}
	return result
}
