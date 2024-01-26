package chat

import (
	"fmt"
	"net/http"
	"strings"
)

var baseUrl = "https://chat.googleapis.com/v1"

// 02
type ChatService struct {
	tokenService *TokenService
}

func NewChatService(tokenService *TokenService) *ChatService {
	return &ChatService{
		tokenService: tokenService,
	}
}

func (s *ChatService) SendTextMessage(space string, message string) error {
	body := SimpleTextMessage(message)
	url := fmt.Sprintf("%v/%v/messages?threadKey=teste", baseUrl, space)
	if err := s.sendRequest(url, body); err != nil {
		return err
	}
	return nil
}

func (s *ChatService) sendRequest(url string, body string) error {

	request, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		return err
	}

	token, err := s.tokenService.GetToken()
	if err != nil {
		return err
	}

	request.Header.Add("Authorization", "Bearer "+string(*token))

	client := &http.Client{}

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
