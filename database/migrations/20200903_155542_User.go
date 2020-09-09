package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20200903_155542 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20200903_155542{}
	m.Created = "20200903_155542"

	migration.Register("User_20200903_155542", m)
}

// Run the migrations
func (m *User_20200903_155542) Up() {
	m.SQL("alter table user add created datetime default now() null;")
	m.SQL("alter table user add role bool default false null;")
	m.SQL("alter table user modify created datetime default CURRENT_TIMESTAMP not null;")
	m.SQL("alter table user modify role int default 5 not null;")
	m.SQL("alter table user add active bool default true not null;")
}

// Reverse the migrations
func (m *User_20200903_155542) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
