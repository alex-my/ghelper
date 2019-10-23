package cache

import (
	"github.com/gomodule/redigo/redis"
)

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
