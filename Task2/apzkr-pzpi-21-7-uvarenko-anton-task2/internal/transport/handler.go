package transport

import (
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/service"

	"firebase.google.com/go/v4/auth"
)

type Handler struct {
	AuthHandler *AuthHandler
}

func NewHandler(service *service.Service, authClient *auth.Client) *Handler {
	return &Handler{
		AuthHandler: NewAuthHandler(service.AuthService, authClient),
	}
}
