package cache

import "time"

// config 配置文件
type config struct {
	host     string
	port     int
	db       int
	password string
	maxIdle  int
	// 最多生成一定数量的连接，0 表示不限制
	maxActive int
	// 如果 maxActive > 0, 此字段设置为 true, 表示会阻塞于此，等待获取连接
	wait bool
	// 空闲的连接一定时间后会被关闭，默认为0，表示不会关闭空闲连接
	idleTimeout time.Duration
}

func defaultConfig() *config {
	return &config{
		host:        "127.0.0.1",
		port:        6379,
		db:          0,
		maxIdle:     20,
		maxActive:   0,
		wait:        false,
		idleTimeout: 0,
	}
}

// Option ..
type Option func(Cache)

// WithHost 地址
func WithHost(host string) Option {
	return func(c Cache) {
		c.config().host = host
	}
}

// WithPort 端口号
func WithPort(port int) Option {
	return func(c Cache) {
		c.config().port = port
	}
}

// WithDB 0-15
func WithDB(db int) Option {
	return func(c Cache) {
		c.config().db = db
	}
}

// WithPassword 密码
func WithPassword(password string) Option {
	return func(c Cache) {
		c.config().password = password
	}
}

// WithMaxIdle ..
func WithMaxIdle(maxIdle int) Option {
	return func(c Cache) {
		c.config().maxIdle = maxIdle
	}
}

// WithMaxActive ..
func WithMaxActive(maxActive int) Option {
	return func(c Cache) {
		c.config().maxActive = maxActive
	}
}

// WithWait ..
func WithWait(wait bool) Option {
	return func(c Cache) {
		c.config().wait = wait
	}
}

// WithIdleTimeout ..
func WithIdleTimeout(timeout time.Duration) Option {
	return func(c Cache) {
		c.config().idleTimeout = timeout
	}
}
