// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"zhihu/application/applet/internal/config"
	"zhihu/application/user/rpc/user"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  user.User
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		BizRedis: redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
