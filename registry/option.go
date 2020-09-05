package registry

import (
	"crypto/tls"
	"time"
)

// Config 关于服务的一些配置
type Config struct {
	// Addrs 支持多个地址
	Addrs []string

	// DialTimeout 连接超时时间
	DialTimeout time.Duration

	// TLSConfig tls 配置
	TLSConfig *tls.Config

	// TTL 设置存活时间，超过这个时间会被移除，时间从注册开始计时
	TTL time.Duration

	// Interval 每隔一段时间就重新注册，用以保证不会因为超过 TTL 而被移除
	Interval time.Duration

	// ReadTimeout 读操作超时
	ReadTimeout time.Duration

	// WriteTimeout 写操作超时
	WriteTimeout time.Duration

	// Username 用户名
	Username string

	// Password 密码
	Password string
}

// Option ..
type Option func(config *Config)

// DefaultConfig 默认地址
func DefaultConfig() *Config {
	return &Config{
		Addrs:        []string{"127.0.0.1:2379"},
		DialTimeout:  10 * time.Second,
		TTL:          30 * time.Second,
		Interval:     15 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

// WithAddrs 设置地址
func WithAddrs(addrs ...string) Option {
	return func(config *Config) {
		config.Addrs = addrs
	}
}

// WithDialTimeout 设置连接超时时间
func WithDialTimeout(timeout time.Duration) Option {
	return func(config *Config) {
		config.DialTimeout = timeout
	}
}

// WithTLSConfig 设置 tls 配置
func WithTLSConfig(tls *tls.Config) Option {
	return func(config *Config) {
		config.TLSConfig = tls
	}
}

// WithTTL 置存活时间，超过这个时间会被移除，时间从注册开始计时
func WithTTL(ttl time.Duration) Option {
	return func(config *Config) {
		config.TTL = ttl
	}
}

// WithInterval 每隔一段时间就重新注册，用以保证不会因为超过 TTL 而被移除
func WithInterval(interval time.Duration) Option {
	return func(config *Config) {
		config.Interval = interval
	}
}

// WithReadTimeout 设置读操作超时
func WithReadTimeout(timeout time.Duration) Option {
	return func(config *Config) {
		config.ReadTimeout = timeout
	}
}

// WithWriteTimeout 设置写操作超时
func WithWriteTimeout(timeout time.Duration) Option {
	return func(config *Config) {
		config.WriteTimeout = timeout
	}
}

// WithUsername 设置用户名
func WithUsername(username string) Option {
	return func(config *Config) {
		config.Username = username
	}
}

// WithPassword 设置密码
func WithPassword(password string) Option {
	return func(config *Config) {
		config.Password = password
	}
}
