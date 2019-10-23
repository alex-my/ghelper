package cache

import (
	"github.com/gomodule/redigo/redis"
)

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
