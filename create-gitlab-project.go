package main

import (
	"io/ioutil"
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

	opt := &gitlab.CreateProjectOptions{
		Name:        gitlab.String("CreateThroughAPI"),
		Description: gitlab.String("Project Created From API"),
	}

	log.Println("Creating Project ...")
	project, _, err := git.Projects.CreateProject(opt, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Project Result: %v \n", project)

	log.Println("Creating Project Variables ...")
	varOpt := &gitlab.CreateProjectVariableOptions{
		Key:   gitlab.String("API_PROJECT_VARIABLES"),
		Value: gitlab.String("It works !!!!"),
	}
	projectVariable, _, err := git.ProjectVariables.CreateVariable(project.ID, varOpt, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Project Variables Result: %v \n", projectVariable)

	content, err := ioutil.ReadFile(".gitlab-ci.yml")
	if err != nil {
		log.Fatal(err)
	}

	createFileOpt := &gitlab.CreateFileOptions{
		Branch:        gitlab.String("master"),
		Content:       gitlab.String(string(content)),
		CommitMessage: gitlab.String("Initial Commit"),
	}

	log.Println("Uploading .gitlab-ci.yml file ...")
	fileInfo, _, err := git.RepositoryFiles.CreateFile(4, ".gitlab-ci.yml", createFileOpt)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Project File Result: %v \n", fileInfo)

}
