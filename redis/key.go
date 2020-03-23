package redis

import "github.com/go-redis/redis"

type key struct {
	rc  *Client
	key string
}

func (p *key) ChangeKey(key string) {
	p.key = key
}

func (p *key) Key() string {
	return p.key
}

func (p *key) Del() *redis.IntCmd {
	return p.rc.Del(p.key)
}

func (p *key) Exists() *redis.IntCmd {
	return p.rc.Exists(p.key)
}

func (p *key) Client() *Client {
	return p.rc
}
