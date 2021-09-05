package models

import "github.com/beego/beego/v2/adapter/validation"

type Cluster struct {
	Id          int64  `orm:"pk;auto" json:"id,omitempty"`
	ClusterName string `valid:"Required" orm:"index;unique;size(200);column(clustername);" json:"clustername,omitempty"`
	AliasName   string `orm:"size(255);column(aliasname)" json:"aliasname"`
	KubeConfig  string `valid:"Required" orm:"null;type(text)" json:"kubeconfig,omitempty"`
}

type clusterModel struct{}

func (*clusterModel) valid(c *Cluster) error {
	valid := validation.Validation{}
	b, err := valid.Valid(c)
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

func (c *clusterModel) AddCluster(m *Cluster) (id int64, err error) {
	if err := c.valid(m); err != nil {
		return 0, err
	}
	id, err = Ormer().InsertOrUpdate(m)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (*clusterModel) GetClusterById(id int64) (v *Cluster, err error) {
	v = &Cluster{Id: id}
	if err = Ormer().Read(v); err != nil {
		return nil, err
	}
	return v, nil
}

func (*clusterModel) GetAllNum(scontent ...string) (num int64, err error) {
	query := map[string]interface{}{}
	if scontent != nil {
		query["clustername__icontains"] = scontent
	}
	qs := Ormer().QueryTable(new(Cluster))
	qs = BuildFilter(qs, query)
	num, err = qs.Count()
	if err != nil {
		return -1, err
	}
	return num, nil
}

func (*clusterModel) GetClusters(pers int, offset int, scontent ...string) ([]*Cluster, error) {
	var clusters []*Cluster
	query := map[string]interface{}{}
	qs := Ormer().QueryTable(new(Cluster))
	if scontent != nil {
		query["clustername__icontains"] = scontent
	}
	qs = BuildFilter(qs, query)
	_, _ = qs.Limit(pers, offset).RelatedSel().All(&clusters)
	return clusters, nil
}

func (*clusterModel) DeleteById(m *Cluster) error {
	_, err := Ormer().Delete(m, "id")
	if err != nil {
		return err
	}
	return nil
}
