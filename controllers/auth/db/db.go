package db

import (
	"Mustang/models"
	"Mustang/utils/encode"
	"fmt"
	"github.com/astaxie/beego"
)

type DBAuth struct{}

var GlobalUserSalt = beego.AppConfig.String("GlobalUserSalt")

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
