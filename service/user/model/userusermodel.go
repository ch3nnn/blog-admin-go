package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserUserModel = (*customUserUserModel)(nil)

type (
	// UserUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserUserModel.
	UserUserModel interface {
		userUserModel
	}

	customUserUserModel struct {
		*defaultUserUserModel
	}
)

// NewUserUserModel returns a model for the database table.
func NewUserUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserUserModel {
	return &customUserUserModel{
		defaultUserUserModel: newUserUserModel(conn, c),
	}
}
