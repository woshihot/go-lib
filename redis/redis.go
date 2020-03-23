//Package redis redis client
package redis

import (
	"github.com/go-redis/redis"
	"time"
)

//RedisCreate create a redis client
func RedisCreate(address string) *Client {
	return &Client{redis.NewClient(&redis.Options{
		Addr:               address,
		PoolSize:           200,
		MaxRetries:         10,
		ReadTimeout:        time.Second * 10,
		IdleTimeout:        30 * time.Second,
		IdleCheckFrequency: 30 * time.Second})}

}

type Client struct {
	*redis.Client
}

//NewHash create an redis hash
func (r *Client) NewHash(k string) *Hash {

	return &Hash{key: key{r, k}}
}

//NewValue create an redis value
func (r *Client) NewValue(k string) *Value {

	return &Value{key: key{r, k}}
}

//NewValue create an redis list
func (r *Client) NewList(k string) *List {

	return &List{key: key{r, k}}
}
