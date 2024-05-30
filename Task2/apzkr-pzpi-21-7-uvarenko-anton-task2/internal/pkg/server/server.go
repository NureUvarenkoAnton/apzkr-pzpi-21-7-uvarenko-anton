package server

import (
	"net/http"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/jwt"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/middleware"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/transport"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func New(handler *transport.Handler, jwtHandler jwt.JWT, melody *melody.Melody) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: setUpRoutes(handler, jwtHandler, melody),
	}
}

func setUpRoutes(handler *transport.Handler, jwtHandler jwt.JWT, melody *melody.Melody) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	authRouts := router.Group("/auth")
	{
		authRouts.POST("/register", handler.AuthHandler.RegisterUser)
		authRouts.POST("/login", handler.AuthHandler.Login)
	}

	wsOpenConnection := router.Group("/")
	wsOpenConnection.Use(middleware.TokenVerifier(jwtHandler, []core.UsersUserType{core.UsersUserTypePet, core.UsersUserTypeDefault, core.UsersUserTypeWalker}))
	wsOpenConnection.GET("/ws", handler.PositionHandler.HandleOpenPetConnection)
	melody.HandleMessage(handler.PositionHandler.HandleMessage)

	loginPetRouter := router.Group("/")
	loginPetRouter.Use(middleware.TokenVerifier(jwtHandler, []core.UsersUserType{core.UsersUserTypeDefault}))
	loginPetRouter.POST("/loginpet", handler.AuthHandler.LoginPet)

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

	usersDefaultRouts := router.Group("/users")
	usersDefaultRouts.Use(middleware.TokenVerifier(jwtHandler, []core.UsersUserType{
		core.UsersUserTypeAdmin,
		core.UsersUserTypeDefault,
		core.UsersUserTypeWalker,
	}))
	{
		usersDefaultRouts.GET("/", handler.UserHandler.GetUsers)
	}

	userAdminRouts := router.Group("/users")
	userAdminRouts.Use(middleware.TokenVerifier(jwtHandler, []core.UsersUserType{core.UsersUserTypeAdmin}))
	{
		userAdminRouts.PUT("/ban", handler.UserHandler.SetBanState)
	}

	walkRouts := router.Group("/walk")
	walkRouts.Use(middleware.TokenVerifier(jwtHandler, []core.UsersUserType{
		core.UsersUserTypeAdmin,
		core.UsersUserTypeWalker,
		core.UsersUserTypeDefault,
	}))
	{
		walkRouts.POST("/", handler.WalkHalder.CreateWalkRequest)
		walkRouts.PUT("/", handler.WalkHalder.UpdateWalkState)
		walkRouts.GET("/", handler.WalkHalder.GetWalksByParams)
		walkRouts.GET("/:id", handler.WalkHalder.GetWalkInfoById)
		walkRouts.GET("/self", handler.WalkHalder.GetWalksBySelfId)
	}

	return router
}
