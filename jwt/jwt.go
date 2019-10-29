package jwt

import (
	"errors"
	"fmt"
	"strconv"
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

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString(opt.Secret)
}

// Verify 验证 token，并获取自定义内容
func Verify(s string) (map[string]interface{}, error) {
	parse, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid jwt method")
		}

		return opt.Secret, nil
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
		iexp, err := strconv.ParseInt(exp.(string), 10, 64)
		if err != nil {
			return nil, errors.New("invalid exp")
		}

		if iexp < time.Now().Unix() {
			return nil, fmt.Errorf("token expired at %d", iexp)
		}

		payload := make(map[string]interface{}, len(claims))
		for k, v := range claims {
			payload[k] = v
		}

		return payload, nil
	}

	return nil, errors.New("token not match MapClaims")
}
