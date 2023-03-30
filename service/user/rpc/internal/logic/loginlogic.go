package logic

import (
	"context"
	"github.com/ch3nnn/blog-admin-go/service/user/model"
	"google.golang.org/grpc/status"

	"github.com/ch3nnn/blog-admin-go/service/user/rpc/internal/svc"
	"github.com/ch3nnn/blog-admin-go/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"
)

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

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	// 查询用户是否存在
	userUser, err := l.svcCtx.UserUser.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	// TODO 判断密码是否正确
	loginResponse := user.LoginResponse{}
	if err := copier.Copy(&loginResponse, &userUser); err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &loginResponse, nil
}
