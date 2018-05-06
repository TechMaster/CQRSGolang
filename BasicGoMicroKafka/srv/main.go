package main

import (
	proto "github.com/TechMaster/microKafka/BasicGoMicroKafka/srv/proto"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"context"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/kafka"
	"github.com/micro/go-micro/server"
)

// All methods of Sub will be executed when
// a message is received
type Sub struct{}

// Method can be of any name
func (s *Sub) Process(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.1] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

// Alternatively a function can be used
func subEv(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub.2] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

func main() {
	//

	kafkaBroker := kafka.NewBroker(broker.Addrs("localhost:29092"))


	// create a service
	service := micro.NewService(
		micro.Name("go.micro.srv.pubsub"),
		micro.Broker(kafkaBroker),
	)


	// parse command line
	service.Init()


	// register subscriber
	micro.RegisterSubscriber("Foo", service.Server(), new(Sub))

	// register subscriber with queue, each message is delivered to a unique subscriber
	micro.RegisterSubscriber("Bar", service.Server(), subEv, server.SubscriberQueue("queue.pubsub"))
	//micro.RegisterSubscriber("Bar", service.Server(), new(Sub))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
