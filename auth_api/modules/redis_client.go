package modules

import (
	"github.com/gomodule/redigo/redis"
)

func RedisSet(c redis.Conn, key, value string) {
	c.Do("SET", key, value)
}

func RedisGet(c redis.Conn, key string) string  {
	s, _ := redis.String(c.Do("GET", key))
	return s
}
