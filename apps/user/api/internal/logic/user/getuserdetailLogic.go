package user

import (
	"context"
	"data_source_management_center/apps/user/api/internal/svc"
	"data_source_management_center/apps/user/api/internal/types"
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

	return
}
