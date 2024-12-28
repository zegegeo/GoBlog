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

func UpdateArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := article_logic.NewUpdateArticleLogic(r.Context(), svcCtx)
		l.Logger.Info("开始更新文章")

		// 从URL路径中获取最后一个部分作为ID
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

		var req types.UpdateArticleReq
		req.Id = id
		if err := httpx.Parse(r, &req); err != nil {
			l.Logger.Errorf("解析请求数据失败: %v", err)
			httpx.Error(w, err)
			return
		}

		err = l.UpdateArticle(&req)
		if err != nil {
			l.Logger.Errorf("更新文章失败: %v", err)
			httpx.Error(w, err)
		} else {
			l.Logger.Info("更新文章成功")
			httpx.Ok(w)
		}
	}
}
