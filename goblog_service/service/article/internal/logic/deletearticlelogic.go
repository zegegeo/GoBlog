package logic

import (
	"context"
	"errors"
	"fmt"

	"GoBlog/service/article/article"
	"GoBlog/service/article/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteArticleLogic) DeleteArticle(in *article.DeleteArticleRequest) (*article.DeleteArticleResponse, error) {
	l.Logger.Infof("删除文章请求: %+v", in)

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
		return nil, errors.New("无权删除此文章")
	}

	// 删除文章
	err = l.svcCtx.Articles.Delete(in.Id)
	if err != nil {
		l.Logger.Errorf("删除数据库记录失败: %v", err)
		return nil, fmt.Errorf("删除文章失败: %v", err)
	}

	l.Logger.Info("文章删除成功")
	return &article.DeleteArticleResponse{}, nil
}
