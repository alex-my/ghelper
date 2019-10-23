package cache

import (
	"github.com/gomodule/redigo/redis"
)

// ZAdd 将一个或多个 member 元素及其 score 值加入到有序集 key 当中
// 返回被成功添加的新成员的数量，不包括那些被更新的、已经存在的成员
// 第一个v 是 key
// ZADD key score member [[score member] [score member] ...]
func (c *cache) ZAdd(v ...interface{}) (int, error) {
	if len(v) == 0 {
		return 0, nil
	}
	if len(v)%2 == 0 {
		return 0, ErrInvalidParamCount
	}

	return redis.Int(c.DO("ZADD", v...))
}

// ZCard 返回有序集 key 的基数
func (c *cache) ZCard(key string) (int, error) {
	return redis.Int(c.DO("ZCARD", key))
}

// ZCount 返回有序集 key 中， score 值在 min 和 max 之间(默认包括 score 值等于 min 或 max )的成员的数量
func (c *cache) ZCount(key string, min int, max int) (int, error) {
	return redis.Int(c.DO("ZCOUNT", key, min, max))
}

// ZIncrby 为有序集 key 的成员 member 的 score 值加上增量 increment
// increment 可以为负值
// 返回 member 成员的新 score 值
func (c *cache) ZIncrby(key, member string, incrment int) (int, error) {
	return redis.Int(c.DO("ZINCRBY", key, incrment, member))
}

// ZRange 返回有序集 key 中，指定下标区间内的成员
func (c *cache) ZRange(key string, start, stop int) ([]string, error) {
	return redis.Strings(c.DO("ZRANGE", key, start, stop))
}

// ZRangeWithScores 返回有序集 key 中，指定下标区间内的成员和 score值
func (c *cache) ZRangeWithScores(key string, start, stop int) ([]string, error) {
	return redis.Strings(c.DO("ZRANGE", key, start, stop, "WITHSCORES"))
}

// ZScore 返回有序集 key 中，成员 member 的 score 值
func (c *cache) ZScore(key, member string) (int, error) {
	return redis.Int(c.DO("ZSCORE", key, member))
}

// ZRank 返回有序集 key 中成员 member 的下标
// 其中有序集成员按 score 值递增(从小到大)顺序排列
func (c *cache) ZRank(key, member string) (int, error) {
	return redis.Int(c.DO("ZRANK", key, member))
}

// ZRangeByScore 返回有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成
// 有序集成员按 score 值递增(从小到大)次序排列
func (c *cache) ZRangeByScore(key string, min, max, limit, offset, count int) ([]string, error) {
	return redis.Strings(c.DO("ZRANGEBYSCORE", key, min, max, "WITHSCORES", limit, offset, count))
}

// ZRevRank 返回有序集 key 中成员 member 的排名。其中有序集成员按 score 值递减(从大到小)排序
func (c *cache) ZRevRank(key, member string) (int, error) {
	return redis.Int(c.DO("ZREVRANK", key, member))
}

// ZRevRangeByScore 返回有序集 key 中， score 值介于 max 和 min 之间(默认包括等于 max 或 min )的所有的成员
// 有序集成员按 score 值递减(从大到小)的次序排列
func (c *cache) ZRevRangeByScore(key string, max, min, limit, offset, count int) ([]string, error) {
	return redis.Strings(c.DO("ZREVRANGEBYSCORE", key, max, min, "WITHSCORES", limit, offset, count))
}

// ZRevRange 返回有序集 key 中，指定下标区间内的成员
func (c *cache) ZRevRange(key string, start, stop int) ([]string, error) {
	return redis.Strings(c.DO("ZREVRANGE", key, start, stop, "WITHSCORES"))
}

// ZRemRangeByScore 移除有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员
// 移除有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员
// 返回被移除成员的数量
func (c *cache) ZRemRangeByScore(key string, min, max int) (int, error) {
	return redis.Int(c.DO("ZREMRANGEBYSCORE", key, min, max))
}

// ZRemRangeByRank 移除有序集 key 中，指定下标区间内的所有成员
// 返回被移除成员的数量
func (c *cache) ZRemRangeByRank(key string, start, stop int) (int, error) {
	return redis.Int(c.DO("ZREMRANGEBYRANK", key, start, stop))
}

// ZRem 移除有序集 key 中的一个或多个成员，不存在的成员将被忽略
// 第一个 v 为 key
// 返回被成功移除的成员的数量，不包括被忽略的成员
// ZREM key member [member ...]
func (c *cache) ZRem(v ...interface{}) (int, error) {
	return redis.Int(c.DO("ZREM", v...))
}
