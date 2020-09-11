package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

func BuildFilter(qs orm.QuerySeter, query map[string]interface{}) orm.QuerySeter {
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	return qs
}
