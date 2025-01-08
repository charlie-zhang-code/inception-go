package logic

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"status/proto/model"
	"time"

	"status/rpc/internal/svc"
	"status/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberLogoutStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberLogoutStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLogoutStatusLogic {
	return &UpdateMemberLogoutStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户登录状态
func (l *UpdateMemberLogoutStatusLogic) UpdateMemberLogoutStatus(in *pb.UpdateMemberLogoutStatusReq) (*pb.UpdateMemberLogoutStatusResp, error) {
	updates := make(map[string]interface{})
	var selectColumns []string

	// 检查并添加生日更新
	if in.LogoutAt != "" {
		logoutAt, err := time.Parse(time.RFC3339, in.LogoutAt)
		if err != nil {
			return nil, fmt.Errorf("invalid birthday format: %v", err)
		}
		updates["logout_at"] = logoutAt
		selectColumns = append(selectColumns, "logout_at")
	}

	if len(updates) == 0 {
		return &pb.UpdateMemberLogoutStatusResp{}, errors.New("no fields to update")
	}

	// 使用事务执行更新操作
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.SysMemberLoginStatus{}).
			Select(selectColumns).
			Where("id = ?", in.Id).
			Updates(updates).Error; err != nil {
			return fmt.Errorf("failed to update member: %v", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &pb.UpdateMemberLogoutStatusResp{}, nil
}
