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

// 密钥  必须为32 24 16 长度的
const key = "8f95b53c3c0f598266b801a4b621423d"

// Token过期时间 6小时
// const TokenTime int64 = 6 * 60 * 60 * 180
const TokenTime int64 = 24 * 60 * 60 * 180

// Token自动登录过期时间 半年
const TokenAutoTime int64 = 24 * 60 * 60 * 180

// Token更新时间 30分钟
// const TokenUpdateTime int64 = 30 * 60

// ase解密函数
func Decrypt(decryptString *string) (string, error) {
	// 16进制转bytes
	// b, err :=hex.DecodeString(decryptString)
	//    base64的字符串转为byte  slice
	b, err := base64.StdEncoding.DecodeString(*decryptString)
	if err != nil {
		return "", err
	}
	keyByte := []byte(key)
	//  创建解码算法
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	// 获取block块长度
	blockSize := block.BlockSize()
	// 获取blockMode
	// blockMode := cipher.NewCBCDecrypter(block, commonIV)
	blockMode := cipher.NewCBCDecrypter(block, keyByte[:blockSize])
	// 创建接收数据的bytes
	origData := make([]byte, len(b))
	// origData := crypted
	// CryptBlocks加密或解密一些块。 src的长度必须是块大小的倍数。 Dst和src可能指向相同的内存。
	blockMode.CryptBlocks(origData, b)
	// origData = PKCS5UnPadding(origData)
	origData = ZeroUnPadding(origData)
	return string(origData), nil
}

// ase 加密函数
func Encryption(encryptionString string) (string, error) {
	plaintext := []byte(encryptionString)
	keyByte := []byte(key)
	//	创建加密算法
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	// 获取block块长度
	blockSize := block.BlockSize()
	// plaintext = PKCS5Padding(plaintext, blockSize)
	plaintext = ZeroPadding(plaintext, block.BlockSize())
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
	return base64.URLEncoding.EncodeToString(crypted), err
}

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
	token, err := Encryption(str)
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
	str, err := Decrypt(token)
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

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
