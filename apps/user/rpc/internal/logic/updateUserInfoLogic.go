package logic

import (
	"context"
	"data_source_management_center/apps/user/model"
	"data_source_management_center/apps/user/rpc/internal/svc"
	"data_source_management_center/apps/user/rpc/pb"
	"data_source_management_center/common/ctxdata"
	"errors"

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
	queryUser, err := l.svcCtx.UserModel.FindOne(context.Background(), ctxdata.GetUidFromCtx(l.ctx))
	if err != nil {
		return nil, errors.New("network busy")
	}

	updateUser := new(model.User)
	if in.User.Info != "" {
		updateUser.Info = in.User.Info
	} else {
		updateUser.Info = queryUser.Info
	}

	//err := l.svcCtx.UserModel.Update(context.Background(), updateUser)
	if err != nil {
		return nil, UpdateUserError
	}

	return &pb.UpdateUserInfoResp{
		Res: "update success",
	}, nil
}
