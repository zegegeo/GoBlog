package comment_handler

import (
	"net/http"
	"strconv"

	"GoBlog/api/internal/logic/comment"
	"GoBlog/api/internal/svc"
	"GoBlog/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从路径参数获取评论ID
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		req := &types.DeleteCommentReq{Id: id}
		l := comment_logic.NewDeleteCommentLogic(r.Context(), svcCtx)
		err = l.DeleteComment(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
