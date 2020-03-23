package redis

import "github.com/go-redis/redis"

type List struct {
	key
}

//func (p *List) BlPop(time time.Duration)*redis.StringSliceCmd {
//	return p.rc.BLPop(time, p.Key())
//}
//
//func (p *List) BRPop(time time.Duration)*redis.StringSliceCmd {
//	return p.rc.BRPop(time, p.Key())
//}

//LLen 获取列表长度
func (p *List) LLen() *redis.IntCmd {
	return p.rc.LLen(p.Key())
}

//LPush 将一个或多个值插入到列表头部。
//如果 key 不存在，一个空列表会被创建并执行 LPUSH 操作
//当 key 存在但不是列表类型时，返回一个错误
func (p *List) LPush(value ...interface{}) *redis.IntCmd {
	return p.rc.LPush(p.Key(), value)
}

//LPushX 将一个值插入到已存在的列表头部，列表不存在时操作无效
func (p *List) LPushX(value ...interface{}) *redis.IntCmd {
	return p.rc.LPushX(p.Key(), value)
}

//LSet 通过索引来设置元素的值
func (p *List) LSet(index int64, value interface{}) *redis.StatusCmd {
	return p.rc.LSet(p.Key(), index, value)
}

//RPushX 用于将一个值插入到已存在的列表尾部。如果列表不存在，操作无效
func (p *List) RPushX(value interface{}) *redis.IntCmd {
	return p.rc.RPushX(p.Key(), value)
}

//RPush 用于将一个或多个值插入到列表的尾部(最右边)
//如果列表不存在，一个空列表会被创建并执行。
//当列表存在但不是列表类型时，返回一个错误
func (p *List) RPush(value ...interface{}) *redis.IntCmd {
	return p.rc.RPush(p.Key(), value)
}

//LIndex 通过索引获取列表中的元素。
//你也可以使用负数下标，以 -1 表示列表的最后一个元素
func (p *List) LIndex(index int64) *redis.StringCmd {
	return p.rc.LIndex(p.Key(), index)
}

//LPop 移出并获取列表的第一个元素
func (p *List) LPop() *redis.StringCmd {
	return p.rc.LPop(p.Key())
}

//RPop 移除列表的最后一个元素，返回值为移除的元素
func (p *List) RPop() *redis.StringCmd {
	return p.rc.RPop(p.Key())
}
