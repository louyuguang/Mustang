package main

import (
	_ "Mustang/controllers"
	_ "Mustang/routers"
	"Mustang/utils"
	"Mustang/models"
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
