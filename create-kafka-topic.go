package main

import (
	"log"
	"time"

	// "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/Shopify/sarama"
)

func main() {
	// Set broker configuration
	ips := []string{"localhost:9092"}
	for _, ip := range ips {
		broker := sarama.NewBroker(ip)

		// Additional configurations. Check sarama doc for more info
		config := sarama.NewConfig()

		// Open broker connection with configs defined above
		broker.Open(config)

		// check if the connection was OK
		connected, err := broker.Connected()
		if err != nil {
			log.Print(err.Error())
		}
		log.Print(connected)

		// Setup the Topic details in CreateTopicRequest struct
		topics := []string{"WorkloadResponse", "PipelineCreationRequest", "DeployApplicationRequest",
			"ApplicationDeploymentState", "PipelineCreatedOrAlreadyPresent", "ApplicationPipelineState",
			"ApplicationPipelineInvoked", "DeployApplicationRequest", "Workload"}
		for _, topic := range topics {
			topicDetail := &sarama.TopicDetail{}
			topicDetail.NumPartitions = int32(4)
			topicDetail.ReplicationFactor = int16(3)
			topicDetail.ConfigEntries = make(map[string]*string)

			topicDetails := make(map[string]*sarama.TopicDetail)
			topicDetails[topic] = topicDetail

			request := sarama.CreateTopicsRequest{
				Timeout:      time.Second * 15,
				TopicDetails: topicDetails,
			}

			// Send request to Broker
			log.Printf("Creating topic [%s] on broker [%s] \n", topic, ip)
			response, err := broker.CreateTopics(&request)

			// handle errors if any
			if err != nil {
				// log.Printf("%#v", &err)
				log.Printf(" An error occured while creating topic [%s] on broker [%s] \n %v", topic, ip, err)
				return
			}
			// t := response.TopicErrors

			log.Printf("The response is %v", response)
			log.Printf("Topic created [%s] on broker [%s] \n", topic, ip)
		}

		// close connection to broker
		broker.Close()
	}

}