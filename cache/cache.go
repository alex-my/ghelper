package cache

import (
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var (
	// ErrInvalidConn 无法获取与 redis-server 的连接
	ErrInvalidConn = errors.New("invalid conn")
)

// Cache 缓存
type Cache interface {
	config() *config
	Open() error
	Close() error
	Conn() (Conn, error)
	DO(cmd string, args ...interface{}) (interface{}, error)
	String(key string) (string, error)
	SetString(key, value string, ex string) (interface{}, error)
}

// Conn 连接
type Conn interface {
	Send(cmd string, args ...interface{}) error
	Flush() error
	Close() error
}

type cache struct {
	conf *config
	pool *redis.Pool
}

// NewCache ..
func NewCache(opts ...Option) Cache {
	return newCache(opts...)
}

func newCache(opts ...Option) Cache {
	c := &cache{
		conf: defaultConfig(),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *cache) config() *config {
	return c.conf
}

// Open ..
func (c *cache) Open() error {
	return c.initRedis()
}

// Close ..
func (c *cache) Close() error {
	if c.pool != nil {
		return c.pool.Close()
	}

	return nil
}

// Conn 获取一条连接
func (c *cache) Conn() (Conn, error) {
	conn := c.pool.Get()
	if conn == nil {
		return nil, ErrInvalidConn
	}

	return conn, nil
}

// DO ..
func (c *cache) DO(cmd string, args ...interface{}) (interface{}, error) {
	conn := c.pool.Get()
	if conn == nil {
		return nil, ErrInvalidConn
	}
	defer conn.Close()

	return conn.Do(cmd, args...)
}

// String 获取字符串
func (c *cache) String(key string) (string, error) {
	out, err := redis.String(c.DO("GET", key))

	return out, err
}

// SetString 设置字符串
// ex 为空时，表示永远
func (c *cache) SetString(key, value string, ex string) (interface{}, error) {
	if ex == "" {
		return c.DO("SET", key, value)
	}
	return c.DO("SET", key, value, "EX", ex)
}

func (c *cache) initRedis() error {
	conf := c.conf

	pool := &redis.Pool{
		MaxIdle:     conf.maxIdle,
		MaxActive:   conf.maxActive,
		IdleTimeout: conf.idleTimeout,
	}

	// 创建新连接
	pool.Dial = func() (redis.Conn, error) {
		options := redis.DialPassword(conf.password)
		addr := fmt.Sprintf("%s:%d", conf.host, conf.port)
		conn, err := redis.Dial("tcp", addr, options)
		if err != nil {
			return nil, err
		}

		if _, err := conn.Do("SELECT", conf.db); err != nil {
			conn.Close()
			return nil, err
		}

		return conn, nil
	}

	c.pool = pool

	return nil
}
