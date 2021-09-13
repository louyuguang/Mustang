package models

import (
	"sync"

	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/adapter/validation"
)

var (
	once                   sync.Once
	globalOrm              orm.Ormer
	UserModel              *userModel
	RoleModel              *roleModel
	ClusterModel           *clusterModel
	DeployModel            *deployModel
	EnvClusterBindingModel *envModel
)

func init() {
	orm.RegisterModel(
		new(User),
		new(Role),
		new(Cluster),
		new(Deploy),
		new(EnvClusterBinding),
	)

	UserModel = &userModel{}
	RoleModel = &roleModel{}
	ClusterModel = &clusterModel{}
	DeployModel = &deployModel{}
	EnvClusterBindingModel = &envModel{}
}

func Ormer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}

func valid(m interface{}) error {
	valid := validation.Validation{}
	b, err := valid.Valid(m)
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
