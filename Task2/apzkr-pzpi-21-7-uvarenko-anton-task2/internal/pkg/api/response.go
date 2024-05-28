package api

import "NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"

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
