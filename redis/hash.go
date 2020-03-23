package redis

import "github.com/go-redis/redis"

type Hash struct {
	key
}

//HGet 获取存储在哈希表中指定字段的值
func (h *Hash) HGet(field string) *redis.StringCmd {
	return h.rc.HGet(h.Key(), field)
}

//HExists 查看哈希表中，指定的字段是否存在
func (h *Hash) HExists(field string) *redis.BoolCmd {
	return h.rc.HExists(h.Key(), field)
}

//HDel 删除一个或多个哈希表字段
func (h *Hash) HDel(fields ...string) *redis.IntCmd {
	return h.rc.HDel(h.Key(), fields...)
}

//HGetAll 获取在哈希表中指定的所有字段和值
func (h *Hash) HGetAll() *redis.StringStringMapCmd {
	return h.rc.HGetAll(h.Key())
}

//HKeys 获取所有哈希表中的字段
func (h *Hash) HKeys() *redis.StringSliceCmd {
	return h.rc.HKeys(h.Key())
}

//HLen 获取哈希表中字段的数量
func (h *Hash) HLen() *redis.IntCmd {
	return h.rc.HLen(h.Key())
}

//HSet 将哈希表中的字段 field 的值设为 value
func (h *Hash) HSet(field string, value interface{}) *redis.BoolCmd {
	return h.rc.HSet(h.Key(), field, value)
}

//HSetNX 只有在字段 field 不存在时，设置哈希表字段的值
func (h *Hash) HSetNX(field string, value interface{}) *redis.BoolCmd {
	return h.rc.HSetNX(h.Key(), field, value)
}

//HVals 获取哈希表中所有值
func (h *Hash) HVals() *redis.StringSliceCmd {
	return h.rc.HVals(h.Key())
}
