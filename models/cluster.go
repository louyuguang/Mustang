package models

type Cluster struct {
	Id                int64              `orm:"pk;auto" json:"id,omitempty"`
	ClusterName       string             `valid:"Required" orm:"index;unique;size(200);column(clustername);" json:"clustername,omitempty"`
	AliasName         string             `valid:"Required" orm:"size(255);column(aliasname)" json:"aliasname"`
	KubeConfig        string             `valid:"Required" orm:"null;type(text)" json:"kubeconfig,omitempty"`
	EnvClusterBinding *EnvClusterBinding `orm:"rel(fk);column(env_id);default(0)" json:"envClusterBinding"`
}

type clusterModel struct{}

func (c *clusterModel) Add(m *Cluster) (id int64, err error) {
	if err := valid(m); err != nil {
		return 0, err
	}
	id, err = Ormer().Insert(m)
	return
}

func (c *clusterModel) UpdateById(m *Cluster) (err error) {
	if err := valid(m); err != nil {
		return err
	}
	v := Cluster{Id: m.Id}
	if err = Ormer().Read(&v, "Id"); err == nil {
		_, err = Ormer().Update(m)
		return err
	}
	return
}

func (*clusterModel) GetById(id int64) (v *Cluster, err error) {
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

func (*clusterModel) GetAllWithoutBinding() ([]*Cluster, error) {
	var clusters []*Cluster
	qs := Ormer().QueryTable(new(Cluster))
	_, err := qs.Filter("env_id", 0).All(&clusters)
	if err != nil {
		return nil, err
	}
	return clusters, nil
}

func (*clusterModel) DeleteById(m *Cluster) error {
	_, err := Ormer().Delete(m, "id")
	if err != nil {
		return err
	}
	return nil
}
