package logic

import (
	"context"
	"data_source_management_center/apps/user/cmd/rpc/internal/svc"
	"data_source_management_center/apps/user/cmd/rpc/pb"
	"data_source_management_center/apps/user/model"
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

var UpdateUserError = errors.New("用户更新失败")

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {
	queryUser, err := l.svcCtx.UserModel.FindOne(l.ctx, in.User.Id)
	if err != nil && err != sqlx.ErrNotFound {
		return nil, errors.New(fmt.Sprintf("UpdateUserInfo FindOneById db err,id:%d ,err: %+v", in.User.Id, err))
	}
	if queryUser == nil {
		return nil, errors.New("用户不存在")
	}

	u := new(model.User)
	_ = copier.Copy(u, queryUser)
	if strings.TrimSpace(in.User.Username) != "" {
		u.Username = in.User.Username
	}
	if strings.TrimSpace(in.User.Email) != "" {
		u.Email = in.User.Email
	}
	if strings.TrimSpace(in.User.Info) != "" {
		u.Info = in.User.Info
	}
	if strings.TrimSpace(in.User.Sex) != "" {
		u.Sex = in.User.Sex
	}

	err = l.svcCtx.UserModel.Update(l.ctx, u)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserInfoResp{
		Res: "update success",
	}, nil
}
