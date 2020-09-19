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
	//paramByte, err := hack.FormToJson(c.Ctx.Input.RequestBody)
	//if err != nil {
	//	logs.Error("%v", err)
	//	c.Data["json"] = map[string]interface{}{"status": -1, "error": "post data invalid"}
	//	c.ServeJSON()
	//	return
	//}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &userInfo); err != nil {
		//logs.Error("get login form body error. %v", err)
		//c.Data["json"] = map[string]interface{}{"status": -1, "error": "Get frontend body error."}
		//c.ServeJSON()
		//return
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}
	if userInfo.UserName == "" || userInfo.Password == "" {
		//c.Data["json"] = map[string]interface{}{"status": -1, "error": "username or password cannot be empty!"}
		//c.ServeJSON()
		//return
		c.Fail("username or password cannot be empty!")
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
		//c.Data["json"] = map[string]interface{}{"status": 1, "msg": "用户名或者密码不正确"}
		//c.ServeJSON()
		//return
		c.Fail( "用户名或者密码不正确")
		return
	}

	user, err = models.UserModel.EnsureUser(user)
	if err != nil {
		//msg = "Internal server error."
		//logs.Error(msg)
		//c.Data["json"] = map[string]interface{}{"status": -1, "msg": "内部未知错误"}
		//c.ServeJSON()
		//return
		c.CustomAbort(http.StatusInternalServerError, err.Error())
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
