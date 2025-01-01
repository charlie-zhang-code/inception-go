package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysMemberGroupModel = (*customSysMemberGroupModel)(nil)

type (
	// SysMemberGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMemberGroupModel.
	SysMemberGroupModel interface {
		sysMemberGroupModel
		withSession(session sqlx.Session) SysMemberGroupModel
	}

	customSysMemberGroupModel struct {
		*defaultSysMemberGroupModel
	}
)

// NewSysMemberGroupModel returns a model for the database table.
func NewSysMemberGroupModel(conn sqlx.SqlConn) SysMemberGroupModel {
	return &customSysMemberGroupModel{
		defaultSysMemberGroupModel: newSysMemberGroupModel(conn),
	}
}

func (m *customSysMemberGroupModel) withSession(session sqlx.Session) SysMemberGroupModel {
	return NewSysMemberGroupModel(sqlx.NewSqlConnFromSession(session))
}
