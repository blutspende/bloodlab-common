package utils

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
)

var byteSeparator = byte(0x0A) // ASCII Line Feed (LF) character

func ConvertBytes2Dto1D(twoDim [][]byte) []byte {
	return bytes.Join(twoDim, []byte{byteSeparator})
}
func ConvertBytes2Dto1DWithCheck(twoDim [][]byte) ([]byte, error) {
	result := []byte{}
	for i, row := range twoDim {
		if slices.Contains(row, byteSeparator) {
			return []byte{}, fmt.Errorf("2D byte array contains invalid separator character LF (0x0A)")
		}
		if i > 0 {
			result = append(result, byteSeparator)
		}
		result = append(result, row...)
	}
	return result, nil
}
func ConvertBytes1Dto2D(oneDim []byte) [][]byte {
	return bytes.Split(oneDim, []byte{byteSeparator})
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
