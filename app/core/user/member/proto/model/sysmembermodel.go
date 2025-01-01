package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysMemberModel = (*customSysMemberModel)(nil)

type (
	// SysMemberModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMemberModel.
	SysMemberModel interface {
		sysMemberModel
		withSession(session sqlx.Session) SysMemberModel
	}

	customSysMemberModel struct {
		*defaultSysMemberModel
	}
)

// NewSysMemberModel returns a model for the database table.
func NewSysMemberModel(conn sqlx.SqlConn) SysMemberModel {
	return &customSysMemberModel{
		defaultSysMemberModel: newSysMemberModel(conn),
	}
}

func (m *customSysMemberModel) withSession(session sqlx.Session) SysMemberModel {
	return NewSysMemberModel(sqlx.NewSqlConnFromSession(session))
}
