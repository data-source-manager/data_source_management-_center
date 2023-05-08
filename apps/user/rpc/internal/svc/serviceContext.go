package svc

import (
	"data_source_management_center/apps/user/model"
	"data_source_management_center/apps/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Db.DataSource)
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.MustNewRedis(c.Redis.RedisConf),
		UserModel:   model.NewUserModel(conn, c.Cache),
	}
}
