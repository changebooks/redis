package redis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func (x *Executor) HGet(key string, field string) (value interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	if field == "" {
		err = errors.New("field can't be empty")
		return
	}

	value, err, closeErr = x.Execute("HGET", key, field)
	return
}

func (x *Executor) HSet(key string, field string, value interface{}) (result interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	if field == "" {
		err = errors.New("field can't be empty")
		return
	}

	if value == nil {
		err = errors.New("value can't be nil")
		return
	}

	result, err, closeErr = x.Execute("HSET", key, field, value)
	return
}

func (x *Executor) HDel(key string, fields ...string) (affectedRows interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	if fields == nil {
		err = errors.New("fields can't be nil")
		return
	}

	if len(fields) == 0 {
		err = errors.New("fields can't be empty")
		return
	}

	var args []interface{}
	args = append(args, key)

	for _, f := range fields {
		if f != "" {
			args = append(args, f)
		}
	}

	if len(args) <= 1 {
		err = errors.New("field must be contained in args")
		return
	}

	affectedRows, err, closeErr = x.Execute("HDEL", args...)
	return
}

func (x *Executor) HExists(key string, field string) (result interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	if field == "" {
		err = errors.New("field can't be empty")
		return
	}

	result, err, closeErr = x.Execute("HEXISTS", key, field)
	return
}

func (x *Executor) HMGet(key string, fields ...string) (values []interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	if fields == nil {
		err = errors.New("fields can't be nil")
		return
	}

	if len(fields) == 0 {
		err = errors.New("fields can't be empty")
		return
	}

	var args []interface{}
	args = append(args, key)

	for n, f := range fields {
		if f == "" {
			err = fmt.Errorf("field-%d can't be empty", n)
			return
		}

		args = append(args, f)
	}

	reply, err, closeErr := x.Execute("HMGET", args...)
	values, err = redis.Values(reply, err)
	return
}

func (x *Executor) HMSet(key string, data map[string]interface{}) (result interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	if data == nil {
		err = errors.New("data can't be nil")
		return
	}

	if len(data) == 0 {
		err = errors.New("data can't be empty")
		return
	}

	var args []interface{}
	args = append(args, key)

	for field, value := range data {
		if field == "" {
			err = errors.New("field can't be empty")
			return
		}

		if value == nil {
			err = fmt.Errorf("%s's value can't be nil", field)
			return
		}

		args = append(args, field, value)
	}

	result, err, closeErr = x.Execute("HMSET", args...)
	return
}

func (x *Executor) HGetAll(key string) (values []interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	reply, err, closeErr := x.Execute("HGETALL", key)
	values, err = redis.Values(reply, err)
	return
}

func (x *Executor) HKeys(key string) (fields []interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	reply, err, closeErr := x.Execute("HKEYS", key)
	fields, err = redis.Values(reply, err)
	return
}

func (x *Executor) HVals(key string) (values []interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	reply, err, closeErr := x.Execute("HVALS", key)
	values, err = redis.Values(reply, err)
	return
}
