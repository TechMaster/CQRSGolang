package main

import (
	"log"
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/kafka"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/metadata"

	proto "github.com/TechMaster/microKafka/RestGRPCKafka/proto"
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

func main() {

}
