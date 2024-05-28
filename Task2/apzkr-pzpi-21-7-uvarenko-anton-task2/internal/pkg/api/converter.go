package api

import "NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"

func DbUserToAPIUser(user core.User) UserResponse {
	return UserResponse{
		Id:        user.ID,
		Name:      user.Name.String,
		Email:     user.Email.String,
		UserType:  user.UserType.UsersUserType,
		IsBanned:  user.IsBanned.Bool,
		IsDeleted: user.IsDeleted.Bool,
	}
}

func SliceDbUserToAPIUser(users []core.User) []UserResponse {
	var result []UserResponse

	for _, user := range users {
		result = append(result, DbUserToAPIUser(user))
	}

	return result
}
