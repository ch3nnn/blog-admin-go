package logic

import (
	"context"
	"github.com/ch3nnn/blog-admin-go/service/user/model"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"

	"github.com/ch3nnn/blog-admin-go/service/user/rpc/internal/svc"
	"github.com/ch3nnn/blog-admin-go/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	userUser, err := l.svcCtx.UserUser.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	var userInfoResponse user.UserInfoResponse
	if err := copier.Copy(&userInfoResponse, &userUser); err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &userInfoResponse, nil
}
