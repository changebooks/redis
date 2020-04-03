package redis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func (x *Executor) Get(key string) (value interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	value, err, closeErr = x.Execute("GET", key)
	return
}

func (x *Executor) Set(key string, value interface{}) (result interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	if value == nil {
		err = errors.New("value can't be nil")
		return
	}

	result, err, closeErr = x.Execute("SET", key, value)
	return
}

func (x *Executor) SetEx(key string, value interface{}, seconds int) (result interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	if value == nil {
		err = errors.New("value can't be nil")
		return
	}

	if seconds <= 0 {
		err = errors.New("seconds can't be less or equal than 0")
		return
	}

	result, err, closeErr = x.Execute("SETEX", key, seconds, value)
	return
}

func (x *Executor) Del(keys ...string) (affectedRows interface{}, err error, closeErr error) {
	if keys == nil {
		err = errors.New("keys can't be nil")
		return
	}

	var args []interface{}
	for _, k := range keys {
		if k != "" {
			args = append(args, k)
		}
	}

	if len(args) == 0 {
		err = errors.New("args can't be empty")
		return
	}

	affectedRows, err, closeErr = x.Execute("DEL", args...)
	return
}

func (x *Executor) Exists(key string) (result interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	result, err, closeErr = x.Execute("EXISTS", key)
	return
}

// -1 : Forever
// -2 : Nothing
func (x *Executor) Ttl(key string) (seconds interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	seconds, err, closeErr = x.Execute("TTL", key)
	return
}

func (x *Executor) MGet(keys ...string) (values []interface{}, err error, closeErr error) {
	if keys == nil {
		err = errors.New("keys can't be nil")
		return
	}

	if len(keys) == 0 {
		err = errors.New("keys can't be empty")
		return
	}

	var args []interface{}
	for n, k := range keys {
		if k == "" {
			err = fmt.Errorf("key-%d can't be empty", n)
			return
		}

		args = append(args, k)
	}

	reply, err, closeErr := x.Execute("MGET", args...)
	values, err = redis.Values(reply, err)
	return
}

func (x *Executor) MSet(data map[string]interface{}) (result interface{}, err error, closeErr error) {
	if data == nil {
		err = errors.New("data can't be nil")
		return
	}

	if len(data) == 0 {
		err = errors.New("data can't be empty")
		return
	}

	var args []interface{}
	for key, value := range data {
		if key == "" {
			err = errors.New("key can't be empty")
			return
		}

		if value == nil {
			err = fmt.Errorf("%s's value can't be nil", key)
			return
		}

		args = append(args, key, value)
	}

	result, err, closeErr = x.Execute("MSET", args...)
	return
}
