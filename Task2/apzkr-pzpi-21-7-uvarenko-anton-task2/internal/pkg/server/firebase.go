package server

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func NewFirebaseAuth() *auth.Client {
	opt := option.WithCredentialsFile("./keys/pet-walker-1232f-firebase-adminsdk-iz2lc-9142fc8727.json")
	config := &firebase.Config{ProjectID: os.Getenv("FIREBASE_PROJECT_ID")}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}

	authClinet, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error retrieving auth client: %v\n", err)
	}
	return authClinet
}
