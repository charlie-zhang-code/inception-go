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
	sysMemberGroupFieldNames          = builder.RawFieldNames(&SysMemberGroup{})
	sysMemberGroupRows                = strings.Join(sysMemberGroupFieldNames, ",")
	sysMemberGroupRowsExpectAutoSet   = strings.Join(stringx.Remove(sysMemberGroupFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysMemberGroupRowsWithPlaceHolder = strings.Join(stringx.Remove(sysMemberGroupFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	sysMemberGroupModel interface {
		Insert(ctx context.Context, data *SysMemberGroup) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysMemberGroup, error)
		FindOneByMemberIdGroupId(ctx context.Context, memberId int64, groupId int64) (*SysMemberGroup, error)
		Update(ctx context.Context, data *SysMemberGroup) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysMemberGroupModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysMemberGroup struct {
		Id       int64 `db:"id"`        // 用户用户组关联唯一标识符，主键
		MemberId int64 `db:"member_id"` // 用户唯一标识符
		GroupId  int64 `db:"group_id"`  // 用户组唯一标识符
	}
)

func newSysMemberGroupModel(conn sqlx.SqlConn) *defaultSysMemberGroupModel {
	return &defaultSysMemberGroupModel{
		conn:  conn,
		table: "`sys_member_group`",
	}
}

func (m *defaultSysMemberGroupModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysMemberGroupModel) FindOne(ctx context.Context, id int64) (*SysMemberGroup, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysMemberGroupRows, m.table)
	var resp SysMemberGroup
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

func (m *defaultSysMemberGroupModel) FindOneByMemberIdGroupId(ctx context.Context, memberId int64, groupId int64) (*SysMemberGroup, error) {
	var resp SysMemberGroup
	query := fmt.Sprintf("select %s from %s where `member_id` = ? and `group_id` = ? limit 1", sysMemberGroupRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, memberId, groupId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysMemberGroupModel) Insert(ctx context.Context, data *SysMemberGroup) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sysMemberGroupRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.MemberId, data.GroupId)
	return ret, err
}

func (m *defaultSysMemberGroupModel) Update(ctx context.Context, newData *SysMemberGroup) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysMemberGroupRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.MemberId, newData.GroupId, newData.Id)
	return err
}

func (m *defaultSysMemberGroupModel) tableName() string {
	return m.table
}