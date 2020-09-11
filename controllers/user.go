package controllers

import (
	"Mustang/models"
	"Mustang/utils/logs"
	"net/http"
)

//type Result struct {
//	Data []interface{} `json:"data"`
//}

type UserController struct {
	BaseController
}

func (c *UserController) Prepare() {
	c.BaseController.Prepare()
}

func (c *UserController) Add() {
	id := c.GetIDFromURL()
	if c.Ctx.Input.Method() == "GET" {
		roles, err := models.RoleModel.GetAllRoles()
		if err != nil {
			logs.Error("get all roles error")
			return
		}
		if id != 0 {
			user, err := models.UserModel.GetUserById(id)
			if err != nil {
				logs.Error("get by id (%d) error.%v", id, err)
				return
			}
			c.Data["UserAdd"] = user
		}
		c.Data["Roles"] = roles
		c.TplName = "user/add.html"
		return
	}
}

func (c *UserController) Detail() {

}

func (c *UserController) List() {
	scontent := c.GetString("scontent")
	pers := 10
	cnt, err := models.UserModel.GetAllNum(scontent);
	if err != nil || cnt == -1 {
		logs.Error("get all users nums error")
		return
	}

	pager := c.SetPaginator(pers, cnt)
	users, err := models.UserModel.GetUsers(pers, pager.Offset(), scontent)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}
	c.Data["Users"] = users
	c.Data["Scontent"] = scontent
	c.TplName = "user/list.html"
}

func (c *UserController) Delete() {

}
