package server

import (
	"net/http"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/jwt"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/middleware"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/transport"

	"github.com/gin-gonic/gin"
)

func New(handler *transport.Handler, jwtHandler jwt.JWT) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: setUpRoutes(handler, jwtHandler),
	}
}

func setUpRoutes(handler *transport.Handler, jwtHandler jwt.JWT) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	authRouts := router.Group("/auth")
	{
		authRouts.POST("/register", handler.AuthHandler.RegisterUser)
		authRouts.POST("/login", handler.AuthHandler.Login)
	}

	profileRouts := router.Group("/profile")
	profileRouts.Use(middleware.TokenVerifier(jwtHandler, []core.UsersUserType{
		core.UsersUserTypeAdmin,
		core.UsersUserTypeDefault,
		core.UsersUserTypeWalker,
	}))
	{
		profileRouts.POST("/pet", handler.ProfileHandler.AddPet)
		profileRouts.PUT("/pet", handler.ProfileHandler.UpdatePet)
		profileRouts.GET("/pets", handler.ProfileHandler.GetOwnerPets)
		profileRouts.PUT("/user", handler.ProfileHandler.UpdateUser)
	}

	return router
}
