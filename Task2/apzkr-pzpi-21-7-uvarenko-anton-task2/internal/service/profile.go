package service

import (
	"context"
	"database/sql"
	"errors"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg"

	"github.com/go-sql-driver/mysql"
)

type ProfileService struct {
	userRepo iProfileUserRepo
}

func NewProfileService(userRepo iProfileUserRepo) *ProfileService {
	return &ProfileService{
		userRepo: userRepo,
	}
}

type iProfileUserRepo interface {
	GetPetById(ctx context.Context, id int64) (core.Pet, error)
	UpdatePet(ctx context.Context, arg core.UpdatePetParams) error
	GetAllPetsByOwnerId(ctx context.Context, ownerID sql.NullInt64) ([]core.Pet, error)
	AddPet(ctx context.Context, arg core.AddPetParams) error
	UpdateUser(ctx context.Context, arg core.UpdateUserParams) error
}

func (s ProfileService) AddPet(ctx context.Context, pet core.AddPetParams) error {
	err := s.userRepo.AddPet(ctx, pet)
	if err != nil {
		err, ok := err.(*mysql.MySQLError)
		if !ok {
			pkg.PrintErr(pkg.ErrDbInternal, err)
			return pkg.ErrDbInternal
		}
		// if err.Number ==

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return pkg.ErrDbInternal
	}
	return err
}

func (s ProfileService) GetPetById(ctx context.Context, id int64) (*core.Pet, error) {
	pet, err := s.userRepo.GetPetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return nil, pkg.ErrDbInternal
	}

	return &pet, nil
}

func (s ProfileService) GetAllPetsByOwnerId(ctx context.Context, ownerID sql.NullInt64) ([]core.Pet, error) {
	pets, err := s.userRepo.GetAllPetsByOwnerId(ctx, ownerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return nil, pkg.ErrDbInternal
	}

	return pets, nil
}

func (s ProfileService) UpdatePet(ctx context.Context, pet core.UpdatePetParams) error {
	// dbPet, err := s.userRepo.GetPetById(ctx, pet.ID)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		return pkg.ErrNotFound
	// 	}
	//
	// 	pkg.PrintErr(pkg.ErrDbInternal, err)
	// 	return err
	// }

	err := s.userRepo.UpdatePet(ctx, pet)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return pkg.ErrDbInternal

	}

	return nil
}

func (s ProfileService) UpdateUserData(ctx context.Context, userData core.UpdateUserParams) error {
	err := s.userRepo.UpdateUser(ctx, userData)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pkg.ErrNotFound
		}

		err, ok := err.(*mysql.MySQLError)
		if !ok {
			pkg.PrintErr(pkg.ErrDbInternal, err)
			return pkg.ErrDbInternal
		}
		if err.Number == 1062 {
			return pkg.ErrEmailDuplicate
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return pkg.ErrDbInternal
	}
	return nil
}
