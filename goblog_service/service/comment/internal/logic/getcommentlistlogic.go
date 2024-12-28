package logic

import (
	"context"

	"GoBlog/service/comment/comment"
	"GoBlog/service/comment/internal/svc"
	"GoBlog/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListLogic) GetCommentList(in *comment.GetCommentListRequest) (*comment.GetCommentListResponse, error) {
	// 获取评论列表
	comments, err := l.svcCtx.Comments.FindByArticleId(in.ArticleId, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}

	// 获取评论总数
	total, err := l.svcCtx.Comments.CountByArticleId(in.ArticleId)
	if err != nil {
		return nil, err
	}

	// 构建响应
	commentList := make([]*comment.CommentInfo, 0)
	for _, c := range comments {
		// 获取评论者信息
		userResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserRequest{
			UserId: c.UserId,
		})
		if err != nil {
			return nil, err
		}

		commentList = append(commentList, &comment.CommentInfo{
			Id:        c.Id,
			ArticleId: c.ArticleId,
			UserId:    c.UserId,
			Username:  userResp.User.Username,
			Content:   c.Content,
			CreatedAt: c.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &comment.GetCommentListResponse{
		List:  commentList,
		Total: total,
	}, nil
}
