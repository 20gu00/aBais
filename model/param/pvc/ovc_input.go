package param

type GetPvcInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type GetPvcDetailInput struct {
	PvcName   string `form:"pvc_name"`
	Namespace string `form:"namespace"`
	Cluster   string `form:"cluster"`
}
type UpdatePvcInput struct {
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}

type DeletePvcInput struct {
	PvcName   string `json:"pvc_name"`
	Namespace string `json:"namespace"`
	Cluster   string `json:"cluster"`
}
