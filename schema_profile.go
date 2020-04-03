package redis

import "errors"

func NewSchema(p *Profile) (*Schema, error) {
	if p == nil {
		return nil, errors.New("profile can't be nil")
	}

	host := p.GetHost()
	proto := p.GetProto()
	port := p.GetPort()
	password := p.GetPassword()
	database := p.GetDatabase()
	connectTimeout := p.GetConnectTimeout()
	readTimeout := p.GetReadTimeout()
	writeTimeout := p.GetWriteTimeout()
	maxActive := p.GetMaxActive()
	maxIdle := p.GetMaxIdle()
	maxConnLifetime := p.GetMaxConnLifetime()
	idleTimeout := p.GetIdleTimeout()
	wait := p.GetWait()

	builder := &SchemaBuilder{}
	builder.
		SetProto(proto).
		SetHost(host).
		SetPort(port).
		SetPassword(password).
		SetDatabase(database).
		SetConnectTimeout(connectTimeout).
		SetReadTimeout(readTimeout).
		SetWriteTimeout(writeTimeout).
		SetMaxActive(maxActive).
		SetMaxIdle(maxIdle).
		SetMaxConnLifetime(maxConnLifetime).
		SetIdleTimeout(idleTimeout).
		SetWait(wait)

	return builder.Build()
}
