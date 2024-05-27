package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg"
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
	GetUserByUserType(ctx context.Context, userType core.NullUsersUserType) ([]core.User, error)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]core.User, error) {
	users, err := s.userRepo.GetAllUsers(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return nil, fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}
	return users, nil
}

func (s *UserService) GetUsersByUserType(ctx context.Context, userType core.NullUsersUserType) ([]core.User, error) {
	users, err := s.userRepo.GetUserByUserType(ctx, userType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return nil, fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}

	return users, nil
}

// func (s *UserService) BanUser(ctx context.Context)
