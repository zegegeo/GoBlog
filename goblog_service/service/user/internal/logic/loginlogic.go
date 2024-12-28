package logic

import (
	"context"
	"errors"

	"GoBlog/service/user/internal/svc"
	"GoBlog/service/user/model"
	"GoBlog/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
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

// 用户登录
func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	l.Logger.Infof("开始登录验证，用户名: %s", in.Username)

	// 查找用户
	userModel, err := l.svcCtx.Users.FindByUsername(in.Username)
	if err != nil {
		l.Logger.Errorf("查找用户失败: %v", err)
		return nil, err
	}
	if userModel == nil {
		l.Logger.Errorf("用户不存在: %s", in.Username)
		return nil, errors.New("用户不存在")
	}

	l.Logger.Infof("找到用户，开始验证密码")
	l.Logger.Infof("存储的密码哈希: %s", userModel.Password)
	l.Logger.Infof("输入的密码: %s", in.Password)

	// 基本验证
	if len(in.Password) == 0 {
		l.Logger.Errorf("输入的密码为空")
		return nil, errors.New("密码不能为空")
	}

	if len(userModel.Password) == 0 {
		l.Logger.Errorf("数据库中的密码哈希为空")
		return nil, errors.New("系统错误")
	}

	// 验证密码
	if !model.ComparePassword(userModel.Password, in.Password) {
		l.Logger.Errorf("密码验证失败 - 存储的哈希: %s, 输入的密码: %s",
			userModel.Password, in.Password)
		return nil, errors.New("密码错误")
	}

	l.Logger.Infof("密码验证成功")
	return &user.LoginResponse{
		User: &user.UserInfo{
			UserId:    userModel.Id,
			Username:  userModel.Username,
			CreatedAt: userModel.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
