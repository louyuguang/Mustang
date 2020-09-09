package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var Cfg = beego.AppConfig

func InitDB() {
	//Init database
	RunMode := Cfg.String("runmode")
	if RunMode == "dev" {
		orm.Debug = true
	}
	db := Cfg.String("db")
	db_host := Cfg.String(RunMode+"::db_host")
	db_user := Cfg.String(RunMode+"::db_user")
	db_port := Cfg.String(RunMode+"::db_port")
	db_password := Cfg.String(RunMode+"::db_password")
	db_name := Cfg.String(RunMode+"::db_name")

	if db == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_password, db_host, db_port, db_name)
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", dsn)
	}
}
