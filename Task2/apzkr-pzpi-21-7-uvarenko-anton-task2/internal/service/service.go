package service

import "firebase.google.com/go/v4/auth"

type Service struct {
	AuthService *AuthService
}

func NewService(authRepo iAuthUserRepo, authClient *auth.Client) *Service {
	return &Service{
		AuthService: NewAuthService(authRepo, authClient),
	}
}
