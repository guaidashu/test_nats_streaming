package nats_streaming_libs

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"os"
	"sync"
	"time"
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
			// stan.NatsURL("nats://127.0.0.1:4222"),
			// stan.NatsURL("nats://127.0.0.1:4222"),
			// stan.NatsURL("nats://127.0.0.1:4222"),
			stan.NatsURL(os.Getenv("NATS_URL_1")),
			stan.NatsURL(os.Getenv("NATS_URL_2")),
			stan.NatsURL(os.Getenv("NATS_URL_3")),
			func(options *stan.Options) error {
				options.AckTimeout = time.Second * 60
				options.MaxPubAcksInflight = 1000000
				return nil
			},
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

func (n *NatsPublisher) Publish(topic string, data []byte) (string, error) {
	return n.conn.PublishAsync(topic, data, func(s string, err error) {
	})
}

func (n *NatsPublisher) Subscribe(topic string, handler stan.MsgHandler, opts ...stan.SubscriptionOption) (stan.Subscription, error) {
	opts = append(opts, stan.SetManualAckMode())
	return n.conn.Subscribe(topic, handler, opts...)
}
