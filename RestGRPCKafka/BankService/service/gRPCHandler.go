package service

import (
	"fmt"
	"context"
	"time"
	"github.com/micro/go-micro"
	"github.com/satori/go.uuid"
	proto "github.com/TechMaster/microKafka/RestGRPCKafka/proto"
	protobuf "github.com/golang/protobuf/proto"
	google_protobuf2 "github.com/gogo/protobuf/types"
	"reflect"
)

//
type BankService struct {
	BankAccountPub micro.Publisher
}

const serviceDomain = "BankService"
/*
1- Receive Command
2- Validate Command:
	- may check
	- if success then write to event store
	- if fail then report error
*/
func (bs *BankService) CreateAccount(ctx context.Context, req *proto.NewAccount, resp *proto.Response ) error {
	//Input data validation
	if len(req.Name) < 4 || len(req.Name) > 32 {
		resp.ErrorCode = 100
		resp.ErrorDesc = "Account name is invalid"
		resp.Account = nil
		return nil
	}

	//Insert into database
	id := uuid.Must(uuid.NewV4()).String()
	var account proto.Account
	account.Id = id
	account.Name = req.Name
	account.Balance = req.Balance
	//Try to save to database
	if err := Db.Insert(&account); err != nil {
		fmt.Println(err)
		resp.ErrorCode = 200
		resp.ErrorDesc = "Fail to insert: " + err.Error()
		return nil
	}

	//Use to return to client
	resp.ErrorCode = 0
	resp.ErrorDesc = ""
	resp.Account = &account

	//Append into Event Store
	//Later we must reconstruct current state of object by select of event records that has
	//same event.id (aggreagate ID)
	var event proto.Event
	event.Id = id
	event.Type = "BankService.AccountCreated"
	event.Timestamp = time.Now().Unix()
	serialized, err := protobuf.Marshal(&account)
	if err != nil {
		fmt.Println(err)
	}

	//See this tutorial https://medium.com/@pokstad/sending-any-any-thing-in-golang-with-protobuf-3-95f84838028d
	event.Payload = &google_protobuf2.Any{
		TypeUrl:"github.com/TechMaster/microKafka/RestGRPCKafka/proto/" + reflect.TypeOf(account).String(),
		Value: serialized,
	}

	//Append event record into back end database table events
	if err := Db.Insert(&event); err != nil {
		resp.ErrorCode = 300
		resp.ErrorDesc = "Fail to insert event store: " + err.Error()
		return nil
	}


	//Publish event to Kafka
	if err := bs.BankAccountPub.Publish(context.Background(), &event); err != nil {
		fmt.Println("error publishing %v", err)
	}
	return nil
}