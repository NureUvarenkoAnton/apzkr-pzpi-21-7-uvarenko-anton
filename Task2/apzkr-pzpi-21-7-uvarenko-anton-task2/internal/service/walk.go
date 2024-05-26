package service

import (
	"context"
	"database/sql"
	"errors"

	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/core"
	"NureUvarenkoAnton/apzkr-pzpi-21-7-uvarenko-anton/Task2/apzkr-pzpi-21-7-uvarenko-anton-task2/internal/pkg"

	"github.com/go-sql-driver/mysql"
)

type WalkService struct {
	walkRepo iWalkRepo
}

func NewWalkService(walkRepo iWalkRepo) *WalkService {
	return &WalkService{
		walkRepo: walkRepo,
	}
}

type iWalkRepo interface {
	CreateWalk(ctx context.Context, arg core.CreateWalkParams) error
	GetWalksByWalkerId(ctx context.Context, walkerID sql.NullInt64) ([]core.Walk, error)
	UpdateWalkState(ctx context.Context, arg core.UpdateWalkStateParams) error
	GetWalksByOwnerId(ctx context.Context, ownerID sql.NullInt64) ([]core.Walk, error)
}

func (s *WalkService) CreateWalk(ctx context.Context, walkParams core.CreateWalkParams) error {
	err := s.walkRepo.CreateWalk(ctx, walkParams)
	if err != nil {
		err, ok := err.(*mysql.MySQLError)
		if !ok {
			pkg.PrintErr(pkg.ErrDbInternal, err)
			return pkg.ErrDbInternal
		}

		return pkg.ErrDbInternal
	}
	return nil
}

func (s *WalkService) GetWalksByWalkerId(ctx context.Context, walkerID sql.NullInt64) ([]core.Walk, error) {
	walks, err := s.walkRepo.GetWalksByWalkerId(ctx, walkerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return nil, pkg.ErrDbInternal
	}
	return walks, nil
}

func (s *WalkService) GetWalksByOwnerId(ctx context.Context, ownerID sql.NullInt64) ([]core.Walk, error) {
	walks, err := s.walkRepo.GetWalksByOwnerId(ctx, ownerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return nil, pkg.ErrDbInternal
	}

	return walks, nil
}

func (s *WalkService) UpdateWalkState(ctx context.Context, params core.UpdateWalkStateParams) error {
	err := s.walkRepo.UpdateWalkState(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return pkg.ErrDbInternal
	}
	return nil
}
