package redis

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

type Executor struct {
	schema *Schema
	pool   *redis.Pool
}

func NewExecutor(schema *Schema) (*Executor, error) {
	if schema == nil {
		return nil, errors.New("schema can't be nil")
	}

	pool, err := NewPool(schema)
	if err != nil {
		return nil, err
	}

	return &Executor{
		schema: schema,
		pool:   pool,
	}, nil
}

func (x *Executor) Execute(commandName string, args ...interface{}) (reply interface{}, err error, closeErr error) {
	if commandName == "" {
		err = errors.New("commandName can't be empty")
		return
	}

	conn := x.pool.Get()
	defer func() {
		closeErr = conn.Close()
	}()

	reply, err = conn.Do(commandName, args...)
	return
}

func (x *Executor) Script(script *redis.Script, args ...interface{}) (reply interface{}, err error, closeErr error) {
	if script == nil {
		err = errors.New("script can't be nil")
		return
	}

	if args == nil {
		err = errors.New("args can't be nil")
		return
	}

	if len(args) == 0 {
		err = errors.New("args can't be empty")
		return
	}

	conn := x.pool.Get()
	defer func() {
		closeErr = conn.Close()
	}()

	reply, err = script.Do(conn, args...)
	return
}

func (x *Executor) GetSchema() *Schema {
	return x.schema
}

func (x *Executor) GetPool() *redis.Pool {
	return x.pool
}
