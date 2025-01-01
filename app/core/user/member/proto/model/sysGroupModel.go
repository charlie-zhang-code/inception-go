package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysGroupModel = (*customSysGroupModel)(nil)

type (
	// SysGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysGroupModel.
	SysGroupModel interface {
		sysGroupModel
		withSession(session sqlx.Session) SysGroupModel
	}

	customSysGroupModel struct {
		*defaultSysGroupModel
	}
)

// NewSysGroupModel returns a model for the database table.
func NewSysGroupModel(conn sqlx.SqlConn) SysGroupModel {
	return &customSysGroupModel{
		defaultSysGroupModel: newSysGroupModel(conn),
	}
}

func (m *customSysGroupModel) withSession(session sqlx.Session) SysGroupModel {
	return NewSysGroupModel(sqlx.NewSqlConnFromSession(session))
}
