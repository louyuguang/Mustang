package envclusterbinding

import (
	"Mustang/controllers/base"
	"Mustang/models"
	"Mustang/utils/logs"
	"encoding/json"
	"fmt"
	"strconv"
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
	env.Namespace = c.GetString("namespace")
	env.ClusterIds = strings.Join(r, ",")
	_, err := models.EnvClusterBindingModel.Add(&env)
	if err != nil {
		c.Fail(err)
		return
	}
	c.Success(fmt.Sprintf("环境%s创建成功！", env.EnvName))
}

// @router /delete [post]
func (c *EnvClusterBindingController) Delete() {
	var env *models.EnvClusterBinding
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &env); err != nil {
		c.Fail(err)
		return
	}
	if err := models.EnvClusterBindingModel.DeleteById(env); err != nil {
		c.Fail(err)
		return
	}
	c.Success("删除成功！")
}

// @router /update/:id [get,post]
func (c *EnvClusterBindingController) Update() {
	id := c.GetIDFromURL()
	//GET
	if c.Ctx.Input.Method() == "GET" {
		if id != 0 {
			env, err := models.EnvClusterBindingModel.GetById(id)
			if err != nil {
				logs.Error("get by id (%d) error.%v", id, err)
				c.Fail(err)
				return
			}
			clusters, err := models.ClusterModel.GetAllWithoutBinding()
			if err != nil {
				c.Fail(err)
				return
			}
			c.Data["Clusters"] = clusters
			c.Data["EnvAdd"] = env
		}
		c.TplName = "envclusterbindingcontroller/add.tpl"
		return
	}
	//POST
	var env models.EnvClusterBinding
	r := c.GetStrings("clusterIds")
	env.EnvName = c.GetString("envName")
	env.Namespace = c.GetString("namespace")
	env.ClusterIds = strings.Join(r, ",")
	env.Id = id
	if err := models.EnvClusterBindingModel.UpdateById(&env); err != nil {
		c.Fail(err)
		return
	}
	c.Success("环境信息更新成功！")
}

// @router /list [get]
func (c *EnvClusterBindingController) List() {
	type EnvAndCluster struct {
		Env      *models.EnvClusterBinding
		Clusters []*models.Cluster
	}
	var ec []EnvAndCluster
	scontent := c.GetString("scontent")
	pers := 10
	cnt, err := models.EnvClusterBindingModel.GetAllNum(scontent)
	if err != nil || cnt == -1 {
		logs.Error("get all cluster nums error")
		c.Fail(err)
		return
	}

	pager := c.SetPaginator(pers, cnt)
	envs, err := models.EnvClusterBindingModel.GetEnvs(pers, pager.Offset(), scontent)
	if err != nil {
		c.Fail(err)
		return
	}
	for _, v := range envs {
		var clusters []*models.Cluster
		for _, id := range strings.Split(v.ClusterIds, ",") {
			id, _ := strconv.Atoi(id)
			cluster, _ := models.ClusterModel.GetById(int64(id))
			clusters = append(clusters, cluster)
		}
		ec = append(ec, EnvAndCluster{v, clusters})
	}
	c.Data["Envs"] = ec
	c.Data["Scontent"] = scontent
}
