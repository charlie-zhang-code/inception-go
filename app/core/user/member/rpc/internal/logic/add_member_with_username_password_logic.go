package logic

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"member/proto/model"
	"strings"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMemberWithUsernamePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberWithUsernamePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberWithUsernamePasswordLogic {
	return &AddMemberWithUsernamePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加用户
func (l *AddMemberWithUsernamePasswordLogic) AddMemberWithUsernamePassword(in *pb.AddMemberWithUsernamePasswordReq) (*pb.AddMemberWithUsernamePasswordResp, error) {
	// 随机标识
	identify := strings.ReplaceAll(uuid.New().String(), "-", "")

	// 默认密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return &pb.AddMemberWithUsernamePasswordResp{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	entity := &model.SysMember{
		Identify: identify,
		Password: string(hashedPassword),
		Username: in.Username,
	}

	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(entity).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return &pb.AddMemberWithUsernamePasswordResp{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &pb.AddMemberWithUsernamePasswordResp{
		Code:    200,
		Message: "success",
	}, nil
}
