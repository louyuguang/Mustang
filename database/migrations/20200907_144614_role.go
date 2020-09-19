package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Role_20200907_144614 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Role_20200907_144614{}
	m.Created = "20200907_144614"

	migration.Register("Role_20200907_144614", m)
}

// Run the migrations
func (m *Role_20200907_144614) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("create table role ( id int null, role varchar(255) null )");
	m.SQL("alter table usercontroller change role role_id int default 5 not null;");
	m.SQL("alter table usercontroller add constraint user___fk_1 foreign key (role_id) references role (id);");
	m.SQL("alter table usercontroller alter column role_id set default 3;");
}

// Reverse the migrations
func (m *Role_20200907_144614) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
