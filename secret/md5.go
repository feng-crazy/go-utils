package secret

import (
	"crypto/md5"
	"encoding/hex"
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

func Md5EnCode(string string) string {
	h := md5.New()
	h.Write([]byte(string)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}
