package nats_streaming_libs

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"sync"
)

type (
	NatsPublisher struct {
		conn stan.Conn
	}
)

var (
	_natsPublisher     *NatsPublisher
	_natsPublisherOnce sync.Once
)

func NewNatsPublisher(stanClusterID, clientID string) *NatsPublisher {
	_natsPublisherOnce.Do(func() {
		var err error
		_natsPublisher = &NatsPublisher{}
		// _natsPublisher.conn, err = stan.Connect(stanClusterID, clientID,
		// 	stan.NatsURL("nats://139.155.70.249:4222"),
		// 	stan.NatsURL("nats://139.155.70.249:4223"),
		// 	stan.NatsURL("nats://139.155.70.249:4224"),
		// )

		_natsPublisher.conn, err = stan.Connect(stanClusterID, clientID,
			stan.NatsURL("nats://127.0.0.1:4222"),
			stan.NatsURL("nats://127.0.0.1:4223"),
			stan.NatsURL("nats://127.0.0.1:4224"),
		)

		// _natsPublisher.conn, err = stan.Connect("test-cluster",
		// 	"client-124",
		// 	stan.NatsURL("nats://127.0.0.1:4222"),
		// )
		if err != nil {
			fmt.Println(err)
		}
	})

	return _natsPublisher
}

func (n *NatsPublisher) Publish(topic string, data []byte) error {
	return n.conn.Publish(topic, data)
}

func (n *NatsPublisher) Subscribe(topic string, handler stan.MsgHandler, opts ...stan.SubscriptionOption) (stan.Subscription, error) {
	return n.conn.Subscribe(topic, handler, opts...)
}
