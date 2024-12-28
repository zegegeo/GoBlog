package article_logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"GoBlog/api/internal/middleware"
	"GoBlog/api/internal/svc"
	"GoBlog/api/internal/types"
	"GoBlog/service/article/article"
)

type UpdateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleLogic) UpdateArticle(req *types.UpdateArticleReq) error {
	l.Logger.Info("开始更新文章")
	l.Logger.Infof("更新请求数据: %+v", req)

	// 从上下文获取用户ID
	userId, ok := l.ctx.Value(middleware.UserKey).(int64)
	if !ok {
		l.Logger.Error("获取用户ID失败")
		return errors.New("获取用户信息失败")
	}

	// 调用article rpc服务
	_, err := l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, &article.UpdateArticleRequest{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
		UserId:  userId, // 添加用户ID用于验证权限
	})

	if err != nil {
		l.Logger.Errorf("更新文章失败: %v", err)
		return fmt.Errorf("更新文章失败: %v", err)
	}

	l.Logger.Info("更新文章成功")
	return nil
}
