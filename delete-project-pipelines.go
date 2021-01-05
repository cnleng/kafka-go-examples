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

	projectID := 2

	pipelines, _, err := git.Pipelines.ListProjectPipelines(projectID, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, pipeline := range pipelines {
		log.Printf("Deleting Pipeline : %d \n", pipeline.ID)
		response, err := git.Pipelines.DeletePipeline(projectID, pipeline.ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Pipeline Delete Result: %v \n", response)
	}

}
