package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

// md5 加密
func Md5Encryption(encryptionString string) string {
	h := md5.New()
	io.WriteString(h, encryptionString)

	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	/*
		//salt1+MD5 pwd+salt2
		io.WriteString(h, saltKey1)
		io.WriteString(h, pwmd5)
		io.WriteString(h, saltKey2)

		last := fmt.Sprintf("%x", h.Sum(nil))
	*/
	return pwmd5
}
