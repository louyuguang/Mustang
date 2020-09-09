package models

type Role struct {
	Id       int    `orm:"pk;auto" json:"id,omitempty"`
	Rolename string `orm:"unique;size(200)" json:"rolename,omitempty"`
	User     []*User  `orm:"reverse(many)" json:"role_id, omitempty"`
}

//var Roles = []Role{{1, "超级管理员"}, {2, "系统管理员"}, {3, "数据库管理员"}, {4, "研发工程师"}}
