package comment_logic

import (
	"context"
	"errors"

	"GoBlog/api/internal/svc"
	"GoBlog/api/internal/types"
	"GoBlog/service/comment/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentReq) error {
	userId, ok := l.ctx.Value("userId").(int64)
	if !ok {
		return errors.New("未登录或登录已过期")
	}

	// 调用comment rpc服务
	_, err := l.svcCtx.CommentRpc.DeleteComment(l.ctx, &comment.DeleteCommentRequest{
		Id:     req.Id,
		UserId: userId, // 用于验证评论所有权
	})
	if err != nil {
		return err
	}

	return nil
}
