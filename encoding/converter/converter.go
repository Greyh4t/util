package converter

import (
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func GBKToUTF8(in []byte) []byte {
	if !utf8.Valid(in) {
		out, err := simplifiedchinese.GB18030.NewDecoder().Bytes(in)
		if err == nil {
			return out
		}
	}
	return in
}

func UTF8ToGBK(in []byte) []byte {
	if utf8.Valid(in) {
		out, err := simplifiedchinese.GB18030.NewEncoder().Bytes(in)
		if err == nil {
			return out
		}
	}

	return in
}
