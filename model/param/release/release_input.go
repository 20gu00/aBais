package param

type ListReleaseInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Cluster    string `form:"cluster"`
}

type DetailReleaseInput struct {
	Release   string `form:"release"`
	Namespace string `form:"namespace"`
	Cluster   string `form:"cluster"`
}

type InstallReleaseInput struct {
	Release   string `json:"release"`
	Chart     string `json:"chart"`
	Namespace string `json:"namespace"`
	Cluster   string `json:"cluster"`
}

type UninstallReleaseInput struct {
	Release   string `json:"release"`
	Namespace string `json:"namespace"`
	Cluster   string `json:"cluster"`
}

type DeleteChartFileInput struct {
	Chart string `json:"chart"`
}

type ListChartInput struct {
	Name  string `form:"name"`
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
}
