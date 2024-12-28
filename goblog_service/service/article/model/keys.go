package model

import "fmt"

const (
	prefixArticleCache = "article#%d"       // 单篇文章缓存
	prefixUserArticles = "user#%d#articles" // 用户文章列表缓存
)

// 生成文章缓存key
func GetArticleCacheKey(id int64) string {
	return fmt.Sprintf(prefixArticleCache, id)
}

// 生成用户文章列表缓存key
func GetUserArticlesCacheKey(userId int64) string {
	return fmt.Sprintf(prefixUserArticles, userId)
}
