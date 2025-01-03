// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: article.proto

package articleclient

import (
	"context"

	"GoBlog/service/article/article"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ArticleInfo            = article.ArticleInfo
	CreateArticleRequest   = article.CreateArticleRequest
	CreateArticleResponse  = article.CreateArticleResponse
	DeleteArticleRequest   = article.DeleteArticleRequest
	DeleteArticleResponse  = article.DeleteArticleResponse
	GetArticleListRequest  = article.GetArticleListRequest
	GetArticleListResponse = article.GetArticleListResponse
	GetArticleRequest      = article.GetArticleRequest
	GetArticleResponse     = article.GetArticleResponse
	UpdateArticleRequest   = article.UpdateArticleRequest
	UpdateArticleResponse  = article.UpdateArticleResponse

	Article interface {
		// 创建文章
		CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error)
		// 获取文章列表
		GetArticleList(ctx context.Context, in *GetArticleListRequest, opts ...grpc.CallOption) (*GetArticleListResponse, error)
		// 获取文章详情
		GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleResponse, error)
		// 更新文章
		UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleResponse, error)
		// 删除文章
		DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error)
	}

	defaultArticle struct {
		cli zrpc.Client
	}
)

func NewArticle(cli zrpc.Client) Article {
	return &defaultArticle{
		cli: cli,
	}
}

// 创建文章
func (m *defaultArticle) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error) {
	client := article.NewArticleClient(m.cli.Conn())
	return client.CreateArticle(ctx, in, opts...)
}

// 获取文章列表
func (m *defaultArticle) GetArticleList(ctx context.Context, in *GetArticleListRequest, opts ...grpc.CallOption) (*GetArticleListResponse, error) {
	client := article.NewArticleClient(m.cli.Conn())
	return client.GetArticleList(ctx, in, opts...)
}

// 获取文章详情
func (m *defaultArticle) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleResponse, error) {
	client := article.NewArticleClient(m.cli.Conn())
	return client.GetArticle(ctx, in, opts...)
}

// 更新文章
func (m *defaultArticle) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleResponse, error) {
	client := article.NewArticleClient(m.cli.Conn())
	return client.UpdateArticle(ctx, in, opts...)
}

// 删除文章
func (m *defaultArticle) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error) {
	client := article.NewArticleClient(m.cli.Conn())
	return client.DeleteArticle(ctx, in, opts...)
}
