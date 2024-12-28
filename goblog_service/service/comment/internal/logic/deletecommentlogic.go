package logic

import (
	"context"
	"database/sql"
	"errors"

	"GoBlog/service/comment/comment"
	"GoBlog/service/comment/internal/svc"
	"GoBlog/service/comment/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(in *comment.DeleteCommentRequest) (*comment.DeleteCommentResponse, error) {
	// 获取评论信息
	commentInfo, err := l.svcCtx.Comments.FindById(in.Id)
	if err != nil {
		return nil, err
	}
	if commentInfo == nil {
		return nil, errors.New("评论不存在")
	}

	// 删除评论
	err = l.svcCtx.Comments.Delete(in.Id, in.UserId)
	if err == sql.ErrNoRows {
		return nil, errors.New("无权删除该评论")
	}
	if err != nil {
		return nil, err
	}

	// 删除相关缓存
	l.svcCtx.Redis.Del(l.svcCtx.Config.Redis.Key + model.GetCommentCacheKey(in.Id))
	l.svcCtx.Redis.Del(l.svcCtx.Config.Redis.Key + model.GetArticleCommentsCacheKey(commentInfo.ArticleId))

	return &comment.DeleteCommentResponse{}, nil
}
