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

	subscribe := nats_streaming_libs.NewNatsPublisher("diss-cluster", "test-cluster124")
	sub, err := subscribe.Subscribe("test-1", func(msg *stan.Msg) {
		fmt.Println(string(msg.Data))
		// 确认消息
		e := msg.Ack()
		if e != nil {
			fmt.Println(e)
		}
	}, stan.AckWait(time.Second*20))
	if err != nil {
		fmt.Println(err)
	}

	<-c

	_ = sub.Close()
	_ = sub.Unsubscribe()
}
