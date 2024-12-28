package logic

import (
	"context"
	"errors"

	"GoBlog/service/user/internal/svc"
	"GoBlog/service/user/model"
	"GoBlog/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

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

// 用户注册
func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 检查用户名是否已存在
	existUser, err := l.svcCtx.Users.FindByUsername(in.Username)
	if err != nil {
		return nil, err
	}
	if existUser != nil {
		return nil, errors.New("用户名已存在")
	}

	hashedPassword, err := model.EncryptPassword(in.Password)
	if err != nil {
		return nil, err
	}

	// 插入新用户
	userId, err := l.svcCtx.Users.Insert(in.Username, hashedPassword)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{
		UserId: userId,
	}, nil
}
