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

	kubeID, _, _ := git.InstanceVariables.GetVariable("INSTANCE_VARIABLE_TEST2", nil)
	// kubeID, response, err := git.InstanceVariables.ListVariables(nil, nil)
	// if err != nil {
	// 	log.Printf("Failed to retrieve variables: %v", err)
	// }
	if kubeID == nil {
		log.Printf("Failed to retrieve variables: %v", kubeID)
	}
	log.Println(kubeID)

	// log.Println(&response)

}
