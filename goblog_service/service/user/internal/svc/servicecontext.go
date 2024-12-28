package svc

import (
	"GoBlog/service/user/internal/config"
	"GoBlog/service/user/model"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ServiceContext struct {
	Config config.Config
	Users  *model.Users
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 连接数据库
	db, err := sql.Open("mysql", c.Mysql.DataSource)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		Users:  model.NewUsers(db),
	}
}
