package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	opt = defaultOption()
)

// Init 初始化
func Init(opts ...Option) {
	if len(opts) > 0 {
		opt = opts[0]
	}
}

// Token 生成 token
func Token(data ...map[string]interface{}) (string, error) {
	return createToken(opt.Secret, data...)
}

// TokenWithKey 生成 token
func TokenWithKey(key []byte, data ...map[string]interface{}) (string, error) {
	return createToken(key, data...)
}

// Verify 验证 token，并获取自定义内容
func Verify(s string) (map[string]interface{}, error) {
	return verify(opt.Secret, s)
}

// VerifyWithKey 验证 token，并获取自定义内容
func VerifyWithKey(key []byte, s string) (map[string]interface{}, error) {
	return verify(key, s)
}

func createToken(key []byte, data ...map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{
		// 签发时间
		"iat": time.Now().Unix(),
		// 过期时间
		"exp": time.Now().Add(opt.Exp).Unix(),
	}

	if len(data) > 0 {
		for k, v := range data[0] {
			claims[k] = v
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}

func verify(key []byte, s string) (map[string]interface{}, error) {
	parse, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid jwt method")
		}

		return key, nil
	})

	if err != nil {
		return nil, errors.New("token validate failed")
	}

	if claims, ok := parse.Claims.(jwt.MapClaims); ok && parse.Valid {
		// 验证是否超时
		exp, ok := claims["exp"]
		if !ok {
			return nil, errors.New("no exp in token")
		}
		f64, ok := exp.(float64)
		if !ok {
			return nil, errors.New("invalid exp type")
		}

		i64exp := int64(f64)

		if i64exp < time.Now().Unix() {
			return nil, fmt.Errorf("token expired at %d", i64exp)
		}

		payload := make(map[string]interface{}, len(claims))
		for k, v := range claims {
			payload[k] = v
		}

		return payload, nil
	}

	return nil, errors.New("token not match MapClaims")
}
