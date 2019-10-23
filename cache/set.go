package cache

import (
	"github.com/gomodule/redigo/redis"
)

// SAdd 将一个或多个 member 元素加入到集合 key 当中
// 已经存在于集合的 member 元素将被忽略
// 被添加到集合中的新元素的数量，不包括被忽略的元素
// 第一个 v 为key，其它为 member
// SADD key member [member ...]
func (c *cache) SAdd(v ...interface{}) (int, error) {
	return redis.Int(c.DO("SADD", v...))
}

// SCard 返回集合 key 的基数(集合中元素的数量)
func (c *cache) SCard(key string) (int, error) {
	return redis.Int(c.DO("SCARD", key))
}

// SDiff 返回一个集合的全部成员，该集合是所有给定集合之间的差集
// 不存在的 key 被视为空集
func (c *cache) SDiff(key ...interface{}) ([]string, error) {
	return redis.Strings(c.DO("SDIFF", key...))
}

// SDiffStore 类似于 SDiff 命令，但它将结果保存到 destination 集合，而不是简单地返回结果集
// 如果 destination 集合已经存在，则将其覆盖。
// 第一个 key 为 destination
// SINTERSTORE destination key [key ...]
// 返回结果集中的成员数量
func (c *cache) SDiffStore(key ...interface{}) (int, error) {
	return redis.Int(c.DO("SDIFFSTORE", key...))
}

// SUnion 返回一个集合的全部成员，该集合是所有给定集合的并集
// 不存在的 key 被视为空集
func (c *cache) SUnion(key ...interface{}) ([]string, error) {
	return redis.Strings(c.DO("SUNION", key...))
}

// SUnionStore 类似于 SUnion 命令，但它将结果保存到 destination 集合，而不是简单地返回结果集
// 如果 destination 集合已经存在，则将其覆盖。
// 第一个 key 为 destination
// SINTERSTORE destination key [key ...]
// 返回结果集中的成员数量
func (c *cache) SUnionStore(key ...interface{}) (int, error) {
	return redis.Int(c.DO("SUNIONSTORE", key...))
}

// SInter 返回一个集合的全部成员，该集合是所有给定集合的交集
// 不存在的 key 被视为空集
func (c *cache) SInter(key ...interface{}) ([]string, error) {
	return redis.Strings(c.DO("SINTER", key...))
}

// SInterStore 类似于 SINTER 命令，但它将结果保存到 destination 集合，而不是简单地返回结果集
// 如果 destination 集合已经存在，则将其覆盖。
// 第一个 key 为 destination
// SINTERSTORE destination key [key ...]
// 返回结果集中的成员数量
func (c *cache) SInterStore(key ...interface{}) (int, error) {
	return redis.Int(c.DO("SINTERSTORE", key...))
}

// SIsMember 判断 member 元素是否集合 key 的成员
func (c *cache) SIsMember(key, member string) (bool, error) {
	return redis.Bool(c.DO("SISMEMBER", key, member))
}

// SMembers 返回集合 key 中的所有成员
func (c *cache) SMembers(key string) ([]string, error) {
	return redis.Strings(c.DO("SMEMBERS", key))
}

// SPop 返回集合中的一个随机元素，该元素会从集合中移除
// 返回被移除的随机元素
func (c *cache) SPop(key string) (string, error) {
	return redis.String(c.DO("SPOP", key))
}

// SRandMember 从集合中随机返回 count 个元素，这些元素不会从集合中移除
// 如果 count 为正数，且小于集合基数，那么命令返回一个包含 count 个元素的数组，数组中的元素各不相同。如果 count 大于等于集合基数，那么返回整个集合
// 如果 count 为负数，那么命令返回一个数组，数组中的元素可能会重复出现多次，而数组的长度为 count 的绝对值
func (c *cache) SRandMember(key string, count int) ([]string, error) {
	return redis.Strings(c.DO("SRANDMEMBER", key, count))
}

// SRem 移除集合 key 中的一个或多个 member 元素，不存在的 member 元素会被忽略
// 返回被成功移除的元素的数量，不包括被忽略的元素
// 第一个 v 为集合 key，其它为 member
// SREM key member [member ...]
func (c *cache) SRem(v ...interface{}) (int, error) {
	return redis.Int(c.DO("SREM", v...))
}
