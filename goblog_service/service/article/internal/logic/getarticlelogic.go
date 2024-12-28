package logic

import (
	"context"
	"errors"

	"GoBlog/service/article/article"
	"GoBlog/service/article/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleLogic) GetArticle(in *article.GetArticleRequest) (*article.GetArticleResponse, error) {
	l.Logger.Infof("获取文章, ID: %d", in.Id)

	articleModel, err := l.svcCtx.Articles.FindById(in.Id)
	if err != nil {
		l.Logger.Errorf("查询数据库失败: %v", err)
		return nil, err
	}
	if articleModel == nil {
		l.Logger.Errorf("文章不存在, ID: %d", in.Id)
		return nil, errors.New("文章不存在")
	}

	l.Logger.Info("查询成功，返回文章信息")
	return &article.GetArticleResponse{
		Article: &article.ArticleInfo{
			Id:        articleModel.Id,
			UserId:    articleModel.UserId,
			Title:     articleModel.Title,
			Content:   articleModel.Content,
			CreatedAt: articleModel.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: articleModel.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
