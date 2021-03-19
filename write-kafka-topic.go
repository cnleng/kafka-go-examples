package main

//"encoding/json"
import (
	"encoding/json"
	"fmt"

	//"time"
	//"os/exec"

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

	//correlationId := os.Args[2]
	// topic := os.Args[1]
	// topic := "ClusterCredentialsRequest"
	topic := "PROVISION_CLUSTER_COMPLETED"
	// brokers := []string{"192.168.204.128:9092"}
	brokers := []string{"100.80.250.216:9092"}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5
	//config.Producer.Flush.Frequency = 500 * time.Millisecond

	metadata := map[string]interface{}{
		"serverAddr": "http://20.62.168.104:8200",
		"secretPath": "secret/data/cncp/orchestration/01EYKSNGX3F6ZZPTSWYKVR04JG",
	}
	clusterProvisionCompletedEvent := map[string]interface{}{
		"correlationId":       "01EY97B2HDJ4E065NWSX092S98",
		"messageId":           "01EYVC0B46S11S8V1Y54XNGFMK",
		"workloadId":          "01EY97B2HDJ4E065NWSX092S98",
		"clusterUrl":          "01EZ8DHG6YCT7J5YNP2FNARMNC",
		"clusterId":           "01EYKSNGX3F6ZZPTSWYKVR04JG",
		"platform":            "OnPrem",
		"credentialsStore":    "HashiCorp Vault",
		"credentialsMetadata": metadata,
		"status":              "Complete",
		"statusMessage":       "ghjghjgjhg",
	}

	// clusterCredentialsRequest := map[string]interface{}{
	// 	"EventID":       0,
	// 	"CorrelationID": "01EY97B2HDJ4E065NWSX092S98",
	// 	"ManifestId":    "01EY",
	// 	"WorkloadID":    "01",
	// 	"ClusterID":     "01EYKSNGX3F6ZZPTSWYKVR04JG",
	// 	"PipelineID":    "FD@#2",
	// }

	message, err := json.Marshal(clusterProvisionCompletedEvent)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Initializing Producer ...")
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		// Log error
		panic(err)
	}
	//fmt.Println("Init done.")
	defer func() {
		if err := producer.Close(); err != nil {
			// Log error
			panic(err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		//Partition: 1,
		Value: sarama.StringEncoder(message),
	}
	fmt.Printf("Pushing data Kafka Topic %v\n", topic)
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}

	// Log Message
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

}
