package codef

import (
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/unicode/utf32"
)

func UTF16Convert(b []byte) string {
	// decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	// bs2, err := decoder.Bytes(b[:])
	// if err != nil{
	// 	return ""
	// }
	// return string(bs2)

	decoder := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewDecoder()
	bs2, err := decoder.Bytes(b[:])
	if err != nil {
		return ""
	}
	return string(bs2)
}

func UTF8Convert(b []byte) string {

	// d := b
	// rs := make([]rune, 0)
	// for len(d) > 0 {
	// 	r, size := utf16.DecodeRune(d)
	// 	fmt.Printf("%c %v\n", r, size)
	// 	rs = append(rs, r)
	// 	b = b[size:]
	// }
	// s := string(rs)
	// return s

	// decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	// bs2, err := decoder.Bytes(b[:])
	// if err != nil{
	// 	return ""
	// }
	// return string(bs2)

	decoder := unicode.UTF8.NewDecoder()
	bs2, err := decoder.Bytes(b[:])
	if err != nil {
		return ""
	}
	return string(bs2)
}

func UTF32Convert(b []byte) string {
	// decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	// bs2, err := decoder.Bytes(b[:])
	// if err != nil{
	// 	return ""
	// }
	// return string(bs2)

	decoder := utf32.UTF32(utf32.LittleEndian, utf32.IgnoreBOM).NewDecoder()
	bs2, err := decoder.Bytes(b[:])
	if err != nil {
		return ""
	}
	return string(bs2)

	// decoder := utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM).NewDecoder()
	// bs2, err := decoder.Bytes(b[:])
	// if err != nil{
	// 	return ""
	// }
	// return string(bs2)
}
