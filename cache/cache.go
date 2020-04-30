package cache

import (
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// redigo 目前不支持 redis 集群
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
	Conn() Conn

	Convert
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

// Convert 数据转换
type Convert interface {
	Int(reply interface{}, err error) (int, error)
	Int64(reply interface{}, err error) (int64, error)
	Uint64(reply interface{}, err error) (uint64, error)
	Float64(reply interface{}, err error) (float64, error)
	String(reply interface{}, err error) (string, error)
	Bytes(reply interface{}, err error) ([]byte, error)
	Bool(reply interface{}, err error) (bool, error)
	Float64s(reply interface{}, err error) ([]float64, error)
	Int64s(reply interface{}, err error) ([]int64, error)
	Ints(reply interface{}, err error) ([]int, error)
	ByteSlices(reply interface{}, err error) ([][]byte, error)
	Strings(reply interface{}, err error) ([]string, error)
	StringMap(reply interface{}, err error) (map[string]string, error)
	IntMap(reply interface{}, err error) (map[string]int, error)
	Int64Map(reply interface{}, err error) (map[string]int64, error)
	Values(reply interface{}, err error) ([]interface{}, error)
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
	Set(key string, value interface{}) error
	SetEx(key string, value interface{}, seconds string) error
	PSetEx(key string, value interface{}, milliseconds string) error
	MGet(key ...interface{}) ([]string, error)
	MSet(v ...interface{}) error
	Append(key string, value interface{}) (int, error)
	Strlen(key string) (int, error)
	Incr(key string) (int64, error)
	Incrby(key string, increment int64) (int64, error)
	Decr(key string) (int64, error)
	Decrby(key string, decrement int64) (int64, error)
	GetSet(key string, value interface{}) (string, error)
}

// Hash 哈希表
type Hash interface {
	HGet(key, field string) (string, error)
	HSet(key, field string, value interface{}) error
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

	// 执行结束后，没有错误，没有关闭连接，没有超过 MaxIdle 情况下，activeConn 会放入 idle 队列
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

func (c *cache) Int(reply interface{}, err error) (int, error) {
	return redis.Int(reply, err)
}

func (c *cache) Int64(reply interface{}, err error) (int64, error) {
	return redis.Int64(reply, err)
}

func (c *cache) Uint64(reply interface{}, err error) (uint64, error) {
	return redis.Uint64(reply, err)
}

func (c *cache) Float64(reply interface{}, err error) (float64, error) {
	return redis.Float64(reply, err)
}

func (c *cache) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

func (c *cache) Bytes(reply interface{}, err error) ([]byte, error) {
	return redis.Bytes(reply, err)
}

func (c *cache) Bool(reply interface{}, err error) (bool, error) {
	return redis.Bool(reply, err)
}

func (c *cache) Float64s(reply interface{}, err error) ([]float64, error) {
	return redis.Float64s(reply, err)
}

func (c *cache) Int64s(reply interface{}, err error) ([]int64, error) {
	return redis.Int64s(reply, err)
}

func (c *cache) Ints(reply interface{}, err error) ([]int, error) {
	return redis.Ints(reply, err)
}

func (c *cache) ByteSlices(reply interface{}, err error) ([][]byte, error) {
	return redis.ByteSlices(reply, err)
}

func (c *cache) Strings(reply interface{}, err error) ([]string, error) {
	return redis.Strings(reply, err)
}

func (c *cache) StringMap(reply interface{}, err error) (map[string]string, error) {
	return redis.StringMap(reply, err)
}

func (c *cache) IntMap(reply interface{}, err error) (map[string]int, error) {
	return redis.IntMap(reply, err)
}

func (c *cache) Int64Map(reply interface{}, err error) (map[string]int64, error) {
	return redis.Int64Map(reply, err)
}

func (c *cache) Values(reply interface{}, err error) ([]interface{}, error) {
	return redis.Values(reply, err)
}
