package param

type DeleteClusterRoleBindingInput struct {
	CLusterRoleBindingName string `json:"clusterrolebinding_name"`
	Cluster                string `json:"cluster"`
}

type GetClusterRoleBindingInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type UpdateClusterRoleBindingInput struct {
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}

type GetClusterRoleBindingDetailInput struct {
	ClusterRoleBindingName string `form:"clusterroleBinding_name"`
	Namespace              string `form:"namespace"`
	Cluster                string `form:"cluster"`
}
