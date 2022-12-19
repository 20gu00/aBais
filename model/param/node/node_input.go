package param

type GetNodeInput struct {
	FilterName string `form:"filter_name"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type GetNodeDetailInput struct {
	NodeName string `form:"node_name"`
	Cluster  string `form:"cluster"`
}
