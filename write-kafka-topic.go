package main

//"encoding/json"
import (
	"encoding/json"
	"fmt"

	//"encoding/json"
	//"os/exec"
	"os"

	"github.com/Shopify/sarama"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type WorkLoadRequestEvent struct {
	CorrelationId           string
	ManifestId              string
	WorkloadPackageId       string
	WorkloadPackageLocation string
	Artifacts               string
	SecurityContext         string
	SecurityPermissions     string
	DataMountPath           string
	DataPVC                 string
	ClusterDomain           string
	Ack                     string
}

func main() {

	correlationId := os.Args[2]
	topic := os.Args[1]
	brokers := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	event := &WorkLoadRequestEvent{CorrelationId: correlationId, ManifestId: "0", WorkloadPackageId: "1", WorkloadPackageLocation: "1", Artifacts: "1", SecurityContext: ",", SecurityPermissions: "1", DataMountPath: "", DataPVC: "0", ClusterDomain: "1", Ack: "1"}
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
