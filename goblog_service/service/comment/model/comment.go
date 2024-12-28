package model

import (
	"database/sql"
	"time"
)

type Comment struct {
	Id        int64     `db:"id"`
	ArticleId int64     `db:"article_id"`
	UserId    int64     `db:"user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Comments struct {
	db *sql.DB
}

func NewComments(db *sql.DB) *Comments {
	return &Comments{db: db}
}

// 创建评论
func (c *Comments) Insert(articleId, userId int64, content string) (int64, error) {
	result, err := c.db.Exec("INSERT INTO comments (article_id, user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		articleId, userId, content, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// 获取文章的评论列表
func (c *Comments) FindByArticleId(articleId int64, page, pageSize int64) ([]*Comment, error) {
	offset := (page - 1) * pageSize
	rows, err := c.db.Query(`
		SELECT id, article_id, user_id, content, created_at, updated_at 
		FROM comments 
		WHERE article_id = ? 
		ORDER BY created_at DESC 
		LIMIT ? OFFSET ?`,
		articleId, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.Id,
			&comment.ArticleId,
			&comment.UserId,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

// 获取评论总数
func (c *Comments) CountByArticleId(articleId int64) (int64, error) {
	var count int64
	err := c.db.QueryRow("SELECT COUNT(*) FROM comments WHERE article_id = ?", articleId).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 获取单个评论
func (c *Comments) FindById(id int64) (*Comment, error) {
	var comment Comment
	err := c.db.QueryRow(`
		SELECT id, article_id, user_id, content, created_at, updated_at 
		FROM comments 
		WHERE id = ?`,
		id).Scan(
		&comment.Id,
		&comment.ArticleId,
		&comment.UserId,
		&comment.Content,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// 删除评论
func (c *Comments) Delete(id, userId int64) error {
	result, err := c.db.Exec("DELETE FROM comments WHERE id = ? AND user_id = ?", id, userId)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
