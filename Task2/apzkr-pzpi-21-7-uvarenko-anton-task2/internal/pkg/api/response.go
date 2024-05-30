package api

import (
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
)

type TokenResponse struct {
	Token string `json:"token"`
}

type ResponseWSMessage struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Data  any    `json:"data"`
}

type AvgRatingResponse struct {
	AvgRating int `json:"avgRating"`
}

type UserResponse struct {
	Id        int64              `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	UserType  core.UsersUserType `json:"userType"`
	IsBanned  bool               `json:"isBanned"`
	IsDeleted bool               `json:"isDelted"`
}

type WalkInfoResponse struct {
	WalkId      int64  `json:"walkId,omitempty"`
	WalkState   string `json:"walkState,omitempty"`
	StartTime   string `json:"startTime,omitempty"`
	FinishTime  string `json:"finishTime,omitempty"`
	OwnerId     int64  `json:"ownerId,omitempty"`
	OwnerName   string `json:"ownerName,omitempty"`
	OwnerEmail  string `json:"ownerEmail,omitempty"`
	WalkerId    int64  `json:"walkerId,omitempty"`
	WalkerName  string `json:"walkerName,omitempty"`
	WalkerEmail string `json:"walkerEmail,omitempty"`
	PetId       int64  `json:"petId,omitempty"`
	PetName     string `json:"petName,omitempty"`
}
