package logic

import (
	"context"
	"data_source_management_center/apps/user/model"
	"data_source_management_center/apps/user/rpc/internal/svc"
	"data_source_management_center/apps/user/rpc/pb"
	"data_source_management_center/common/tools"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

var UserExistError = errors.New("用户已存在")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	if len(strings.TrimSpace(in.Username)) == 0 || len(strings.TrimSpace(in.Password)) == 0 {
		return nil, UserNameOrPwdEmptyError
	}
	userFindRes, err := l.svcCtx.UserModel.FindOneByUserName(context.Background(), in.Username)
	if err != nil && err != sqlc.ErrNotFound {
		return nil, err
	}

	if userFindRes != nil {
		return nil, UserExistError
	}

	user := new(model.User)

	user.Username = in.Username
	user.Password = tools.Md5ByString(in.Password)
	if strings.TrimSpace(in.Sex) != "" {
		user.Sex = sql.NullString{
			String: in.Sex,
			Valid:  true,
		}
	}
	if strings.TrimSpace(in.Email) != "" {
		user.Email = sql.NullString{
			String: in.Email,
			Valid:  true,
		}
	}
	if strings.TrimSpace(in.Info) != "" {
		user.Info = sql.NullString{
			String: in.Info,
			Valid:  true,
		}
	}

	insertRes, err := l.svcCtx.UserModel.Insert(context.Background(), user)
	if err != nil {
		return nil, err
	}

	userId, err := insertRes.LastInsertId()
	createTokenService := NewGenerateTokenLogic(context.Background(), l.svcCtx)
	genTokenResp, err := createTokenService.GenerateToken(&pb.GenerateTokenReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResp{
		AccessExpire: genTokenResp.AccessExpire,
		AccessToken:  genTokenResp.AccessToken,
		RefreshAfter: genTokenResp.RefreshAfter,
	}, nil
}
