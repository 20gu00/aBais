package param

type DeleteRoleBindingInput struct {
	RoleBindingName string `json:"rolebinding_name"`
	Namespace       string `json:"namespace"`
	Cluster         string `json:"cluster"`
}

type GetRoleBindingInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type UpdateRoleBindingInput struct {
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}

type GetRoleBindingDetailInput struct {
	RoleBindingName string `form:"roleBinding_name"`
	Namespace       string `form:"namespace"`
	Cluster         string `form:"cluster"`
}
