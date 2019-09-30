# 说明

封装加密，解密，编码，解码

# API

- **func Md5(str string) string | func Md5Byte(b []byte) string**

  - 功能: `md5`加密(对指定信息生成信息摘要)
  - 示例:

    ```go
    Md5("123456")
    // => e10adc3949ba59abbe56e057f20f883e
    ```

- **func Sha1(str string) string | func Sha1Byte(b []byte) string**

  - 功能: `sha1`加密(对指定信息生成信息摘要)
  - 示例:

    ```go
    Sha1("123456")
    // => 7c4a8d09ca3762af61e59520943dc26494f8941b
    ```

- **func Sha224(str string) string | func Sha224Byte(b []byte) string**

  - 功能: `sha224`加密(对指定信息生成信息摘要)
  - 示例:

    ```go
    Sha224("123456")
    // => f8cdb04495ded47615258f9dc6a3f4707fd2405434fefc3cbf4ef4e6
    ```

- **func Sha256(str string) string | func Sha256Byte(b []byte) string**

  - 功能: `sha256`加密(对指定信息生成信息摘要)
  - 示例:

    ```go
    Sha256("123456")
    // => 8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92
    ```

* **func Sha384(str string) string | func Sha384Byte(b []byte) string**

  - 功能: `sha384`加密(对指定信息生成信息摘要)
  - 示例:

    ```go
    Sha384("123456")
    // => 0a989ebc4a77b56a6e2bb7b19d995d185ce44090c13e2984b7ecc6d446d4b61ea9991b76a4c2f04b1b4d244841449454
    ```

- **func Sha512(str string) string | func Sha512Byte(b []byte) string**

  - 功能: `sha512`加密(对指定信息生成信息摘要)
  - 示例:

    ```go
    Sha512("123456")
    // => ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413
    ```

- **func HmacMd5(str, key string) string | func HmacMd5Byte(b, key []byte) string**

  - 功能: `hmacmd5`
  - 参数:
    - **key**: 密钥
  - 示例:

    ```go
    HmacMd5("123456", "abcdef")
    // => c6bdcc80c381536a3e85f2ee5f71cebb
    ```

- **func HmacSha1(str, key string) string | func HmacSha1Byte(b, key []byte) string**

  - 功能: `hmacsha1`
  - 参数:
    - **key**: 密钥
  - 示例:

    ```go
    HmacSha1("123456", "abcdef")
    // => b8466fbb9634771d25d8ddd1242484bdb748b179
    ```

- **func HmacSha256(str, key string) string | func HmacSha256Byte(b, key []byte) string**

  - 功能: `hmacsha256`
  - 参数:
    - **key**: 密钥
  - 示例:

    ```go
    HmacSha256("123456", "abcdef")
    // => ec4a11a5568e5cfdb5fbfe7152e8920d7bad864a0645c57fe49046a3e81ec91d
    ```

- **func HmacSha512(str, key string) string | func HmacSha512Byte(b, key []byte) string**

  - 功能: `hmacsha512`
  - 参数:
    - **key**: 密钥
  - 示例:

    ```go
    HmacSha512("123456", "abcdef")
    // => 130a4caafb11b798dd7528628d21f742feaad266e66141cc2ac003f0e6437cb5749245af8a3018d354e4b55e14703a5966808438afe4aae516d2824b014b5902
    ```

- **URLEncode(str string) string**

  - 功能: 对 `url` 进行编码
  - 示例:

    ```go
    URLEncode("www.keylala.cn?name=alex&age=18&say=你好")
    // => www.keylala.cn%3Fname%3Dalex%26age%3D18%26say%3D%E4%BD%A0%E5%A5%BD
    ```

- **func URLDecode(str string) string**

  - 功能: 对 `url` 进行解码，解码失败返回空字符串
  - 示例:

    ```go
    URLDecode("www.keylala.cn%3Fname%3Dalex%26age%3D18%26say%3D%E4%BD%A0%E5%A5%BD")
    // => www.keylala.cn?name=alex&age=18&say=你好
    ```

- **func Base64Encode(str string) string | func Base64EncodeByte(b []byte) string**

  - 功能: `base64` 编码
  - 示例:

    ```go
    Base64Encode("https://www.keylala.cn/json?str=hello world")
    // => aHR0cHM6Ly93d3cua2V5bGFsYS5jbi9qc29uP3N0cj1oZWxsbyB3b3JsZA==
    ```

- **func Base64Decode(str string) ([]byte, error)**

  - 功能: `base64` 解码
  - 示例:

    ```go
    b, _ := Base64Decode("aHR0cHM6Ly93d3cua2V5bGFsYS5jbi9qc29uP3N0cj1oZWxsbyB3b3JsZA==")

    if bytes.Equal(b, []byte("https://www.keylala.cn/json?str=hello world")) {
        fmt.Println("success")
    }
    ```

- **aes cbc 加密**

  - 功能: `AES CBC` 模式加密
  - 函数:
    - **func AesCBCEncode(key, iv, plaintext []byte) ([]byte, error)**
    - **func AesCBCEncodeHex(key, iv, plaintext []byte) (string, error)**: 加密结果转为 `16` 进制
    - **func AesCBCEncodeBase64(key, iv, plaintext []byte) (string, error)**: 加密结果经过 `base64` 编码
  - 参数:
    - **key**: 加密密钥，长度(`len(key)`)需要为 `16，24，32`
    - **iv**: 向量(密钥偏移量)，长度(`len(iv)`)必须为 `16`
    - **plaintext**: 待加密的内容
  - 示例:

    ```go
    plaintext := []byte("abcdefg")

    key := []byte("1234567890123456")
    iv := []byte("1234567890123456")

    ciphertextHex, _ := AesCBCEncodeHex(key, iv, plaintext)
    // ciphertextHex => ae5d9a1e7e4260832cba80647b1e788d

    ciphertextBase64, _ := AesCBCEncodeBase64(key, iv, plaintext)
    // ciphertextHex => rl2aHn5CYIMsuoBkex54jQ==
    ```

- **aes cbc 解密**

  - 功能: `AES CBC`模式解密
  - 函数:
    - **func AesCBCDecode(key, iv, ciphertext []byte) ([]byte, error)**
    - **func AesCBCDecodeHex(key, iv []byte, ciphertext string) ([]byte, error)**: `ciphertext` 为 `base64` 编码的字符串
    - **func AesCBCDecodeBase64(key, iv []byte, ciphertext string) ([]byte, error)**: `ciphertext` 为 16 进制字符串
  - 参数:
    - **key**: 加密密钥，长度(`len(key)`)需要为 `16，24，32`
    - **iv**: 向量(密钥偏移量)，长度(`len(iv)`)必须为 `16`
    - **ciphertext**: 待解密的内容
  - 示例:

    ```go
    ciphertextHex := "ae5d9a1e7e4260832cba80647b1e788d"
    ciphertextBase64 := "rl2aHn5CYIMsuoBkex54jQ=="

    key := []byte("1234567890123456")
    iv := []byte("1234567890123456")

    plaintextHex, _ := AesCBCDecodeHex(key, iv, ciphertextHex)
    // plaintextHex => abcdefg

    plaintextBase64, _ := AesCBCDecodeBase64(key, iv, ciphertextBase64)
    // plaintextBase64 => abcdefg
    ```
