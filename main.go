package main

import (
	"Mustang/models"
	"Mustang/pkg/machinery"
	_ "Mustang/routers"
	"Mustang/utils"
	"encoding/gob"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func init() {
	utils.InitDB()
	machinery.CreateMachinery()
	gob.Register(&models.Role{})
}
func main() {
	orm.RunCommand()
	beego.Run()
}
