package redis

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/huangjiasingle/suyi/pkg/config/generic"
)

var Client *redis.Client

func Init(opt *generic.RedisOptions) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", opt.IP, opt.Port), // redis port
		Username: opt.Username,
		Password: opt.Password,
		DB:       opt.DB,
		PoolSize: opt.PoolSzie,
	})
	_, err := Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
