package main

//"encoding/json"
import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

type DeployApplicationEvent struct {
	CorrelationId           string `json:"correlationId"`
	ManifestId              string `json:"manifestId"`
	WorkloadPackageId       string `json:"workloadPackageId"`
	WorkloadPackageLocation string `json:"workloadPackageLocation"`
	PipelineURL             string `json:"pipelineURL"`
	EventId                 string `json:"eventId"`
}

func main() {

	correlationId := "dfsdafasdfsdafsdfdsf"
	topic := os.Args[1]
	brokers := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	event := &DeployApplicationEvent{
		CorrelationId:           correlationId,
		ManifestId:              "0",
		EventId:                 "1",
		WorkloadPackageId:       "1",
		WorkloadPackageLocation: "1",
		PipelineURL:             "/foo/bar/2",
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
