package models

import (
	"Mustang/utils/encode"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       int64      `orm:"pk;auto" json:"id,omitempty"`
	Role     *Role      `orm:"rel(fk);default(4);column(role_id)" json:"role"`
	UserName string     `orm:"index;unique;size(200);column(username);" json:"username,omitempty"`
	Password string     `orm:"size(255)" json:"password"`
	RealName string     `orm:"size(255)" json:"realname,omitempty"`
	Email    string     `orm:"size(200)" json:"email,omitempty"`
	Active   bool       `orm:"default(true)" json:"is_active"`
	Created  *time.Time `orm:"auto_now_add;type(datetime)" json:"createTime,omitempty"`
}

type userModel struct{}

var GlobalUserSalt = beego.AppConfig.String("GlobalUserSalt")

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
	_, _ = qs.Limit(pers, offset).RelatedSel().All(&users)
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
		//oldUser.Email = m.Email
		_, err := Ormer().Update(oldUser)
		if err != nil {
			return nil, err
		}
	}
	return oldUser, err
}

func (*userModel) AddUser(m *User) (id int64, err error) {
	if m.Id != 0 {
		user := &User{Id: m.Id}
		if err := Ormer().Read(user, "id"); err != nil {
			return 0, err
		}
		if m.Password == "" {
			m.Password = user.Password
		} else {
			m.Password = encode.EncodePassword(m.Password, GlobalUserSalt)
		}
		m.UserName = user.UserName
	} else {
		m.Password = encode.EncodePassword(m.Password, GlobalUserSalt)
	}
	id, err = Ormer().InsertOrUpdate(m)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (*userModel) DeleteById(m *User) error {
	_, err := Ormer().Delete(m, "id")
	if err != nil {
		return err
	}
	return nil
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