package param

type GetPodsInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type GetPodDetailInput struct {
	PodName   string `form:"pod_name"`
	Namespace string `form:"namespace"`
	Cluster   string `form:"cluster"`
}

type DeletePodInput struct {
	PodName   string `json:"pod_name"`
	Namespace string `json:"namespace"`
	Cluster   string `json:"cluster"`
}

type UpdatePodInput struct {
	PodName   string `json:"pod_name"`
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}

type GetPodContainers struct {
	PodName   string `form:"pod_name"`
	Namespace string `form:"namespace"`
	Cluster   string `form:"cluster"`
}

type GetPodLog struct {
	ContainerName string `form:"container_name"`
	PodName       string `form:"pod_name"`
	Namespace     string `form:"namespace"`
	Cluster       string `form:"cluster"`
}

type GetPodNumPerNamespace struct {
	Cluster string `form:"cluster"`
}
