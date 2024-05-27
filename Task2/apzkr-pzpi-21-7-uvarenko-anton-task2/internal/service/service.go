package service

import (
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/jwt"
)

type Service struct {
	AuthService    *AuthService
	ProfileService *ProfileService
}

func NewService(
	authRepo iAuthRepo,
	jwtHandler jwt.JWT,
	profileRepo iProfileRepo,
) *Service {
	return &Service{
		AuthService:    NewAuthService(authRepo, jwtHandler),
		ProfileService: NewProfileService(profileRepo),
	}
}
