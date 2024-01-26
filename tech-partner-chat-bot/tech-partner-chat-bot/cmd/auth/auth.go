package main

import (
	"fmt"
	"log"

	"ciandt.com/techPartnerReponseHandler/pkg/chat"
)

func main() {

	// 01
	service, err := chat.NewTokenService()

	if err != nil {
		log.Fatal("Failed to create token service")
		return
	}
	// 05
	token, err := service.GetToken()

	fmt.Println(*token)
	fmt.Println(service.GetToken())
	fmt.Println(service.GetToken())

}
