package param

type ListReleaseInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Cluster    string `form:"cluster"`
}
