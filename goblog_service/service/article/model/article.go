package model

import (
	"database/sql"
	"time"
)

type Article struct {
	Id        int64     `db:"id"`
	UserId    int64     `db:"user_id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Articles struct {
	db *sql.DB
}

func NewArticles(db *sql.DB) *Articles {
	return &Articles{db: db}
}

// 创建文章
func (a *Articles) Insert(userId int64, title, content string) (int64, error) {
	result, err := a.db.Exec("INSERT INTO articles (user_id, title, content) VALUES (?, ?, ?)",
		userId, title, content)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// 查询文章列表
func (a *Articles) Find(offset, limit int) ([]*Article, error) {
	var articles []*Article
	rows, err := a.db.Query(
		"SELECT id, user_id, title, content, created_at, updated_at FROM articles ORDER BY created_at DESC LIMIT ? OFFSET ?",
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var article Article
		err := rows.Scan(
			&article.Id,
			&article.UserId,
			&article.Title,
			&article.Content,
			&article.CreatedAt,
			&article.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}
	return articles, nil
}

// 获取单篇文章
func (a *Articles) FindById(id int64) (*Article, error) {
	var article Article
	err := a.db.QueryRow("SELECT id, user_id, title, content, created_at, updated_at FROM articles WHERE id = ?", id).
		Scan(&article.Id, &article.UserId, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// 更新文章
func (a *Articles) Update(id int64, title, content string) error {
	_, err := a.db.Exec("UPDATE articles SET title = ?, content = ? WHERE id = ?",
		title, content, id)
	return err
}

// 删除文章
func (a *Articles) Delete(id int64) error {
	_, err := a.db.Exec("DELETE FROM articles WHERE id = ?", id)
	return err
}

// 获取文章总数
func (a *Articles) Count() (int64, error) {
	var count int64
	err := a.db.QueryRow(
		"SELECT COUNT(*) FROM articles",
	).Scan(&count)
	return count, err
}
