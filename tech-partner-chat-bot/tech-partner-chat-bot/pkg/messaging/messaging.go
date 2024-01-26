package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"ciandt.com/techPartnerReponseHandler/pkg/chat"
	pubsub "cloud.google.com/go/pubsub"
)

type Messaging struct {
	sub         *pubsub.Subscription
	context     context.Context
	chatService *chat.ChatService
}

func CreateNewMessaging(chatService *chat.ChatService) (*Messaging, error) {
	client, err := pubsub.NewClient(context.Background(), "itau-techpartner")
	if err != nil {
		return nil, err
	}

	sub := client.Subscription("tech-partner-incoming-messages-sub")
	return &Messaging{
		sub:         sub,
		context:     context.Background(),
		chatService: chatService,
	}, nil
}

func (m *Messaging) handleMessage(ctx context.Context, msg *pubsub.Message) {
	defer msg.Ack()

	fmt.Println("Mensagem recebida")

	parsed := &IncomingBotMessage{}
	err := json.Unmarshal(msg.Data, parsed)
	if err != nil {
		fmt.Errorf("Error when parsing message %v", err)
		return
	}

	err = m.chatService.SendTextMessage(parsed.Space.Name, fmt.Sprintf("A mensagem que voce enviou foi: %v", parsed.Message.Text))
	if err != nil {
		fmt.Errorf("Error when parsing message %v", err)
		return
	}

	fmt.Println("Mensagem processada com sucesso")
}

func (m *Messaging) Start() {
	m.sub.Receive(m.context, m.handleMessage)
}

func (m *Messaging) Stop() {
	m.context.Done()
}
