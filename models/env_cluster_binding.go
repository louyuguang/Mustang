package models

import "Mustang/utils/logs"

type EnvClusterBinding struct {
	Id        int64      `orm:"pk;auto" json:"id,omitempty"`
	Namespace string     `valid:"Required" orm:"size(255);column(namespace)" json:"namespace"`
	EnvName   string     `valid:"Required" orm:"size(255);column(envname)" json:"envName"`
	Deploy    []*Deploy  `orm:"reverse(many)" json:"deployIds"`
	Clusters  []*Cluster `orm:"rel(m2m);rel_table(env_cluster)" json:"clusterIds" `
}

type envModel struct{}

func (*envModel) Add(m *EnvClusterBinding) (id int64, err error) {
	if err := valid(m); err != nil {
		return 0, err
	}
	id, err = Ormer().Insert(m)
	m2m := Ormer().QueryM2M(m, "Clusters")
	for _, cluster := range m.Clusters {
		num, err := m2m.Add(cluster)
		if err != nil {
			return 0, err
		}
		logs.Info("env, cluster binding num: %d", num)
	}
	return
}

func (c *envModel) UpdateById(m *EnvClusterBinding) (err error) {
	if err := valid(m); err != nil {
		return err
	}
	v := &EnvClusterBinding{Id: m.Id}
	if err = Ormer().Read(v, "Id"); err == nil {
		_, err = Ormer().Update(m)
		m2m := Ormer().QueryM2M(m, "Clusters")
		num, err := m2m.Clear()
		if err != nil {
			return err
		}
		logs.Info("env, cluster binding delete num: %d", num)
		for _, cluster := range m.Clusters {
			num, err := m2m.Add(cluster)
			if err != nil {
				return err
			}
			logs.Info("env, cluster binding num: %d", num)
		}
	}
	return
}

func (*envModel) GetById(id int64) (v *EnvClusterBinding, err error) {
	v = &EnvClusterBinding{Id: id}
	if err = Ormer().Read(v); err != nil {
		return
	}
	_, err = Ormer().LoadRelated(v, "Clusters")
	if err != nil {
		return
	}
	return
}

func (*envModel) GetAllNum(scontent ...string) (num int64, err error) {
	query := map[string]interface{}{}
	if scontent != nil {
		query["envName__icontains"] = scontent
	}
	qs := Ormer().QueryTable(new(EnvClusterBinding))
	qs = BuildFilter(qs, query)
	num, err = qs.Count()
	if err != nil {
		return -1, err
	}
	return num, nil
}

func (*envModel) GetEnvs(pers int, offset int, scontent ...string) ([]*EnvClusterBinding, error) {
	var envs []*EnvClusterBinding
	query := map[string]interface{}{}
	qs := Ormer().QueryTable(new(EnvClusterBinding))
	if scontent != nil {
		query["envName__icontains"] = scontent
	}
	qs = BuildFilter(qs, query)
	_, _ = qs.Limit(pers, offset).RelatedSel().All(&envs)
	for i, env := range envs {
		Ormer().LoadRelated(env, "Clusters")
		envs[i] = env
	}
	return envs, nil
}

func (*envModel) DeleteById(m *EnvClusterBinding) error {
	_, err := Ormer().Delete(m, "id")
	if err != nil {
		return err
	}
	return nil
}
