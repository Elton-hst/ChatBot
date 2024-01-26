package main

import (
	"log"

	"ciandt.com/techPartnerReponseHandler/pkg/chat"
	"ciandt.com/techPartnerReponseHandler/pkg/messaging"
)

func main() {
	ts, err := chat.NewTokenService()
	if err != nil {
		log.Fatalf("Could not create token service %v", err)
		return
	}
	cs := chat.NewChatService(ts)
	msg, err := messaging.CreateNewMessaging(cs)

	if err != nil {
		log.Fatalf("Error creating messaging %v", err)
		return
	}

	msg.Start()
}
