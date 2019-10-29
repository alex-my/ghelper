package jwt

import (
	"time"
)

// Option ..
type Option struct {
	// Secret 签名密钥
	Secret []byte
	// Exp 过期时间
	Exp time.Duration
}

func defaultOption() Option {
	return Option{
		Secret: []byte("123456"),
		Exp:    time.Minute * 5,
	}
}
