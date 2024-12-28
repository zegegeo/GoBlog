package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type User struct {
	Id        int64     `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{db: db}
}

func (u *Users) FindByUsername(username string) (*User, error) {
	var user User
	err := u.db.QueryRow("SELECT id, username, password, created_at FROM users WHERE username = ?", username).
		Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *Users) FindById(id int64) (*User, error) {
	var user User
	err := u.db.QueryRow("SELECT id, username, password, created_at FROM users WHERE id = ?", id).
		Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *Users) Insert(username, password string) (int64, error) {
	// 检查参数
	if len(username) == 0 || len(password) == 0 {
		return 0, fmt.Errorf("用户名或密码不能为空")
	}

	result, err := u.db.Exec(
		"INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)",
		username,
		password,
		time.Now(),
	)
	if err != nil {
		// 检查是否是唯一键冲突
		if strings.Contains(err.Error(), "Duplicate entry") {
			return 0, fmt.Errorf("用户名已存在")
		}
		return 0, err
	}
	return result.LastInsertId()
}
