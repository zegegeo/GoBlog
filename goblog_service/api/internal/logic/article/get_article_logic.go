package article_logic

import (
	"context"
	"errors"
	"fmt"

	"GoBlog/api/internal/svc"
	"GoBlog/api/internal/types"
	"GoBlog/service/article/article"
	"GoBlog/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleLogic) GetArticle(req *types.ArticleReq) (*types.ArticleResp, error) {
	l.Logger.Infof("获取文章详情, ID: %d", req.Id)

	// 检查 RPC 客户端
	if l.svcCtx.ArticleRpc == nil {
		l.Logger.Error("ArticleRpc client is nil")
		return nil, errors.New("系统错误")
	}

	// 参数验证
	if req == nil {
		return nil, errors.New("请求参数不能为空")
	}
	if req.Id <= 0 {
		return nil, errors.New("无效的文章ID")
	}

	// 调用article rpc服务
	l.Logger.Info("开始调用 Article RPC 服务")
	resp, err := l.svcCtx.ArticleRpc.GetArticle(l.ctx, &article.GetArticleRequest{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("RPC调用失败: %v", err)
		return nil, fmt.Errorf("获取文章失败: %v", err)
	}

	// 检查文章是否存在
	if resp == nil || resp.Article == nil {
		return nil, errors.New("文章不存在")
	}

	// 获取作者信息
	userResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserRequest{
		UserId: resp.Article.UserId,
	})
	if err != nil {
		l.Logger.Errorf("获取作者信息失败: %v", err)
		// 不要因为获取作者信息失败而影响整个文章获取
		return &types.ArticleResp{
			Article: types.Article{
				Id:        resp.Article.Id,
				UserId:    resp.Article.UserId,
				Title:     resp.Article.Title,
				Content:   resp.Article.Content,
				CreatedAt: resp.Article.CreatedAt,
				UpdatedAt: resp.Article.UpdatedAt,
				Author: types.User{
					Username: "未知作者",
				},
			},
		}, nil
	}

	l.Logger.Info("RPC调用成功，开始转换响应")
	return &types.ArticleResp{
		Article: types.Article{
			Id:        resp.Article.Id,
			UserId:    resp.Article.UserId,
			Title:     resp.Article.Title,
			Content:   resp.Article.Content,
			ViewCount: resp.Article.ViewCount,
			CreatedAt: resp.Article.CreatedAt,
			UpdatedAt: resp.Article.UpdatedAt,
			Author: types.User{
				Id:        userResp.User.UserId,
				Username:  userResp.User.Username,
				CreatedAt: userResp.User.CreatedAt,
			},
		},
	}, nil
}
