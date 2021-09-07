package models

type EnvClusterBinding struct {
	Id         int64     `orm:"pk;auto" json:"id,omitempty"`
	ClusterIds string    `orm:"size(255);column(clusterids)" json:"clusterIds"`
	Namespace  string    `valid:"Required" orm:"size(255);column(namespace)" json:"namespace"`
	EnvName    string    `valid:"Required" orm:"size(255);column(envname)" json:"envName"`
	Deploy     []*Deploy `orm:"reverse(many)" json:"env_id,omitempty"`
}

type envModel struct{}

func (*envModel) Add(m *EnvClusterBinding) (id int64, err error) {
	if err := valid(m); err != nil {
		return 0, err
	}
	id, err = Ormer().Insert(m)
	//if clusters, err := ClusterModel.GetAllWithoutBinding(); err != nil {
	//	return 0, err
	//}
	return
}

//func (c *clusterModel) UpdateById(m *Cluster) (err error) {
//	if err := valid(m); err != nil {
//		return err
//	}
//	v := Cluster{Id: m.Id}
//	if err = Ormer().Read(&v, "Id"); err == nil {
//		_, err = Ormer().Update(m)
//		return err
//	}
//	return
//}
//
//func (*clusterModel) GetById(id int64) (v *Cluster, err error) {
//	v = &Cluster{Id: id}
//	if err = Ormer().Read(v); err != nil {
//		return nil, err
//	}
//	return v, nil
//}
//
//func (*clusterModel) GetAllNum(scontent ...string) (num int64, err error) {
//	query := map[string]interface{}{}
//	if scontent != nil {
//		query["clustername__icontains"] = scontent
//	}
//	qs := Ormer().QueryTable(new(Cluster))
//	qs = BuildFilter(qs, query)
//	num, err = qs.Count()
//	if err != nil {
//		return -1, err
//	}
//	return num, nil
//}
//
//func (*clusterModel) GetClusters(pers int, offset int, scontent ...string) ([]*Cluster, error) {
//	var clusters []*Cluster
//	query := map[string]interface{}{}
//	qs := Ormer().QueryTable(new(Cluster))
//	if scontent != nil {
//		query["clustername__icontains"] = scontent
//	}
//	qs = BuildFilter(qs, query)
//	_, _ = qs.Limit(pers, offset).RelatedSel().All(&clusters)
//	return clusters, nil
//}
//
//func (*clusterModel) DeleteById(m *Cluster) error {
//	_, err := Ormer().Delete(m, "id")
//	if err != nil {
//		return err
//	}
//	return nil
//}