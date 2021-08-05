package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/xanzy/go-gitlab"
)

// listClusters
func listClusters() {
	apiUrl := "http://100.80.250.204/api/v4"
	accessToken := "zobMrk5HLnYibjhsnUxe"
	git, err := gitlab.NewClient(accessToken, gitlab.WithBaseURL(apiUrl))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	clusters, _, err := git.InstanceCluster.ListClusters(nil)
	if err != nil {
		log.Println(err)
	}

	for index := range clusters {
		fmt.Println(clusters[index])
	}
}

// addCluster
func addCluster() {
	apiUrl := "http://100.80.250.207/api/v4"
	accessToken := "CYStz23c3Eue9zRdyak8"
	git, err := gitlab.NewClient(accessToken, gitlab.WithBaseURL(apiUrl))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	caCert, _ := ioutil.ReadFile("caCert.pem")
	token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJnaXRsYWItdG9rZW4tcWJ4bGwiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZ2l0bGFiIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiMDFlMDM5OTUtMjQwNS00YjQyLTk1NzktYTQwNThmZTc5NDgzIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Omt1YmUtc3lzdGVtOmdpdGxhYiJ9.DXca9pDtCsMBdOa8mSwnDS0vxj9lNIN-bAFNbE1qc-1QHYX9HTIycuuIsJvdoRYWS74IOYtDafL-F-jZgJ4go4xbqUPuXQ7GbIMo7FaFhrBvddlk7vZL_6WHnis-Q8GaZlEco9zr4IVXAC4qIccdRlobE-ssMGHsaMYoYTton3qr_rBPfC5M_p8YMAzXfDfxX1V8XeaBxDDXy8XAW45A9XBekS4QM3qNyotN8oHhnUGWfxbnGC0CGdX7p_3jHneFFxQytPG-2sqAnaP3PlotkjcWNTcfiQukS6I5yIEkrzDYD_-PUk07dLM0GPa9lCCcHCi95b6m7EPKdownxqX0eQ"

	platformKubernetesOpt := &gitlab.AddPlatformKubernetesOptions{
		APIURL: gitlab.String("https://100.80.250.202:6443"),
		Token:  gitlab.String(token),
		CaCert: gitlab.String(string(caCert)),
		// AuthorizationType: gitlab.String("rbac"),
	}

	clusterOpt := &gitlab.AddClusterOptions{
		Name:               gitlab.String("Cluster II"),
		Enabled:            gitlab.Bool(true),
		PlatformKubernetes: platformKubernetesOpt,
	}

	clusters, _, err := git.InstanceCluster.AddCluster(clusterOpt, nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(clusters)
}

// deleteCluster
func deleteCluster() {
	apiUrl := "http://100.80.250.204/api/v4"
	accessToken := "zobMrk5HLnYibjhsnUxe"
	git, err := gitlab.NewClient(accessToken, gitlab.WithBaseURL(apiUrl))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	_, err = git.InstanceCluster.DeleteCluster(1, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Cluster deleted")
}

// main
func main() {
	// addCluster()
	deleteCluster()
}
