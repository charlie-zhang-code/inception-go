package logic

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"member/proto/model"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberStatusLogic {
	return &UpdateMemberStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户状态
func (l *UpdateMemberStatusLogic) UpdateMemberStatus(in *pb.UpdateMemberStatusReq) (*pb.UpdateMemberStatusResp, error) {
	updates := make(map[string]interface{})
	var selectColumns []string

	// 检查并添加状态更新
	updates["status"] = int32(in.Status)
	selectColumns = append(selectColumns, "status")

	// 使用事务执行更新操作
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.SysMember{}).
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

	return &pb.UpdateMemberStatusResp{}, nil
}
