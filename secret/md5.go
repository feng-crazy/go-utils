package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

// md5 加密
func Md5Encryption(encryptionString string) string {
	h := md5.New()
	_, _ = io.WriteString(h, encryptionString)

	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	return pwmd5
}
