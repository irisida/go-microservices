package repositories

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateRepoResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
