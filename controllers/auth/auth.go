package auth

import (
	"Mustang/controllers/auth/db"
	"Mustang/controllers/base"
	"Mustang/models"
	"Mustang/utils/logs"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthController struct {
	base.ResultHandlerController
}

func (c *AuthController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Logout", c.Logout)
}

type Authenticator interface {
	// Authenticate ...
	Authenticate(m models.AuthModel) (*models.User, error)
}

type UserLoginForm struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// @router /login [get,post]
func (c *AuthController) Login() {
	if c.Ctx.Input.Method() == "GET" {
		return
	}
	var msg string
	var userInfo UserLoginForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &userInfo); err != nil {
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}
	if userInfo.UserName == "" || userInfo.Password == "" {
		c.Fail("username or password cannot be empty!")
		return
	}
	var authenticator Authenticator

	authModel := models.AuthModel{
		UserName: userInfo.UserName,
		Password: userInfo.Password,
	}

	authenticator = &db.DBAuth{}
	user, err := authenticator.Authenticate(authModel)
	if err != nil {
		msg = fmt.Sprintf("try to login in with usercontroller (%s) error %v ", authModel.UserName, err)
		logs.Warning(msg)
		c.Fail( "用户名或者密码不正确")
		return
	}

	user, err = models.UserModel.EnsureUser(user)
	if err != nil {
		c.Fail(err)
		return
	}
	c.SetSession("userId", user.Id)
	c.SetSession("userName", user.UserName)
	c.SetSession("role", user.Role)
	c.Success(nil)
}

// @router /logout [get]
func (c *AuthController) Logout() {
	c.DelSession("userId")
	c.DelSession("userName")
	c.DelSession("role")
	c.Ctx.Redirect(302, "/login")
}
