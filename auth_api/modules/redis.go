package modules

import "github.com/gomodule/redigo/redis"

func GetRedis(r redis.Conn, key string) string {
	res, _ := redis.String(r.Do("GET", key))
	return res
}

func SetRedis(r redis.Conn, key, value string) {
	r.Do("SET", key, value)
}
