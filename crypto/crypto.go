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

// Md5 对指定内容生成信息摘要
// eg: Md5("123456") -> e10adc3949ba59abbe56e057f20f883e
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

// Md5Byte 对指定内容生成信息摘要
func Md5Byte(b []byte) string {
	m := md5.New()
	m.Write(b)
	return hex.EncodeToString(m.Sum(nil))
}

// Sha1 对指定内容生成信息摘要
// eg: Sha1("123456") -> 7c4a8d09ca3762af61e59520943dc26494f8941b
func Sha1(str string) string {
	s := sha1.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// Sha1Byte 对指定内容生成信息摘要
func Sha1Byte(b []byte) string {
	s := sha1.New()
	s.Write(b)
	return hex.EncodeToString(s.Sum(nil))
}

// Sha224 对指定内容生成信息摘要
// eg: Sha224("123456") -> f8cdb04495ded47615258f9dc6a3f4707fd2405434fefc3cbf4ef4e6
func Sha224(str string) string {
	s := sha256.New224()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// Sha224Byte 对指定内容生成信息摘要
func Sha224Byte(b []byte) string {
	s := sha256.New224()
	s.Write(b)
	return hex.EncodeToString(s.Sum(nil))
}

// Sha256 对指定内容生成信息摘要
// eg: Sha256("123456") -> 8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92
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

// Sha384 对指定内容生成信息摘要
// eg: Sha384("123456")
// 	-> 0a989ebc4a77b56a6e2bb7b19d995d185ce44090c13e2984b7ecc6d446d4b61ea9991b76a4c2f04b1
//     b4d244841449454
func Sha384(str string) string {
	s := sha512.New384()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// Sha384Byte 对指定内容生成信息摘要
func Sha384Byte(b []byte) string {
	s := sha512.New384()
	s.Write(b)
	return hex.EncodeToString(s.Sum(nil))
}

// Sha512 对指定内容生成信息摘要
// eg: Sha512("123456")
// 	-> ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c
// 	   17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413
func Sha512(str string) string {
	s := sha512.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// Sha512Byte 对指定内容生成信息摘要
func Sha512Byte(b []byte) string {
	s := sha512.New()
	s.Write(b)
	return hex.EncodeToString(s.Sum(nil))
}

// HmacMd5 对指定内容生成信息摘要
// eg: HmacMd5("123456", "abcdef")
// 	-> c6bdcc80c381536a3e85f2ee5f71cebb
func HmacMd5(str, key string) string {
	mac := hmac.New(md5.New, []byte(key))
	mac.Write([]byte(str))
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacMd5Byte 对指定内容生成信息摘要
func HmacMd5Byte(b, key []byte) string {
	mac := hmac.New(md5.New, key)
	mac.Write(b)
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacSha1 对指定内容生成信息摘要
// eg: HmacSha1("123456", "abcdef")
//	-> b8466fbb9634771d25d8ddd1242484bdb748b179
func HmacSha1(str, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(str))
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacSha1Byte 对指定内容生成信息摘要
func HmacSha1Byte(b, key []byte) string {
	mac := hmac.New(sha1.New, key)
	mac.Write(b)
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacSha256 对指定内容生成信息摘要
// eg: HmacSha256("123456", "abcdef")
//	-> ec4a11a5568e5cfdb5fbfe7152e8920d7bad864a0645c57fe49046a3e81ec91d
func HmacSha256(str, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str))
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacSha256Byte 对指定内容生成信息摘要
func HmacSha256Byte(b, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(b)
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacSha512 对指定内容生成信息摘要
// eg: HmacSha512("123456", "abcdef")
//	-> 130a4caafb11b798dd7528628d21f742feaad266e66141cc2ac003f0e6437cb57
// 		49245af8a3018d354e4b55e14703a5966808438afe4aae516d2824b014b5902
func HmacSha512(str, key string) string {
	mac := hmac.New(sha512.New, []byte(key))
	mac.Write([]byte(str))
	return hex.EncodeToString(mac.Sum(nil))
}

// HmacSha512Byte 对指定内容生成信息摘要
func HmacSha512Byte(b, key []byte) string {
	mac := hmac.New(sha512.New, key)
	mac.Write(b)
	return hex.EncodeToString(mac.Sum(nil))
}

// URLEncode ..
// 相当于 JS encodeURIComponent
// eg: URLEncode("www.keylala.cn?name=alex&age=18&say=你好")
//  -> "www.keylala.cn%3Fname%3Dalex%26age%3D18%26say%3D%E4%BD%A0%E5%A5%BD"
func URLEncode(str string) string {
	return url.QueryEscape(str)
}

// URLDecode ..
// 相当于 JS decodeURIComponent
// eg: URLDecode(www.keylala.cn%3Fname%3Dalex%26age%3D18%26say%3D%E4%BD%A0%E5%A5%BD)
//  -> "www.keylala.cn?name=alex&age=18&say=你好"
func URLDecode(str string) string {
	data, err := url.QueryUnescape(str)
	if err != nil {
		return ""
	}
	return data
}

// Base64Encode base64 编码
// eg: Base64Encode("https://www.keylala.cn/json?str=hello world")
// 	-> aHR0cHM6Ly93d3cua2V5bGFsYS5jbi9qc29uP3N0cj1oZWxsbyB3b3JsZA==
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
// eg:
// plaintext := []byte("abcdefg")
//
// key := []byte("1234567890123456")
// iv := []byte("1234567890123456")
//
// ciphertextHex, _ := AesCBCEncodeHex(key, iv, plaintext)
//
// output:
// ciphertextHex -> ae5d9a1e7e4260832cba80647b1e788d
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
// eg:
// ciphertextHex := "ae5d9a1e7e4260832cba80647b1e788d"
// ciphertextBase64 := "rl2aHn5CYIMsuoBkex54jQ=="
//
// key := []byte("1234567890123456")
// iv := []byte("1234567890123456")
//
// plaintextHex, _ := AesCBCDecodeHex(key, iv, ciphertextHex)
//
// output:
// // plaintextHex => abcdefg
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
