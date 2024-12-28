package comment_logic

import (
	"context"

	"GoBlog/api/internal/svc"
	"GoBlog/api/internal/types"
	"GoBlog/service/comment/comment"
	"GoBlog/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentListLogic) GetCommentList(req *types.CommentListReq) (*types.CommentListResp, error) {
	// 调用comment rpc服务
	resp, err := l.svcCtx.CommentRpc.GetCommentList(l.ctx, &comment.GetCommentListRequest{
		ArticleId: req.ArticleId,
		Page:      req.Page,
		PageSize:  req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	// 转换响应格式
	comments := make([]types.Comment, 0)
	for _, c := range resp.List {
		// 获取评论作者信息
		userResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserRequest{
			UserId: c.UserId,
		})
		if err != nil {
			return nil, err
		}

		comments = append(comments, types.Comment{
			Id:        c.Id,
			ArticleId: c.ArticleId,
			UserId:    c.UserId,
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
			Author: types.User{
				Id:        userResp.User.UserId,
				Username:  userResp.User.Username,
				CreatedAt: userResp.User.CreatedAt,
			},
		})
	}

	return &types.CommentListResp{
		Total:    resp.Total,
		Comments: comments,
	}, nil
}
