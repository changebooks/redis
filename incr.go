package redis

import "errors"

func (x *Executor) Incr(key string) (value interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	value, err, closeErr = x.Execute("INCR", key)
	return
}

func (x *Executor) Decr(key string) (value interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	value, err, closeErr = x.Execute("DECR", key)
	return
}

func (x *Executor) IncrBy(key string, increment int) (value interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	value, err, closeErr = x.Execute("INCRBY", key, increment)
	return
}

func (x *Executor) DecrBy(key string, decrement int) (value interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	value, err, closeErr = x.Execute("DECRBY", key, decrement)
	return
}

func (x *Executor) IncrByFloat(key string, increment float64) (value interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	value, err, closeErr = x.Execute("INCRBYFLOAT", key, increment)
	return
}
