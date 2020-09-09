package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Role_20200907_145226 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Role_20200907_145226{}
	m.Created = "20200907_145226"

	migration.Register("Role_20200907_145226", m)
}

// Run the migrations
func (m *Role_20200907_145226) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Role_20200907_145226) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
