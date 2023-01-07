package param

type GetPvInput struct {
	FilterName string `form:"filter_name"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type GetPvDetailInput struct {
	PvName  string `form:"pv_name"`
	Cluster string `form:"cluster"`
}

type DeletePvInput struct {
	PvName  string `json:"pv_name"`
	Cluster string `json:"cluster"`
}

type UpdatePvInput struct {
	//Namespace string `json:"namespace"`
	Content string `json:"content"`
	Cluster string `json:"cluster"`
}
