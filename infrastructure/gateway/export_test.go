package gateway

import "github.com/gomodule/redigo/redis"

func (r *RedisGateway) SetPool(pool *redis.Pool) {
	r.pool = pool
}
