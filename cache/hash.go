package cache

import (
	"github.com/gomodule/redigo/redis"
)

// HGet 返回哈希表 key 中给定域 field 的值
func (c *cache) HGet(key, field string) (string, error) {
	return redis.String(c.DO("HGET", key, field))
}

// HSet 将哈希表 key 中的域 field 的值设为 value
func (c *cache) HSet(key, field string, value interface{}) error {
	_, err := c.DO("HSET", key, field, value)
	return err
}

// HMGet v ...interface{}
// 如果给定的域不存在于哈希表，那么返回一个 nil 值
// HMGET key field [field ...]
func (c *cache) HMGet(v ...interface{}) ([]string, error) {
	return redis.Strings(c.DO("HMGET", v...))
}

// HMSet 同时将多个 field-value (域-值)对设置到哈希表 key 中
// HMSET key field value [field value ...]
func (c *cache) HMSet(v ...interface{}) error {
	if len(v) == 0 {
		return nil
	}
	if len(v)%2 == 0 {
		return ErrInvalidParamCount
	}

	_, err := c.DO("HMSET", v...)
	return err
}

// HGetAll 返回哈希表 key 中，所有的域和值
func (c *cache) HGetAll(key string) ([]string, error) {
	return redis.Strings(c.DO("HGETALL", key))
}

// HExists 查看哈希表 key 中，给定域 field 是否存在
func (c *cache) HExists(key, field string) (bool, error) {
	return redis.Bool(c.DO("HEXISTS", key))
}

// HDel 删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略
// 返回被成功移除的域的数量，不包括被忽略的域
// HDEL key field [field ...]
func (c *cache) HDel(v ...interface{}) (int, error) {
	return redis.Int(c.DO("HDel", v...))
}

// HLen 返回哈希表 key 中域的数量
func (c *cache) HLen(key string) (int, error) {
	return redis.Int(c.DO("HLEN", key))
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
