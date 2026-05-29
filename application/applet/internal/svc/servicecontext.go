// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"zhihu/application/applet/internal/config"
	"zhihu/application/user/rpc/user"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  user.User
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	userCli:= zrpc.MustNewClient(c.UserRPC)
	return &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUser(userCli),
		BizRedis: redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
