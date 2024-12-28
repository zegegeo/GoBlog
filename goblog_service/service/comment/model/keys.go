package model

import "fmt"

const (
	prefixArticleComments = "article#%d#comments" // 文章评论列表缓存
	prefixComment         = "comment#%d"          // 单个评论缓存
)

// 生成文章评论列表缓存key
func GetArticleCommentsCacheKey(articleId int64) string {
	return fmt.Sprintf(prefixArticleComments, articleId)
}

// 生成评论缓存key
func GetCommentCacheKey(id int64) string {
	return fmt.Sprintf(prefixComment, id)
}
