package logic

import (
	"context"

	"GoBlog/service/article/article"
	"GoBlog/service/article/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListLogic {
	return &GetArticleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleListLogic) GetArticleList(in *article.GetArticleListRequest) (*article.GetArticleListResponse, error) {
	l.Logger.Infof("获取文章列表请求: userId=%d, page=%d, pageSize=%d", in.UserId, in.Page, in.PageSize)

	// 计算偏移量
	offset := (in.Page - 1) * in.PageSize

	// 获取文章列表
	articles, err := l.svcCtx.Articles.Find(int(offset), int(in.PageSize))
	if err != nil {
		l.Logger.Errorf("查询文章失败: %v", err)
		return nil, err
	}

	// 获取总数
	total, err := l.svcCtx.Articles.Count()
	if err != nil {
		l.Logger.Errorf("获取文章总数失败: %v", err)
		return nil, err
	}

	l.Logger.Infof("找到 %d 篇文章，总数: %d", len(articles), total)

	// 转换为响应格式
	list := make([]*article.ArticleInfo, 0)
	for _, a := range articles {
		list = append(list, &article.ArticleInfo{
			Id:        a.Id,
			UserId:    a.UserId,
			Title:     a.Title,
			Content:   a.Content,
			CreatedAt: a.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: a.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &article.GetArticleListResponse{
		List:  list,
		Total: total,
	}, nil
}
