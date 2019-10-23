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

// Del 删除给定的一个或多个key
// 返回被删除的 key 的数量
func (c *cache) Del(key ...interface{}) (int, error) {
	return redis.Int(c.DO("DEL", key...))
}

// Exists 检查指定的 key 是否存在
func (c *cache) Exists(key string) (bool, error) {
	return redis.Bool(c.DO("EXISTS", key))
}

// Expire 为 key 设置生存时间，单位 秒
// 返回是否设置成功
func (c *cache) Expire(key, ex string) (bool, error) {
	return redis.Bool(c.DO("EXPIRE", key, ex))
}

// ExpireAt 为 key 设置生存时间，key 存活到 t
// t 为 unix时间戳
// 返回是否设置成功
func (c *cache) ExpireAt(key string, t int) (bool, error) {
	return redis.Bool(c.DO("EXPIREAT", key, t))
}

// PExpire 为 key 设置生存时间，单位 毫秒
// 返回是否设置成功
func (c *cache) PExpire(key, ex string) (bool, error) {
	return redis.Bool(c.DO("PEXPIRE", key, ex))
}

// PExpireAt 为 key 设置生存时间，key 存活到 t
// t 为 毫秒
// 返回是否设置成功
func (c *cache) PExpireAt(key string, t int) (bool, error) {
	return redis.Bool(c.DO("PEXPIREAT", key, t))
}

// TTL 以秒为单位
// 返回给定 key 的剩余生存时间 (秒)，-2 表示 key 不存在, -1 表示永久
func (c *cache) TTL(key string) (int, error) {
	return redis.Int(c.DO("TTL", key))
}

// PTTL 以毫秒为单位
// 返回给定 key 的剩余生存时间 (毫秒)，-2 表示 key 不存在, -1 表示永久
func (c *cache) PTTL(key string) (int, error) {
	return redis.Int(c.DO("PTTL", key))
}

// Get 返回 key 所关联的字符串
func (c *cache) Get(key string) (string, error) {
	return redis.String(c.DO("GET", key))
}

// Set 将字符串 value 关联到 key
func (c *cache) Set(key, value string) {
	c.DO("SET", key, value)
}

// SetEx 将字符串 value 关联到 key，并将 key 的生存时间设置为 seconds (秒)
func (c *cache) SetEx(key, value, seconds string) {
	c.DO("SET", key, value, "EX", seconds)
}

// PSetEx 将字符串 value 关联到 key，并将 key 的生存时间设置为 milliseconds (毫秒)
func (c *cache) PSetEx(key, value, milliseconds string) {
	c.DO("SET", key, value, "PX", milliseconds)
}

// MGet 返回所有(一个或多个)给定 key 的值
// v 必须是 字符串集合
// 如果给定的 key 里面，有某个 key 不存在，那么这个 key 返回特殊值 nil
func (c *cache) MGet(key ...interface{}) []string {
	r, _ := redis.Strings(c.DO("MGET", key...))
	return r
}

// MSet 同时设置一个或多个 key-value 对
// v 必须是 字符串集合
// 这是一个原子性操作
func (c *cache) MSet(v ...interface{}) {
	if len(v) == 0 {
		return
	}

	c.DO("MSET", v...)
}

// Append 命令将 value 追加到 key 原来的值的末尾
// 返回追加之后，字符串总长度
func (c *cache) Append(key, value string) (int, error) {
	return redis.Int(c.DO("APPEND", key, value))
}

// Strlen 返回 key 所储存的字符串值的长度
// 当 key 不存在时，返回 0
func (c *cache) Strlen(key string) (int, error) {
	return redis.Int(c.DO("STRLEN", key))
}

// Incr 将 key 所储存的值加上增量 1
// 将 key 所储存的值加上增量 1
// 加上 1 之后， key 的值
func (c *cache) Incr(key string) (int64, error) {
	return redis.Int64(c.DO("INCR", key))
}

// Incrby 将 key 所储存的值加上增量 increment
// 将 key 所储存的值加上增量 increment
// 加上 increment 之后， key 的值
func (c *cache) Incrby(key string, increment int64) (int64, error) {
	return redis.Int64(c.DO("INCRBY", key, increment))
}

// Decr 将 key 所储存的值减去 1
// 将 key 所储存的值减去 1
// 减去 1 之后， key 的值
func (c *cache) Decr(key string) (int64, error) {
	return redis.Int64(c.DO("DECR", key))
}

// Decrby 将 key 所储存的值减去减量 decrement
// 将 key 所储存的值减去减量 decrement
// 减去 decrement 之后， key 的值
func (c *cache) Decrby(key string, decrement int64) (int64, error) {
	return redis.Int64(c.DO("DECRBY", key, decrement))
}

// GetSet 将给定 key 的值设为 value ，并返回 key 的旧值
func (c *cache) GetSet(key, value string) (string, error) {
	return redis.String(c.DO("GETSET", key, value))
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
