package user

import (
	"context"
	"data_source_management_center/apps/user/cmd/api/internal/svc"
	"data_source_management_center/apps/user/cmd/api/internal/types"
	"data_source_management_center/apps/user/cmd/rpc/pb"
	"data_source_management_center/common/ctxdata"
	"strings"

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
	pbUser := new(pb.User)
	if strings.TrimSpace(req.UserInfo.Info) != "" {
		pbUser.Info = req.UserInfo.Info
	}
	if strings.TrimSpace(req.UserInfo.Sex) != "" {
		pbUser.Sex = req.UserInfo.Sex
	}
	if strings.TrimSpace(req.UserInfo.UserName) != "" {
		pbUser.Username = req.UserInfo.UserName
	}
	if strings.TrimSpace(req.UserInfo.Email) != "" {
		pbUser.Email = req.UserInfo.Email
	}

	pbUser.Id = ctxdata.GetUidFromCtx(l.ctx)
	updateRes, err := l.svcCtx.UserRpc.UpdateUserInfo(l.ctx, &pb.UpdateUserInfoReq{
		User: pbUser,
	})

	return &types.UpdateUserInfoResp{
		UpdateInfo: updateRes.Res,
	}, nil
}
