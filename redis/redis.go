package db

import (
	"github.com/go-redis/redis"
	"github.com/woshihot/go-lib/utils/log"
	"time"
)

var (
	Redis       *Client
	RedisErrTag = "[Redis-error]"
)

func RedisInit(address string) {
	Redis = &Client{redis.NewClient(&redis.Options{
		Addr:               address,
		PoolSize:           200,
		MaxRetries:         10,
		ReadTimeout:        time.Second * 10,
		IdleTimeout:        30 * time.Second,
		IdleCheckFrequency: 30 * time.Second})}

	if nil == Redis {
		log.EF(RedisErrTag, "redis-server not start!")
	}
}

type Client struct {
	*redis.Client
}

func (r *Client) NewHash(key string) *Hash {

	return &Hash{key: key, rc: r}
}

func (r *Client) NewValue(key string) *Value {
	return &Value{key: key, rc: r}
}

//------------------------------------------------------------------------------
//hash map
type Hash struct {
	key    string
	suffix string
	rc     *Client
}

func (h *Hash) Key(key string) *Hash {
	h.key = key
	return h
}

func (h *Hash) Suffix(suffix string) *Hash {
	h.suffix = suffix
	return h
}

func (h *Hash) HGet(field string) *redis.StringCmd {
	return h.rc.HGet(h.key+h.suffix, field)
}

func (h *Hash) HExists(field string) *redis.BoolCmd {
	return h.rc.HExists(h.key+h.suffix, field)
}

func (h *Hash) HDel(fields ...string) *redis.IntCmd {
	return h.rc.HDel(h.key+h.suffix, fields...)

}

func (h *Hash) HGetAll() *redis.StringStringMapCmd {
	return h.rc.HGetAll(h.key + h.suffix)
}

func (h *Hash) HKeys() *redis.StringSliceCmd {
	return h.rc.HKeys(h.key + h.suffix)

}

func (h *Hash) HLen() *redis.IntCmd {
	return h.rc.HLen(h.key + h.suffix)

}

func (h *Hash) HSet(field string, value interface{}) *redis.BoolCmd {
	return h.rc.HSet(h.key+h.suffix, field, value)
}

//------------------------------------------------------------------------------
//string
type Value struct {
	key string
	rc  *Client
}

func (v *Value) Get() *redis.StringCmd {
	return v.rc.Get(v.key)
}
func (v *Value) Set(value interface{}, time time.Duration) *redis.StatusCmd {
	return v.rc.Set(v.key, value, time)
}

func (v *Value) Del() *redis.IntCmd {
	return v.rc.Del(v.key)
}
func (v *Value) Exists() *redis.IntCmd {
	return v.rc.Exists(v.key)
}

//------------------------------------------------------------------------------
//list
//type List model {
//	R
//}
