package server

import (
	"net/http"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/transport"

	"github.com/gin-gonic/gin"
)

func New(handler *transport.Handler) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: setUpRoutes(handler),
	}
}

func setUpRoutes(handler *transport.Handler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	authRouts := router.Group("/auth")
	{
		authRouts.POST("/register", handler.AuthHandler.RegisterUser)
		authRouts.POST("/login", handler.AuthHandler.Login)
	}

	return router
}
