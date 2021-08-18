package redis

import (
	"github.com/dipper-mizar/sky/conf"
	"github.com/go-redis/redis"
)

type InitParams struct {
	Addr     string
	Password string
	Database int
}

type ParamSetter func(*InitParams)

func (params *InitParams) SetAddr(host string) ParamSetter {
	return func(params *InitParams) {
		params.Addr = host
	}
}

func (params *InitParams) SetPassword(password string) ParamSetter {
	return func(params *InitParams) {
		params.Password = password
	}
}

func (params *InitParams) SetDatabase(database int) ParamSetter {
	return func(params *InitParams) {
		params.Database = database
	}
}

func (params *InitParams) Init(functions ...ParamSetter) *redis.Client {
	p := &InitParams{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPassword,
		Database: conf.RedisDatabase,
	}
	for _, function := range functions {
		function(p)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     p.Addr,
		Password: p.Password, // No password set by default.
		DB:       p.Database,
	})
	if _, err := client.Ping().Result(); err != nil {
		// TODO: Write error into logger.
	}
	return client
}
