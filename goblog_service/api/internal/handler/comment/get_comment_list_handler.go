package comment_handler

import (
	"net/http"
	"strconv"

	"GoBlog/api/internal/logic/comment"
	"GoBlog/api/internal/svc"
	"GoBlog/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 从路径参数获取文章ID
		articleIdStr := r.PathValue("article_id")
		articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		req.ArticleId = articleId

		// 设置默认值
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 10
		}

		l := comment_logic.NewGetCommentListLogic(r.Context(), svcCtx)
		resp, err := l.GetCommentList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
