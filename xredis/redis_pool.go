package xredis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/nickxb/pkg/xjson"
	"github.com/nickxb/pkg/xsync"
	"sync"
	"time"
)

var (
	redisPoolLock = new(sync.RWMutex)
	redisPools    = make(map[string]*redis.Pool)
)

type RedisPoolConfig struct {
	Alias          string `json:"alias"`
	Address        string `json:"address"`
	Password       string `json:"password"`
	DB             *int   `json:"db"`
	ConnectTimeout int    `json:"connect_timeout"`
	ReadTimeout    int    `json:"read_timeout"`
	WriteTimeout   int    `json:"write_timeout"`
	Wait           bool   `json:"wait"`
	MaxIdle        int    `json:"max_idle"`
	IdleTimeout    int    `json:"idle_timeout"`
}

func InitRedisPool(configs []*RedisPoolConfig) error {
	for _, c := range configs {
		if _, ok := redisPools[c.Alias]; ok {
			return errors.New("duplicate redis pool: " + c.Alias)
		}
		p, err := createRedisPool(c)
		if err != nil {
			return errors.New(fmt.Sprintf("redis pool %s error %v", xjson.SafeMarshal(c), err))
		}
		xsync.WithLock(redisPoolLock, func() {
			redisPools[c.Alias] = p
		})
	}
	return nil
}

func createRedisPool(c *RedisPoolConfig) (*redis.Pool, error) {

	p := &redis.Pool{
		MaxIdle:     c.MaxIdle,
		IdleTimeout: time.Duration(c.IdleTimeout) * time.Second,
		Wait:        c.Wait,
		Dial: func() (conn redis.Conn, err error) {
			var options []redis.DialOption

			if c.ConnectTimeout != 0 {
				options = append(options, redis.DialConnectTimeout(time.Duration(c.ConnectTimeout)*time.Second))
			}
			if c.ReadTimeout != 0 {
				options = append(options, redis.DialReadTimeout(time.Duration(c.ReadTimeout)*time.Second))
			}
			if c.WriteTimeout != 0 {
				options = append(options, redis.DialWriteTimeout(time.Duration(c.WriteTimeout)*time.Second))
			}
			conn, err = redis.Dial(
				"tcp",
				c.Address,
				options...,
			)

			if err != nil {
				return nil, err
			}

			if c.Password != "" {
				if _, err := conn.Do("AUTH", c.Password); err != nil {
					conn.Close()
					return nil, err
				}
			}

			if c.DB != nil {
				if _, err := conn.Do("SELECT", *c.DB); err != nil {
					conn.Close()
					return nil, err
				}
			}

			return conn, nil
		},

		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}

	conn := p.Get()
	_, err := conn.Do("PING")
	return p, err
}

func GetRedisPool(alias string) *redis.Pool {
	redisPoolLock.RLock()
	defer redisPoolLock.RUnlock()
	return redisPools[alias]
}

func WithConn(poolAlias string, fn func(conn redis.Conn) error) error {
	conn := GetRedisPool(poolAlias).Get()
	defer conn.Close()
	return fn(conn)
}
