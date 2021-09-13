package auth

import (
	"Mustang/controllers/auth/db"
	"Mustang/utils/hack"

	"Mustang/models"
	"Mustang/utils/logs"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
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

type Result struct {
}

// @router /login [get,post]
func (c *AuthController) Login() {
	c.TplExt = "html"
	if c.Ctx.Input.Method() == "GET" {
		return
	}
	var userInfo UserLoginForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &userInfo); err != nil {
		logs.Error("get body error. %v", err)
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Ctx.Output.Body(hack.Slice("Invalid param"))
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
		logs.Warning(fmt.Sprintf("try to login in with usercontroller (%s) error %v ", authModel.UserName, err))
		c.Ctx.Output.SetStatus(http.StatusOK)
		c.Data["json"] = map[string]interface{}{"status": -1, "msg": fmt.Sprintf("Login failed. %v", err)}
		c.ServeJSON()
		return
	}

	user, err = models.UserModel.EnsureUser(user)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusOK)
		c.Data["json"] = map[string]interface{}{"status": -1, "msg": err.Error()}
		c.ServeJSON()
		return
	}
	c.SetSession("userId", user.Id)
	c.SetSession("userName", user.UserName)
	c.SetSession("role", user.Role)
	c.Data["json"] = map[string]interface{}{"status": 0, "msg": ""}
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.ServeJSON()
}

// @router /logout [get]
func (c *AuthController) Logout() {
	c.DelSession("userId")
	c.DelSession("userName")
	c.DelSession("role")
	c.Ctx.Redirect(302, "/login")
}
