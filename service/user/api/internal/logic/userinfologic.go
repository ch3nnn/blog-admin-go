package logic

import (
	"context"
	"github.com/ch3nnn/blog-admin-go/service/user/rpc/user"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"

	"github.com/ch3nnn/blog-admin-go/service/user/api/internal/svc"
	"github.com/ch3nnn/blog-admin-go/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (*types.UserInfoResponse, error) {
	userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: req.Id})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	var resp types.UserInfoResponse
	if err := copier.Copy(&resp, userInfo); err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &resp, nil
}
