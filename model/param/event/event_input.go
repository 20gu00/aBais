package param

type ListEventInput struct {
	Name    string `form:"name"`
	Cluster string `form:"cluster"`
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
}
