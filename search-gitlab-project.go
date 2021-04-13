package main

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func main() {
	apiUrl := "http://100.80.250.204/api/v4"
	accessToken := "zobMrk5HLnYibjhsnUxe"
	git, err := gitlab.NewClient(accessToken, gitlab.WithBaseURL(apiUrl))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// opt := &gitlab.SearchOptions{
	// 	Scope:  gitlab.String("projects"),
	// 	Search: gitlab.String("Project Created From API"),
	// }
	opt := &gitlab.ListProjectsOptions{
		Search: gitlab.String("01EYKSNGX3F6ZZPTSWYKVR04JGgfghfghfghfh"),
	}
	projects, _, err := git.Projects.ListProjects(opt, nil)
	log.Println(projects)

}
