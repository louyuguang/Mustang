package deploy

import (
	"Mustang/controllers/base"
	"Mustang/models"
	"Mustang/pkg/machinery"
	"Mustang/utils/logs"
	"encoding/json"
	"fmt"
	"net/http"
)

type DeployController struct {
	base.BaseController
}

func (c *DeployController) URLMapping() {
	c.Mapping("Add", c.Add)
	//c.Mapping("Detail", c.Detail)
	c.Mapping("List", c.List)
}

func (c *DeployController) Prepare() {
	c.BaseController.Prepare()
	c.Data["LevelOne"] = map[string]string{"content": "部署管理"}
}

// @router /add [get,post]
func (c *DeployController) Add() {
	if c.Ctx.Input.Method() == "GET" {
		return
	}
	var deploy *models.Deploy
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &deploy); err != nil {
		logs.Error("Deploy's form error. %v", err)
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}
	userid := c.GetSession("userId").(int64)
	deploy.User = &models.User{Id: userid}
	_, err := models.DeployModel.Add(deploy)
	if err != nil {
		c.Fail(err)
		return
	}
	c.Success(fmt.Sprintf("部署%s创建成功！", deploy.ProjectName))
}

// @router /exec [post]
func (c *DeployController) Exec() {
	var deploy *models.Deploy
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &deploy); err != nil {
		logs.Error("Deploy's form error. %v", err)
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}
	if err := machinery.TaskDeploy(deploy); err != nil {
		logs.Error("Deploy's exec failed. %v", err)
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}
	c.Success(fmt.Sprintf("部署%s成功下发！", deploy.ProjectName))
}

// @router /list [get]
func (c *DeployController) List() {
	pers := 10
	cnt, err := models.DeployModel.GetAllNum()
	if err != nil || cnt == -1 {
		logs.Error("get all deploys nums error")
		c.Fail(err)
		return
	}

	pager := c.SetPaginator(pers, cnt)
	deploys, err := models.DeployModel.GetAll(pers, pager.Offset())
	if err != nil {
		c.Fail(err)
		return
	}
	c.Data["Deploys"] = deploys
}

// @router /delete [post]
func (c *DeployController) Delete() {
	var deploy *models.Deploy
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &deploy); err != nil {
		c.Fail(err)
		return
	}
	if err := models.DeployModel.DeleteById(deploy); err != nil {
		c.Fail(err)
		return
	}
	c.Success("删除成功！")
}
