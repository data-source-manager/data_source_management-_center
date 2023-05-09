package svc

import (
	"data_source_management_center/apps/user/api/internal/config"
	"data_source_management_center/apps/user/rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpc)),
	}
}
