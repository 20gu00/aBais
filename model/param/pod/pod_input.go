package param

type GetPodsInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type GetPodDetailInput struct {
	PodName   string `form:"pod_name"`
	Namespace string `form:"namespace"`
	Cluster   string `form:"cluster"`
}
