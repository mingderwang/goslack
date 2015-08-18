//go:generate ginger $GOFILE
package main

// @ginger
type GithubUser struct {
	Created int32  `json:"created,omitempty"`
	Id      int32  `json:"id,omitempty"`
	Login   string `json:"login"`
	Plan    struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		Collaborators int    `json:"collaborators"`
		PrivateRepos  int    `json:"private_repos"`
	} `json:"plan"`
}
