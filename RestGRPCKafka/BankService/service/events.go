package service

import (
	proto "github.com/TechMaster/microKafka/RestGRPCKafka/proto"
)

// Base Event
type Event struct {
	id string
	Type  string
}

/* each event expressed as its own struct */
type CreateAccountEvent struct {
	Event
	newAccount *proto.Account
}
