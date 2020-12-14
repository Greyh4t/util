package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(t []byte) string {
	h := md5.New()
	h.Write(t)
	return hex.EncodeToString(h.Sum(nil))
}
