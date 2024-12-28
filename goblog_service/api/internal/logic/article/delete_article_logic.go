package article_logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"GoBlog/api/internal/middleware"
	"GoBlog/api/internal/svc"
	"GoBlog/service/article/article"
)

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(id int64) error {
	l.Logger.Info("开始删除文章")
	l.Logger.Infof("删除文章ID: %d", id)

	// 从上下文获取用户ID
	userId, ok := l.ctx.Value(middleware.UserKey).(int64)
	if !ok {
		l.Logger.Error("获取用户ID失败")
		return errors.New("获取用户信息失败")
	}

	// 调用article rpc服务
	_, err := l.svcCtx.ArticleRpc.DeleteArticle(l.ctx, &article.DeleteArticleRequest{
		Id:     id,
		UserId: userId, // 添加用户ID用于验证权限
	})

	if err != nil {
		l.Logger.Errorf("删除文章失败: %v", err)
		return fmt.Errorf("删除文章失败: %v", err)
	}

	l.Logger.Info("删除文章成功")
	return nil
}
