package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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
	GetWalkById(ctx context.Context, id int64) (core.Walk, error)
	GetWalkInfoByParams(ctx context.Context, arg core.GetWalkInfoByParamsParams) ([]core.WalkInfo, error)
	GetWalkInfoByWalkId(ctx context.Context, walkID int64) (core.WalkInfo, error)
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
	_, err := s.walkRepo.GetWalkById(ctx, params.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}

	err = s.walkRepo.UpdateWalkState(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return pkg.ErrDbInternal
	}
	return nil
}

func (s *WalkService) GetWalksInfoByParams(
	ctx context.Context,
	params core.GetWalkInfoByParamsParams,
) ([]core.WalkInfo, error) {
	info, err := s.walkRepo.GetWalkInfoByParams(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return nil, fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}

	return info, nil
}

func (s *WalkService) GetWalkInfoByWalkId(ctx context.Context, walkId int64) (core.WalkInfo, error) {
	walkInfo, err := s.walkRepo.GetWalkInfoByWalkId(ctx, walkId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return core.WalkInfo{}, pkg.ErrNotFound
		}

		pkg.PrintErr(pkg.ErrDbInternal, err)
		return core.WalkInfo{}, fmt.Errorf("%w: [%w]", pkg.ErrDbInternal, err)
	}

	return walkInfo, nil
}
