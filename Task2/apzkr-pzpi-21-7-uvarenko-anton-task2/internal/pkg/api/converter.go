package api

import (
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
)

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

func DbWalkInfoToAPIWalkInfo(walkInfo core.WalkInfo) WalkInfoResponse {
	result := WalkInfoResponse{
		WalkId:      walkInfo.WalkID,
		OwnerId:     walkInfo.OwnerID,
		OwnerName:   walkInfo.OwnerName.String,
		OwnerEmail:  walkInfo.OwnerEmail.String,
		WalkerId:    walkInfo.WalkerID,
		WalkerName:  walkInfo.WalkerName.String,
		WalkerEmail: walkInfo.WalkerEmail.String,
		PetId:       walkInfo.PetID,
		PetName:     walkInfo.PetName.String,
	}

	if walkInfo.StartTime.Valid {
		result.StartTime = walkInfo.StartTime.Time.String()
	}

	if walkInfo.FinishTime.Valid {
		result.FinishTime = walkInfo.FinishTime.Time.String()
	}

	return result
}

func SliceDbWalkInfoToAPIWalkInfo(walksInfo []core.WalkInfo) []WalkInfoResponse {
	var result []WalkInfoResponse
	for _, walkInfo := range walksInfo {
		result = append(result, DbWalkInfoToAPIWalkInfo(walkInfo))
	}
	return result
}
