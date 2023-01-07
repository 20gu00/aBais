package cluster_role

type DeleteClusterRoleInput struct {
	ClusterRoleName string `json:"clusterrole_name"`
	//Namespace       string `json:"namespace"`
	Cluster string `json:"cluster"`
}

type GetClusterRoleInput struct {
	FilterName string `form:"filter_name"`
	//Namespace  string `form:"namespace"`
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
	Cluster string `form:"cluster"`
}

type UpdateClusterRoleInput struct {
	//Namespace string `json:"namespace"`
	Content string `json:"content"`
	Cluster string `json:"cluster"`
}

type GetClusterRoleDetailInput struct {
	ClusterRoleName string `form:"clusterrole_name"`
	//Namespace       string `form:"namespace"`
	Cluster string `form:"cluster"`
}
