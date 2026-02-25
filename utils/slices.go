package utils

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
)

const lineFeed = byte(0x0A)

var ErrSliceContainsLF = fmt.Errorf("slice contains lineFeed (0x0A)")

func JoinByteSlicesWithLF(byteSlices [][]byte) []byte {
	return bytes.Join(byteSlices, []byte{lineFeed})
}

// JoinSingleLineByteSlicesWithLF concatenates a slice of byte slices into a single byte slice, separating them with a line feed (0x0A).
// It returns an error if any of the input byte slices contains a line feed.
func JoinSingleLineByteSlicesWithLF(byteSlices [][]byte) ([]byte, error) {
	result := make([]byte, 0)
	for i := range byteSlices {
		if slices.Contains(byteSlices[i], lineFeed) {
			return []byte{}, fmt.Errorf("%w: %q", ErrSliceContainsLF, byteSlices[i])
		}
		if i > 0 {
			result = append(result, lineFeed)
		}
		result = append(result, byteSlices[i]...)
	}

	return result, nil
}

func SplitByteSliceByLF(oneDim []byte) [][]byte {
	return bytes.Split(oneDim, []byte{lineFeed})
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
