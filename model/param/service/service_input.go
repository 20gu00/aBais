package param

type GetServiceInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type GetServiceDetailInput struct {
	ServiceName string `form:"service_name"`
	Namespace   string `form:"namespace"`
	Cluster     string `form:"cluster"`
}

type DeleteServiceInput struct {
	ServiceName string `json:"service_name"`
	Namespace   string `json:"namespace"`
	Cluster     string `json:"cluster"`
}

type UpdateServiceInput struct {
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}
