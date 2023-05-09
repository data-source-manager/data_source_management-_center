package user

import (
	"context"
	"data_source_management_center/apps/user/api/internal/svc"
	"data_source_management_center/apps/user/api/internal/types"
	"data_source_management_center/apps/user/rpc/pb"
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
	fmt.Println(fmt.Sprintf("%+v", req))
	pbUser := new(pb.User)
	_ = copier.Copy(pbUser, req)
	updateRes, err := l.svcCtx.UserRpc.UpdateUserInfo(context.Background(), &pb.UpdateUserInfoReq{
		User: pbUser,
	})
	fmt.Println(updateRes)
	_ = copier.Copy(resp, updateRes)
	return
}
