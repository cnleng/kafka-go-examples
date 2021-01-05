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

	variables := map[string]string{
		"KUBE_ACCOUNT":    "kubemaster01",
		"DATA_MOUNT_PATH": "data",
		"NFS_STORAGE_PVC": "nfs-storage-pvc",
	}

	projectID := 2

	log.Println("Checking trigger ...")
	triggers, _, err := git.PipelineTriggers.ListPipelineTriggers(projectID, nil)
	if err != nil {
		log.Fatal(err)
	}

	token := triggers[0].Token
	opt := &gitlab.RunPipelineTriggerOptions{
		Ref:       gitlab.String("master"),
		Token:     gitlab.String(token),
		Variables: variables,
	}

	log.Println("Run trigger ...")
	pipeline, _, err := git.PipelineTriggers.RunPipelineTrigger(2, opt)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Pipeline Result: %v", pipeline)
}
