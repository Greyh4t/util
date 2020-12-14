package stringutil

import (
	"strings"
	"unsafe"
)

// LeftPad 靠左补全字符串
func LeftPad(str, pad string, length int) string {
	return strings.Repeat(pad, length-len(str)) + str
}

// RightPad 靠右补全字符串
func RightPad(str, pad string, length int) string {
	return str + strings.Repeat(pad, length-len(str))
}

// SubStr 截取字符串的子字符串
func SubStr(s string, start, length int) string {
	r := []rune(s)
	l := len(r)

	start = start % l
	if start < 0 {
		start = l + start
	}

	length = length % l
	if length < 0 {
		length = l + length
	}

	end := start + length
	if end <= l {
		return string(r[start:end])
	}

	return ""
}

// SplitList 将数组按照per值分割为多个
func SplitList(list []string, per int) [][]string {
	var (
		rlines [][]string
		count  = len(list) / per
	)

	if len(list)%per != 0 {
		count++
	}

	for i := 0; i < count; i++ {
		start := i * per
		end := (i + 1) * per
		if end > len(list) {
			end = len(list)
		}
		rlines = append(rlines, list[start:end])
	}

	return rlines
}

// StringBytes 不通过copy将string转为[]byte
func StringBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// BytesString 不通过copy将[]byte转为string
func BytesString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
