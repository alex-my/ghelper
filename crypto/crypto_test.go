package crypto

import (
	"strings"
	"testing"
)

func TestMd5(t *testing.T) {
	if Md5("123456") != "e10adc3949ba59abbe56e057f20f883e" {
		t.Error("Md5 error")
	}
	if Md5Byte([]byte("123456")) != "e10adc3949ba59abbe56e057f20f883e" {
		t.Error("Md5Byte error")
	}
}

func TestSha1(t *testing.T) {
	if Sha1("123456") != "7c4a8d09ca3762af61e59520943dc26494f8941b" {
		t.Error("Sha1 error")
	}
	if Sha1Byte([]byte("123456")) != "7c4a8d09ca3762af61e59520943dc26494f8941b" {
		t.Error("Sha1Byte error")
	}
}

func TestSha224(t *testing.T) {
	if Sha224("123456") != "f8cdb04495ded47615258f9dc6a3f4707fd2405434fefc3cbf4ef4e6" {
		t.Error("Sha256 error")
	}
	if Sha224Byte([]byte("123456")) != "f8cdb04495ded47615258f9dc6a3f4707fd2405434fefc3cbf4ef4e6" {
		t.Error("Sha256Byte error")
	}
}

func TestSha256(t *testing.T) {
	if Sha256("123456") != "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92" {
		t.Error("Sha256 error")
	}
	if Sha256Byte([]byte("123456")) != "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92" {
		t.Error("Sha256Byte error")
	}
}

func TestSha384(t *testing.T) {
	if Sha384("123456") != "0a989ebc4a77b56a6e2bb7b19d995d185ce44090c13e2984b7ecc6d446d4b61ea9991b76a4c2f04b1b4d244841449454" {
		t.Error("Sha512 error")
	}
	if Sha384Byte([]byte("123456")) != "0a989ebc4a77b56a6e2bb7b19d995d185ce44090c13e2984b7ecc6d446d4b61ea9991b76a4c2f04b1b4d244841449454" {
		t.Error("Sha512Byte error")
	}
}

func TestSha512(t *testing.T) {
	if Sha512("123456") != "ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413" {
		t.Error("Sha512 error")
	}
	if Sha512Byte([]byte("123456")) != "ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413" {
		t.Error("Sha512Byte error")
	}
}

func TestHmacMd5(t *testing.T) {
	if HmacMd5("123456", "abcdef") != "c6bdcc80c381536a3e85f2ee5f71cebb" {
		t.Error("HmacMd5 error")
	}
	if HmacMd5Byte([]byte("123456"), []byte("abcdef")) != "c6bdcc80c381536a3e85f2ee5f71cebb" {
		t.Error("HmacMd5Byte error")
	}
}

func TestHmacSha1(t *testing.T) {
	if HmacSha1("123456", "abcdef") != "b8466fbb9634771d25d8ddd1242484bdb748b179" {
		t.Error("HmacSha1 error")
	}
	if HmacSha1Byte([]byte("123456"), []byte("abcdef")) != "b8466fbb9634771d25d8ddd1242484bdb748b179" {
		t.Error("HmacSha1Byte error")
	}
}

func TestHmacSha256(t *testing.T) {
	if HmacSha256("123456", "abcdef") != "ec4a11a5568e5cfdb5fbfe7152e8920d7bad864a0645c57fe49046a3e81ec91d" {
		t.Error("HmacSha256 error")
	}
	if HmacSha256Byte([]byte("123456"), []byte("abcdef")) != "ec4a11a5568e5cfdb5fbfe7152e8920d7bad864a0645c57fe49046a3e81ec91d" {
		t.Error("HmacSha256Byte error")
	}
}

func TestHmacSha512(t *testing.T) {
	if HmacSha512("123456", "abcdef") != "130a4caafb11b798dd7528628d21f742feaad266e66141cc2ac003f0e6437cb5749245af8a3018d354e4b55e14703a5966808438afe4aae516d2824b014b5902" {
		t.Error("HmacSha512 error")
	}
	if HmacSha512Byte([]byte("123456"), []byte("abcdef")) != "130a4caafb11b798dd7528628d21f742feaad266e66141cc2ac003f0e6437cb5749245af8a3018d354e4b55e14703a5966808438afe4aae516d2824b014b5902" {
		t.Error("HmacSha512Byte error")
	}
}

func TestURLEncode(t *testing.T) {
	// www.keylala.cn%3Fname%3Dalex%26age%3D18%26say%3D%E4%BD%A0%E5%A5%BD
	if strings.ToLower(URLEncode("www.keylala.cn?name=alex&age=18&say=你好")) != "www.keylala.cn%3fname%3dalex%26age%3d18%26say%3d%e4%bd%a0%e5%a5%bd" {
		t.Error("URLEncode error")
	}
}

func TestURLDecode(t *testing.T) {
	if URLDecode("www.keylala.cn%3Fname%3Dalex%26age%3D18%26say%3D%E4%BD%A0%E5%A5%BD") != "www.keylala.cn?name=alex&age=18&say=你好" {
		t.Error("URLDecode error")
	}
}