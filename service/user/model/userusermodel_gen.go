// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userUserFieldNames          = builder.RawFieldNames(&UserUser{})
	userUserRows                = strings.Join(userUserFieldNames, ",")
	userUserRowsExpectAutoSet   = strings.Join(stringx.Remove(userUserFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userUserRowsWithPlaceHolder = strings.Join(stringx.Remove(userUserFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserUserIdPrefix       = "cache:userUser:id:"
	cacheUserUserUsernamePrefix = "cache:userUser:username:"
)

type (
	userUserModel interface {
		Insert(ctx context.Context, data *UserUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserUser, error)
		FindOneByUsername(ctx context.Context, username string) (*UserUser, error)
		Update(ctx context.Context, data *UserUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserUserModel struct {
		sqlc.CachedConn
		table string
	}

	UserUser struct {
		Id          int64          `db:"id"`
		Password    string         `db:"password"`
		LastLogin   sql.NullTime   `db:"last_login"`
		IsSuperuser int64          `db:"is_superuser"`
		Username    string         `db:"username"`
		FirstName   string         `db:"first_name"`
		LastName    string         `db:"last_name"`
		Email       string         `db:"email"`
		IsStaff     int64          `db:"is_staff"`
		IsActive    int64          `db:"is_active"`
		DateJoined  time.Time      `db:"date_joined"`
		Nickname    sql.NullString `db:"nickname"`
		Text        string         `db:"text"`
		Link        string         `db:"link"`
		AvatarUrl   string         `db:"avatar_url"`
	}
)

func newUserUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserUserModel {
	return &defaultUserUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_user`",
	}
}

func (m *defaultUserUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userUserIdKey := fmt.Sprintf("%s%v", cacheUserUserIdPrefix, id)
	userUserUsernameKey := fmt.Sprintf("%s%v", cacheUserUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userUserIdKey, userUserUsernameKey)
	return err
}

func (m *defaultUserUserModel) FindOne(ctx context.Context, id int64) (*UserUser, error) {
	userUserIdKey := fmt.Sprintf("%s%v", cacheUserUserIdPrefix, id)
	var resp UserUser
	err := m.QueryRowCtx(ctx, &resp, userUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userUserRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserUserModel) FindOneByUsername(ctx context.Context, username string) (*UserUser, error) {
	userUserUsernameKey := fmt.Sprintf("%s%v", cacheUserUserUsernamePrefix, username)
	var resp UserUser
	err := m.QueryRowIndexCtx(ctx, &resp, userUserUsernameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, username); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserUserModel) Insert(ctx context.Context, data *UserUser) (sql.Result, error) {
	userUserIdKey := fmt.Sprintf("%s%v", cacheUserUserIdPrefix, data.Id)
	userUserUsernameKey := fmt.Sprintf("%s%v", cacheUserUserUsernamePrefix, data.Username)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userUserRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Password, data.LastLogin, data.IsSuperuser, data.Username, data.FirstName, data.LastName, data.Email, data.IsStaff, data.IsActive, data.DateJoined, data.Nickname, data.Text, data.Link, data.AvatarUrl)
	}, userUserIdKey, userUserUsernameKey)
	return ret, err
}

func (m *defaultUserUserModel) Update(ctx context.Context, newData *UserUser) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	userUserIdKey := fmt.Sprintf("%s%v", cacheUserUserIdPrefix, data.Id)
	userUserUsernameKey := fmt.Sprintf("%s%v", cacheUserUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userUserRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Password, newData.LastLogin, newData.IsSuperuser, newData.Username, newData.FirstName, newData.LastName, newData.Email, newData.IsStaff, newData.IsActive, newData.DateJoined, newData.Nickname, newData.Text, newData.Link, newData.AvatarUrl, newData.Id)
	}, userUserIdKey, userUserUsernameKey)
	return err
}

func (m *defaultUserUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserUserIdPrefix, primary)
}

func (m *defaultUserUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userUserRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserUserModel) tableName() string {
	return m.table
}
