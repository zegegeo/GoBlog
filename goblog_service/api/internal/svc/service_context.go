package svc

import (
	"GoBlog/api/internal/config"
	"GoBlog/api/internal/middleware"
	"GoBlog/service/article/article"
	"GoBlog/service/comment/comment"
	"GoBlog/service/user/user"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	//中间件
	AuthMiddleware  rest.Middleware //JWT认证中间件
	CorsMiddleware  rest.Middleware
	DebugMiddleware rest.Middleware
	//RPC客户端
	UserRpc        user.UserClient //用户rpc服务
	ArticleRpc     article.ArticleClient
	CommentRpc     comment.CommentClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRpc := zrpc.MustNewClient(c.UserRpc)
	articleRpc := zrpc.MustNewClient(c.ArticleRpc)
	commentRpc := zrpc.MustNewClient(c.CommentRpc)
	return &ServiceContext{
		Config:         c,
		//初始化中间件
		AuthMiddleware: middleware.NewAuthMiddleware(c.Auth.AccessSecret).Handle,
		CorsMiddleware: middleware.NewCorsMiddleware().Handle,
		DebugMiddleware: middleware.NewDebugMiddleware().Handle,
		//初始化RPC客户端
		UserRpc:        user.NewUserClient(userRpc.Conn()),
		ArticleRpc:     article.NewArticleClient(articleRpc.Conn()),
		CommentRpc:     comment.NewCommentClient(commentRpc.Conn()),
	}
}
