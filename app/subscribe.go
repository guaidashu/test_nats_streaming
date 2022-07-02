package app

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"os"
	"os/signal"
	"test_nats_streaming/nats_streaming_libs"
	"time"
)

func Subscribe() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	topic := os.Getenv("NATS_TEST_TOPIC")
	if topic == "" {
		topic = "test-1"
	}
	subscribe := nats_streaming_libs.NewNatsPublisher(0)
	sub, err := subscribe.Subscribe(topic, func(msg *stan.Msg) {
		fmt.Println(string(msg.Data))
		// 确认消息
		e := msg.Ack()
		if e != nil {
			fmt.Println(e)
		}
	}, stan.AckWait(time.Second*60), stan.DurableName("test-group123"))
	if err != nil {
		fmt.Println(err)
	}

	<-c

	_ = sub.Close()
	_ = sub.Unsubscribe()
}
