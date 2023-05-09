package logic

import (
	"context"
	"data_source_management_center/apps/user/rpc/internal/svc"
	"data_source_management_center/apps/user/rpc/pb"
	"data_source_management_center/common/tools"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

var UserPasswordError = errors.New("密码错误")
var UserNameOrPwdEmptyError = errors.New("用户名或者密码为空")
var EmptyUserError = errors.New("用户名不存在")

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	if len(strings.TrimSpace(in.Username)) == 0 || len(strings.TrimSpace(in.Password)) == 0 {
		return nil, UserNameOrPwdEmptyError
	}

	user, err := l.svcCtx.UserModel.FindOneByUserName(context.Background(), in.Username)
	if err == sqlc.ErrNotFound {
		return nil, EmptyUserError
	}
	if err != nil {
		return nil, err
	}

	if user.Password != tools.Md5ByString(in.Password) {
		return nil, UserPasswordError
	}
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)

	tokenResp, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{UserId: user.Id})
	if err != nil {
		return nil, err
	}

	return &pb.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
