package param

type GetDeploymentInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type ScaleDeployment struct {
	DeploymentName string `json:"deployment_name"`
	Namespace      string `json:"namespace"`
	ScaleNum       int    `json:"scale_num"`
	Cluster        string `json:"cluster"`
}
