package svc

import (
	"GoBlog/service/article/internal/config"
	"GoBlog/service/article/model"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config   config.Config
	Articles *model.Articles
	Redis    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := sql.Open("mysql", c.Mysql.DataSource)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	rdb := redis.New(c.RedisConf.Host, func(r *redis.Redis) {
		r.Type = c.RedisConf.Type
		r.Pass = c.RedisConf.Pass
	})

	return &ServiceContext{
		Config:   c,
		Articles: model.NewArticles(db),
		Redis:    rdb,
	}
}
