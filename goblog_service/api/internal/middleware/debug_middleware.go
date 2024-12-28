package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type DebugMiddleware struct{}

func NewDebugMiddleware() *DebugMiddleware {
	return &DebugMiddleware{}
}

func (m *DebugMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Infof("收到请求: Method=%s, URL=%s, Path=%s, RemoteAddr=%s",
			r.Method, r.URL.String(), r.URL.Path, r.RemoteAddr)
		next(w, r)
	}
}
