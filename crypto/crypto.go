package crypto

// 测试站点: https://www.keylala.cn

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
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
