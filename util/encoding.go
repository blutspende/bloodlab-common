package util

import (
	"bytes"
	"fmt"
	enums "github.com/blutspende/bloodlab-common/enums/encoding"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/unicode/utf32"
	"golang.org/x/text/transform"
	"io"
)

func ConvertFromEncodingToUtf8(input []byte, encoding enums.Encoding) (output string, err error) {
	enc, err := findEncoding(encoding)
	if err != nil {
		return "", err
	}
	if enc == nil {
		return string(input), nil
	}
	encoded, err := io.ReadAll(enc.NewDecoder().Reader(bytes.NewReader(input)))
	return string(encoded), err
}

func ConvertFromUtf8ToEncoding(input string, encoding enums.Encoding) (output []byte, err error) {
	enc, err := findEncoding(encoding)
	if err != nil {
		return []byte{}, err
	}
	if enc == nil {
		return []byte(input), nil
	}
	output, _, err = transform.Bytes(enc.NewEncoder(), []byte(input))
	return output, err
}

func ConvertArrayFromUtf8ToEncoding(input []string, encoding enums.Encoding) (output [][]byte, err error) {
	output = make([][]byte, len(input))
	for i, line := range input {
		output[i], err = ConvertFromUtf8ToEncoding(line, encoding)
		if err != nil {
			return nil, err
		}
	}
	return output, nil
}

func findEncoding(encoding enums.Encoding) (encoding.Encoding, error) {
	switch encoding {
	case enums.UTF8:
		return nil, nil
	case enums.UTF16:
		return unicode.UTF16(unicode.LittleEndian, unicode.UseBOM), nil
	case enums.UTF16BE:
		return unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM), nil
	case enums.UTF16LE:
		return unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM), nil
	case enums.UTF32:
		return utf32.UTF32(utf32.LittleEndian, utf32.UseBOM), nil
	case enums.UTF32BE:
		return utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM), nil
	case enums.UTF32LE:
		return utf32.UTF32(utf32.LittleEndian, utf32.IgnoreBOM), nil
	case enums.ASCII:
		return nil, nil
	case enums.Windows874:
		return charmap.Windows874, nil
	case enums.Windows1250:
		return charmap.Windows1250, nil
	case enums.Windows1251:
		return charmap.Windows1251, nil
	case enums.Windows1252:
		return charmap.Windows1252, nil
	case enums.Windows1253:
		return charmap.Windows1253, nil
	case enums.Windows1254:
		return charmap.Windows1254, nil
	case enums.Windows1255:
		return charmap.Windows1255, nil
	case enums.Windows1256:
		return charmap.Windows1256, nil
	case enums.Windows1257:
		return charmap.Windows1257, nil
	case enums.Windows1258:
		return charmap.Windows1258, nil
	case enums.DOS852:
		return charmap.CodePage852, nil
	case enums.DOS855:
		return charmap.CodePage855, nil
	case enums.DOS866:
		return charmap.CodePage866, nil
	case enums.ISO8859_1:
		return charmap.ISO8859_1, nil
	case enums.ISO8859_2:
		return charmap.ISO8859_2, nil
	case enums.ISO8859_3:
		return charmap.ISO8859_3, nil
	case enums.ISO8859_4:
		return charmap.ISO8859_4, nil
	case enums.ISO8859_5:
		return charmap.ISO8859_5, nil
	case enums.ISO8859_6:
		return charmap.ISO8859_6, nil
	case enums.ISO8859_6E:
		return charmap.ISO8859_6E, nil
	case enums.ISO8859_6I:
		return charmap.ISO8859_6I, nil
	case enums.ISO8859_7:
		return charmap.ISO8859_7, nil
	case enums.ISO8859_8:
		return charmap.ISO8859_8, nil
	case enums.ISO8859_8E:
		return charmap.ISO8859_8E, nil
	case enums.ISO8859_8I:
		return charmap.ISO8859_8I, nil
	case enums.ISO8859_9:
		return charmap.ISO8859_9, nil
	case enums.ISO8859_10:
		return charmap.ISO8859_10, nil
	case enums.ISO8859_13:
		return charmap.ISO8859_13, nil
	case enums.ISO8859_14:
		return charmap.ISO8859_14, nil
	case enums.ISO8859_15:
		return charmap.ISO8859_15, nil
	case enums.ISO8859_16:
		return charmap.ISO8859_16, nil
	case enums.IBM037:
		return charmap.CodePage037, nil
	case enums.IBM437:
		return charmap.CodePage437, nil
	case enums.IBM850:
		return charmap.CodePage850, nil
	case enums.IBM858:
		return charmap.CodePage858, nil
	case enums.IBM860:
		return charmap.CodePage860, nil
	case enums.IBM862:
		return charmap.CodePage862, nil
	case enums.IBM863:
		return charmap.CodePage863, nil
	case enums.IBM865:
		return charmap.CodePage865, nil
	case enums.IBM1047:
		return charmap.CodePage1047, nil
	case enums.IBM1140:
		return charmap.CodePage1140, nil
	case enums.KOI8R:
		return charmap.KOI8R, nil
	case enums.KOI8U:
		return charmap.KOI8U, nil
	case enums.Macintosh:
		return charmap.Macintosh, nil
	case enums.MacintoshCyrillic:
		return charmap.MacintoshCyrillic, nil
	default:
		return nil, fmt.Errorf("%s: invalid encoding", encoding)
	}
}
