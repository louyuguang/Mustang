package models

type Role struct {
	Id       int    `orm:"pk;auto" json:"id,omitempty"`
	Rolename string `orm:"unique;size(200)" json:"rolename,omitempty"`
	User     []*User  `orm:"reverse(many)" json:"role_id, omitempty"`
}
//var Roles = []Role{{1, "超级管理员"}, {2, "系统管理员"}, {3, "数据库管理员"}, {4, "研发工程师"}}

type roleModel struct{}

func (*roleModel) GetAllRoles() ([]*Role, error) {
	var roles []*Role
	Ormer().QueryTable(new(Role)).RelatedSel().All(&roles)
	return roles, nil
}

func (*roleModel) GetRole(user *userModel, id int) (*Role, error) {
	role := &Role{Id: id}
	if err := Ormer().Read(&role); err != nil{
		return nil, err
	}
	return role, nil
}