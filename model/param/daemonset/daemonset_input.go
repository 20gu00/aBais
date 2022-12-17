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
