package app

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/astaxie/beego"
)

var Machinery *machinery.Server

//var web beego.AddTemplateExt
var Cfg = beego.AppConfig
