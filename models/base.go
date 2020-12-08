package models

import (
	"github.com/astaxie/beego/orm"
	"sync"
)

var (
	once      sync.Once
	globalOrm orm.Ormer
	UserModel *userModel
	RoleModel *roleModel
	ClusterModel *clusterModel
)

func init() {
	orm.RegisterModel(
		new(User),
		new(Role),
		new(Cluster),
	)

	UserModel = &userModel{}
	RoleModel = &roleModel{}
	ClusterModel = &clusterModel{}
}

func Ormer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}