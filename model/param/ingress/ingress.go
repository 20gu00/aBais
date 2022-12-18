package param

type GetIngressInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type CommonIngressDetailInput struct {
	IngressName string `form:"ingress_name"`
	Namespace   string `form:"namespace"`
	Cluster     string `form:"cluster"`
}

type UpdateIngressInput struct {
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}
