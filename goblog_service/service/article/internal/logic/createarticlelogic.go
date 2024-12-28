package logic

import (
	"context"
	"errors"
	"fmt"

	"GoBlog/service/article/article"
	"GoBlog/service/article/internal/svc"
	"GoBlog/service/article/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateArticleLogic) CreateArticle(in *article.CreateArticleRequest) (*article.CreateArticleResponse, error) {
	l.Logger.Infof("Creating article in RPC: %+v", in)

	// 参数验证
	if in.Title == "" {
		return nil, errors.New("标题不能为空")
	}
	if in.Content == "" {
		return nil, errors.New("内容不能为空")
	}
	if in.UserId <= 0 {
		return nil, errors.New("用户ID无效")
	}

	// 创建文章
	id, err := l.svcCtx.Articles.Insert(in.UserId, in.Title, in.Content)
	if err != nil {
		l.Logger.Errorf("Failed to insert article to database: %v", err)
		return nil, fmt.Errorf("创建文章失败: %v", err)
	}

	l.Logger.Infof("Article created successfully with ID: %d", id)
	l.svcCtx.Redis.Del(model.GetUserArticlesCacheKey(in.UserId))

	return &article.CreateArticleResponse{
		Id: id,
	}, nil
}
