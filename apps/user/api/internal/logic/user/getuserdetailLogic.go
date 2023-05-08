package user

import (
	"context"
	"data_source_management_center/apps/user/api/internal/svc"
	"data_source_management_center/apps/user/api/internal/types"
	"data_source_management_center/apps/user/rpc/usercenter"
	"data_source_management_center/common/ctxdata"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetuserdetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetuserdetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetuserdetailLogic {
	return &GetuserdetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetuserdetailLogic) Getuserdetail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userInfo, err := l.svcCtx.UserRpc.GetUserInfo(context.Background(), &usercenter.GetUserInfoReq{
		Id: ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{UserInfo: types.User{
		UserName: userInfo.User.Username,
		Email:    userInfo.User.Email,
		Info:     userInfo.User.Info,
		Sex:      userInfo.User.Sex,
	}}, nil
}
