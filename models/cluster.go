package models

type Cluster struct {
	Id          int64  `orm:"pk;auto" json:"id,omitempty"`
	ClusterName string `orm:"index;unique;size(200);column(clustername);" json:"clustername,omitempty"`
	AliasName   string `orm:"size(255);column(aliasname)" json:"aliasname"`
	Config      string `orm:"size(10240);column(config)" json:"config"`
}

type clusterModel struct{}
