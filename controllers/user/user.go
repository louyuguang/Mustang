package user

import (
	"Mustang/controllers/base"
	"Mustang/models"
	//"Mustang/utils/hack"
	"Mustang/utils/logs"
	"encoding/json"
	"net/http"
)

type UserController struct {
	base.BaseController
}

func (c *UserController) URLMapping() {
	c.Mapping("Add", c.Add)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Detail", c.Detail)
	c.Mapping("List", c.List)
}

func (c *UserController) Prepare() {
	c.BaseController.Prepare()
}

// @router /add [get,post]
func (c *UserController) Add() {
	if c.Ctx.Input.Method() == "GET" {
		roles, err := models.RoleModel.GetAllRoles()
		if err != nil {
			c.CustomAbort(http.StatusInternalServerError, err.Error())
		}
		c.Data["Roles"] = roles
		return
	}
	//POST, Add a new User
	var user *models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user);err != nil {
		logs.Error("User's form error. %v", err)
		//c.Data["json"] = map[string]interface{}{"status": -1, "error": "User's form error."}
		//c.ServeJSON()
		//return
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}
	_, err := models.UserModel.AddUser(user)
	if err != nil {
		c.Fail(err)
	}
	c.Success("Add User Success!")
}

// @router /update/:id [get,post]
func (c *UserController) Update() {
	id := c.GetIDFromURL()
	//GET
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
		return
	}
	//POST
	var user *models.User
	if id != 0 {
		user.Id = id
	} else {
		//Add a new User
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
		if err != nil {
			logs.Error("User's form error. %v", err)
			c.Data["json"] = map[string]interface{}{"status": -1, "error": "User's form error."}
			c.ServeJSON()
			return
		}
		_, err = models.UserModel.AddUser(user)
		if err != nil {
			//c.Data[""]
			return
		}
	}
}

// @router /detail [get]
func (c *UserController) Detail() {

}

// @router /list [get]
func (c *UserController) List() {
	scontent := c.GetString("scontent")
	pers := 10
	cnt, err := models.UserModel.GetAllNum(scontent)
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
}

// @router /delete [post]
func (c *UserController) Delete() {
	return
}
