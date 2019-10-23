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
	DO(cmd string, args ...interface{}) (interface{}, error)

	Key
	String
}

// Key ...
type Key interface {
	Del(key ...interface{}) (int, error)
	Exists(key string) (bool, error)
	Expire(key, ex string) (bool, error)
	ExpireAt(key string, t int) (bool, error)
	PExpire(key, ex string) (bool, error)
	PExpireAt(key string, t int) (bool, error)
	TTL(key string) (int, error)
	PTTL(key string) (int, error)
}

// String ...
type String interface {
	Get(key string) (string, error)
	Set(key, value string)
	SetEx(key, value, seconds string)
	PSetEx(key, value, milliseconds string)
	MGet(key ...interface{}) []string
	MSet(v ...interface{})
	Append(key, value string) (int, error)
	Strlen(key string) (int, error)
	Incr(key string) (int64, error)
	Incrby(key string, increment int64) (int64, error)
	Decr(key string) (int64, error)
	Decrby(key string, decrement int64) (int64, error)
	GetSet(key, value string) (string, error)
}

// TODO Hash
// TODO List
// TODO Set
// TODO SortedSet
// TODO Pub/Sub
// TODO Transaction
// TODO Server

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

// DO ..
func (c *cache) DO(cmd string, args ...interface{}) (interface{}, error) {
	conn := c.pool.Get()
	if conn == nil {
		return nil, ErrInvalidConn
	}
	defer conn.Close()

	return conn.Do(cmd, args...)
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
