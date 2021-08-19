package gateway

import (
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"go-feedmaker/adapter/repository"
)

var (
	ErrRedisDisconnected = errors.New("gateway is not connected to Redis")
)

type (
	RedisGateway struct {
		Config     RedisConfig
		Dialer     RedisDialer
		connection RedisConnection
		pool       *redis.Pool
	}

	RedisConfig struct {
		Host        string
		Port        string
		ConnTimeout time.Duration `config:"conn_timeout"`
		PoolSize    int           `config:"pool_size"`
	}

	RedisDialer interface {
		Dial(network, addr string, options ...redis.DialOption) (RedisConnection, error)
	}

	RedisClient interface {
		Connection() redis.Conn
	}

	PubSub interface {
		Subscribe(channel ...interface{}) error
		Unsubscribe(channel ...interface{}) error
		Receive() interface{}
	}

	RedisConnection interface {
		redis.Conn
	}
)

func (c RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (r *RedisGateway) dial() (redis.Conn, error) {
	addr := r.Config.Addr()
	connectTimeout := redis.DialConnectTimeout(r.Config.ConnTimeout)
	return r.Dialer.Dial("tcp", addr, connectTimeout)
}

func (r *RedisGateway) Connect() error {
	r.pool = r.makePool()
	conn := r.pool.Get()
	defer conn.Close()
	return r.ping(conn)
}

func (r *RedisGateway) makePool() *redis.Pool {
	return &redis.Pool{
		Dial:    r.dial,
		MaxIdle: r.Config.PoolSize,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			return r.ping(c)
		},
		IdleTimeout: time.Minute,
	}
}

func (r *RedisGateway) ping(conn RedisConnection) error {
	_, err := conn.Do("PING")
	return err
}

func (r *RedisGateway) Connection() repository.Connection {
	return r.pool.Get()
}

func (r *RedisGateway) PubSub() repository.PubSub {
	return &redis.PubSubConn{Conn: r.pool.Get()}
}

func (r *RedisGateway) Disconnect() error {
	if r.pool == nil {
		return ErrRedisDisconnected
	}
	return r.pool.Close()
}
