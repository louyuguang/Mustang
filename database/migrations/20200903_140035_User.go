package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20200903_140035 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20200903_140035{}
	m.Created = "20200903_140035"

	migration.Register("User_20200903_140035", m)
}

// Run the migrations
func (m *User_20200903_140035) Up() {
	m.SQL("CREATE TABLE IF NOT EXISTS `usercontroller` (`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,`username` varchar(255) NOT NULL DEFAULT '' ,`password` varchar(255) NOT NULL DEFAULT '' ,`email` varchar(255) NOT NULL DEFAULT '') ENGINE=InnoDB")

}

// Reverse the migrations
func (m *User_20200903_140035) Down() {
	m.SQL("DROP TABLE usercontroller")

}
