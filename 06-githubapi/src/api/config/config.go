package config

import "os"

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

// GetGithubAccessToken - public accessor for our token
func GetGithubAccessToken() string {
	return githubAccessToken
}
