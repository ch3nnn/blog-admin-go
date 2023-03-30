package svc

import (
	"github.com/ch3nnn/blog-admin-go/service/user/model"
	"github.com/ch3nnn/blog-admin-go/service/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserUser model.UserUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:   c,
		UserUser: model.NewUserUserModel(conn, c.CacheRedis),
	}
}
