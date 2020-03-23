package redis

import (
	"github.com/go-redis/redis"
	"time"
)

//string
type Value struct {
	key
}

//Set 设置指定 key 的值
func (p *Value) Set(value interface{}, time time.Duration) *redis.StatusCmd {
	return p.rc.Set(p.Key(), value, time)
}

//Get 获取指定 key 的值
func (p *Value) Get() *redis.StringCmd {
	return p.rc.Get(p.Key())
}

//SetNx 只有在 key 不存在时设置 key 的值
func (p *Value) SetNx(value interface{}, time time.Duration) *redis.BoolCmd {
	return p.rc.SetNX(p.Key(), value, time)
}

//SetXX 只在键已经存在时，才对键进行设置操作K
func (p *Value) SetXX(value interface{}, time time.Duration) *redis.BoolCmd {
	return p.rc.SetXX(p.key.key, value, time)
}

//Append 如果 key 已经存在并且是一个字符串，将指定的 value 追加到该 key 原来值（value）的末尾
func (p *Value) Append(value string) *redis.IntCmd {
	return p.rc.Append(p.Key(), value)
}
