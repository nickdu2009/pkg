package redisimpl

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/nickxb/pkg/xqueue"
)

type redisImpl struct {
	pool *redis.Pool
}

func New(pool *redis.Pool) xqueue.Queue  {
	return redisImpl{
		pool: pool,
	}
}

func (ri redisImpl) Push(topic, data []byte) error {
	conn := ri.pool.Get()
	defer conn.Close()
	_, err := conn.Do("LPUSH", string(topic), string(data))
	return err
}

func (ri redisImpl) Pull(topic []byte) (data []byte, err error) {
	for {
		data, err = ri.pullOnce(topic)
		if err == redis.ErrNil {
			return nil, xqueue.ErrNil
		} else {
			return
		}
	}
}

func (ri redisImpl) pullOnce(topic []byte) ([]byte, error) {
	conn := ri.pool.Get()
	defer conn.Close()
	arr, err := redis.Values(conn.Do("BRPOP", string(topic), 5))
	if err != nil {
		return nil, err
	}
	if len(arr) != 2 {
		return nil, errors.New("bad values")
	}
	data, ok := arr[1].([]byte)
	if !ok {
		return nil, errors.New("bad values")
	}
	return data, nil
}
