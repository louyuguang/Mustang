package models

type Project struct {
	Id     int64  `orm:"pk;auto" json:"id,omitempty"`
	Name   string `valid:"Required" orm:"index;unique;size(200);column(name);" json:"name,omitempty"`
	GitUrl string `valid:"Required" orm:"size(255);column(git_url);'" json:"gitUrl,omitempty"`
	EnvId  int64  `orm:"column(env_id);default(0)" json:"envId"`
}
