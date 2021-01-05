package main

//"encoding/json"
import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

type PipelineCreationRequestEvent struct {
	CorrelationId           string `json:"correlationId"`
	ManifestId              string `json:"manifestId"`
	EventId                 string `json:"eventId"`
	WorkloadPackageId       string `json:"workloadPackageId"`
	WorkloadPackageLocation string `json:"workloadPackageLocation"`
	Artifacts               string `json:"artifacts"`
	SecurityContext         string `json:"securityContext"`
	SecurityPermissions     string `json:"securityPermissions"`
	DataMountPath           string `json:"dataMountPath"`
	DataPVC                 string `json:"dataPVC"`
	ClusterDomain           string `json:"clusterDomain"`
	Ack                     string `json:"ack"`
}

func main() {

	correlationId := "dfsdafasdfsdafsdfdsf"
	topic := os.Args[1]
	brokers := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	event := &PipelineCreationRequestEvent{
		CorrelationId:           correlationId,
		ManifestId:              "0",
		EventId:                 "1",
		WorkloadPackageId:       "1",
		WorkloadPackageLocation: "1",
		Artifacts:               "1",
		SecurityContext:         ",",
		SecurityPermissions:     "1",
		DataMountPath:           "/foo/bar/bla",
		DataPVC:                 "0",
		ClusterDomain:           "1",
		Ack:                     "1",
	}
	message, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
		return
	}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		// Log error
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			// Log error
			panic(err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}

	// Log Message
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

}
