package base

import (
	"Mustang/models"
	"Mustang/utils/paginator"
	"fmt"
	"strconv"
)

type BaseController struct {
	ResultHandlerController
	User *models.User
}

func (c *BaseController) URLMapping() {
	c.Mapping("Index", c.Index)
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
		c.Fail(err)
		return
	}
	c.Data["User"] = user
}

func (c *BaseController) GetIDFromURL() int64 {
	return c.GetIntParamFromURL(":id")
}

func (c *BaseController) GetIntParamFromURL(param string) int64 {
	paramStr := c.Ctx.Input.Param(param)
	if paramStr != "" {
		paramInt, err := strconv.ParseInt(paramStr, 10, 64)
		if err != nil || paramInt < 0 {
			c.AbortBadRequest(fmt.Sprintf("Invalid %s in URL", param))
		}
		return paramInt
	}
	return 0
}

// @router / [get]
func (c *BaseController) Index() {
	return
}

func (c *BaseController) SetPaginator(pers int, cnt int64) *paginator.Paginator {
	p := paginator.NewPaginator(c.Ctx.Request, pers, cnt)
	c.Data["paginator"] = p
	return p
}