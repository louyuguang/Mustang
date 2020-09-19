package main

import (
	"Mustang/models"
	_ "Mustang/routers"
	"Mustang/utils"
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
)

func init() {
	utils.InitDB()
	gob.Register(&models.Role{})
}
func main() {
	orm.RunCommand()
	beego.Run()
}
