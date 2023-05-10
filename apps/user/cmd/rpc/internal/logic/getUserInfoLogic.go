package logic

import (
	"context"
	"data_source_management_center/apps/user/cmd/rpc/internal/svc"
	"data_source_management_center/apps/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	queryRes, err := l.svcCtx.UserModel.FindOne(context.Background(), in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserInfoResp{User: &pb.User{
		Username: queryRes.Username,
		Email:    queryRes.Email,
		Info:     queryRes.Info,
		Sex:      queryRes.Sex,
	}}, nil
}
