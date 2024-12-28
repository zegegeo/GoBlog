package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type contextKey string

const UserKey contextKey = "userId"

type AuthMiddleware struct {
	accessSecret string
}

func NewAuthMiddleware(accessSecret string) *AuthMiddleware {
	return &AuthMiddleware{
		accessSecret: accessSecret,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("开始验证token")
		auth := r.Header.Get("Authorization")
		logx.Infof("Authorization header: %s", auth)
		if auth == "" {
			httpx.Error(w, ErrUnauthorized)
			return
		}

		authParts := strings.SplitN(auth, " ", 2)
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			httpx.Error(w, ErrInvalidToken)
			return
		}

		logx.Infof("解析token: %s", authParts[1])
		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logx.Error("无效的签名方法")
				return nil, ErrInvalidToken
			}
			return []byte(m.accessSecret), nil
		})

		if err != nil || !token.Valid {
			logx.Errorf("token验证失败: %v", err)
			httpx.Error(w, ErrInvalidToken)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			logx.Error("无法解析token claims")
			httpx.Error(w, ErrInvalidToken)
			return
		}

		logx.Infof("token claims: %+v", claims)
		userId, ok := claims["userId"].(float64)
		if !ok {
			logx.Error("无法获取userId")
			httpx.Error(w, ErrInvalidToken)
			return
		}

		logx.Infof("成功获取userId: %v", userId)
		r = r.WithContext(context.WithValue(r.Context(), UserKey, int64(userId)))
		next(w, r)
	}
}
