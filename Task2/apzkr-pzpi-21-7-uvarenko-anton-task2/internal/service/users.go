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
	SetBanState(ctx context.Context, arg core.SetBanStateParams) error
	DeleteUser(ctx context.Context, id int64) error
	GetUserById(ctx context.Context, id int64) (core.User, error)
	GetUsers(ctx context.Context, arg core.GetUsersParams) ([]core.User, error)
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

func (s *UserService) GetUsers(ctx context.Context, params core.GetUsersParams) ([]core.User, error) {
	// if no paramters provided, then return all users
	if params.ID == 0 &&
		!params.Name.Valid &&
		!params.IsBanned.Valid &&
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

	return users, nil
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
