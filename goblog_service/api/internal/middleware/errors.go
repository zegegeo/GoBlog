package middleware

import "errors"

var (
	ErrUnauthorized = errors.New("未授权的访问")
	ErrInvalidToken = errors.New("无效的令牌")
)
