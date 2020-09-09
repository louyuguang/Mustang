package controllers

import (
	"Mustang/models"
	"github.com/astaxie/beego"
	"net/http"
)

type BaseController struct {
	beego.Controller
	User *models.User
	//IsAdmin bool
	//o orm.Ormer
}

func (c *BaseController) Prepare() {
	userId := c.GetSession("userId")
	if userId == nil {
		url := "/login"
		next := c.Ctx.Request.URL.Path
		if next != "" {
			url += "?next=" + next
		}
		c.Ctx.Redirect(302, url) //若Session中无用户ID则302重定向至登陆页面
	}
	UserName := c.GetSession("userName")
	user, err := models.UserModel.GetUserDetail(UserName.(string))
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}
	c.User = user
	c.Data["User"] = user
}

// @router / [get]
func (c *BaseController) Index() {
	c.TplName = "index.html"
}