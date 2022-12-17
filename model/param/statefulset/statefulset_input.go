package param

type GetStatefulsetInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type GetStatefulsetDetailInput struct {
	StatefulSetName string `form:"statefulset_name"`
	Namespace       string `form:"namespace"`
	Cluster         string `form:"cluster"`
}

type DeleteStatefulsetInput struct {
	StatefulSetName string `json:"statefulset_name"`
	Namespace       string `json:"namespace"`
	Cluster         string `json:"cluster"`
}

type UpdateStatefulsetInput struct {
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}
