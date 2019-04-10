package crypto

// 测试站点: https://www.keylala.cn

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
)

// Md5 ..
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

// Md5Byte ..
func Md5Byte(b []byte) string {
	m := md5.New()
	m.Write(b)
	return hex.EncodeToString(m.Sum(nil))
}

// Sha1 ..
func Sha1(str string) string {
	s := sha1.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// Sha1Byte ..
func Sha1Byte(b []byte) string {
	s := sha1.New()
	s.Write(b)
	return hex.EncodeToString(s.Sum(nil))
}

// Sha224 ..
func Sha224(str string) string {
	s := sha256.New224()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// Sha224Byte ..
func Sha224Byte(b []byte) string {
	s := sha256.New224()
	s.Write(b)
	return hex.EncodeToString(s.Sum(nil))
}

// Sha256 ..
func Sha256(str string) string {
	s := sha256.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// Sha256Byte ..
func Sha256Byte(b []byte) string {
	s := sha256.New()
	s.Write(b)
	return hex.EncodeToString(s.Sum(nil))
}

// Sha384 ..
func Sha384(str string) string {
	s := sha512.New384()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// Sha384Byte ..
func Sha384Byte(b []byte) string {
	s := sha512.New384()
	s.Write(b)
	return hex.EncodeToString(s.Sum(nil))
}

// Sha512 ..
func Sha512(str string) string {
	s := sha512.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// Sha512Byte ..
func Sha512Byte(b []byte) string {
	s := sha512.New()
	s.Write(b)
	return hex.EncodeToString(s.Sum(nil))
}

// HmacMd5 ..
func HmacMd5(str, key string) string {
	mac := hmac.New(md5.New, []byte(key))
	mac.Write([]byte(str))
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacMd5Byte ..
func HmacMd5Byte(b, key []byte) string {
	mac := hmac.New(md5.New, key)
	mac.Write(b)
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacSha1 ..
func HmacSha1(str, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(str))
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacSha1Byte ..
func HmacSha1Byte(b, key []byte) string {
	mac := hmac.New(sha1.New, key)
	mac.Write(b)
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacSha256 ..
func HmacSha256(str, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str))
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacSha256Byte ..
func HmacSha256Byte(b, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(b)
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacSha512 ..
func HmacSha512(str, key string) string {
	mac := hmac.New(sha512.New, []byte(key))
	mac.Write([]byte(str))
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacSha512Byte ..
func HmacSha512Byte(b, key []byte) string {
	mac := hmac.New(sha512.New, key)
	mac.Write(b)
	return hex.EncodeToString(mac.Sum(nil))
}

// URLEncode ..
// 相当于 JS encodeURIComponent
// "www.keylala.cn?name=alex&age=18&say=你好" -> "www.keylala.cn%3Fname%3Dalex%26age%3D18%26say%3D%E4%BD%A0%E5%A5%BD"
func URLEncode(str string) string {
	return url.QueryEscape(str)
}

// URLDecode ..
// 相当于 JS decodeURIComponent
// www.keylala.cn%3Fname%3Dalex%26age%3D18%26say%3D%E4%BD%A0%E5%A5%BD -> "www.keylala.cn?name=alex&age=18&say=你好"
func URLDecode(str string) string {
	data, err := url.QueryUnescape(str)
	if err != nil {
		return ""
	}
	return data
}

// Base64Encode base64 编码
func Base64Encode(str string) string {
	return Base64EncodeByte([]byte(str))
}

// Base64EncodeByte base64 编码
func Base64EncodeByte(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Base64Decode base64 解码
func Base64Decode(str string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// AesCBCEncodeHex ..
// key: 密钥，长度需要为 16, 24, 或者 32
// iv: 向量长度，长度必须为 16
// plaintext: 待加密的内容
// return: 16进制
func AesCBCEncodeHex(key, iv, plaintext []byte) (string, error) {
	result, err := AesCBCEncode(key, iv, plaintext)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(result), nil
}

// AesCBCEncodeBase64 ..
// key: 密钥，长度需要为 16, 24, 或者 32
// iv: 向量长度，长度必须为 16
// plaintext: 待加密的内容
// return: base64
func AesCBCEncodeBase64(key, iv, plaintext []byte) (string, error) {
	result, err := AesCBCEncode(key, iv, plaintext)
	if err != nil {
		return "", err
	}
	return Base64EncodeByte(result), nil
}

// AesCBCEncode ..
// key: 密钥，长度需要为 16, 24, 或者 32
// iv: 向量长度，长度必须为 16
// plaintext: 待加密的内容
func AesCBCEncode(key, iv, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(iv) != aes.BlockSize {
		return nil, fmt.Errorf("the length of iv must be %d", aes.BlockSize)
	}

	blockSize := block.BlockSize()
	plaintext = pkcs7Padding(plaintext, blockSize)

	ciphertext := make([]byte, len(plaintext))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

// AesCBCDecodeHex ..
// key: 密钥，长度需要为 16, 24, 或者 32
// iv: 向量长度，长度必须为 16
// ciphertext: 待解密的内容(格式为 hex)
func AesCBCDecodeHex(key, iv []byte, ciphertext string) ([]byte, error) {
	result, err := hex.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	return AesCBCDecode(key, iv, result)
}

// AesCBCDecodeBase64 ..
// key: 密钥，长度需要为 16, 24, 或者 32
// iv: 向量长度，长度必须为 16
// ciphertext: 待解密的内容(格式为 base64)
func AesCBCDecodeBase64(key, iv []byte, ciphertext string) ([]byte, error) {
	result, err := Base64Decode(ciphertext)
	if err != nil {
		return nil, err
	}
	return AesCBCDecode(key, iv, result)
}

// AesCBCDecode ..
// key: 密钥，长度需要为 16, 24, 或者 32
// iv: 向量长度，长度必须为 16
// ciphertext: 待解密的内容
func AesCBCDecode(key, iv, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(iv) != aes.BlockSize {
		return nil, fmt.Errorf("the length of iv must be %d", aes.BlockSize)
	}

	blockSize := block.BlockSize()

	// 密文大小必须是快的倍数
	if len(ciphertext)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	plaintext := make([]byte, len(ciphertext))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	return pkcs7Unpadding(plaintext), nil
}

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
