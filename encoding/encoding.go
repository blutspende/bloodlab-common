package encoding

import (
	"fmt"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/unicode/utf32"
)

type Encoding string

const UTF8 Encoding = "UTF8"
const UTF16 Encoding = "UTF16"
const UTF16BE Encoding = "UTF16BE"
const UTF16LE Encoding = "UTF16LE"
const UTF32 Encoding = "UTF32"
const UTF32BE Encoding = "UTF32BE"
const UTF32LE Encoding = "UTF32LE"
const ASCII Encoding = "ASCII"
const Windows874 Encoding = "Windows-874"
const Windows1250 Encoding = "Windows1250"
const Windows1251 Encoding = "Windows1251"
const Windows1252 Encoding = "Windows1252"
const Windows1253 Encoding = "Windows1253"
const Windows1254 Encoding = "Windows1254"
const Windows1255 Encoding = "Windows1255"
const Windows1256 Encoding = "Windows1256"
const Windows1257 Encoding = "Windows1257"
const Windows1258 Encoding = "Windows1258"
const DOS852 Encoding = "DOS852"
const DOS855 Encoding = "DOS855"
const DOS866 Encoding = "DOS866"
const ISO8859_1 Encoding = "ISO8859-1"
const ISO8859_2 Encoding = "ISO8859-2"
const ISO8859_3 Encoding = "ISO8859-3"
const ISO8859_4 Encoding = "ISO8859-4"
const ISO8859_5 Encoding = "ISO8859-5"
const ISO8859_6 Encoding = "ISO8859-6"
const ISO8859_6E Encoding = "ISO8859-6E"
const ISO8859_6I Encoding = "ISO8859-6I"
const ISO8859_7 Encoding = "ISO8859-7"
const ISO8859_8 Encoding = "ISO8859-8"
const ISO8859_8E Encoding = "ISO8859-8E"
const ISO8859_8I Encoding = "ISO8859-8I"
const ISO8859_9 Encoding = "ISO8859-9"
const ISO8859_10 Encoding = "ISO8859-10"
const ISO8859_13 Encoding = "ISO8859-13"
const ISO8859_14 Encoding = "ISO8859-14"
const ISO8859_15 Encoding = "ISO8859-15"
const ISO8859_16 Encoding = "ISO8859-16"
const IBM037 Encoding = "IBM037"
const IBM437 Encoding = "IBM437"
const IBM850 Encoding = "IBM850"
const IBM858 Encoding = "IBM858"
const IBM860 Encoding = "IBM860"
const IBM862 Encoding = "IBM862"
const IBM863 Encoding = "IBM863"
const IBM865 Encoding = "IBM865"
const IBM1047 Encoding = "IBM1047"
const IBM1140 Encoding = "IBM1140"
const KOI8R Encoding = "KOI8-R"
const KOI8U Encoding = "KOI8-U"
const Macintosh Encoding = "Macintosh"
const MacintoshCyrillic Encoding = "MacintoshCyrillic"

func (encoding Encoding) GetEncoding() (encoding.Encoding, error) {
	switch encoding {
	case UTF8:
		return nil, nil
	case UTF16:
		return unicode.UTF16(unicode.LittleEndian, unicode.UseBOM), nil
	case UTF16BE:
		return unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM), nil
	case UTF16LE:
		return unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM), nil
	case UTF32:
		return utf32.UTF32(utf32.LittleEndian, utf32.UseBOM), nil
	case UTF32BE:
		return utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM), nil
	case UTF32LE:
		return utf32.UTF32(utf32.LittleEndian, utf32.IgnoreBOM), nil
	case ASCII:
		return nil, nil
	case Windows874:
		return charmap.Windows874, nil
	case Windows1250:
		return charmap.Windows1250, nil
	case Windows1251:
		return charmap.Windows1251, nil
	case Windows1252:
		return charmap.Windows1252, nil
	case Windows1253:
		return charmap.Windows1253, nil
	case Windows1254:
		return charmap.Windows1254, nil
	case Windows1255:
		return charmap.Windows1255, nil
	case Windows1256:
		return charmap.Windows1256, nil
	case Windows1257:
		return charmap.Windows1257, nil
	case Windows1258:
		return charmap.Windows1258, nil
	case DOS852:
		return charmap.CodePage852, nil
	case DOS855:
		return charmap.CodePage855, nil
	case DOS866:
		return charmap.CodePage866, nil
	case ISO8859_1:
		return charmap.ISO8859_1, nil
	case ISO8859_2:
		return charmap.ISO8859_2, nil
	case ISO8859_3:
		return charmap.ISO8859_3, nil
	case ISO8859_4:
		return charmap.ISO8859_4, nil
	case ISO8859_5:
		return charmap.ISO8859_5, nil
	case ISO8859_6:
		return charmap.ISO8859_6, nil
	case ISO8859_6E:
		return charmap.ISO8859_6E, nil
	case ISO8859_6I:
		return charmap.ISO8859_6I, nil
	case ISO8859_7:
		return charmap.ISO8859_7, nil
	case ISO8859_8:
		return charmap.ISO8859_8, nil
	case ISO8859_8E:
		return charmap.ISO8859_8E, nil
	case ISO8859_8I:
		return charmap.ISO8859_8I, nil
	case ISO8859_9:
		return charmap.ISO8859_9, nil
	case ISO8859_10:
		return charmap.ISO8859_10, nil
	case ISO8859_13:
		return charmap.ISO8859_13, nil
	case ISO8859_14:
		return charmap.ISO8859_14, nil
	case ISO8859_15:
		return charmap.ISO8859_15, nil
	case ISO8859_16:
		return charmap.ISO8859_16, nil
	case IBM037:
		return charmap.CodePage037, nil
	case IBM437:
		return charmap.CodePage437, nil
	case IBM850:
		return charmap.CodePage850, nil
	case IBM858:
		return charmap.CodePage858, nil
	case IBM860:
		return charmap.CodePage860, nil
	case IBM862:
		return charmap.CodePage862, nil
	case IBM863:
		return charmap.CodePage863, nil
	case IBM865:
		return charmap.CodePage865, nil
	case IBM1047:
		return charmap.CodePage1047, nil
	case IBM1140:
		return charmap.CodePage1140, nil
	case KOI8R:
		return charmap.KOI8R, nil
	case KOI8U:
		return charmap.KOI8U, nil
	case Macintosh:
		return charmap.Macintosh, nil
	case MacintoshCyrillic:
		return charmap.MacintoshCyrillic, nil
	default:
		return nil, fmt.Errorf("%s: invalid encoding", encoding)
	}
}
