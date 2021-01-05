package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"eos2git.cec.lab.emc.com/OCTO-ICV/iil-gc-contracts/gc"
	"github.com/oklog/ulid"
)

type WorkloadBody struct {
	directiveLocation string
	directiveKey      string
}

type AppliedWorkloadTelemetry struct {
	CorrelationID string `json:"Correlation_id"`
	EventID       string `json:"Event_id"`
	EventType     string `json:"Event_type"`
	WorkloadType  string `json:"Workload_type"`
	WorkloadID    string `json:"Workload_id"`
	Created       string `json:"Created"`
}

// getUULID
func getUULID() string {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
	// Output format: 0000XSNJG0MQJHBF4QX1EFD6Y3
}

func createAppliedWorkloadTelemetry() (evs []AppliedWorkloadTelemetry, err error) {

	evs = []AppliedWorkloadTelemetry{}
	coreID := getUULID()
	for i := 0; i < 10; i++ {

		timestamp := time.Now().Format(time.RFC850)
		messageValue := AppliedWorkloadTelemetry{
			CorrelationID: coreID,
			EventID:       getUULID(),
			EventType:     "appliedWorkloadTelemetry",
			WorkloadType:  "application",
			WorkloadID:    getUULID(),
			Created:       timestamp,
		}
		evs = append(evs, messageValue)
	}
	return evs, nil
}

func main() {
	fmt.Println("Receiving...")
	source := "RegisteredSource123456"
	body := WorkloadBody{"https://mys3.us", "playbook/abcdefg"}
	bodyBytes, _ := json.Marshal(body)
	event := gc.NewReceivedWorkload(source, bodyBytes)
	fmt.Println(event)
}
