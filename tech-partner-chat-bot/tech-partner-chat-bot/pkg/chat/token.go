package chat

import (
	"context"
	"time"

	auth "golang.org/x/oauth2/google"
)

// 03
type TokenService struct {
	creationTime *int64
	token        *string
	credentials  *auth.Credentials
}

// 04
func NewTokenService() (*TokenService, error) {
	ctx := context.Background()
	scopes := []string{
		"https://www.googleapis.com/auth/chat.bot",
	}
	credentials, err := auth.FindDefaultCredentials(ctx, scopes...)
	if err != nil {
		return nil, err
	}

	return &TokenService{credentials: credentials}, nil

}

// 06
func (s *TokenService) GetToken() (*string, error) {
	now := time.Now().Unix()

	if s.token == nil || s.creationTime == nil || now-int64(*s.creationTime) > 900 {

		token, err := s.credentials.TokenSource.Token()

		if err != nil {
			return nil, err
		}

		s.token = &token.AccessToken

		s.creationTime = &now
	}

	return s.token, nil
}
