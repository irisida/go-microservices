package github

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Mesage   string `json:"message"`
}

type GithubErrorResponse struct {
	StatusCode       int    `json:"status_code"`
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
	Errors           []GithubError
}
