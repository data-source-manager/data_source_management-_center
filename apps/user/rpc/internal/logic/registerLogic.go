package logic

import (
	"context"
	"data_source_management_center/apps/user/model"
	"data_source_management_center/apps/user/rpc/internal/svc"
	"data_source_management_center/apps/user/rpc/pb"
	"data_source_management_center/common/tools"
	"database/sql"
	"errors"
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
	if err != nil {
		return nil, err
	}

	if userFindRes != nil {
		return nil, UserExistError
	}

	user := new(model.User)

	user.Username = in.Username
	user.Password = tools.Md5ByString(in.Password)
	user.Sex = sql.NullInt64{
		Int64: in.Sex,
		Valid: true,
	}

	insertRes, err := l.svcCtx.UserModel.Insert(context.Background(), user)
	if err != nil {
		return nil, err
	}

	if v, err := insertRes.RowsAffected(); v != 1 {
		return nil, err
	}

	return &pb.RegisterResp{}, nil
}
