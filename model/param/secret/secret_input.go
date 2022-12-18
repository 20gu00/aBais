package param

type GetSecretInput struct {
	FilterName string `form:"filter_name"`
	Namespace  string `form:"namespace"`
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Cluster    string `form:"cluster"`
}

type GetSecretDetailInput struct {
	SecretName string `form:"secret_name"`
	Namespace  string `form:"namespace"`
	Cluster    string `form:"cluster"`
}

type DeleteSecretInput struct {
	SecretName string `json:"secret_name"`
	Namespace  string `json:"namespace"`
	Cluster    string `json:"cluster"`
}

type UpdateSecretInput struct {
	Namespace string `json:"namespace"`
	Content   string `json:"content"`
	Cluster   string `json:"cluster"`
}
