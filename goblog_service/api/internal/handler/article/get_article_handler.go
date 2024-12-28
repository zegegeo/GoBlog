package article_handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	article_logic "GoBlog/api/internal/logic/article"
	"GoBlog/api/internal/svc"
	"GoBlog/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := article_logic.NewGetArticleLogic(r.Context(), svcCtx)
		l.Logger.Info("开始获取文章详情")

		path := r.URL.Path
		parts := strings.Split(path, "/")
		vars := parts[len(parts)-1]
		l.Logger.Infof("获取到文章ID: %s", vars)

		if vars == "" {
			httpx.Error(w, errors.New("文章ID不能为空"))
			return
		}

		id, err := strconv.ParseInt(vars, 10, 64)
		if err != nil {
			l.Logger.Errorf("解析文章ID失败: %v", err)
			httpx.Error(w, errors.New("无效的文章ID"))
			return
		}

		l.Logger.Infof("解析后的文章ID: %d", id)
		var req types.ArticleReq
		req.Id = id

		resp, err := l.GetArticle(&req)
		if err != nil {
			l.Logger.Errorf("获取文章失败: %v", err)
			httpx.Error(w, err)
		} else {
			l.Logger.Info("获取文章成功")
			httpx.OkJson(w, resp)
		}
	}
}
