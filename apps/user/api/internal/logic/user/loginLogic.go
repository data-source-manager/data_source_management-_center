package user

import (
	"context"
	"data_source_management_center/apps/user/api/internal/svc"
	"data_source_management_center/apps/user/api/internal/types"
	"data_source_management_center/apps/user/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginRes, err := l.svcCtx.UserRpc.Login(l.ctx, &usercenter.LoginReq{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		AccessToken:  loginRes.AccessToken,
		AccessExpire: loginRes.AccessExpire,
		RefreshAfter: loginRes.RefreshAfter,
	}, nil
}
