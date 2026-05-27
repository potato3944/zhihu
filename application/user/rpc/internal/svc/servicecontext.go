package svc

import (
	"zhihu/application/user/rpc/internal/config"
	"zhihu/application/user/rpc/internal/model"
)

type ServiceContext struct {
	UserModel model.UserModel
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
