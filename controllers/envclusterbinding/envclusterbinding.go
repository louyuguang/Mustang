package envclusterbinding

import (
	"Mustang/controllers/base"
	"Mustang/models"
	"Mustang/utils/logs"
	"encoding/json"
	"fmt"
	"strconv"
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
		clusterAll := make(map[int64]string)
		clusters, err := models.ClusterModel.GetAll()
		if err != nil {
			c.Fail(err)
			return
		}
		for _, cluster := range clusters {
			clusterAll[cluster.Id] = cluster.ClusterName
		}
		c.Data["ClusterAll"] = clusterAll
		return
	}

	env := &models.EnvClusterBinding{
		EnvName:   c.GetString("envName"),
		Namespace: c.GetString("namespace"),
	}
	r := c.GetStrings("clusterIds")
	for _, clusterId := range r {
		id, _ := strconv.Atoi(clusterId)
		cluster, err := models.ClusterModel.GetById(int64(id))
		if err != nil {
			c.Fail(err)
			return
		}
		env.Clusters = append(env.Clusters, cluster)
	}
	_, err := models.EnvClusterBindingModel.Add(env)
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
			clusterAll, clusterSelect := make(map[int64]string), make(map[int64]string)
			env, err := models.EnvClusterBindingModel.GetById(id)
			if err != nil {
				logs.Error("get by id (%d) error.%v", id, err)
				c.Fail(err)
				return
			}
			clusters, err := models.ClusterModel.GetAll()
			if err != nil {
				c.Fail(err)
				return
			}
			for _, cluster := range clusters {
				clusterAll[cluster.Id] = cluster.ClusterName
			}
			for _, envCluster := range env.Clusters {
				clusterSelect[envCluster.Id] = envCluster.ClusterName
			}
			c.Data["ClusterAll"] = clusterAll
			c.Data["ClusterSelect"] = clusterSelect
			c.Data["EnvAdd"] = env
		}
		c.TplName = "envclusterbindingcontroller/add.html"
		return
	}
	//POST
	env := &models.EnvClusterBinding{
		Id:        id,
		EnvName:   c.GetString("envName"),
		Namespace: c.GetString("namespace"),
	}
	r := c.GetStrings("clusterIds")
	for _, clusterId := range r {
		id, _ := strconv.Atoi(clusterId)
		cluster, err := models.ClusterModel.GetById(int64(id))
		if err != nil {
			c.Fail(err)
			return
		}
		env.Clusters = append(env.Clusters, cluster)
	}
	if err := models.EnvClusterBindingModel.UpdateById(env); err != nil {
		c.Fail(err)
		return
	}
	c.Success("环境信息更新成功！")
}

// @router /list [get]
func (c *EnvClusterBindingController) List() {
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
	c.Data["Envs"] = envs
	c.Data["Scontent"] = scontent
}
