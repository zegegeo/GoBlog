package logic

import (
	"context"
	"errors"
	"fmt"

	"GoBlog/service/article/article"
	"GoBlog/service/article/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateArticleLogic) UpdateArticle(in *article.UpdateArticleRequest) (*article.UpdateArticleResponse, error) {
	l.Logger.Infof("更新文章请求: %+v", in)

	// 检查文章是否存在
	existingArticle, err := l.svcCtx.Articles.FindById(in.Id)
	if err != nil {
		l.Logger.Errorf("查询文章失败: %v", err)
		return nil, err
	}
	if existingArticle == nil {
		return nil, errors.New("文章不存在")
	}

	// 检查权限
	if existingArticle.UserId != in.UserId {
		return nil, errors.New("无权修改此文章")
	}

	// 更新文章
	err = l.svcCtx.Articles.Update(in.Id, in.Title, in.Content)
	if err != nil {
		l.Logger.Errorf("更新数据库失败: %v", err)
		return nil, fmt.Errorf("更新文章失败: %v", err)
	}

	l.Logger.Info("文章更新成功")
	return &article.UpdateArticleResponse{}, nil
}
