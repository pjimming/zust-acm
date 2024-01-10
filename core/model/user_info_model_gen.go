// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userInfoFieldNames          = builder.RawFieldNames(&UserInfo{})
	userInfoRows                = strings.Join(userInfoFieldNames, ",")
	userInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(userInfoFieldNames, "`id`", "`gmt_created`", "`gmt_updated`"), ",")
	userInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(userInfoFieldNames, "`id`", "`gmt_created`", "`gmt_updated`"), "=?,") + "=?"
)

type (
	userInfoModel interface {
		Insert(ctx context.Context, data *UserInfo) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*UserInfo, error)
		FindOneByUsername(ctx context.Context, username string) (*UserInfo, error)
		Update(ctx context.Context, data *UserInfo) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultUserInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserInfo struct {
		Id         uint64    `db:"id"`          // 序号
		GmtCreated time.Time `db:"gmt_created"` // 创建时间
		GmtUpdated time.Time `db:"gmt_updated"` // 更新时间
		IsDelete   uint64    `db:"is_delete"`   // 是否删除，0-未删除，1-删除，默认为0
		Username   string    `db:"username"`    // 账号
		Email      string    `db:"email"`       // 邮箱
		TitlePhoto string    `db:"title_photo"` // 头像
		Name       string    `db:"name"`        // english name
		Cname      string    `db:"cname"`       // chinese name
		CfId       string    `db:"cf_id"`       // codeforces用户名
		CfRating   int64     `db:"cf_rating"`   // codeforces rating
		CfRank     string    `db:"cf_rank"`     // codeforces rank
		AtcId      string    `db:"atc_id"`      // atcoder用户名
		NowcoderId string    `db:"nowcoder_id"` // 牛客网id
	}
)

func newUserInfoModel(conn sqlx.SqlConn) *defaultUserInfoModel {
	return &defaultUserInfoModel{
		conn:  conn,
		table: "`user_info`",
	}
}

func (m *defaultUserInfoModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserInfoModel) FindOne(ctx context.Context, id uint64) (*UserInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userInfoRows, m.table)
	var resp UserInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserInfoModel) FindOneByUsername(ctx context.Context, username string) (*UserInfo, error) {
	var resp UserInfo
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userInfoRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserInfoModel) Insert(ctx context.Context, data *UserInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.IsDelete, data.Username, data.Email, data.TitlePhoto, data.Name, data.Cname, data.CfId, data.CfRating, data.CfRank, data.AtcId, data.NowcoderId)
	return ret, err
}

func (m *defaultUserInfoModel) Update(ctx context.Context, newData *UserInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.IsDelete, newData.Username, newData.Email, newData.TitlePhoto, newData.Name, newData.Cname, newData.CfId, newData.CfRating, newData.CfRank, newData.AtcId, newData.NowcoderId, newData.Id)
	return err
}

func (m *defaultUserInfoModel) tableName() string {
	return m.table
}