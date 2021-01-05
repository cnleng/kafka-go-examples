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

	log.Println("Getting Pipeline ...")
	pipeline, _, err := git.Pipelines.GetPipeline(2, 55, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Pipeline Result: %v \n", pipeline)
}
