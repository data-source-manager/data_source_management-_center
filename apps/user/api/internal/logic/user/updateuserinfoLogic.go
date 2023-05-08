package user

import (
	"context"
	"data_source_management_center/apps/user/api/internal/svc"
	"data_source_management_center/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateuserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateuserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateuserinfoLogic {
	return &UpdateuserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateuserinfoLogic) Updateuserinfo(req *types.UpdateUserInfoReq) (resp *types.UpdateUserInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
