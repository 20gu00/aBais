package param

type DeleteJobInput struct {
	JobName   string `json:"job_name"`
	Namespace string `json:"namespace"`
	Cluster   string `json:"cluster"`
}

type GetJobInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type UpdateJobInput struct {
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}

type GetJobDetailInput struct {
	JobName   string `form:"job_name"`
	Namespace string `form:"namespace"`
	Cluster   string `form:"cluster"`
}
