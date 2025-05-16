package utils

import (
	"fmt"
	"slices"
	"strings"
)

func ConvertBytes2Dto1D(twoDim [][]byte) ([]byte, error) {
	result := []byte{}
	for i, row := range twoDim {
		if slices.Contains(row, '\u0000') {
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

func JoinEnumsAsString[T ~string](enumList []T, separator string) string {
	items := make([]string, len(enumList))
	for i := range enumList {
		items[i] = string(enumList[i])
	}
	return strings.Join(items, separator)
}

func Partition(totalLength int, partitionLength int, consumer func(low int, high int) error) error {
	if partitionLength <= 0 || totalLength <= 0 {
		return nil
	}
	partitions := totalLength / partitionLength
	var i int
	var err error
	for i = 0; i < partitions; i++ {
		err = consumer(i*partitionLength, i*partitionLength+partitionLength)
		if err != nil {
			return err
		}
	}
	if rest := totalLength % partitionLength; rest != 0 {
		err = consumer(i*partitionLength, i*partitionLength+rest)
		if err != nil {
			return err
		}
	}
	return err
}
