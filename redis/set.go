package redis

import "github.com/go-redis/redis"

type Set struct {
	key
}

//SAdd 向集合添加一个或多个成员
func (p *Set) SAdd(members interface{}) *redis.IntCmd {
	return p.rc.SAdd(p.Key(), members)
}

//SLen 返回集合中元素的数量
func (p *Set) SLen() *redis.IntCmd {
	return p.rc.SCard(p.Key())
}

//SDiff 返回给定集合之间的差集。不存在的集合 key 将视为空集
func (p *Set) SDiff(keys ...string) *redis.StringSliceCmd {
	newKeys := appenKeys(p.Key(), keys)
	return p.rc.SDiff(newKeys...)
}

//SDiff 返回给定集合之间的交集。不存在的集合 key 将视为空集
func (p *Set) SInter(keys ...string) *redis.StringSliceCmd {
	newKeys := appenKeys(p.Key(), keys)
	return p.rc.SInter(newKeys...)
}

//SIsMember 判断成员元素是否是集合的成员
func (p *Set) SIsMember(member interface{}) *redis.BoolCmd {
	return p.rc.SIsMember(p.Key(), member)
}

//SPop 移除并返回集合中的一个随机元素
func (p *Set) SPop() *redis.StringCmd {
	return p.rc.SPop(p.Key())
}

//SRandMember 返回集合中一个随机数
func (p *Set) SRandMember() *redis.StringCmd {
	return p.rc.SRandMember(p.Key())
}

//SRandMemberN 返回集合中一个或多个随机数
func (p *Set) SRandMemberN(count int64) *redis.StringSliceCmd {
	return p.rc.SRandMemberN(p.Key(), count)
}

//SUnion 返回所有给定集合的并集
func (p *Set) SUnion(keys ...string) *redis.StringSliceCmd {
	newKeys := appenKeys(p.Key(), keys)
	return p.rc.SUnion(newKeys...)
}

func appenKeys(key string, keys []string) []string {
	newKey := make([]string, 0)
	newKey = append(newKey, key)
	newKey = append(newKey, keys...)
	return newKey
}
