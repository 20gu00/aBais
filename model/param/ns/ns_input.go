package param

type GetNsInput struct {
	FilterName string `form:"filter_name"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type GetNsDetailInput struct {
	NamespaceName string `form:"namespace_name"`
	Cluster       string `form:"cluster"`
}

type DeleteNsInput struct {
	NamespaceName string `json:"namespace_name"`
	Cluster       string `json:"cluster"`
}
