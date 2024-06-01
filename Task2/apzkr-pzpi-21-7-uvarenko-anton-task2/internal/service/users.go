package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/api"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg/statistics"
)

type UserService struct {
	userRepo iUsersRepo
}

func NewUserService(userRepo iUsersRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

type iUsersRepo interface {
	GetAllUsers(ctx context.Context) ([]core.User, error)
	SetBanState(ctx context.Context, arg core.SetBanStateParams) error
	DeleteUser(ctx context.Context, id int64) error
	GetUserById(ctx context.Context, id int64) (core.User, error)
	GetUsers(ctx context.Context, arg core.GetUsersParams) ([]core.User, error)
	RatingsByRateeId(ctx context.Context, rateeID sql.NullInt32) ([]core.Rating, error)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]api.UserResponse, error) {
	users, err := s.userRepo.GetAllUsers(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return nil, fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}

	var usersResponse []api.UserResponse
	for _, user := range users {
		ratings, err := s.userRepo.RatingsByRateeId(ctx, sql.NullInt32{Int32: int32(user.ID), Valid: true})
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
		}
		var ratingValues []int32
		for _, rating := range ratings {
			ratingValues = append(ratingValues, rating.Value.Int32)
		}

		usersResponse = append(usersResponse, api.UserResponse{
			Id:        user.ID,
			Name:      user.Name.String,
			Email:     user.Email.String,
			UserType:  user.UserType.UsersUserType,
			AvgRating: statistics.AvgWeighted(ratingValues),
			IsBanned:  user.IsBanned.Bool,
			IsDeleted: user.IsBanned.Bool,
		})
	}

	return usersResponse, nil
}

func (s *UserService) GetById(ctx context.Context, id int64, requesterType core.UsersUserType) (api.UserResponse, error) {
	user, err := s.userRepo.GetUsers(ctx, core.GetUsersParams{
		IsBanned:  sql.NullBool{Bool: false, Valid: requesterType != core.UsersUserTypeAdmin},
		IsDeleted: sql.NullBool{Bool: false, Valid: requesterType != core.UsersUserTypeAdmin},
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return api.UserResponse{}, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return api.UserResponse{}, fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}

	if requesterType == core.UsersUserTypeDefault && user[0].UserType.UsersUserType != core.UsersUserTypeWalker {
		return api.UserResponse{}, pkg.ErrForbiden
	}

	ratings, err := s.userRepo.RatingsByRateeId(ctx, sql.NullInt32{Int32: int32(user[0].ID), Valid: true})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		pkg.PrintErr(pkg.ErrDbInternal, err)
		return api.UserResponse{}, fmt.Errorf("%w: %w", pkg.ErrDbInternal, err)
	}
	var ratingValues []int32
	for _, rating := range ratings {
		ratingValues = append(ratingValues, rating.Value.Int32)
	}

	return api.UserResponse{
		Id:        user[0].ID,
		Name:      user[0].Name.String,
		Email:     user[0].Email.String,
		UserType:  user[0].UserType.UsersUserType,
		AvgRating: statistics.AvgWeighted(ratingValues),
		IsBanned:  user[0].IsBanned.Bool,
		IsDeleted: user[0].IsBanned.Bool,
	}, nil
}

func (s *UserService) GetUsers(ctx context.Context, params core.GetUsersParams) ([]api.UserResponse, error) {
	// if no paramters provided, then return all users
	if !params.IsBanned.Valid &&
		!params.UserType.Valid &&
		!params.IsDeleted.Valid {

		return s.GetAllUsers(ctx)
	}

	users, err := s.userRepo.GetUsers(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return nil, fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}

	var usersResponse []api.UserResponse
	for _, user := range users {
		ratings, err := s.userRepo.RatingsByRateeId(ctx, sql.NullInt32{Int32: int32(user.ID), Valid: true})
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
		}
		var ratingValues []int32
		for _, rating := range ratings {
			ratingValues = append(ratingValues, rating.Value.Int32)
		}

		usersResponse = append(usersResponse, api.UserResponse{
			Id:        user.ID,
			Name:      user.Name.String,
			Email:     user.Email.String,
			UserType:  user.UserType.UsersUserType,
			AvgRating: statistics.AvgWeighted(ratingValues),
			IsBanned:  user.IsBanned.Bool,
			IsDeleted: user.IsBanned.Bool,
		})
	}

	return usersResponse, nil
}

func (s *UserService) BanUser(ctx context.Context, id int64) error {
	user, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}

	err = s.userRepo.SetBanState(ctx, core.SetBanStateParams{
		IsBanned: sql.NullBool{Bool: !user.IsBanned.Bool, Valid: true},
		ID:       user.ID,
	})
	if err != nil {
		pkg.PrintErr(pkg.ErrDbInternal, err)
		return fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}

	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	err := s.userRepo.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}
	return nil
}
