package service

import (
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/jwt"
)

type Service struct {
	AuthService    *AuthService
	ProfileService *ProfileService
	UsersService   *UserService
	WalkService    *WalkService
	RatingService  *RatingService
}

func NewService(
	authRepo iAuthRepo,
	jwtHandler jwt.JWT,
	profileRepo iProfileRepo,
	usersRepo iUsersRepo,
	walkRepo iWalkRepo,
	ratingRepo iRatingRepo,
	petTranslator iPetTranlsator,
	waklTranslator iWalkTranslator,
) *Service {
	return &Service{
		AuthService:    NewAuthService(authRepo, jwtHandler),
		ProfileService: NewProfileService(profileRepo, petTranslator),
		UsersService:   NewUserService(usersRepo),
		WalkService:    NewWalkService(walkRepo, waklTranslator),
		RatingService:  NewRatingSrvice(ratingRepo),
	}
}
