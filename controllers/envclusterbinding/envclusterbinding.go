package envclusterbinding

import (
	"Mustang/controllers/base"
	"Mustang/models"
	"fmt"
	"strings"
)

type EnvClusterBindingController struct {
	base.BaseController
}

//func (c *EnvClusterBindingController) URLMapping() {
//	c.Mapping("Add", c.Add)
//}

func (c *EnvClusterBindingController) Prepare() {
	c.BaseController.Prepare()
	c.Data["LevelOne"] = map[string]string{"content": "环境管理"}
}

// @router /add [get,post]
func (c *EnvClusterBindingController) Add() {
	if c.Ctx.Input.Method() == "GET" {
		clusters, err := models.ClusterModel.GetAllWithoutBinding()
		if err != nil {
			c.Fail(err)
			return
		}
		c.Data["Clusters"] = clusters
		return
	}

	var env models.EnvClusterBinding
	r := c.GetStrings("clusterIds")
	env.EnvName = c.GetString("envName")
	env.EnvName = c.GetString("envName")
	env.Namespace = c.GetString("namespace")
	env.ClusterIds = strings.Join(r, ",")
	_, err := models.EnvClusterBindingModel.Add(&env)
	if err != nil {
		c.Fail(err)
		return
	}
	c.Success(fmt.Sprintf("环境%s创建成功！", env.EnvName))
}
