package main

import (
	"os"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/db"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/jwt"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/server"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/service"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/transport"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := db.Connect()
	repo := core.New(db)
	jwtHandler := jwt.NewJWT(os.Getenv("JWT_SECRET"))
	service := service.NewService(
		repo,
		*jwtHandler,
		repo,
	)
	handler := transport.NewHandler(
		service.AuthService,
		service.ProfileService,
	)

	s := server.New(handler, *jwtHandler)
	s.ListenAndServe()
}
