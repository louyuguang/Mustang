package cluster

import (
	"Mustang/controllers/base"
	"Mustang/models"
	"Mustang/utils/logs"
	"encoding/json"
	"fmt"
	"net/http"
)

type ClusterController struct {
	base.BaseController
}

func (c *ClusterController) URLMapping() {
	c.Mapping("Add", c.Add)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Update", c.Update)
	c.Mapping("List", c.List)
}

func (c *ClusterController) Prepare() {
	c.BaseController.Prepare()
	c.Data["LevelOne"] = map[string]string{"content": "集群管理"}
}

// @router /add [get,post]
func (c *ClusterController) Add() {
	if c.Ctx.Input.Method() == "GET" {
		return
	}
	//POST, Add a new User
	var cluster *models.Cluster
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &cluster); err != nil {
		logs.Error("Cluster's form error. %v", err)
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}
	_, err := models.ClusterModel.Add(cluster)
	if err != nil {
		c.Fail(err)
		return
	}
	c.Success(fmt.Sprintf("集群%s创建成功！", cluster.ClusterName))
}

// @router /delete [post]
func (c *ClusterController) Delete() {
	var cluster *models.Cluster
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &cluster); err != nil {
		c.Fail(err)
		return
	}
	if err := models.ClusterModel.DeleteById(cluster); err != nil {
		c.Fail(err)
		return
	}
	c.Success("删除成功！")
}

// @router /update/:id [get,post]
func (c *ClusterController) Update() {
	id := c.GetIDFromURL()
	//GET
	if c.Ctx.Input.Method() == "GET" {
		if id != 0 {
			cluster, err := models.ClusterModel.GetById(id)
			if err != nil {
				logs.Error("get by id (%d) error.%v", id, err)
				c.Fail(err)
				return
			}
			c.Data["ClusterAdd"] = cluster
		}
		c.TplName = "clustercontroller/add.html"
		return
	}
	//POST
	var cluster *models.Cluster
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cluster)
	cluster.Id = id
	if err != nil {
		logs.Error("Cluster's form error. %v", err)
		c.Fail(err)
		return
	}
	err = models.ClusterModel.UpdateById(cluster)
	if err != nil {
		c.Fail(err)
		return
	}
	c.Success("集群信息更新成功！")
}

// @router /list [get]
func (c *ClusterController) List() {
	scontent := c.GetString("scontent")
	pers := 10
	cnt, err := models.ClusterModel.GetAllNum(scontent)
	if err != nil || cnt == -1 {
		logs.Error("get all cluster nums error")
		c.Fail(err)
		return
	}

	pager := c.SetPaginator(pers, cnt)
	clusters, err := models.ClusterModel.GetClusters(pers, pager.Offset(), scontent)
	if err != nil {
		c.Fail(err)
		return
	}
	c.Data["Clusters"] = clusters
	c.Data["Scontent"] = scontent
}
