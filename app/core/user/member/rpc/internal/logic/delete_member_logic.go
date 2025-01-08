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

type DeleteMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberLogic {
	return &DeleteMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除用户
func (l *DeleteMemberLogic) DeleteMember(in *pb.DeleteMemberReq) (*pb.DeleteMemberResp, error) {
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if len(in.Ids) == 0 {
			return fmt.Errorf("no IDs provided for soft deletion")
		}
		// 软删除
		if err := tx.Model(&model.SysMember{}).
			Where("id IN (?)", in.Ids).
			UpdateColumn("deleted", 1).Error; err != nil {
			return err
		}
		return nil
	})
	return &pb.DeleteMemberResp{
		Code:    200,
		Message: "success",
	}, err
}
