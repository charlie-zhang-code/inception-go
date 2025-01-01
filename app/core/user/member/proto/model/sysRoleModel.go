package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysRoleModel = (*customSysRoleModel)(nil)

type (
	// SysRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleModel.
	SysRoleModel interface {
		sysRoleModel
		withSession(session sqlx.Session) SysRoleModel
	}

	customSysRoleModel struct {
		*defaultSysRoleModel
	}
)

// NewSysRoleModel returns a model for the database table.
func NewSysRoleModel(conn sqlx.SqlConn) SysRoleModel {
	return &customSysRoleModel{
		defaultSysRoleModel: newSysRoleModel(conn),
	}
}

func (m *customSysRoleModel) withSession(session sqlx.Session) SysRoleModel {
	return NewSysRoleModel(sqlx.NewSqlConnFromSession(session))
}