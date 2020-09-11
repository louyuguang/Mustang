package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"sync"
)

var (
	once      sync.Once
	globalOrm orm.Ormer
	UserModel *userModel
	RoleModel *roleModel
)

func init() {
	orm.RegisterModel(
		new(User),
		new(Role),
	)

	UserModel = &userModel{}
}

func Ormer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}

func CountObjects(qs orm.QuerySeter) (int64, error) {
	cnt, err := qs.Count()
	if err != nil {
		beego.Error("models.CountObjects ", err)
		return 0, err
	}
	return cnt, err
}