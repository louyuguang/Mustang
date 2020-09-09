package models

const (
	//资源类型
	PermissionUser    = "User"
	PermissionCluster = "Cluster"
	PermissionEnv     = "Env"
	PermissionProject = "Project"
	PermissionConfig  = "Config"
	PermissionDeploy  = "Deploy"
	//动作
	PermissionCreate = "Add"
	PermissionUpdate = "Update"
	PermissionRead   = "Get"
	PermissionDelete = "Delete"
)




func init() {
	Permission := make(map[int][]string)
	Permission[1] =	[]string{PermissionCluster, PermissionEnv, PermissionProject, PermissionConfig, PermissionDeploy}
}
