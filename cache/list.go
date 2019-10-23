package cache

import (
	"github.com/gomodule/redigo/redis"
)

// LPush 将一个或多个值 value 插入到列表 key 的表头(最左边)
// LPUSH key value [value ...]
func (c *cache) LPush(v ...interface{}) (int, error) {
	return redis.Int(c.DO("LPUSH", v...))
}

// LPop 移除并返回列表 key 的头元素
func (c *cache) LPop(key string) (string, error) {
	return redis.String(c.DO("LPOP", key))
}

// RPush 将一个或多个值 value 插入到列表 key 的表尾(最右边)
// RPUSH key value [value ...]
func (c *cache) RPush(v ...interface{}) (int, error) {
	return redis.Int(c.DO("RPUSH", v...))
}

// RPop 移除并返回列表 key 的尾元素
func (c *cache) RPop(key string) (string, error) {
	return redis.String(c.DO("RPOP", key))
}

// RPopLPush 在一个原子时间内，执行以下两个动作
// 将列表 source 中的最后一个元素(尾元素)弹出，并返回给客户端
// 将列表 source 中的最后一个元素(尾元素)弹出，并返回给客户端
// source 和 destination 可以相同
func (c *cache) RPopLPush(source, destination string) (string, error) {
	return redis.String(c.DO("RPOPLPUSH", source, destination))
}

// LTrim 对一个列表进行修剪
// 让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除
// 如果 start > stop 会清空列表
func (c *cache) LTrim(key string, start, stop int) error {
	_, err := c.DO("LTRIM", key, start, stop)
	return err
}

// LSet 让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除
func (c *cache) LSet(key string, index int, value string) error {
	_, err := c.DO("LSET", key, index, value)
	return err
}

// LRem 根据参数 count 的值，移除列表中与参数 value 相等的元素
// count > 0 : 从表头开始向表尾搜索，移除与 value 相等的元素，数量为 count
// count < 0 : 从表尾开始向表头搜索，移除与 value 相等的元素，数量为 count 的绝对值
// count = 0 : 移除表中所有与 value 相等的值
// 返回被移除元素的数量
func (c *cache) LRem(key string, count int, value string) (int, error) {
	return redis.Int(c.DO("LREM", key, count, value))
}

// LRange 返回列表 key 中指定区间内的元素，区间以偏移量 start 和 stop 指定
func (c *cache) LRange(key string, start, stop int) ([]string, error) {
	return redis.Strings(c.DO("LRANGE", key, start, stop))
}

// LLen 返回列表 key 的长度
func (c *cache) LLen(key string) (int, error) {
	return redis.Int(c.DO("LLEN", key))
}

// LInsertBefore 将值 value 插入到列表 key 当中，位于值 pivot 之前
func (c *cache) LInsertBefore(key, pivot, value string) (int, error) {
	return redis.Int(c.DO("LINSERT", key, "BEFORE", pivot, value))
}

// LInsertAfter 将值 value 插入到列表 key 当中，位于值 pivot 之后
func (c *cache) LInsertAfter(key, pivot, value string) (int, error) {
	return redis.Int(c.DO("LINSERT", key, "AFTER", pivot, value))
}

// LIndex 返回列表 key 中，下标为 index 的元素
func (c *cache) LIndex(key string, index int) (string, error) {
	return redis.String(c.DO("LINDEX", key, index))
}
