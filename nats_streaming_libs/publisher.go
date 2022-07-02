package nats_streaming_libs

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"os"
	"sync"
)

var ClientNum = 10

type (
	NatsPublisher struct {
		conn stan.Conn
	}
)

var (
	_natsPublisher     []*NatsPublisher
	_natsPublisherOnce sync.Once
)

func NewNatsPublisher(id int) *NatsPublisher {
	_natsPublisherOnce.Do(func() {
		var err error
		_natsPublisher = make([]*NatsPublisher, 0, ClientNum)

		for i := 0; i < ClientNum; i++ {
			c := &NatsPublisher{}
			c.conn, err = stan.Connect("diss-cluster", fmt.Sprintf("test-cluster-%v", i),
				// c.conn, err = stan.Connect("test-cluster", fmt.Sprintf("test-cluster-id%v", i),
				stan.NatsURL(os.Getenv("NATS_URL_1")),
				stan.NatsURL(os.Getenv("NATS_URL_2")),
				stan.NatsURL(os.Getenv("NATS_URL_3")),
				// stan.NatsURL("nats://127.0.0.1:4222"),
			)
			if err != nil {
				fmt.Println(err)
				continue
			}

			_natsPublisher = append(_natsPublisher, c)
		}
	})

	return _natsPublisher[id]
}

func (n *NatsPublisher) PublishAsync(topic string, data []byte) (string, error) {
	return n.conn.PublishAsync(topic, data, func(s string, err error) {
		if err != nil {
			fmt.Println(err)
		}
	})
}

func (n *NatsPublisher) Publish(topic string, data []byte) error {
	return n.conn.Publish(topic, data)
}

func (n *NatsPublisher) Subscribe(topic string, handler stan.MsgHandler, opts ...stan.SubscriptionOption) (stan.Subscription, error) {
	opts = append(opts, stan.SetManualAckMode())
	return n.conn.Subscribe(topic, handler, opts...)
}
