package cache

import "time"

// config 配置文件
type config struct {
	host     string
	port     int
	password string
	db       int
	// 池中最大空闲数量
	maxIdle int
	// 空闲的连接一定时间后会被关闭，默认为0，表示不会关闭空闲连接
	idleTimeout time.Duration
	// 池中的最大连接数，0 表示不限制，一般推荐配置大一些
	maxActive int
	// 如果 maxActive > 0：
	// 设置为 true， 数量达到 maxActive 时，会阻塞于此，等待获取连接
	// 设置为 false， 数量达到 maxActive 时，直接报错
	wait bool
	// 读取操作时，超时时间
	dialReadTimeout time.Duration
	// 写入操作时，超时时间
	dialWriteTimeout time.Duration
	// 连接 redis 服务器时的超时时间
	dialConnectTimeout time.Duration
}

func defaultConfig() *config {
	return &config{
		host:               "127.0.0.1",
		port:               6379,
		db:                 0,
		maxIdle:            20,
		idleTimeout:        0,
		maxActive:          100,
		wait:               true,
		dialReadTimeout:    time.Duration(500) * time.Millisecond,
		dialWriteTimeout:   time.Duration(500) * time.Millisecond,
		dialConnectTimeout: time.Duration(500) * time.Millisecond,
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

// WithPassword 密码
func WithPassword(password string) Option {
	return func(c Cache) {
		c.config().password = password
	}
}

// WithDB 0-15
func WithDB(db int) Option {
	return func(c Cache) {
		c.config().db = db
	}
}

// WithMaxIdle ..
func WithMaxIdle(maxIdle int) Option {
	return func(c Cache) {
		c.config().maxIdle = maxIdle
	}
}

// WithIdleTimeout ..
func WithIdleTimeout(timeout time.Duration) Option {
	return func(c Cache) {
		c.config().idleTimeout = timeout
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

// WithDialReadTimeout ..
func WithDialReadTimeout(timeout time.Duration) Option {
	return func(c Cache) {
		c.config().dialReadTimeout = timeout
	}
}

// WithDialWriteTimeout ..
func WithDialWriteTimeout(timeout time.Duration) Option {
	return func(c Cache) {
		c.config().dialWriteTimeout = timeout
	}
}

// WithDialConnectTimeout ..
func WithDialConnectTimeout(timeout time.Duration) Option {
	return func(c Cache) {
		c.config().dialConnectTimeout = timeout
	}
}
