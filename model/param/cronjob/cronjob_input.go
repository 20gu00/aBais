package param

type DeleteCronJobInput struct {
	CronJobName string `json:"cronjob_name"`
	Namespace   string `json:"namespace"`
	Cluster     string `json:"cluster"`
}

type GetCronJobInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type UpdateCronJobInput struct {
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}

type GetCronJobDetailInput struct {
	CronJobName string `form:"cronjob_name"`
	Namespace   string `form:"namespace"`
	Cluster     string `form:"cluster"`
}
