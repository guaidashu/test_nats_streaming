package app

import (
	"fmt"
	"os"
	"test_nats_streaming/nats_streaming_libs"
)

func Publisher(num int) {
	var index = 1

	topic := os.Getenv("NATS_TEST_TOPIC")
	if topic == "" {
		topic = "test-1"
	}
	for {
		publisher := nats_streaming_libs.NewNatsPublisher(index % nats_streaming_libs.ClientNum)
		err := publisher.Publish(topic, []byte(fmt.Sprintf("{\"id\":\"%v\"}", index)))
		if err != nil {
			fmt.Println(err)
		}

		index = index + 1
		if index > num {
			break
		}
	}

	fmt.Println(index)
}

func PublishAsync(num int) {
	var index = 1

	topic := os.Getenv("NATS_TEST_TOPIC")
	if topic == "" {
		topic = "test-1"
	}
	for {
		publisher := nats_streaming_libs.NewNatsPublisher(index % nats_streaming_libs.ClientNum)
		_, err := publisher.PublishAsync(topic, []byte(fmt.Sprintf("{\"id\":\"%v\"}", index)))
		if err != nil {
			fmt.Println(err)
		}

		index = index + 1
		if index > num {
			break
		}
	}

	fmt.Println(index)
}
