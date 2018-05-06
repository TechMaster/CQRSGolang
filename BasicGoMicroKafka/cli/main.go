package main

import (
	"fmt"
	"time"

	proto "github.com/TechMaster/microKafka/BasicGoMicroKafka/srv/proto"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/pborman/uuid"
	"context"
	"github.com/micro/go-plugins/broker/kafka"
	"github.com/micro/go-micro/broker"
)

// send events using the publisher
func sendEv(topic string, p micro.Publisher) {
	t := time.NewTicker(time.Second)

	for _ = range t.C {
		// create new event
		ev := &proto.Event{
			Id:        uuid.NewUUID().String(),
			Timestamp: time.Now().Unix(),
			Message:   fmt.Sprintf("Messaging you all day on %s", topic),
		}

		log.Logf("publishing %+v\n", ev)

		// publish an event
		if err := p.Publish(context.Background(), ev); err != nil {
			log.Logf("error publishing %v", err)
		}
	}
}


func main() {
	//Create Kafka broker
	kafkaBroker := kafka.NewBroker(broker.Addrs("localhost:29092"))

	// create a service
	service := micro.NewService(
		micro.Name("go.micro.cli.pubsub"),
		micro.Broker(kafkaBroker),
	)
	// parse command line
	service.Init()

	// create publisher
	pub1 := micro.NewPublisher("Foo", service.Client())
	pub2 := micro.NewPublisher("Bar", service.Client())

	// pub to topic 1
	go sendEv("Foo", pub1)
	// pub to topic 2
	go sendEv("Bar", pub2)

	// block forever
	select {}
}
