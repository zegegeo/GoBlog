package article_handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	article_logic "GoBlog/api/internal/logic/article"
	"GoBlog/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := article_logic.NewDeleteArticleLogic(r.Context(), svcCtx)
		l.Logger.Info("开始删除文章")

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

		err = l.DeleteArticle(id)
		if err != nil {
			l.Logger.Errorf("删除文章失败: %v", err)
			httpx.Error(w, err)
		} else {
			l.Logger.Info("删除文章成功")
			httpx.Ok(w)
		}
	}
}
