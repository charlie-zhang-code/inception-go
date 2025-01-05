package logic

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"member/proto/model"
	"time"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLogic {
	return &UpdateMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户
func (l *UpdateMemberLogic) UpdateMember(in *pb.UpdateMemberReq) (*pb.UpdateMemberResp, error) {
	updates := make(map[string]interface{})
	var selectColumns []string

	// 检查并添加用户名更新
	if in.Username != "" {
		updates["username"] = in.Username
		selectColumns = append(selectColumns, "username")
	}

	// 检查并添加昵称更新
	if in.Nickname != "" {
		updates["nickname"] = in.Nickname
		selectColumns = append(selectColumns, "nickname")
	}

	// 检查并添加头像更新
	if in.Avatar != "" {
		updates["avatar"] = in.Avatar
		selectColumns = append(selectColumns, "avatar")
	}

	// 检查并添加说说更新
	if in.Quote != "" {
		updates["quote"] = in.Quote
		selectColumns = append(selectColumns, "quote")
	}

	// 检查并添加生日更新
	if in.Birthday != "" {
		birthday, err := time.Parse(time.DateOnly, in.Birthday)
		if err != nil {
			return nil, fmt.Errorf("invalid birthday format: %v", err)
		}
		updates["birthday"] = birthday
		selectColumns = append(selectColumns, "birthday")
	}

	// 检查并添加性别更新
	updates["gender"] = int32(in.Gender)
	selectColumns = append(selectColumns, "gender")

	// 检查并添加邮箱更新
	if in.Email != "" {
		updates["email"] = in.Email
		selectColumns = append(selectColumns, "email")
	}

	// 检查并添加手机号更新
	if in.Telephone != "" {
		updates["telephone"] = in.Telephone
		selectColumns = append(selectColumns, "telephone")
	}

	// 检查并添加状态更新
	updates["status"] = int32(in.Status)
	selectColumns = append(selectColumns, "status")

	// 检查并添加删除更新
	updates["deleted"] = int32(in.Deleted)
	selectColumns = append(selectColumns, "deleted")

	// 检查并添加备注更新
	if in.Remark != "" {
		updates["remark"] = in.Remark
		selectColumns = append(selectColumns, "remark")
	}

	if len(updates) == 0 {
		return &pb.UpdateMemberResp{}, errors.New("no fields to update")
	}

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

	return &pb.UpdateMemberResp{}, nil
}
