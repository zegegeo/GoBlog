package user_logic

import (
	"context"
	"time"

	"GoBlog/api/internal/svc"
	"GoBlog/api/internal/types"
	"GoBlog/service/user/user"

	"github.com/golang-jwt/jwt"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginReply, error) {
	// 1.调用user的rpc服务验证用户身份
	resp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	// 生成JWT token
	auth := l.svcCtx.Config.Auth
	token, err := l.getJwtToken(auth.AccessSecret, auth.AccessExpire, resp.User.UserId)
	if err != nil {
		l.Logger.Errorf("Failed to generate token: %v", err)
		return nil, err
	}
	l.Logger.Infof("Generated token: %s", token)

	return &types.LoginReply{
		Token: token,
		User: types.User{
			Id:        resp.User.UserId,
			Username:  resp.User.Username,
			CreatedAt: resp.User.CreatedAt,
		},
	}, nil
}

// 生成JWT token
func (l *LoginLogic) getJwtToken(secretKey string, expire int64, userId int64) (string, error) {
	iat := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + expire
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	l.Logger.Infof("Generating token with claims: %+v", claims)
	return token.SignedString([]byte(secretKey))
}
