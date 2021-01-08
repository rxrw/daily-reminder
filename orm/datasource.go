package orm

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine

//GetEngine 获取数据库实例
func GetEngine() *xorm.Engine {
	if engine != nil {
		return engine
	}
	var err error
	engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		os.Getenv("MYSQL_1_USER"),
		os.Getenv("MYSQL_1_PASSWORD"),
		os.Getenv("MYSQL_1_HOST"),
		os.Getenv("MYSQL_1_PORT"),
		os.Getenv("MYSQL_1_DBNAME"),
		os.Getenv("DB_CHARSET")))

	if err != nil {
		log.Printf("%v", err)
	}
	engine.SetMapper(names.GonicMapper{})
	engine.Sync2(new(Lottery), new(HotSearch), new(Weather), new(UserInfo))
	return engine
}
