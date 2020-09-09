package auth

import (
	"Mustang/controllers/auth/db"
	"Mustang/models"
	"Mustang/utils/logs"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

var GlobalUserSalt  = beego.AppConfig.String("GlobalUserSalt")

type AuthController struct {
	beego.Controller
}

type Authenticator interface {
	// Authenticate ...
	Authenticate(m models.AuthModel) (*models.User, error)
}

type LoginResult struct {
	Token string `json:"token"`
}

type UserLoginForm struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (c *AuthController) Login() {
	if c.Ctx.Input.Method() == "GET" {
		c.TplName = "login.html"
		return
	}
	var msg string
	var userInfo UserLoginForm
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		logs.Error("get login form body error. %v", err)
		c.Data["json"] = map[string]interface{}{"status": -1, "error": "Get frontend body error."}
		c.ServeJSON()
		return
	}
	if userInfo.UserName == "" || userInfo.Password == "" {
		c.Data["json"] = map[string]interface{}{"status": -1, "error": "username or password cannot be empty!"}
		c.ServeJSON()
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
		msg = fmt.Sprintf("try to login in with user (%s) error %v ", authModel.UserName, err)
		logs.Warning(msg)
		c.Data["json"] = map[string]interface{}{"status": 1, "msg": "用户名或者密码不正确"}
		c.ServeJSON()
		return
	}

	user, err = models.UserModel.EnsureUser(user)
	if err != nil {
		msg = "Internal server error."
		logs.Error(msg)
		c.Data["json"] = map[string]interface{}{"status": -1, "msg": "内部未知错误"}
		c.ServeJSON()
		return
	}
	c.SetSession("userId", user.Id)
	c.SetSession("userName", user.UserName)
	c.SetSession("role", user.Role)
	c.Data["json"] = map[string]interface{}{"status": 0, "error": ""}
	c.ServeJSON()
}

func (c *AuthController) Logout() {
	c.DelSession("userId")
	c.DelSession("userName")
	c.DelSession("role")
	c.Ctx.Redirect(302,"/login")
}
