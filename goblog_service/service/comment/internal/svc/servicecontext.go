package svc

import (
	"GoBlog/service/comment/internal/config"
	"GoBlog/service/comment/model"
	"GoBlog/service/user/userclient"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	Comments *model.Comments
	Redis    *redis.Redis
	UserRpc  userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 连接MySQL
	db, err := sql.Open("mysql", c.Mysql.DataSource)
	if err != nil {
		panic(err)
	}

	// 创建Redis客户端
	rdb := redis.New(c.RedisConf.Host, func(r *redis.Redis) {
		r.Type = c.RedisConf.Type
		r.Pass = c.RedisConf.Pass
	})

	// 创建用户服务客户端
	userRpc := userclient.NewUser(zrpc.MustNewClient(c.UserRpc))

	return &ServiceContext{
		Config:   c,
		Comments: model.NewComments(db),
		Redis:    rdb,
		UserRpc:  userRpc,
	}
}
