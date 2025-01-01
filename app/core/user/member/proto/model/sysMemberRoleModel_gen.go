// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysMemberRoleFieldNames          = builder.RawFieldNames(&SysMemberRole{})
	sysMemberRoleRows                = strings.Join(sysMemberRoleFieldNames, ",")
	sysMemberRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(sysMemberRoleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysMemberRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(sysMemberRoleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	sysMemberRoleModel interface {
		Insert(ctx context.Context, data *SysMemberRole) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysMemberRole, error)
		FindOneByMemberIdRoleId(ctx context.Context, memberId int64, roleId int64) (*SysMemberRole, error)
		Update(ctx context.Context, data *SysMemberRole) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysMemberRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysMemberRole struct {
		Id       int64 `db:"id"`        // 角色用户关联唯一标识符，主键
		MemberId int64 `db:"member_id"` // 用户唯一标识符
		RoleId   int64 `db:"role_id"`   // 角色唯一标识符
	}
)

func newSysMemberRoleModel(conn sqlx.SqlConn) *defaultSysMemberRoleModel {
	return &defaultSysMemberRoleModel{
		conn:  conn,
		table: "`sys_member_role`",
	}
}

func (m *defaultSysMemberRoleModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysMemberRoleModel) FindOne(ctx context.Context, id int64) (*SysMemberRole, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysMemberRoleRows, m.table)
	var resp SysMemberRole
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysMemberRoleModel) FindOneByMemberIdRoleId(ctx context.Context, memberId int64, roleId int64) (*SysMemberRole, error) {
	var resp SysMemberRole
	query := fmt.Sprintf("select %s from %s where `member_id` = ? and `role_id` = ? limit 1", sysMemberRoleRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, memberId, roleId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysMemberRoleModel) Insert(ctx context.Context, data *SysMemberRole) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysMemberRoleRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.MemberId, data.RoleId)
	return ret, err
}

func (m *defaultSysMemberRoleModel) Update(ctx context.Context, newData *SysMemberRole) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysMemberRoleRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.MemberId, newData.RoleId, newData.Id)
	return err
}

func (m *defaultSysMemberRoleModel) tableName() string {
	return m.table
}
