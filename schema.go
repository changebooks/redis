package redis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strings"
	"sync"
	"time"
)

type Schema struct {
	TestOnBorrow    func(c redis.Conn, t time.Time) error
	proto           string        // 协议，如：tcp，缺省：define.Proto
	host            string        // 域名或Ip
	port            int           // 端口，如：6379，缺省：define.Port
	password        string        // 密码
	database        int           // 数据库，缺省：0
	connectTimeout  time.Duration // 连接超时，毫秒
	readTimeout     time.Duration // 读超时，毫秒
	writeTimeout    time.Duration // 写超时，毫秒
	maxActive       int           // 最大连接数
	maxIdle         int           // 最大空闲连接数，< maxActive
	maxConnLifetime time.Duration // 连接最大生命周期，缺省：0-不设置，永不过期
	idleTimeout     time.Duration // 空闲连接超时，小于服务端的连接超时，缺省：0-不设置，永不过期
	wait            bool          // > maxActive？true：等待、false：报错
	address         string        // 域名或Ip + 端口
}

func (x *Schema) ToString() string {
	return fmt.Sprintf("proto:           %v\n"+
		"host:            %v\n"+
		"port:            %v\n"+
		"password:        %v\n"+
		"database:        %v\n"+
		"connectTimeout:  %v\n"+
		"readTimeout:     %v\n"+
		"writeTimeout:    %v\n"+
		"maxActive:       %v\n"+
		"maxIdle:         %v\n"+
		"maxConnLifetime: %v\n"+
		"idleTimeout:     %v\n"+
		"wait:            %v\n"+
		"address:         %v\n",
		x.GetProto(), x.GetHost(), x.GetPort(), x.GetPassword(), x.GetDatabase(),
		x.GetConnectTimeout(), x.GetReadTimeout(), x.GetWriteTimeout(),
		x.GetMaxActive(), x.GetMaxIdle(), x.GetMaxConnLifetime(), x.GetIdleTimeout(),
		x.GetWait(), x.GetAddress(),
	)
}

func (x *Schema) GetProto() string {
	return x.proto
}

func (x *Schema) GetHost() string {
	return x.host
}

func (x *Schema) GetPort() int {
	return x.port
}

func (x *Schema) GetPassword() string {
	return x.password
}

func (x *Schema) GetDatabase() int {
	return x.database
}

func (x *Schema) GetConnectTimeout() time.Duration {
	return x.connectTimeout
}

func (x *Schema) GetReadTimeout() time.Duration {
	return x.readTimeout
}

func (x *Schema) GetWriteTimeout() time.Duration {
	return x.writeTimeout
}

func (x *Schema) GetMaxActive() int {
	return x.maxActive
}

func (x *Schema) GetMaxIdle() int {
	return x.maxIdle
}

func (x *Schema) GetMaxConnLifetime() time.Duration {
	return x.maxConnLifetime
}

func (x *Schema) GetIdleTimeout() time.Duration {
	return x.idleTimeout
}

func (x *Schema) GetWait() bool {
	return x.wait
}

func (x *Schema) GetAddress() string {
	return x.address
}

type SchemaBuilder struct {
	mu              sync.Mutex // ensures atomic writes; protects the following fields
	proto           string
	host            string
	port            int
	password        string
	database        int
	connectTimeout  time.Duration
	readTimeout     time.Duration
	writeTimeout    time.Duration
	maxActive       int
	maxIdle         int
	maxConnLifetime time.Duration
	idleTimeout     time.Duration
	wait            bool
	address         string
}

func (x *SchemaBuilder) Build() (*Schema, error) {
	if x.host == "" {
		return nil, errors.New("host can't be empty")
	}

	if x.port < 0 {
		return nil, errors.New("port can't be less than 0")
	}

	if x.database < 0 {
		return nil, errors.New("database can't be less than 0")
	}

	if x.connectTimeout < 0 {
		return nil, errors.New("connect timeout can't be less than 0")
	}

	if x.readTimeout < 0 {
		return nil, errors.New("read timeout can't be less than 0")
	}

	if x.writeTimeout < 0 {
		return nil, errors.New("write timeout can't be less than 0")
	}

	if x.maxActive < 0 {
		return nil, errors.New("max active can't be less than 0")
	}

	if x.maxIdle < 0 {
		return nil, errors.New("max idle can't be less than 0")
	}

	if x.maxConnLifetime < 0 {
		return nil, errors.New("max conn lifetime can't be less than 0")
	}

	if x.idleTimeout < 0 {
		return nil, errors.New("idle timeout can't be less than 0")
	}

	if x.maxIdle > 0 && x.maxActive > 0 && x.maxIdle > x.maxActive {
		return nil, errors.New("max idle can't be greater than max active")
	}

	proto := x.proto
	if proto == "" {
		proto = Proto
	}

	port := x.port
	if port == 0 {
		port = Port
	}

	return &Schema{
		proto:           proto,
		host:            x.host,
		port:            port,
		password:        x.password,
		database:        x.database,
		connectTimeout:  x.connectTimeout,
		readTimeout:     x.readTimeout,
		writeTimeout:    x.writeTimeout,
		maxActive:       x.maxActive,
		maxIdle:         x.maxIdle,
		maxConnLifetime: x.maxConnLifetime,
		idleTimeout:     x.idleTimeout,
		wait:            x.wait,
		address:         fmt.Sprintf("%s:%d", x.host, port),
	}, nil
}

func (x *SchemaBuilder) SetProto(s string) *SchemaBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.proto = s
	return x
}

func (x *SchemaBuilder) SetHost(s string) *SchemaBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.host = s
	return x
}

func (x *SchemaBuilder) SetPort(p int) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.port = p
	return x
}

func (x *SchemaBuilder) SetPassword(s string) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.password = s
	return x
}

func (x *SchemaBuilder) SetDatabase(n int) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.database = n
	return x
}

func (x *SchemaBuilder) SetConnectTimeout(d time.Duration) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.connectTimeout = d
	return x
}

func (x *SchemaBuilder) SetReadTimeout(d time.Duration) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.readTimeout = d
	return x
}

func (x *SchemaBuilder) SetWriteTimeout(d time.Duration) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.writeTimeout = d
	return x
}

func (x *SchemaBuilder) SetMaxActive(n int) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.maxActive = n
	return x
}

func (x *SchemaBuilder) SetMaxIdle(n int) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.maxIdle = n
	return x
}

func (x *SchemaBuilder) SetMaxConnLifetime(d time.Duration) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.maxConnLifetime = d
	return x
}

func (x *SchemaBuilder) SetIdleTimeout(d time.Duration) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.idleTimeout = d
	return x
}

func (x *SchemaBuilder) SetWait(b bool) *SchemaBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.wait = b
	return x
}
