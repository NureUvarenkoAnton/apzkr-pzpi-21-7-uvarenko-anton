package transport

import "github.com/olahol/melody"

type Handler struct {
	AuthHandler     *AuthHandler
	ProfileHandler  *ProfileHandler
	PositionHandler *PositionHandler
	UserHandler     *UserHandler
	WalkHalder      *WalkHalder
}

func NewHandler(
	authService iAuthService,
	profileService iUserProfileService,
	melody *melody.Melody,
	userService iUserService,
	walkService iWalkService,
) *Handler {
	return &Handler{
		AuthHandler:     NewAuthHandler(authService),
		ProfileHandler:  NewProfileHandler(profileService, melody),
		PositionHandler: NewPositionHandler(melody, profileService),
		UserHandler:     NewUserHandler(userService),
		WalkHalder:      NewWalkHandler(walkService),
	}
}
