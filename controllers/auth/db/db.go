package db

import (
	//"Mustang/controllers/auth"
	"Mustang/models"
	"Mustang/utils/encode"
	"fmt"
	"github.com/astaxie/beego"
)

type DBAuth struct{}

//func init() {
//	auth.Register(models.AuthTypeDB, &DBAuth{})
//}
var GlobalUserSalt  = beego.AppConfig.String("GlobalUserSalt")


type CurrentInfo struct {
	User   *models.User      `json:"user"`
	Config map[string]string `json:"config"`
}

func (*DBAuth) Authenticate(m models.AuthModel) (*models.User, error) {
	username := m.UserName
	password := m.Password
	user, err := models.UserModel.GetUserByName(username)
	if err != nil {
		return nil, err
	}

	if user.Password == "" {
		return nil, fmt.Errorf("user dons't support db login!")
	}

	passwordHashed := encode.EncodePassword(password, GlobalUserSalt)

	if passwordHashed != user.Password {
		return nil, fmt.Errorf("username or password error!")
	}
	return user, nil
}
