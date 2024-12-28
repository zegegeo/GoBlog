package article_handler

import (
	"net/http"

	"GoBlog/api/internal/logic/article"
	"GoBlog/api/internal/svc"
	"GoBlog/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetArticleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 10
		}

		l := article_logic.NewGetArticleListLogic(r.Context(), svcCtx)
		resp, err := l.GetArticleList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
