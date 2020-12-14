package xor

// XOR 使用异或算法加密或解密字符串
func XOR(src string, key string) string {
	r := ""
	var x int
	var srcR = []rune(src)
	for i := 0; i < len(srcR); i++ {
		x = int(srcR[i])
		for j := 0; j < len(key); j++ {
			x ^= int(key[j])
		}
		r += string(rune(x))
	}
	return r
}
