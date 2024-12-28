package article_logic

import (
	"context"
	"errors"

	"GoBlog/api/internal/svc"
	"GoBlog/api/internal/types"
	"GoBlog/service/article/article"

	"GoBlog/api/internal/middleware"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateArticleLogic) CreateArticle(req *types.CreateArticleReq) (*types.CreateArticleResp, error) {
	l.Logger.Info("开始创建文章")

	// 从上下文获取用户ID
	userId, ok := l.ctx.Value(middleware.UserKey).(int64)
	if !ok {
		l.Logger.Error("获取用户ID失败")
		return nil, errors.New("获取用户信息失败")
	}
	l.Logger.Infof("获取到用户ID: %d", userId)

	// 调用article rpc服务
	l.Logger.Infof("创建文章请求: title=%s, content=%s", req.Title, req.Content)
	resp, err := l.svcCtx.ArticleRpc.CreateArticle(l.ctx, &article.CreateArticleRequest{
		UserId:  userId,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		l.Logger.Errorf("创建文章失败: %v", err)
		return nil, err
	}

	l.Logger.Infof("文章创建成功, ID: %d", resp.Id)
	return &types.CreateArticleResp{
		Id: resp.Id,
	}, nil
}
