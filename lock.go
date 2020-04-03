package redis

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

func (x *Executor) Lock(name string, token string, seconds int) (result interface{}, err error, closeErr error) {
	if name == "" {
		err = errors.New("name can't be empty")
		return
	}

	if token == "" {
		err = errors.New("token can't be empty")
		return
	}

	if seconds <= 0 {
		err = errors.New("seconds can't be less or equal than 0")
		return
	}

	result, err, closeErr = x.Execute("SET", name, token, "EX", seconds, "NX")
	return
}

var unlockScript = redis.NewScript(1, UnlockScript)

func (x *Executor) Unlock(name string, token string) (result interface{}, err error, closeErr error) {
	if name == "" {
		err = errors.New("name can't be empty")
		return
	}

	if token == "" {
		err = errors.New("token can't be empty")
		return
	}

	result, err, closeErr = x.Script(unlockScript, name, token)
	return
}
