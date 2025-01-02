package logic

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"member/proto/model"
	"member/rpc/internal/svc"
	"member/rpc/pb"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberLogic {
	return &AddMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加用户
func (l *AddMemberLogic) AddMember(in *pb.AddMemberReq) (*pb.AddMemberResp, error) {
	layout := "2006-01-02"
	birthday, err := time.Parse(layout, in.Birthday)
	if err != nil {
		return nil, err
	}

	// 随机标识
	identify := strings.ReplaceAll(uuid.New().String(), "-", "")

	// 默认密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(l.svcCtx.Config.Default.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	entity := &model.SysMember{
		Identify:  identify,
		Password:  string(hashedPassword),
		Username:  in.Username,
		Nickname:  in.Nickname,
		Avatar:    in.Avatar,
		Quote:     in.Quote,
		Birthday:  birthday,
		Gender:    int32(in.Gender),
		Email:     in.Email,
		Telephone: in.Telephone,
		Status:    int32(in.Status),
		Deleted:   int32(in.Deleted),
		Remark:    in.Remark,
		CreateBy:  in.CreateBy,
	}

	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(entity).Error; err != nil {
			return fmt.Errorf("failed to create member: %v", err)
		}
		return nil
	})

	return &pb.AddMemberResp{}, err
}
