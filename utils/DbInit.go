package utils

import (
	"Mustang/app"
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() {
	//Init database
	cfg := app.Cfg
	RunMode := cfg.String("runmode")
	if RunMode == "dev" {
		orm.Debug = true
	}
	db := cfg.String("db")
	dbHost := cfg.String("db_host")
	dbUser := cfg.String("db_user")
	dbPort := cfg.String("db_port")
	dbPassword := cfg.String("db_password")
	dbName := cfg.String("db_name")

	if db == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbName)
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", dsn)
	}
}
