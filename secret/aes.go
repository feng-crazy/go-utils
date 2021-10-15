package secret

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strconv"
	"strings"
	"time"
)

// ase解密函数
func DecryptCBCPKCS7(key string, decryptByte []byte) ([]byte, error) {
	// 16进制转bytes
	// b, err :=hex.DecodeString(decryptString)
	//    base64的字符串转为byte  slice
	dBuf := make([]byte, base64.StdEncoding.DecodedLen(len(decryptByte)))
	n, err := base64.StdEncoding.Decode(dBuf, decryptByte)
	if err != nil {
		return nil, err
	}
	dBuf = dBuf[:n]

	keyByte := []byte(key)
	//  创建解码算法
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}
	// 获取block块长度
	blockSize := block.BlockSize()
	// 获取blockMode
	// blockMode := cipher.NewCBCDecrypter(block, commonIV)
	blockMode := cipher.NewCBCDecrypter(block, keyByte[:blockSize])
	// 创建接收数据的bytes
	origData := make([]byte, len(dBuf))
	// origData := crypted
	// CryptBlocks加密或解密一些块。 src的长度必须是块大小的倍数。 Dst和src可能指向相同的内存。
	blockMode.CryptBlocks(origData, dBuf)
	origData = PKCS7UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

// ase 加密函数
func EncryptCBCPKCS7(key string, encryptionString []byte) ([]byte, error) {
	plaintext := []byte(encryptionString)
	keyByte := []byte(key)
	//	创建加密算法
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}
	// 获取block块长度
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)
	// plaintext = ZeroPadding(plaintext, block.BlockSize())
	// plaintext = PKCS5Padding(plaintext)
	// NewCBCEncrypter返回一个BlockMode，它使用给定的Block以密码块链接模式加密。 iv的长度必须与块的大小相同。commonIV
	// blockMode := cipher.NewCBCEncrypter(block, commonIV)
	blockMode := cipher.NewCBCEncrypter(block, keyByte[:blockSize])
	// 创建接收byte  slice
	crypted := make([]byte, len(plaintext))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, plaintext)
	// 转16进制 字符串
	// return hex.EncodeToString(crypted), err
	// 加密 后的字符串转为base64的字符串 返回

	buf := make([]byte, base64.StdEncoding.EncodedLen(len(crypted)))
	base64.StdEncoding.Encode(buf, crypted)
	return buf, err
}

// ase解密函数
func DecryptionCBCPKCS7(key, decryptString string) (string, error) {
	origData, err := DecryptCBCPKCS7(key, []byte(decryptString))
	if err != nil {
		return "", err
	}
	return string(origData), nil
}

// ase 加密函数
func EncryptionCBCPKCS7(key, encryptionString string) (string, error) {
	crypted, err := EncryptCBCPKCS7(key, []byte(encryptionString))
	if err != nil {
		return "", err
	}
	return string(crypted), err
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte) []byte {
	padding := 8 - len(ciphertext)%8
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 使用PKCS7进行填充，IOS也是7
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize // 需要填充的数目
	// 只要少于256就能放到一个byte中，默认的blockSize=16(即采用16*8=128, AES-128长的密钥)
	// 最少填充1个byte，如果原文刚好是blocksize的整数倍，则再填充一个blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding) // 生成填充的文本
	return append(ciphertext, padtext...)

}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]

}

// 密钥  必须为32 24 16 长度的
const key = "8f95b53c3c0f598266b801a4b621423d"

// Token过期时间 6小时
// const TokenTime int64 = 6 * 60 * 60 * 180
const TokenTime int64 = 24 * 60 * 60 * 180

// Token自动登录过期时间 半年
const TokenAutoTime int64 = 24 * 60 * 60 * 180

// Token更新时间 30分钟
// const TokenUpdateTime int64 = 30 * 60

// 生成用户Token
func CreateUserToken(userId string, auto bool) (string, int64, error) {
	var str string
	var tokentime int64
	// 是否自动登录
	if auto { // 是
		// 根据用户ID及Token过期时间生成字符串
		tokentime = time.Now().Unix() + TokenAutoTime
		str = userId + strconv.FormatInt(tokentime, 10) + "+"
	} else { // 否
		// 根据用户ID及Token过期时间生成字符串
		tokentime = time.Now().Unix() + TokenTime
		str = userId + strconv.FormatInt(tokentime, 10)
	}

	// ASE加密生成Token
	token, err := EncryptionCBCPKCS7(key, str)
	return token, tokentime, err
}

/*
	验证Token有效时间，Token有效,判断是否需要更新Token。并返回验证状态，用户ID
	返回：1、验证状态 2、UserID 3、error
	验证状态：
		0：正常		返回用户ID
		1：验证失败	返回错误信息
		2：过期		返回用户ID
*/
func TokenValidTimeVerification(token *string) (int8, string, error) {

	// ASE解密Token
	str, err := DecryptionCBCPKCS7(key, *token)
	str = strings.Replace(str, " ", "", -1)
	if err != nil { // 解密失败
		return 0, "", err
	}

	rs := []rune(str)
	// TokenAutoTime
	// 获取用户ID
	userId := string(rs[:24])

	// 获取Token过期时间
	tokenTime, err := strconv.ParseInt(string(rs[24:34]), 10, 64)
	if err != nil { // 转换失败
		return 1, "", err
	}

	// 获取当前时间
	nowTime := time.Now().Unix()

	// 与当前时间进行对比
	if tokenTime > nowTime {
		// Token有效
		return 0, userId, err
	} else {
		// Token过期
		return 2, userId, err
	}
}
