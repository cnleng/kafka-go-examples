package main

import (
	// "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	// Set broker configuration
	ips := []string{"100.80.250.185:32400", "100.80.250.185:32401", "100.80.250.185:32402"}
	config := sarama.NewConfig()
	clusterAdmin, _ := sarama.NewClusterAdmin(ips, config)

	// Setup the Topic details in CreateTopicRequest struct
	// topics := []string{"CNCP_SECURITY_MANAGER_CREDENTIAL_METADATA", "CNCP_CLUSTERS_INSTALLED", "CNCP_SECURITY_MANAGER_CLUSTERS_STATUS",
	// 	"CNCP_PIPELINE_INSTALLED", "CNCP_PIPELINE_INSTALLED_URL", "CNCP_SECURITY_MANAGER_ACQUIRED_CREDENTIALS"}
	topics := []string{"CNCP_SECURITY_MANAGER_ACQUIRED_CREDENTIALS"}

	// topics := []string{"WorkloadResponse", "PipelineCreationRequest", "DeployApplicationRequest",
	// 	"ApplicationDeploymentState", "PipelineCreatedOrAlreadyPresent", "ApplicationPipelineState",
	// 	"ApplicationPipelineInvoked", "DeployApplicationRequest", "Workload",
	// 	"ClusterCredentialsRequest",
	// 	"ClusterCredentialsAcquired"}

	// topic := "WorkloadResponse"
	// log.Printf("Deleting topic [%s] on broker [%v] \n", topic, ips)
	// err := clusterAdmin.DeleteTopic("WorkloadResponse")
	// if err != nil {
	// 	// log.Printf("%#v", &err)
	// 	log.Printf(" An error occured while delete topic [%s] on broker [%v] \n %v", topic, ips, err)
	// 	return
	// }
	// log.Printf("Topic deleted [%s] on broker [%v] \n", topic, ips)

	for _, topic := range topics {
		// Send request to Broker
		log.Printf("Deleting topic [%s] on broker [%v] \n", topic, ips)

		err := clusterAdmin.DeleteTopic(topic)

		// handle errors if any
		if err != nil {
			// log.Printf("%#v", &err)
			log.Printf(" An error occured while delete topic [%s] on broker [%v] \n %v", topic, ips, err)
			//return
		}

		log.Printf("Topic deleted [%s] on broker [%v] \n", topic, ips)
	}

}
