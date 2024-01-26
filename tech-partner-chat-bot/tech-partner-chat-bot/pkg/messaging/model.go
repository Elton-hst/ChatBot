package messaging

type IncomingBotMessage struct {
	Message struct {
		Text string
	}
	Space struct {
		Name string
	}
	User struct {
		Email string
	}
}
