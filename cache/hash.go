package cache

import (
	"github.com/gomodule/redigo/redis"
)

// HGet 返回哈希表 key 中给定域 field 的值
func (c *cache) HGet(key, field string) string {
	r, _ := redis.String(c.DO("HGET", key, field))
	return r
}

// HSet 将哈希表 key 中的域 field 的值设为 value
func (c *cache) HSet(key, field, value string) {
	c.DO("HSET", key, field, value)
}

// HMGet v ...interface{}
// 如果给定的域不存在于哈希表，那么返回一个 nil 值
// HMGET key field [field ...]
func (c *cache) HMGet(v ...interface{}) []string {
	r, _ := redis.Strings(c.DO("HMGET", v...))
	return r
}

// HMSet 同时将多个 field-value (域-值)对设置到哈希表 key 中
// HMSET key field value [field value ...]
func (c *cache) HMSet(v ...interface{}) {
	if len(v) == 0 {
		return
	}
	if len(v)%2 == 0 {
		return
	}

	c.DO("HMSET", v...)
}

// HGetAll 返回哈希表 key 中，所有的域和值
func (c *cache) HGetAll(key string) []string {
	r, _ := redis.Strings(c.DO("HGETALL", key))
	return r
}

// HExists 查看哈希表 key 中，给定域 field 是否存在
func (c *cache) HExists(key, field string) bool {
	r, _ := redis.Bool(c.DO("HEXISTS", key))
	return r
}

// HDel 删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略
// 返回被成功移除的域的数量，不包括被忽略的域
// HDEL key field [field ...]
func (c *cache) HDel(v ...interface{}) int {
	r, _ := redis.Int(c.DO("HDel", v...))
	return r
}

// HLen 返回哈希表 key 中域的数量
func (c *cache) HLen(key string) int {
	r, _ := redis.Int(c.DO("HLEN", key))
	return r
}

// HIncrby 为哈希表 key 中的域 field 的值加上增量 increment
// 增量也可以为负数
// 返回操作之后，哈希表 key 中域 field 的值
func (c *cache) HIncrby(key, field string, increment int) (int, error) {
	return redis.Int(c.DO("HINCRBY", key, field, increment))
}

// HIncrby 为哈希表 key 中的域 field 的值加上浮点数增量 increment
// 增量也可以为负数
// 返回操作之后，哈希表 key 中域 field 的值
func (c *cache) HIncrbyFloat(key, field string, increment float64) (float64, error) {
	return redis.Float64(c.DO("HINCRBYFLOAT", key, field, increment))
}
