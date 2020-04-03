package redis

import (
	"errors"
	"strconv"
	"time"
)

type Profile struct {
	host            string        // 域名或Ip
	proto           string        // 协议，如：tcp，缺省：define.Proto
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
}

func NewProfile(data map[string]string) (*Profile, error) {
	if data == nil {
		return nil, errors.New("data can't be nil")
	}

	port := 0
	if data[ProfilePort] != "" {
		if p, err := strconv.ParseInt(data[ProfilePort], 10, 32); err == nil {
			port = int(p)
		} else {
			return nil, err
		}
	}

	database := 0
	if data[ProfileDatabase] != "" {
		if n, err := strconv.ParseInt(data[ProfileDatabase], 10, 32); err == nil {
			database = int(n)
		} else {
			return nil, err
		}
	}

	var connectTimeout time.Duration = 0
	if data[ProfileConnectTimeout] != "" {
		if d, err := strconv.ParseInt(data[ProfileConnectTimeout], 10, 64); err == nil {
			connectTimeout = time.Duration(d)
		} else {
			return nil, err
		}
	}

	var readTimeout time.Duration = 0
	if data[ProfileReadTimeout] != "" {
		if d, err := strconv.ParseInt(data[ProfileReadTimeout], 10, 64); err == nil {
			readTimeout = time.Duration(d)
		} else {
			return nil, err
		}
	}

	var writeTimeout time.Duration = 0
	if data[ProfileWriteTimeout] != "" {
		if d, err := strconv.ParseInt(data[ProfileWriteTimeout], 10, 64); err == nil {
			writeTimeout = time.Duration(d)
		} else {
			return nil, err
		}
	}

	maxActive := 0
	if data[ProfileMaxActive] != "" {
		if n, err := strconv.ParseInt(data[ProfileMaxActive], 10, 32); err == nil {
			maxActive = int(n)
		} else {
			return nil, err
		}
	}

	maxIdle := 0
	if data[ProfileMaxIdle] != "" {
		if n, err := strconv.ParseInt(data[ProfileMaxIdle], 10, 32); err == nil {
			maxIdle = int(n)
		} else {
			return nil, err
		}
	}

	var maxConnLifetime time.Duration = 0
	if data[ProfileMaxConnLifetime] != "" {
		if d, err := strconv.ParseInt(data[ProfileMaxConnLifetime], 10, 64); err == nil {
			maxConnLifetime = time.Duration(d)
		} else {
			return nil, err
		}
	}

	var idleTimeout time.Duration = 0
	if data[ProfileIdleTimeout] != "" {
		if d, err := strconv.ParseInt(data[ProfileIdleTimeout], 10, 64); err == nil {
			idleTimeout = time.Duration(d)
		} else {
			return nil, err
		}
	}

	wait := false
	if data[ProfileWait] != "" {
		if b, err := strconv.ParseBool(data[ProfileWait]); err == nil {
			wait = b
		} else {
			return nil, err
		}
	}

	host := data[ProfileHost]
	proto := data[ProfileProto]
	password := data[ProfilePassword]

	return &Profile{
		host:            host,
		proto:           proto,
		port:            port,
		password:        password,
		database:        database,
		connectTimeout:  connectTimeout,
		readTimeout:     readTimeout,
		writeTimeout:    writeTimeout,
		maxActive:       maxActive,
		maxIdle:         maxIdle,
		maxConnLifetime: maxConnLifetime,
		idleTimeout:     idleTimeout,
		wait:            wait,
	}, nil
}

func (x *Profile) GetHost() string {
	return x.host
}

func (x *Profile) GetProto() string {
	return x.proto
}

func (x *Profile) GetPort() int {
	return x.port
}

func (x *Profile) GetPassword() string {
	return x.password
}

func (x *Profile) GetDatabase() int {
	return x.database
}

func (x *Profile) GetConnectTimeout() time.Duration {
	return x.connectTimeout
}

func (x *Profile) GetReadTimeout() time.Duration {
	return x.readTimeout
}

func (x *Profile) GetWriteTimeout() time.Duration {
	return x.writeTimeout
}

func (x *Profile) GetMaxActive() int {
	return x.maxActive
}

func (x *Profile) GetMaxIdle() int {
	return x.maxIdle
}

func (x *Profile) GetMaxConnLifetime() time.Duration {
	return x.maxConnLifetime
}

func (x *Profile) GetIdleTimeout() time.Duration {
	return x.idleTimeout
}

func (x *Profile) GetWait() bool {
	return x.wait
}
