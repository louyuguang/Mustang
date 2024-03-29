package user

import (
	"Mustang/controllers/base"
	"Mustang/models"
	"Mustang/utils/logs"
	"encoding/json"
	"fmt"
)

type UserController struct {
	base.BaseController
}

func (c *UserController) URLMapping() {
	c.Mapping("Add", c.Add)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Detail", c.Detail)
	c.Mapping("List", c.List)
	c.Mapping("Update", c.Update)
}

func (c *UserController) Prepare() {
	c.BaseController.Prepare()
	c.Data["LevelOne"] = map[string]string{"content": "用户管理"}
}

// @router /add [get,post]
func (c *UserController) Add() {
	if c.Ctx.Input.Method() == "GET" {
		roles, err := models.RoleModel.GetAllRoles()
		if err != nil {
			c.Fail(err)
			return
		}
		c.Data["Roles"] = roles
		return
	}
	//POST, Add a new User
	var user *models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		logs.Error("User's form error. %v", err)
		c.Fail(err)
		return
	}
	_, err := models.UserModel.AddUser(user)
	if err != nil {
		c.Fail(err)
		return
	}
	c.Success(fmt.Sprintf("用户%s创建成功！", user.UserName))
}

// @router /update/:id [get,post]
func (c *UserController) Update() {
	id := c.GetIDFromURL()
	//GET
	if c.Ctx.Input.Method() == "GET" {
		roles, err := models.RoleModel.GetAllRoles()
		if err != nil {
			c.Fail(err)
			return
		}
		if id != 0 {
			user, err := models.UserModel.GetUserById(id)
			if err != nil {
				logs.Error("get by id (%d) error.%v", id, err)
				c.Fail(err)
				return
			}
			c.Data["UserAdd"] = user
		}
		c.Data["Roles"] = roles
		return
	}
	//POST
	var user *models.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Error("User's form error. %v", err)
		c.Fail(err)
		return
	}
	if err = models.UserModel.UpdateById(user); err != nil {
		c.Fail(err)
		return
	}

	c.Success("用户信息更新成功！")
}

// @router /detail/:id [get]
func (c *UserController) Detail() {

}

// @router /list [get]
func (c *UserController) List() {
	scontent := c.GetString("scontent")
	pers := 10
	cnt, err := models.UserModel.GetAllNum(scontent)
	if err != nil || cnt == -1 {
		logs.Error("get all users nums error")
		c.Fail(err)
		return
	}

	pager := c.SetPaginator(pers, cnt)
	users, err := models.UserModel.GetUsers(pers, pager.Offset(), scontent)
	if err != nil {
		c.Fail(err)
		return
	}
	c.Data["Users"] = users
	c.Data["Scontent"] = scontent
}

// @router /delete [post]
func (c *UserController) Delete() {
	var user *models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		c.Fail(err)
		return
	}
	if err := models.UserModel.DeleteById(user); err != nil {
		c.Fail(err)
		return
	}
	c.Success("删除成功！")
}
