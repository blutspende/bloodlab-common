package encoding

import (
	"bytes"
	"io"

	"golang.org/x/text/transform"
)

func ConvertFromEncodingToUTF8(input []byte, encoding Encoding) (output string, err error) {
	enc, err := encoding.GetEncoding()
	if err != nil {
		return "", err
	}
	if enc == nil {
		return string(input), nil
	}
	encoded, err := io.ReadAll(enc.NewDecoder().Reader(bytes.NewReader(input)))
	return string(encoded), err
}

func ConvertFromUTF8ToEncoding(input string, encoding Encoding) (output []byte, err error) {
	enc, err := encoding.GetEncoding()
	if err != nil {
		return []byte{}, err
	}
	if enc == nil {
		return []byte(input), nil
	}
	output, _, err = transform.Bytes(enc.NewEncoder(), []byte(input))
	return output, err
}

func ConvertArrayFromUTF8ToEncoding(input []string, encoding Encoding) (output [][]byte, err error) {
	output = make([][]byte, len(input))
	for i, line := range input {
		output[i], err = ConvertFromUTF8ToEncoding(line, encoding)
		if err != nil {
			return nil, err
		}
	}
	return output, nil
}
