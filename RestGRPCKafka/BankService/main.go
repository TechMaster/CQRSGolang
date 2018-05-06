package main

import (
	"log"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/kafka"
	"github.com/micro/go-micro/broker"

	proto "github.com/TechMaster/microKafka/RestGRPCKafka/proto"
	srv "github.com/TechMaster/microKafka/RestGRPCKafka/BankService/service"

)

func main() {

	//Create Kafka broker
	kafkaBroker := kafka.NewBroker(broker.Addrs("localhost:29092"))


	service := micro.NewService(
		micro.Name("BankService"),
		micro.Broker(kafkaBroker),
	)

	service.Init()

	bankService := new(srv.BankService)

	proto.RegisterBankServiceHandler(service.Server(), bankService)

	bankService.BankAccountPub = micro.NewPublisher("BankAccount", service.Client())

	if  err := srv.ConnectDb(); err != nil {
		log.Fatal(err)
	} else {
		srv.InitSchema()
	}

	defer srv.Db.Close()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

