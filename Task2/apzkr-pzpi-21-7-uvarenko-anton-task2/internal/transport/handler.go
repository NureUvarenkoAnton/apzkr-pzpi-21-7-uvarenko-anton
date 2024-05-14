package transport

type Handler struct {
	AuthHandler    *AuthHandler
	ProfileHandler *ProfileHandler
}

func NewHandler(
	authService iAuthService,
	profileService iUserProfileService,
) *Handler {
	return &Handler{
		AuthHandler:    NewAuthHandler(authService),
		ProfileHandler: NewProfileHandler(profileService),
	}
}
