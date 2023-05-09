package logic

import (
	"context"
	"data_source_management_center/apps/user/rpc/internal/svc"
	"data_source_management_center/apps/user/rpc/pb"
	"data_source_management_center/common/ctxdata"
	"errors"
	"fmt"

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
	queryUser, err := l.svcCtx.UserModel.FindOne(context.Background(), ctxdata.GetUidFromCtx(l.ctx))
	if err != nil {
		return nil, errors.New("network busy")
	}

	if err := l.svcCtx.UserModel.Trans(context.Background(), func(ctx context.Context, session sqlx.Session) error {
		delSql := fmt.Sprintf("delete from user where id=?")
		prepare, err := session.Prepare(delSql)
		if err != nil {
			return err
		}
		defer prepare.Close()
		if _, err := prepare.ExecCtx(ctx, prepare, queryUser.Id); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.UpdateUserInfoResp{
		Res: "update success",
	}, nil
}
