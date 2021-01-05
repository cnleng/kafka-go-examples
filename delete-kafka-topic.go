package main

import (
	// "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	// Set broker configuration
	ips := []string{"100.80.250.206:32400", "100.80.250.215:32401", "100.80.250.209:32402"}
	config := sarama.NewConfig()
	clusterAdmin, _ := sarama.NewClusterAdmin(ips, config)

	// Setup the Topic details in CreateTopicRequest struct
	// topics := []string{"WorkloadResponse", "PipelineCreationRequest", "DeployApplicationRequest",
	// 	"ApplicationDeploymentState", "PipelineCreatedOrAlreadyPresent", "ApplicationPipelineState",
	// 	"ApplicationPipelineInvoked", "DeployApplicationRequest", "Workload"}

	topic := "WorkloadResponse"
	log.Printf("Deleting topic [%s] on broker [%v] \n", topic, ips)
	err := clusterAdmin.DeleteTopic("WorkloadResponse")
	if err != nil {
		// log.Printf("%#v", &err)
		log.Printf(" An error occured while delete topic [%s] on broker [%v] \n %v", topic, ips, err)
		return
	}
	log.Printf("Topic deleted [%s] on broker [%v] \n", topic, ips)

	// for _, topic := range topics {
	// 	// Send request to Broker
	// 	log.Printf("Deleting topic [%s] on broker [%v] \n", topic, ips)

	// 	err := clusterAdmin.DeleteTopic(topic)

	// 	// handle errors if any
	// 	if err != nil {
	// 		// log.Printf("%#v", &err)
	// 		log.Printf(" An error occured while delete topic [%s] on broker [%v] \n %v", topic, ips, err)
	// 		return
	// 	}

	// 	log.Printf("Topic deleted [%s] on broker [%v] \n", topic, ips)
	// }

}
