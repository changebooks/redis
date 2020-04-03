package redis

import "errors"

func (x *Executor) HIncrBy(key string, field string, increment int) (value interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	if field == "" {
		err = errors.New("field can't be empty")
		return
	}

	value, err, closeErr = x.Execute("HINCRBY", key, field, increment)
	return
}

func (x *Executor) HIncrByFloat(key string, field string, increment float64) (value interface{}, err error, closeErr error) {
	if key == "" {
		err = errors.New("key can't be empty")
		return
	}

	if field == "" {
		err = errors.New("field can't be empty")
		return
	}

	value, err, closeErr = x.Execute("HINCRBYFLOAT", key, field, increment)
	return
}
