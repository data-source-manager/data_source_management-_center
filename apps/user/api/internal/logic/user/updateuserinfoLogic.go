package user

import (
	"context"
	"data_source_management_center/apps/user/api/internal/svc"
	"data_source_management_center/apps/user/api/internal/types"
	"data_source_management_center/apps/user/rpc/pb"
	"data_source_management_center/common/ctxdata"
	"fmt"
	"github.com/jinzhu/copier"
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
	fmt.Println("api获取的用户id:", ctxdata.GetUidFromCtx(l.ctx))
	pbUser := new(pb.User)
	_ = copier.Copy(pbUser, req)
	updateRes, err := l.svcCtx.UserRpc.UpdateUserInfo(l.ctx, &pb.UpdateUserInfoReq{
		User: pbUser,
	})
	_ = copier.Copy(resp, updateRes)
	return
}
