package app

import (
	"fmt"
	"test_nats_streaming/nats_streaming_libs"
)

func Publisher() {
	publisher := nats_streaming_libs.NewNatsPublisher("diss-cluster", "test-cluster-123")
	var index = 1

	for {
		err := publisher.Publish("test-1", []byte("{\"id\":\"111111111\"}"))
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