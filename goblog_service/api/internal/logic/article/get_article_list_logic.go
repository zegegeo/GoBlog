package article_logic

import (
	"context"

	"GoBlog/api/internal/svc"
	"GoBlog/api/internal/types"
	"GoBlog/service/article/article"
	"GoBlog/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListLogic {
	return &GetArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleListLogic) GetArticleList(req *types.ArticleListReq) (*types.ArticleListResp, error) {
	l.Logger.Info("开始获取文章列表")

	// 获取文章列表
	articles, err := l.svcCtx.ArticleRpc.GetArticleList(l.ctx, &article.GetArticleListRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		l.Logger.Errorf("获取文章列表失败: %v", err)
		return nil, err
	}

	// 转换响应格式
	list := make([]types.Article, 0)
	for _, a := range articles.List {
		// 获取作者信息
		userResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserRequest{
			UserId: a.UserId,
		})

		var author types.User
		if err != nil {
			l.Logger.Errorf("获取作者(ID:%d)信息失败: %v", a.UserId, err)
			author = types.User{
				Username: "未知作者",
			}
		} else {
			author = types.User{
				Id:        userResp.User.UserId,
				Username:  userResp.User.Username,
				CreatedAt: userResp.User.CreatedAt,
			}
		}

		list = append(list, types.Article{
			Id:        a.Id,
			UserId:    a.UserId,
			Title:     a.Title,
			Content:   a.Content,
			ViewCount: a.ViewCount,
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
			Author:    author,
		})
	}

	l.Logger.Infof("获取到 %d 篇文章", len(list))
	return &types.ArticleListResp{
		Articles: list,
		Total:    articles.Total,
	}, nil
}
