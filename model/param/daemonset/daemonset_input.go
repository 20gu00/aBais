package param

type GetDaemonsetInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type GetDaemonsetDetailInput struct {
	DaemonSetName string `form:"daemonset_name"`
	Namespace     string `form:"namespace"`
	Cluster       string `form:"cluster"`
}

type DeleteDaemonsetInput struct {
	DaemonSetName string `json:"daemonset_name"`
	Namespace     string `json:"namespace"`
	Cluster       string `json:"cluster"`
}

type UpdateDaemonsetInput struct {
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}
