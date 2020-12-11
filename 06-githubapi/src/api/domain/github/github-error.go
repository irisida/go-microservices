package github

// GithubError struct
type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Mesage   string `json:"message"`
}

// GithubErrorResponse - struct
type GithubErrorResponse struct {
	StatusCode       int    `json:"status_code"`
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
	Errors           []GithubError
}
