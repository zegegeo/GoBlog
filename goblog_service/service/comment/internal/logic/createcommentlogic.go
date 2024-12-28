package logic

import (
	"context"

	"GoBlog/service/comment/comment"
	"GoBlog/service/comment/internal/svc"
	"GoBlog/service/comment/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentLogic) CreateComment(in *comment.CreateCommentRequest) (*comment.CreateCommentResponse, error) {
	// 创建评论
	id, err := l.svcCtx.Comments.Insert(in.ArticleId, in.UserId, in.Content)
	if err != nil {
		return nil, err
	}

	// 删除文章评论列表缓存
	l.svcCtx.Redis.Del(l.svcCtx.Config.Redis.Key + model.GetArticleCommentsCacheKey(in.ArticleId))

	return &comment.CreateCommentResponse{
		Id: id,
	}, nil
}
