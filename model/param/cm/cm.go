package param

type DeleteCmInput struct {
	ConfigMapName string `json:"configmap_name"`
	Namespace     string `json:"namespace"`
	Cluster       string `json:"cluster"`
}

type GetCmInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type GetCmDetailInput struct {
	ConfigMapName string `form:"configmap_name"`
	Namespace     string `form:"namespace"`
	Cluster       string `form:"cluster"`
}

type UpdateCmInput struct {
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}
