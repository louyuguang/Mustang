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

func (*userModel) GetAllNum(scontent ...string) (num int64, err error) {
	query := map[string]interface{}{}
	if scontent != nil {
		query["username__icontains"] = scontent
	}
	qs := Ormer().QueryTable(new(User))
	qs = BuildFilter(qs, query)
	num, err = qs.Count()
	if err != nil {
		return -1, err
	}
	return num, nil
}

func (*userModel) GetUsers(pers int, offset int, scontent ...string) ([]*User, error) {
	var users []*User
	query := map[string]interface{}{}
	qs := Ormer().QueryTable(new(User))
	if scontent != nil {
		query["username__icontains"] = scontent
	}
	qs = BuildFilter(qs, query)
	qs.Limit(pers, offset).RelatedSel().All(&users)
	return users, nil
}

func (*userModel) GetUserById(id int64) (v *User, err error) {
	v = &User{Id: id}
	if err = Ormer().Read(v); err != nil {
		return nil, err
	}
	_, err = Ormer().LoadRelated(v, "Role")
	if err == nil {
		return v, nil
	}
	return nil, err
}

func (*userModel) GetUserByName(username string) (v *User, err error) {
	v = &User{UserName: username}
	if err = Ormer().Read(v, "username"); err != nil {
		return nil, err
	}
	_, err = Ormer().LoadRelated(v, "Role")
	if err == nil {
		return v, nil
	}
	return nil, err
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
