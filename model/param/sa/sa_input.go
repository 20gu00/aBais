package param

type DeleteSaInput struct {
	SaName    string `json:"sa_name"`
	Namespace string `json:"namespace"`
	Cluster   string `json:"cluster"`
}

type GetSaInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type UpdateSaInput struct {
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}

type GetSaDetailInput struct {
	SaName    string `form:"sa_name"`
	Namespace string `form:"namespace"`
	Cluster   string `form:"cluster"`
}
