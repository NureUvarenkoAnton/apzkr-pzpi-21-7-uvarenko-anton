package transport

import "github.com/olahol/melody"

type Handler struct {
	AuthHandler     *AuthHandler
	ProfileHandler  *ProfileHandler
	PositionHandler *PositionHandler
}

func NewHandler(
	authService iAuthService,
	profileService iUserProfileService,
	melody *melody.Melody,
) *Handler {
	return &Handler{
		AuthHandler:     NewAuthHandler(authService),
		ProfileHandler:  NewProfileHandler(profileService, melody),
		PositionHandler: NewPositionHandler(melody, profileService),
	}
}
