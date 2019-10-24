package sign

import (
	"bytes"
	"errors"
	"sort"
	"sync"

	"github.com/alex-my/ghelper/crypto"
	"github.com/alex-my/ghelper/logger"
)

var (
	// ErrorInvalidSignParams 无效的签名参数
	ErrorInvalidSignParams = errors.New("invalid sign param")
)

var (
	// ignoreField 不参与签名的字段
	ignoreField = "sign"
	// secretKey 密钥
	secretKey = "abcdefg"
	// calcFunc 签名函数 func(signStr string, secret string) string
	calcFunc = calc
	// log 日志
	log = logger.NewLogger()
	// debug 是否打印日志
	debug = true
)

// SetIgnoreField 设置忽略的字段
func SetIgnoreField(field string) {
	ignoreField = field
}

// SetSecretKey 设置签名密钥
func SetSecretKey(key string) {
	secretKey = key
}

// SetCalcFunc 设置签名函数
func SetCalcFunc(f func(string, string) string) {
	if f == nil {
		return
	}
	calcFunc = f
}

// SetLog 设置日志
func SetLog(l logger.Logger) {
	log = l
}

// SetDebug 是否打印日志
func SetDebug(b bool) {
	debug = b
}

// Signs 计算签名，使用全局密钥 secretKey
func Signs(values map[string][]string) (string, error) {
	if values == nil {
		return "", ErrorInvalidSignParams
	}

	v := map[string]string{}
	for name, value := range values {
		if len(value) > 0 {
			v[name] = value[0]
		}
	}

	return sign(v, secretKey)
}

// Sign 计算签名，使用全局密钥 secretKey
func Sign(values map[string]string) (string, error) {
	if values == nil {
		return "", ErrorInvalidSignParams
	}

	return sign(values, secretKey)
}

// SWithKey 计算签名
// key 密钥
func SWithKey(values map[string][]string, key string) (string, error) {
	if values == nil {
		return "", ErrorInvalidSignParams
	}

	v := map[string]string{}
	for name, value := range values {
		if len(value) > 0 {
			v[name] = value[0]
		}
	}

	return sign(v, key)
}

// WithKey 计算签名
// key 密钥
func WithKey(values map[string]string, key string) (string, error) {
	if values == nil {
		return "", ErrorInvalidSignParams
	}

	return sign(values, secretKey)
}

// sign 计算签名过程
func sign(values map[string]string, key string) (string, error) {

	// 所有参数按照字母顺序从小到大排列，参数值为空不参与签名
	// 所有参数形成如 key1=value1&key2=value2 的形式
	size := len(values)
	keys := make([]string, 0, size)
	for key, value := range values {
		if key != ignoreField && value != "" {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	b := buffer()
	defer releaseBuffer(b)
	b.Reset()

	size = len(keys)
	for index, key := range keys {
		if key == "" {
			continue
		}
		value := values[key]

		b.WriteString(key)
		b.WriteString("=")
		b.WriteString(value)

		if index != size-1 {
			b.WriteString("&")
		}
	}

	signStr := b.String()

	return calcFunc(signStr, key), nil
}

// calc 默认的签名函数
func calc(s string, k string) string {
	out := crypto.Md5(s + k)
	if debug {
		log.Debugf("signStr: %s, calcSign: %s", s, out)
	}
	return out
}

var bufferPool *sync.Pool

func buffer() *bytes.Buffer {
	buff := bufferPool.Get().(*bytes.Buffer)
	buff.Reset()
	return buff
}

func releaseBuffer(buff *bytes.Buffer) {
	bufferPool.Put(buff)
}

func init() {
	bufferPool = &sync.Pool{}
	bufferPool.New = func() interface{} {
		return &bytes.Buffer{}
	}
}
