package randutil

import (
	"math/rand"
)

// Seed 设置随机种子
func Seed(seed int64) {
	rand.Seed(seed)
}

// Num 获取min与max之间的随机数
func Num(min, max int) int {
	if min == max {
		return min
	}

	if min > max {
		min, max = max, min
	}

	return rand.Intn(max-min+1) + min
}

// Str 获取长度为n的随机字符串，若letters为空，则使用默认值
func Str(n int, letters string) string {
	if letters == "" {
		letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
