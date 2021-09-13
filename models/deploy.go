package models

import (
	"time"
)

type Deploy struct {
	Id                int64              `orm:"pk;auto" json:"id,omitempty"`
	User              *User              `orm:"rel(fk);column(user_id);" json:"user"`
	ProjectName       string             `orm:"size(255);column(project_name)" json:"projectName"`
	GitUrl            string             `orm:"size(255);column(git_url)" json:"gitUrl"`
	Image             string             `orm:"size(255);column(image)" json:"image"`
	Port              int64              `orm:"size(255);column(port)" json:"port,string"`
	Created           *time.Time         `orm:"auto_now_add;type(datetime)" json:"createTime,omitempty"`
	EnvClusterBinding *EnvClusterBinding `valid:"Required" orm:"rel(fk);default(0);column(env_id)" json:"envId"`
}

type deployModel struct{}

func (*deployModel) GetAllNum() (num int64, err error) {
	qs := Ormer().QueryTable(new(Deploy))
	num, err = qs.Count()
	if err != nil {
		return -1, err
	}
	return num, nil
}

func (*deployModel) GetAll(pers int, offset int) ([]*Deploy, error) {
	var deploys []*Deploy
	qs := Ormer().QueryTable(new(Deploy))
	_, _ = qs.Limit(pers, offset).RelatedSel().All(&deploys)
	return deploys, nil
}

func (*deployModel) GetById(id int64) (d *Deploy, err error) {
	d = &Deploy{Id: id}
	if err = Ormer().Read(d); err != nil {
		return nil, err
	}
	if _, err = Ormer().LoadRelated(d, "EnvClusterBinding"); err != nil {
		return nil, err
	}
	return d, err
}

func (*deployModel) Add(m *Deploy) (id int64, err error) {
	id, err = Ormer().Insert(m)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (*deployModel) DeleteById(m *Deploy) error {
	_, err := Ormer().Delete(m, "id")
	if err != nil {
		return err
	}
	return nil
}

func (*deployModel) UpdateById(m *Deploy) (err error) {
	d := &Deploy{Id: m.Id}
	if err = Ormer().Read(d); err == nil {
		_, err = Ormer().Update(m)
		return
	}
	return
}
