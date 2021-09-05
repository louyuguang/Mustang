package models

import (
	"Mustang/utils/encode"
	"errors"
	"time"

	"github.com/beego/beego/v2/adapter/validation"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64      `orm:"pk;auto" json:"id,omitempty"`
	Role     *Role      `valid:"Required" orm:"rel(fk);default(4);column(role_id)" json:"role"`
	UserName string     `valid:"Required" orm:"index;unique;size(200);column(username);" json:"username,omitempty"`
	Password string     `orm:"size(255)" json:"password"`
	RealName string     `valid:"Required" orm:"size(255)" json:"realname,omitempty"`
	Email    string     `valid:"Required; Email" orm:"size(200)" json:"email,omitempty"`
	Active   bool       `valid:"Required" orm:"default(true)" json:"is_active"`
	Created  *time.Time `orm:"auto_now_add;type(datetime)" json:"createTime,omitempty"`
	Deploy   []*Deploy  `orm:"reverse(many)" json:"user_id,omitempty"`
}

type userModel struct{}

var GlobalUserSalt = beego.AppConfig.String("GlobalUserSalt")

func (*userModel) valid(u *User) error {
	valid := validation.Validation{}
	b, err := valid.Valid(u)
	if err != nil {
		return err
	}
	if !b {
		for _, err := range valid.Errors {
			return err
		}
	}
	return nil
}

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

func (u *userModel) AddUser(m *User) (id int64, err error) {
	m.Password = encode.EncodePassword(m.Password, GlobalUserSalt)
	if err := u.valid(m); err != nil {
		return 0, err
	}
	if m.Role.Id == 0 {
		return 0, errors.New("role id 不能为 0")
	}
	id, err = Ormer().Insert(m)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userModel) UpdateUserById(m *User) (err error) {
	v := &User{Id: m.Id}
	if err = Ormer().Read(v); err != nil {
		return
	}
	if m.Password != "" {
		v.Password = encode.EncodePassword(m.Password, GlobalUserSalt)
	}
	v.RealName = m.RealName
	v.Email = m.Email
	v.Role = m.Role
	v.Active = m.Active
	if err := u.valid(m); err != nil {
		return err
	}
	if m.Role.Id == 0 {
		return errors.New("role id 不能为 0")
	}
	_, err = Ormer().Update(v)
	return
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
	_, err = Ormer().LoadRelated(user, "Role")
	if err != nil {
		return nil, err
	}
	return user, nil
}
