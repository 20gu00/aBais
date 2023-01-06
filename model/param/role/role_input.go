package param

type DeleteRoleInput struct {
	RoleName  string `json:"role_name"`
	Namespace string `json:"namespace"`
	Cluster   string `json:"cluster"`
}

type GetRoleInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type UpdateRoleInput struct {
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}

type GetRoleDetailInput struct {
	RoleName  string `form:"Role_name"`
	Namespace string `form:"namespace"`
	Cluster   string `form:"cluster"`
}
