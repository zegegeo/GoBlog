package comment_logic

import (
	"context"
	"errors"

	"GoBlog/api/internal/svc"
	"GoBlog/api/internal/types"
	"GoBlog/service/comment/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentReq) error {
	userId, ok := l.ctx.Value("userId").(int64)
	if !ok {
		return errors.New("未登录或登录已过期")
	}

	// 调用comment rpc服务
	_, err := l.svcCtx.CommentRpc.CreateComment(l.ctx, &comment.CreateCommentRequest{
		ArticleId: req.ArticleId,
		UserId:    userId,
		Content:   req.Content,
	})
	if err != nil {
		return err
	}

	return nil
}
