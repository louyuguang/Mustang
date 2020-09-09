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
	if c.Ctx.Input.Method() == "GET" {
		id, err := c.GetInt64("id")
		if err != nil {
			logs.Error("id's type is not int!%s", err)
			return
		}
		user, err := models.UserModel.GetUserById(id)
		if err != nil {
			logs.Error("get by id (%d) error.%v", id, err)
			return
		}
		c.Data["UserAdd"] = user
		c.TplName = "user/add.html"
		return
	}
}

func (c *UserController) Detail() {

}

func (c *UserController) List() {
	users, err := models.UserModel.GetAllUsers()
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}
	c.Data["Users"] = users
	c.TplName = "user/list.html"
}