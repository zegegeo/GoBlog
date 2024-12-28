package comment_handler

import (
	"net/http"
	"strconv"

	"GoBlog/api/internal/logic/comment"
	"GoBlog/api/internal/svc"
	"GoBlog/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateCommentReq
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

		l := comment_logic.NewCreateCommentLogic(r.Context(), svcCtx)
		err = l.CreateComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
