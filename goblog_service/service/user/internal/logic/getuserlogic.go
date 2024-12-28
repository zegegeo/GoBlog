package logic

import (
	"context"
	"errors"

	"GoBlog/service/user/internal/svc"
	"GoBlog/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.GetUserResponse, error) {
	// 查找用户
	userModel, err := l.svcCtx.Users.FindById(in.UserId)
	if err != nil {
		return nil, err
	}
	if userModel == nil {
		return nil, errors.New("用户不存在")
	}

	return &user.GetUserResponse{
		User: &user.UserInfo{
			UserId:    userModel.Id,
			Username:  userModel.Username,
			CreatedAt: userModel.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
