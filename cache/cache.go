package cache

import (
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 命令的文字注释来自于 http://doc.redisfans.com，稍有修改

var (
	// ErrInvalidConn 无法获取与 redis-server 的连接
	ErrInvalidConn = errors.New("invalid conn")
	// ErrInvalidParamCount 参数数量错误
	ErrInvalidParamCount = errors.New("invalid param count")
)

// Cache 缓存
type Cache interface {
	config() *config
	Open() error
	Close() error
	DO(cmd string, args ...interface{}) (interface{}, error)

	Key
	String
	Hash
	List
	Set
	SortedSet
}

// Conn ..
type Conn interface {
	redis.Conn
}

// Key 键
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

// String 字符串
type String interface {
	Get(key string) (string, error)
	Set(key, value string) error
	SetEx(key, value, seconds string) error
	PSetEx(key, value, milliseconds string) error
	MGet(key ...interface{}) ([]string, error)
	MSet(v ...interface{}) error
	Append(key, value string) (int, error)
	Strlen(key string) (int, error)
	Incr(key string) (int64, error)
	Incrby(key string, increment int64) (int64, error)
	Decr(key string) (int64, error)
	Decrby(key string, decrement int64) (int64, error)
	GetSet(key, value string) (string, error)
}

// Hash 哈希表
type Hash interface {
	HGet(key, field string) (string, error)
	HSet(key, field, value string) error
	HMGet(v ...interface{}) ([]string, error)
	HMSet(v ...interface{}) error
	HGetAll(key string) ([]string, error)
	HExists(key, field string) (bool, error)
	HDel(v ...interface{}) (int, error)
	HLen(key string) (int, error)
	HIncrby(key, field string, increment int) (int, error)
	HIncrbyFloat(key, field string, increment float64) (float64, error)
}

// List 列表
type List interface {
	LPush(v ...interface{}) (int, error)
	LPop(key string) (string, error)
	RPush(v ...interface{}) (int, error)
	RPop(key string) (string, error)
	RPopLPush(source, destination string) (string, error)
	LTrim(key string, start, stop int) error
	LSet(key string, index int, value string) error
	LRem(key string, count int, value string) (int, error)
	LRange(key string, start, stop int) ([]string, error)
	LLen(key string) (int, error)
	LInsertBefore(key, pivot, value string) (int, error)
	LInsertAfter(key, pivot, value string) (int, error)
	LIndex(key string, index int) (string, error)
}

// Set 集合
type Set interface {
	SAdd(v ...interface{}) (int, error)
	SCard(key string) (int, error)
	SDiff(key ...interface{}) ([]string, error)
	SDiffStore(key ...interface{}) (int, error)
	SUnion(key ...interface{}) ([]string, error)
	SUnionStore(key ...interface{}) (int, error)
	SInter(key ...interface{}) ([]string, error)
	SInterStore(key ...interface{}) (int, error)
	SIsMember(key, member string) (bool, error)
	SMembers(key string) ([]string, error)
	SPop(key string) (string, error)
	SRandMember(key string, count int) ([]string, error)
	SRem(v ...interface{}) (int, error)
}

// SortedSet 有序集合
type SortedSet interface {
	ZAdd(v ...interface{}) (int, error)
	ZCard(key string) (int, error)
	ZCount(key string, min int, max int) (int, error)
	ZIncrby(key, member string, incrment int) (int, error)
	ZRange(key string, start, stop int) ([]string, error)
	ZRangeWithScores(key string, start, stop int) ([]string, error)
	ZScore(key, member string) (int, error)
	ZRank(key, member string) (int, error)
	ZRangeByScore(key string, min, max, limit, offset, count int) ([]string, error)
	ZRevRank(key, member string) (int, error)
	ZRevRangeByScore(key string, max, min, limit, offset, count int) ([]string, error)
	ZRevRange(key string, start, stop int) ([]string, error)
	ZRemRangeByScore(key string, min, max int) (int, error)
	ZRemRangeByRank(key string, start, stop int) (int, error)
	ZRem(v ...interface{}) (int, error)
}

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

// Conn 获取 redigo Conn
func (c *cache) Conn() Conn {
	conn := c.pool.Get()
	return conn
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
