package services

import (
	"database/sql"

	"github.com/canyinghao/go-sevice-template/pkg"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var db *sql.DB

func InitServices(c *pkg.Config) {

	// 初始化数据库
	pg, err := sql.Open("postgres", "user="+c.Pgsql.User+" password="+c.Pgsql.Password+" dbname="+c.Pgsql.Dbname+" host="+c.Pgsql.Host+" port="+c.Pgsql.Port+" sslmode=disable")
	checkErr(err)
	db = pg
	zap.L().Info("init db success")

	// 初始化redis

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
