package app

import (
	"fmt"
	"os"
	"test_nats_streaming/nats_streaming_libs"
)

func Publisher() {
	publisher := nats_streaming_libs.NewNatsPublisher("diss-cluster", "test-cluster-123")
	var index = 1

	topic := os.Getenv("NATS_TEST_TOPIC")
	if topic == "" {
		topic = "test-1"
	}
	for {
		err := publisher.Publish(topic, []byte("{\"id\":\"111111111\"}"))
		if err != nil {
			fmt.Println(err)
		}

		index = index + 1
		if index > 300 {
			break
		}
	}

	fmt.Println(index)
}
