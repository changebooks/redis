package redis

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"time"
)

func NewPool(s *Schema) (*redis.Pool, error) {
	if s == nil {
		return nil, errors.New("schema can't be nil")
	}

	maxActive := s.GetMaxActive()
	maxIdle := s.GetMaxIdle()
	maxConnLifetime := s.GetMaxConnLifetime()
	idleTimeout := s.GetIdleTimeout()
	wait := s.GetWait()

	proto := s.GetProto()
	address := s.GetAddress()
	password := s.GetPassword()
	database := s.GetDatabase()
	connectTimeout := s.GetConnectTimeout()
	readTimeout := s.GetReadTimeout()
	writeTimeout := s.GetWriteTimeout()

	return &redis.Pool{
		MaxActive:       maxActive,
		MaxIdle:         maxIdle,
		MaxConnLifetime: maxConnLifetime,
		IdleTimeout:     idleTimeout,
		Wait:            wait,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				proto,
				address,
				redis.DialPassword(password),
				redis.DialDatabase(database),
				redis.DialConnectTimeout(connectTimeout),
				redis.DialReadTimeout(readTimeout),
				redis.DialWriteTimeout(writeTimeout),
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if s.TestOnBorrow != nil {
				return s.TestOnBorrow(c, t)
			} else {
				return nil
			}
		},
	}, nil
}
