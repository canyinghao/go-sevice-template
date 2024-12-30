package services

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/canyinghao/go-sevice-template/pkg"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var db *sql.DB

var ctx = context.Background()

var clients map[string]*redis.Client

func InitServices(c *pkg.Config) {

	// 初始化数据库
	pg, err := sql.Open("postgres", "user="+c.Pgsql.User+" password="+c.Pgsql.Password+" dbname="+c.Pgsql.Dbname+" host="+c.Pgsql.Host+" port="+strconv.Itoa(c.Pgsql.Port)+" sslmode=disable")
	checkErr(err)
	db = pg
	// 判断是否连接成功
	err = db.Ping()
	checkErr(err)
	// 初始化redis
	clients = make(map[string]*redis.Client)
	for k := range c.Redis {
		address := fmt.Sprintf("%s:%d", c.Redis[k].Host, c.Redis[k].Port)
		rdb := redis.NewClient(&redis.Options{
			Addr:            address,
			Password:        c.Redis[k].Password,
			DB:              c.Redis[k].Db,
			DialTimeout:     3000,
			ReadTimeout:     3000,
			WriteTimeout:    3000,
			MaxIdleConns:    5,
			MaxActiveConns:  30,
			ConnMaxIdleTime: 60000,
		})
		// 判断是否连接成功
		// _, err := rdb.Ping(ctx).Result()
		// checkErr(err)
		clients[k] = rdb
	}

	zap.L().Info("init db success")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
