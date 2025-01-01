package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysMemberRoleModel = (*customSysMemberRoleModel)(nil)

type (
	// SysMemberRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMemberRoleModel.
	SysMemberRoleModel interface {
		sysMemberRoleModel
		withSession(session sqlx.Session) SysMemberRoleModel
	}

	customSysMemberRoleModel struct {
		*defaultSysMemberRoleModel
	}
)

// NewSysMemberRoleModel returns a model for the database table.
func NewSysMemberRoleModel(conn sqlx.SqlConn) SysMemberRoleModel {
	return &customSysMemberRoleModel{
		defaultSysMemberRoleModel: newSysMemberRoleModel(conn),
	}
}

func (m *customSysMemberRoleModel) withSession(session sqlx.Session) SysMemberRoleModel {
	return NewSysMemberRoleModel(sqlx.NewSqlConnFromSession(session))
}
