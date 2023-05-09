package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var cacheUserNamePrefix = "cache:user:username:"

func (m *defaultUserModel) FindOneByUserName(ctx context.Context, username string) (*User, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserNamePrefix, username)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, username)
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

func (m *defaultUserModel) UpdateOptionalFiled(ctx context.Context, data *User, removeFiled []string) error {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Username, data.Password, data.Sex, data.Email, data.Info, data.Id)
	}, userIdKey)
	return err
}
