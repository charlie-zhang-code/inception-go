package logic

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"status/proto/model"
	"status/rpc/internal/svc"
	"status/rpc/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMemberLoginStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberLoginStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberLoginStatusLogic {
	return &AddMemberLoginStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加用户登录状态
func (l *AddMemberLoginStatusLogic) AddMemberLoginStatus(in *pb.AddMemberLoginStatusReq) (*pb.AddMemberLoginStatusResp, error) {
	loginAt, err := time.Parse(time.RFC3339, in.LoginAt)
	if err != nil {
		return nil, err
	}

	entity := &model.SysMemberLoginStatus{
		MemberID:    in.MemberId,
		Client:      in.Client,
		LoginIP:     in.LoginIp,
		LoginRegion: in.LoginRegion,
		LoginOs:     in.LoginOs,
		LoginAt:     loginAt,
	}

	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(entity).Error; err != nil {
			return fmt.Errorf("failed to create status: %v", err)
		}
		return nil
	})

	return &pb.AddMemberLoginStatusResp{}, nil
}
