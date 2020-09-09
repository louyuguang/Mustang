package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       int64       `orm:"pk;auto" json:"id,omitempty"`
	Role     *Role     `orm:"rel(fk);default(3);column(role_id)"`
	UserName string    `orm:"index;unique;size(200);column(username)" json:"name,omitempty"`
	Password string    `orm:"size(255)" json:"-"`
	RealName string    `orm:"size(255)"`
	Email    string    `orm:"size(200)" json:"email,omitempty"`
	Active   bool      `orm:"default(true)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)" json:"createTime,omitempty"`
}

//var (
//	UserModel *User
//)

type userModel struct{}

func (*userModel) GetAllUsers() ([]*User, error) {
	var users []*User
	_, err := Ormer().QueryTable(new(User)).RelatedSel().All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (*userModel) GetUserById(id int64) (v *User, err error) {
	v = &User{Id: id}
	if err = Ormer().Read(v); err != nil {
		return nil, err
	}
	return v, nil
}

func (*userModel) GetUserByName(username string) (v *User, err error) {
	v = &User{UserName: username}
	if err = Ormer().Read(v, "username"); err != nil {
		return nil, err
	}
	Ormer().LoadRelated(v, "Role")
	return v, nil
}

func (*userModel) EnsureUser(m *User) (*User, error) {
	oldUser := &User{UserName: m.UserName}
	err := Ormer().Read(oldUser, "username")
	//
	if err != nil {
		if err == orm.ErrNoRows {
			_, err := UserModel.AddUser(m)
			if err != nil {
				return nil, err
			}
			oldUser = m
		} else {
			return nil, err
		}
	} else {
		oldUser.Email = m.Email
		oldUser.Active = m.Active
		oldUser.Role = m.Role
		_, err := Ormer().Update(oldUser)
		if err != nil {
			return nil, err
		}
	}
	Ormer()
	return oldUser, err
}

func (*userModel) AddUser(m *User) (id int64, err error) {
	id, err = Ormer().Insert(m)
	if err != nil {
		return
	}
	return id, nil
}

func (*userModel) GetUserDetail(name string) (user *User, err error) {
	user = &User{UserName: name}

	err = Ormer().Read(user, "username")
	if err != nil {
		return nil, err
	}
	Ormer().LoadRelated(user, "Role")
	return user, nil
}

//func (*userModel) EditUser(id int64) (v *User, err error) {
//	if c.Ctx.Input.Method() == "GET" {
//		c.TplName = "user/add.html"
//		return
//	}
//	return
//}
