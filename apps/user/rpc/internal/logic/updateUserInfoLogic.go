package logic

import (
	"context"
	"data_source_management_center/apps/user/model"
	"data_source_management_center/apps/user/rpc/internal/svc"
	"data_source_management_center/apps/user/rpc/pb"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"strings"

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
	_, err := l.svcCtx.UserModel.FindOne(l.ctx, in.User.Id)
	if err != nil {
		return nil, errors.New("network busy")
	}
	if err == sqlx.ErrNotFound {
		return nil, errors.New("用户不存在")
	}

	if err := l.svcCtx.UserModel.Trans(context.Background(), func(ctx context.Context, session sqlx.Session) error {
		delSql := fmt.Sprintf("delete from user where id=?")
		prepare, err := session.Prepare(delSql)
		if err != nil {
			return err
		}
		defer prepare.Close()
		if _, err := prepare.Exec(in.User.Id); err != nil {
			return err
		}
		u := new(model.User)
		_ = copier.Copy(u, in.User)
		if strings.TrimSpace(in.User.Username) != "" {
			u.Username = in.User.Username
		}
		if strings.TrimSpace(in.User.Email) != "" {
			u.Email = in.User.Email
		}
		if strings.TrimSpace(in.User.Info) != "" {
			u.Info = in.User.Info
		}
		if strings.TrimSpace(in.User.Sex) != "" {
			u.Sex = in.User.Sex
		}

		insertSql := "insert into user (password,username,sex,email,info) values (?,?, ?, ?, ?)"
		fmt.Println(insertSql)
		preInsert, err := session.Prepare(insertSql)
		if err != nil {
			return err
		}
		if _, err := preInsert.Exec(u.Password, u.Username, u.Sex, u.Email, u.Info); err != nil {
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
